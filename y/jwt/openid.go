package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type ShopOpenidCustomClaims struct {
	Openid   string `json:"openid"`
	ClientIp string `json:"client_ip"`
	Client   string `json:"client"`
	jwt.StandardClaims
}

func (j *JWT) ShopCreateOpenidToken(openid string, clientIp string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, ShopOpenidCustomClaims{
		Openid:   openid,
		ClientIp: clientIp,
		Client:   j.client,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	})
	return token.SignedString(j.getSignKey())
}

func (j *JWT) ShopParseOpenidToken(tokenString string) (*ShopOpenidCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &ShopOpenidCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.getSignKey(), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
		return nil, err
	}
	if claims, ok := token.Claims.(*ShopOpenidCustomClaims); ok && token.Valid {
		if j.client != "" && claims.Client != j.client {
			return nil, fmt.Errorf("insufficient permissions %v ip:%v", claims.Client, claims.ClientIp)
		}
		return claims, nil
	}
	return nil, TokenInvalid
}
