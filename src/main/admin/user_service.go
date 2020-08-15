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

func (us *UserService) Register(name, pw string) (bool, error) {
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

	// 取出旧Token
	tokenCache, err := ReCli.Cli.Get(fmt.Sprintf(AdminTokenKey, name)).Result()
	if err == redis.Nil {
	} else if err != nil {
		return "", err
	}
	if tokenCache != "" {
		err := ReCli.Cli.Del(fmt.Sprintf(AdminInfoKey, tokenCache)).Err()
		if err != nil {
			log.Printf("用户缓存清除失败 %s", err.Error())
		}
	}

	// 生成新Token并缓存
	sum1 := sha256.Sum256([]byte(pw + strconv.Itoa(int(time.Now().UnixNano()))))
	token := fmt.Sprintf("%x", sum1)

	expiration := time.Minute * time.Duration(TokenFreshTime)
	err = ReCli.Cli.Set(fmt.Sprintf(AdminTokenKey, name), token, expiration).Err()
	if err != nil {
		return "", err
	}
	// 缓存用户信息
	userJson, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	err = ReCli.Cli.Set(fmt.Sprintf(AdminInfoKey, token), userJson, expiration).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

// 根据token查询用户缓存
func (us *UserService) UserInfo(token string) (string, error) {
	if token == "" {
		return "", badReqErr("用户名/Token参数错误")
	}

	userCache, err := ReCli.Cli.Get(fmt.Sprintf(AdminInfoKey, token)).Result()
	if err == redis.Nil {
		return "", unAuthErr("Token无效 ")
	} else if err != nil {
		return "", err
	}
	if userCache == "" {
		return "", unAuthErr("Token无效 " + token)
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
