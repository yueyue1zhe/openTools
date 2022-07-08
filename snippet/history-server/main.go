package history_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("web/index.html")
	r.NoRoute(func(c *gin.Context) {
		if strings.Index(c.Request.RequestURI, "/web") == 0 {
			c.HTML(http.StatusOK, "index.html", gin.H{})
			return
		}
		c.Writer.WriteString("hello")
	})
	r.Static("/web", "./web")
	//r.GET("/", func(context *gin.Context) {
	//	context.HTML(http.StatusOK, "index.html", gin.H{})
	//})
	r.Run(":8091")
}
