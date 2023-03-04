package rgapi

import "fmt"

//func (w *RgApi) GetWechatPay() *WechatPay {
//	return &WechatPay{}
//}

func (w *RgApi) PayTransactionsJsapi() {
	body := map[string]string{
		"description":  "测试支付",               //商品描述
		"out_trade_no": "20250205100100000",  //商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
		"total":        "1",                  //订单总金额，单位为分。
		"openid":       "wx52s1fq9kgf3h5t89", //用户在直连商户appid下的唯一标识。
		"pay_type":     "wechat",
		"method":       "payTransactionsJsapi",
		"notify_url":   "http://www.baidu.com",
	}
	r, err := w.Client.R().SetBody(body).Post(w.buildSignQueryUri(body, payBasePath))
	if err != nil {
		//return ac, fmt.Errorf("err:%v", err.Error())
	}
	fmt.Println(r.String())
	//var res commonResult[AccessTokenResult]
	//_ = r.UnmarshalJson(&res)
	//if res.Data.ExpiresIn <= 0 && repeat {
	//	return w.GetAccessToken(false)
	//}
	//return res.Data, nil
}
