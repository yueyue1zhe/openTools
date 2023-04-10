package ginutil

import (
	"e.coding.net/zhechat/magic/taihao/library/logutil"
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
func HtmlRenderByStringPrintErr(c *gin.Context, name string, html string, data any) {
	if err := HtmlRenderByString(c, name, html, data); err != nil {
		_, _ = c.Writer.WriteString(err.Error())
	}
}
func HtmlRenderByStringLogErr(c *gin.Context, name string, html string, data any) {
	if err := HtmlRenderByString(c, name, html, data); err != nil {
		logutil.Error("页面加载失败:", name, err.Error())
	}
}
