package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ErrorToast    = 1
	ErrorRedirect = 2
	ErrorAuthFail = 40019
)

type Response struct {
	Errno   int64       `json:"errno"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ReqRedirect(c *gin.Context, tips, path string) {
	Req(c, ErrorRedirect, path, tips)
}

func ReqToast(c *gin.Context, message string) {
	Req(c, 1, nil, message)
}

func ReqFail(c *gin.Context, errno int64, message string) {
	Req(c, errno, nil, message)
}

func ReqOk(c *gin.Context, data interface{}) {
	Req(c, 0, data, "")
	c.Abort()
}

func Req(c *gin.Context, errno int64, data interface{}, message string) {
	c.JSON(http.StatusOK, Response{
		Errno:   errno,
		Message: message,
		Data:    data,
	})
}
