package admin

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/joeyscat/ok-short/common"
	"github.com/joeyscat/ok-short/model"
	. "github.com/joeyscat/ok-short/store"
	"log"
	"strconv"
	"time"
)

type UserService struct {
}

const (
	TokenFreshTime = 60 // minutes
)

// 管理员
type UserVO struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

func (us *UserService) Register(name, pw string) (bool, error) {
	var u model.User
	MyDB.Where("name = ?", &name).First(&u)
	if u.Name != "" {
		return false, common.StatusError{
			Code: common.UserAlreadyExists,
			Err:  errors.New(common.BSText(common.UserAlreadyExists)),
		}
	}

	sum := sha256.Sum256([]byte(pw))
	passHash := fmt.Sprintf("%x", sum)

	user := model.User{
		Name:     name,
		Password: passHash,
	}
	MyDB.Create(&user)
	log.Printf("%+v", user)

	return !MyDB.NewRecord(user), nil
}

// 校验用户名密码，检验成功则生成token缓存到redis，同时缓存一份用户信息
func (us *UserService) Login(name, pw string) (string, error) {
	sum := sha256.Sum256([]byte(pw))
	passHash := fmt.Sprintf("%x", sum)

	var u model.User
	MyDB.Where("name = ? and password = ?", name, passHash).First(&u)
	if u.Name == "" || u.Name != name {
		return "", bsError(common.UserAccOrPassIncorrect)
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
	userJson, err := json.Marshal(fromUserModel(&u))
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
	userCache, err := ReCli.Cli.Get(fmt.Sprintf(AdminInfoKey, token)).Result()
	if err == redis.Nil {
		return "", bsError(common.TokenInvalid)
	} else if err != nil {
		return "", err
	}
	if userCache == "" {
		return "", bsError(common.TokenInvalid)
	}

	return userCache, nil
}

func (l *UserService) QueryAdminUserList(page, limit uint32) (*[]UserVO, uint32, error) {
	var users []model.User
	var totalCount uint32
	offset := (page - 1) * limit
	MyDB.Offset(offset).Limit(limit).Find(&users).Offset(0).Count(&totalCount)

	var userList []UserVO
	for _, user := range users {
		userList = append(userList, *fromUserModel(&user))
	}

	return &userList, totalCount, nil
}

func bsError(code int) error {
	return common.StatusError{
		Code: code,
		Err:  errors.New(common.BSText(code)),
	}
}

func fromUserModel(u *model.User) *UserVO {
	return &UserVO{
		Name:      u.Name,
		Email:     u.Email,
		AvatarURL: u.AvatarURL,
	}
}
