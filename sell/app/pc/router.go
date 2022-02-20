package pc

import (
	"github.com/gin-gonic/gin"
	"yueyue-sell/app/pc/page"
)

func Register(router *gin.Engine) {
	router.GET("/", page.Home)
}
