package page

import (
	"github.com/gin-gonic/gin"
	"yueyue-sell/app"
)

func Home(c *gin.Context) {
	app.ReqOk(c, "hello")
}
