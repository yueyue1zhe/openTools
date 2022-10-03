package ginutil

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

func HtmlRenderByString(c *gin.Context, name string, html string, data any) error {
	t, err := template.New(name).Parse(html)
	if err != nil {
		return err
	}
	return t.Execute(c.Writer, data)
}
