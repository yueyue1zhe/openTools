package jwt

import (
	"errors"
)

type JWT struct {
	client string
}

func NewJWT(client string) *JWT {
	return &JWT{
		client: client,
	}
}

const (
	SignKeyRaw     = "newTheodore"
	ClientAdminApi = "admin-api"
	ClientAppApi   = "app-api"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

func (j *JWT) getSignKey() []byte {
	var raw string
	switch j.client {
	case ClientAppApi:
		raw = SignKeyRaw + ClientAppApi
	case ClientAdminApi:
		raw = SignKeyRaw + ClientAdminApi
	default:
		raw = SignKeyRaw
	}
	return []byte(raw)
}
