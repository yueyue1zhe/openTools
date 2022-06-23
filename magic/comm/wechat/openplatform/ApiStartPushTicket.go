package openplatform

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

//请求开启开放平台票据推送

func ApiStartPushTicket() error {
	uri := "https://api.weixin.qq.com/cgi-bin/component/api_start_push_ticket"
	req := map[string]string{
		"component_appid":  openPlatform.AppID,
		"component_secret": openPlatform.AppSecret,
	}
	body, err := util.PostJSON(uri, req)
	if err != nil {
		return fmt.Errorf("请求异常：%v", err.Error())
	}
	var ret struct {
		util.CommonError
	}
	if err := json.Unmarshal(body, &ret); err != nil {
		return fmt.Errorf("解码异常：%v", err.Error())
	}
	if ret.ErrCode != 0 {
		return fmt.Errorf("errcode=%v , errmsg=%v", ret.ErrCode, ret.ErrMsg)
	}
	fmt.Println(ret.ErrMsg)
	return nil
}
