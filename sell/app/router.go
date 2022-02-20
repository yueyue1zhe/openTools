package app

import (
	"github.com/gin-gonic/gin"
	"yueyue-sell/app/pc"
)

func Register(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		ReqToast(c, "what ?")
	})
	pc.Register(router)
}
