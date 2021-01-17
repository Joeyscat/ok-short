package service

import (
	"github.com/joeyscat/ok-short/internal/model"
)

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

func (svc *Service) CheckAuth(param *AuthRequest) (err error) {
	_, err = model.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	return
}
