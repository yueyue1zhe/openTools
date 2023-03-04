package rgapi

import (
	"fmt"
	"github.com/imroc/req/v3"
	"time"
)

type RgApi struct {
	AppID     string
	AppSecret string
	Client    *req.Client

	useAccountType int
}

func GetClient(appID string, appSecret string) *RgApi {
	client := req.C()
	client.SetCommonContentType("application/x-www-form-uriencoded").
		SetTimeout(time.Minute).
		OnAfterResponse(onAfterResponse)
	return &RgApi{
		AppID:     appID,
		AppSecret: appSecret,
		Client:    client,
	}
}

func (w *RgApi) SetUseAccountType(acType int) *RgApi {
	w.useAccountType = acType
	return w
}

func (w *RgApi) GetAccountList() (list []AccountListResultDataItem, err error) {
	path := "/open/api/account/list"
	r, err := w.Client.R().Post(w.buildSignQueryUri("", path))
	if err != nil {
		return list, fmt.Errorf("err:%v", err.Error())
	}
	var res commonResult[[]AccountListResultDataItem]
	_ = r.UnmarshalJson(&res)
	return res.Data, nil
}

func (w *RgApi) GetAccessToken(repeat bool) (ac AccessTokenResult, err error) {
	path := "/open/api/account/getAccessToken"
	r, err := w.Client.R().Post(w.buildSignQueryUri("", path))
	if err != nil {
		return ac, fmt.Errorf("err:%v", err.Error())
	}
	var res commonResult[AccessTokenResult]
	_ = r.UnmarshalJson(&res)
	if res.Data.ExpiresIn <= 0 && repeat {
		return w.GetAccessToken(false)
	}
	return res.Data, nil
}

func (w *RgApi) WechatPay(notifyUri string) {

}
