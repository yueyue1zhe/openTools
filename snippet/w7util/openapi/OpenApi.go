package openapi

import (
	"crypto/md5"
	"fmt"
	"github.com/imroc/req/v3"
	"sort"
	"strings"
	"time"
)

type OpenApi struct {
	AppID     string
	AppSecret string

	Client *req.Client
}

const baseUrl = "https://openapi.w7.cc"

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

func GetClient(appid, appSecret string) *OpenApi {
	client := req.C()
	client.SetCommonContentType("application/x-www-form-uriencoded").
		SetTimeout(time.Minute).
		OnAfterResponse(onAfterResponse)
	return &OpenApi{
		AppID:     appid,
		AppSecret: appSecret,
		Client:    client,
	}
}

type getLoginUrlResult struct {
	Url string `json:"uri"`
}

func (w *OpenApi) GetLoginUrl(redirectUrl string) (string, error) {
	uri := baseUrl + "/we7/open/oauth/login-uri/index"
	r, err := w.Client.R().
		SetFormData(map[string]string{
			"appid":    w.AppID,
			"redirect": redirectUrl,
		}).Post(uri)
	if err != nil {
		return "", fmt.Errorf("网络请求异常：%v", err.Error())
	}
	var res getLoginUrlResult
	if err := r.UnmarshalJson(&res); err != nil {
		return "", fmt.Errorf("数据解析异常：%v", err.Error())
	}
	return res.Url, nil
}

type CodeToAccessTokenResult struct {
	//access_token, 有效期两小时
	AccessToken string `json:"access_token"`
	ExpireTime  int64  `json:"expire_time"` //截止的时间戳
}

func (w *OpenApi) CodeToAccessToken(code string) (*CodeToAccessTokenResult, error) {
	uri := baseUrl + "/we7/open/oauth/access-token/code"
	data := map[string]string{
		"appid": w.AppID,
		"code":  code,
	}
	data["sign"] = w.getSign(data)
	r, err := w.Client.R().SetFormData(data).Post(uri)
	if err != nil {
		return nil, err
	}
	var res CodeToAccessTokenResult
	if err := r.UnmarshalJson(&res); err != nil {
		return nil, fmt.Errorf("数据解析异常")
	}
	return &res, nil
}

type UserInfoResult struct {
	OpenID         string `json:"open_id"`  //用户openid
	Nickname       string `json:"nickname"` //昵称
	Avatar         string `json:"avatar"`
	RoleIdentity   string `json:"role_identity"`   //角色
	ComponentAppid string `json:"component_appid"` //站点id
	FounderOpenid  string `json:"founder_openid"`  //创始人openid
}

func (w *OpenApi) UserInfo(accessToken string) (*UserInfoResult, error) {
	uri := baseUrl + "/we7/open/oauth/user/info"
	data := map[string]string{"access_token": accessToken}
	r, err := w.Client.R().SetFormData(data).Post(uri)
	if err != nil {
		return nil, err
	}
	var res UserInfoResult
	if err := r.UnmarshalJson(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (w *OpenApi) getSign(data map[string]string) string {
	var dataParams []string
	var keys []string
	for k := range data {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for _, key := range keys {
		dataParams = append(dataParams, fmt.Sprintf("%v=%v", key, data[key]))
	}
	srcCode := md5.Sum([]byte(strings.Join(dataParams, "&") + w.AppSecret))
	return fmt.Sprintf("%x", srcCode)
}
