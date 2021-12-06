package wechat

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
)

var (
	Wc     *wechat.Wechat
	Memory *cache.Memory
)

func Register() {
	Wc = wechat.NewWechat()
	Memory = cache.NewMemory()
}

type wxApiRes struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type TransferParams struct {
	OrderNo  string
	Openid   string
	Fee      float64
	Title    string
	ClientIp string
	UniAcid  uint
}

type TplNoticeParam struct {
	Openid     string
	TemplateId string
	Url        string
	Data       map[string]string
}

type PayConfParams struct {
	Openid    string
	OrderNo   string
	Fee       float64
	Title     string
	NotifyURL string
	ClientIp  string
}

type PayConfRes struct {
	Timestamp string `json:"timestamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}
