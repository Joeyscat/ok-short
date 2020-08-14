package admin

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/joeyscat/ok-short/common"
	. "github.com/joeyscat/ok-short/store"
	"log"
	"net/http"
	"strconv"
	"time"
)

type UserService struct {
	R RedisCli
}

const (
	TokenFreshTime = 60 // minutes
)

// 管理员
type User struct {
	gorm.Model
	Name      string
	Password  string
	Email     string `gorm:"type:varchar(100);unique_index"`
	AvatarURL string
}

// 将 User 的表名设置为 `ok_link_admin_user`
func (User) TableName() string {
	return "ok_link_admin_user"
}

func (us *UserService) Registry(name, pw string) (bool, error) {
	if name == "" || pw == "" {
		return false, badReqErr("用户名/密码不能为空")
	}
	var u User
	MyDB.Where("name = ?", &name).First(&u)
	if u.Name != "" {
		return false, common.StatusError{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("用户已存在 %s", name),
		}
	}

	sum := sha256.Sum256([]byte(pw))
	passHash := fmt.Sprintf("%x", sum)

	user := User{
		Name:     name,
		Password: passHash,
	}
	MyDB.Create(&user)
	log.Printf("%+v", user)

	return !MyDB.NewRecord(user), nil
}

// 校验用户名密码，检验成功则生成token缓存到redis，同时缓存一份用户信息
func (us *UserService) Login(name, pw string) (string, error) {
	if name == "" || pw == "" {
		return "", badReqErr("用户名/密码输入错误")
	}

	sum := sha256.Sum256([]byte(pw))
	passHash := fmt.Sprintf("%x", sum)

	var u User
	MyDB.Where("name = ? and password = ?", name, passHash).First(&u)
	if u.Name == "" || u.Name != name {
		return "", badReqErr("用户名/密码输入错误")
	}
	// 生成Token
	sum1 := sha256.Sum256([]byte(pw + strconv.Itoa(int(time.Now().UnixNano()))))
	token := fmt.Sprintf("%x", sum1)
	// 缓存Token
	expiration := time.Minute * time.Duration(TokenFreshTime)
	err := us.R.Cli.Set(fmt.Sprintf(AdminTokenKey, name), token, expiration).Err()
	if err != nil {
		return "", err
	}
	// 缓存用户信息
	userJson, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	err = us.R.Cli.Set(fmt.Sprintf(AdminInfoKey, name), userJson, expiration).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

// 先查询token缓存，确认传入的token是正确的，然后根据token查询用户缓存
func (us *UserService) UserInfo(name, token string) (string, error) {
	if name == "" || token == "" {
		return "", badReqErr("用户名/Token参数错误")
	}
	tokenCache, err := us.R.Cli.Get(fmt.Sprintf(AdminTokenKey, name)).Result()
	if err == redis.Nil {
		return "", unAuthErr("用户未登录 " + name)
	} else if err != nil {
		return "", unAuthErr(fmt.Sprintf("获取Token缓存失败 %s", err.Error()))
	}
	if token != tokenCache {
		return "", unAuthErr("Token信息错误")
	}

	userCache, err := us.R.Cli.Get(fmt.Sprintf(AdminInfoKey, name)).Result()
	if err != nil {
		return "", err
	}
	if userCache == "" {
		return "", unAuthErr("用户未登录 " + name)
	}

	return userCache, nil
}

func badReqErr(msg string) error {
	return common.StatusError{
		Code: http.StatusBadRequest,
		Err:  fmt.Errorf(msg),
	}
}

func unAuthErr(msg string) error {
	return common.StatusError{
		Code: http.StatusUnauthorized,
		Err:  fmt.Errorf(msg),
	}
}
