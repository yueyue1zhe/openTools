package cortana

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/imroc/req"
	"time"
)

type Loader struct {
	Conf Conf

	client *req.Client
}
type Conf struct {
	AppID     string
	AppSecret string
	ApiUrl    string //请求域名 附件管理域名
	Url       string //访问域名 配置时通信获取 由cortana-attach服务端返回
}

func New(conf Conf) (*Loader, error) {
	if conf.AppID == "" || conf.AppSecret == "" || conf.ApiUrl == "" {
		return nil, fmt.Errorf("请完整配置cortana参数")
	}
	loader := &Loader{Conf: conf}

	client := req.C()
	client.SetCommonContentType("application/x-www-form-uriencoded").
		SetTimeout(time.Minute * 6).
		OnAfterResponse(onAfterResponse)
	loader.client = client

	return loader, nil
}

func (l *Loader) GetVisitUrl() (url string, err error) {
	uri := l.Conf.ApiUrl + fmt.Sprintf("/openapi/prefix-url?appid=%v", l.Conf.AppID)
	r, err := l.client.R().Get(uri)
	if err != nil {
		return "", err
	}
	var result commonResult[string]
	if err := r.UnmarshalJson(&result); err != nil {
		return "", fmt.Errorf("解析异常：%v", err.Error())
	}
	if result.Errno != 0 {
		return "", fmt.Errorf(result.Message)
	}
	return result.Data, nil
}

type onErrorResp struct {
	Error string `json:"error"`
}

func onAfterResponse(client *req.Client, resp *req.Response) error {
	if resp.Err != nil {
		if dump := resp.Dump(); dump != "" {
			resp.Err = fmt.Errorf("%s\nraw content:\n%s", resp.Err.Error(), resp.Dump())
		}
		return nil
	}
	if !resp.IsSuccess() {
		var respErr onErrorResp
		if err := resp.UnmarshalJson(&respErr); err != nil {
			resp.Err = fmt.Errorf("bad response, raw content:\n%s", resp.String())
		} else {
			resp.Err = fmt.Errorf(respErr.Error)
		}
		return nil
	}
	return nil
}

type commonResult[T any] struct {
	RequestId string `json:"request_id"`
	Errno     int64  `json:"errno"`
	Message   string `json:"message"`
	Data      T      `json:"data"`
}

func (l *Loader) hmac(data []byte) string {
	h := hmac.New(sha1.New, []byte(l.Conf.AppSecret))
	h.Write(data)
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
