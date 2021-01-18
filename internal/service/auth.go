package service

import (
	"errors"
	"github.com/joeyscat/ok-short/internal/model"
	"strings"
)

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (svc Service) CreateAuth(param *AuthRequest) error {
	if strings.Contains(param.AppKey, " ") || strings.Contains(param.AppSecret, " ") ||
		param.AppKey == "" || param.AppSecret == "" {
		return errors.New("app_key or app_secret cannot be empty or contains whitespace")
	}

	_, err := model.CreateAuth(param.AppKey, param.AppSecret)
	return err
}

func (svc *Service) CheckAuth(param *AuthRequest) (err error) {
	_, err = model.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	return
}
