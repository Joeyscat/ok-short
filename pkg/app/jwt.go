package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"net/http"
	"time"
)

type Claims struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}

type Jwt struct {
	Secret string        // 加密的key
	Expire time.Duration // 过期时间
	Issuer string        // 签发人
}

var jwt_ Jwt

func SetJWT(secret, issuer string, expire time.Duration) {
	jwt_ = Jwt{
		Secret: secret,
		Expire: expire,
		Issuer: issuer,
	}
}

func getJWTSecret() []byte {
	return []byte(jwt_.Secret)
}

func GenerateToken(uid string) (string, error) {

	nowTime := time.Now()
	expireTime := nowTime.Add(jwt_.Expire)
	claims := Claims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    jwt_.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(getJWTSecret())
	return token, err
}

func ParseFromRequest(r *http.Request, extractor request.Extractor) (*jwt.Token, error) {
	return request.ParseFromRequest(
		r,
		extractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwt_.Secret), nil
		})
}
