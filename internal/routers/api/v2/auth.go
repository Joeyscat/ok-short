package v2

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joeyscat/ok-short/global"
	"github.com/joeyscat/ok-short/internal/service"
	"github.com/joeyscat/ok-short/pkg/app"
	"github.com/joeyscat/ok-short/pkg/errcode"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
	"strings"
)

func Login(e echo.Context) error {
	var u service.User
	if err := e.Bind(&u); err != nil {
		return err
	}
	response := app.NewResponse(e)
	if strings.Contains(u.Username, " ") || strings.Contains(u.Password, " ") ||
		u.Username == "" || u.Password == "" {
		return response.ToErrorResponse(errcode.InvalidParams)
	}
	userId, ok := findUserId(&u)
	if !ok {
		return response.ToErrorResponse(errcode.UnauthorizedLoginFail)
	}
	token, err := app.GenerateToken(userId)
	if err != nil {
		global.Logger.Warn(e.Request().Context(), err.Error())
		return response.ToErrorResponse(errcode.TokenGenerateError)
	}
	return response.ToResponse(map[string]interface{}{
		"token": token,
	})
}

func findUserId(u *service.User) (string, bool) {
	if u.Username != "jojo" || u.Password != "1234" {
		return "", false
	}
	return "xx", true
}

func Auth(username string, password string, c echo.Context) (bool, error) {
	if username == "" {
		return false, nil
	}

	extractor := basicAuthExtractor{content: username}
	token, err := app.ParseFromRequest(c.Request(), extractor)

	if err != nil {
		if err.Error() == "Token is expired" {
			return false, errcode.UnauthorizedTokenTimeout
		}
		return false, err
	}
	uid := getUidFromClaims("uid", token.Claims)
	if uid == "xx" {
		return true, nil
	}

	return false, nil
}

func getUidFromClaims(key string, claims jwt.Claims) string {
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)
			if fmt.Sprintf("%s", k.Interface()) == key {
				return fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return ""
}

type basicAuthExtractor struct {
	content string
}

func (e basicAuthExtractor) ExtractToken(r *http.Request) (string, error) {
	return e.content, nil
}

//func GetAuth(e echo.Context) error {
//	param := service.AuthRequest{}
//	response := app.NewResponse(e)
//	valid, errs := app.BindAndValid(e, &param)
//	if !valid {
//		global.Logger.Errorf(e.Request().Context(), "app.BindAndValid errs: %v", errs)
//		return response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
//	}
//
//	svc := service.New(e.Request().Context())
//	err := svc.CheckAuth(&param)
//	if err != nil {
//		global.Logger.Errorf(e.Request().Context(), "svc.CheckAuth err: %v", err)
//		return response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
//	}
//
//	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
//	if err != nil {
//		global.Logger.Errorf(e.Request().Context(), "app.GenerateToken err: %v", err)
//		return response.ToErrorResponse(errcode.TokenGenerateError)
//	}
//
//	return response.ToResponse(map[string]interface{}{
//		"token": token,
//	})
//}
