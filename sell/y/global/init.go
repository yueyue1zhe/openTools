package global

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

type Global struct {
	SiteIp     string
	SiteRoot   string
	SiteDomain string

	AttachmentUrl string

	IaRoot         string
	AttachmentRoot string

	ClientIp string
}

const AttachmentDir = "attachment/"

func NewGlobal() *Global {
	return &Global{
		SiteIp:         curIp(),
		IaRoot:         iaRoot(),
		AttachmentRoot: iaRoot() + "/" + AttachmentDir,
	}
}
func (y *Global) Cur(c *gin.Context) *Global {
	y.SiteDomain = getSiteDomain(c)
	y.SiteRoot = y.SiteDomain + "/"
	y.AttachmentUrl = y.SiteRoot + AttachmentDir
	y.ClientIp = y.clientIp(c)
	return y
}
func (y *Global) clientIp(c *gin.Context) string {
	ip := c.GetHeader("X-Real-IP")
	return ip
}

func iaRoot() string {
	if len(os.Args) > 1 && os.Args[1] == "-test.v" {
		p, _ := os.Getwd()
		return p
	}
	ePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return path.Dir(ePath)
}

func curIp() string {
	responseClient, errClient := http.Get("http://ip.dhcp.cn/?ip")
	if errClient != nil {
		return "127.0.0.1"
	}
	defer responseClient.Body.Close()
	body, _ := ioutil.ReadAll(responseClient.Body)
	return fmt.Sprintf("%s", string(body))
}

func getSiteDomain(c *gin.Context) string {
	scheme := c.GetHeader("X-Forwarded-Proto")
	if scheme == "" {
		scheme = "http"
		if c.Request.TLS != nil {
			scheme = "https"
		}
	}
	return fmt.Sprintf("%v://%v", scheme, c.Request.Host)
}
