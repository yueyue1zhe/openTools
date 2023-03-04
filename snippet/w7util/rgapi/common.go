package rgapi

import (
	"crypto/sha1"
	"e.coding.net/zhechat/magic/taihao/function/commutil"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	baseUri     = "https://rgapi.w7.cc"
	payBasePath = "/open/pay/create"
)

const (
	AccountTypeNone        = iota
	AccountTypeWechat                 //微信公众号
	AccountTypeMiniProgram            //微信小程序
	AccountTypeApp         = iota + 1 //App
	AccountTypeAli                    //支付宝小程序
	AccountTypeBaiDu                  //百度小程序
	AccountTypeTouTiao                //字节跳动小程序
	AccountTypeWork                   //企业微信
)

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
	if resp.IsSuccessState() {
		var result commonResult[any]
		if err := resp.UnmarshalJson(&result); err != nil {
			resp.Err = fmt.Errorf("请求解析异常:%v", err.Error())
		} else {
			if result.Code != 0 {
				resp.Err = fmt.Errorf("请求异常：%v", result.Message)
			}
		}
		return nil
	}
	return nil
}

func (w *RgApi) buildSignQueryUri(body interface{}, path string) string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := commutil.StringRandom(8)

	rawData := make(map[string]string)
	var useBody string
	if body != "" {
		bodyJs, _ := json.Marshal(&body)
		useBody = string(bodyJs)
	}

	rawData["Body"] = useBody
	rawData["AppSecret"] = w.AppSecret
	rawData["TimeStamp"] = timestamp
	rawData["Nonce"] = nonce
	rawData["Uri"] = path
	var keys []string
	for k := range rawData {
		keys = append(keys, rawData[k])
	}
	sort.Strings(keys)
	o := sha1.New()
	o.Write([]byte(strings.Join(keys, "")))
	sign := hex.EncodeToString(o.Sum(nil))

	query := commutil.MapToQuery(map[string]interface{}{
		"sign":  sign,
		"appid": w.AppID,
		"type":  w.useAccountType,
		"time":  timestamp,
		"nonce": nonce,
	})
	return baseUri + path + "?" + query
}

func (w *RgApi) getSign(body interface{}) string {
	//map[string]string{
	//		"Body":      "",
	//		"AppSecret": w.AppSecret,
	//		"TimeStamp": timestamp,
	//		"Nonce":     nonce,
	//		"Uri":       path,
	//	}
	rawData := make(map[string]string)
	bodyJs, _ := json.Marshal(&body)
	rawData["Body"] = string(bodyJs)
	rawData["AppSecret"] = w.AppSecret
	rawData["TimeStamp"] = w.AppSecret
	rawData["Nonce"] = w.AppSecret
	rawData["Uri"] = w.AppSecret
	var keys []string
	for k := range rawData {
		keys = append(keys, rawData[k])
	}
	sort.Strings(keys)
	fmt.Println(keys)
	o := sha1.New()
	o.Write([]byte(strings.Join(keys, "")))
	return hex.EncodeToString(o.Sum(nil))
}
