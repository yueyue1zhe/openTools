package noticebotutil

import (
	"fmt"
	"github.com/imroc/req/v3"
)

type WeCom struct {
	uri string
}

func WeComGet(uri string) *WeCom {
	return &WeCom{uri: uri}
}

type weComType struct {
	MsgType string        `json:"msgtype"`
	Text    weComTextType `json:"text"`
}
type weComTextType struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

func (w *WeCom) Text(content string) error {
	return w.send(weComType{
		MsgType: "text",
		Text: weComTextType{
			Content: content,
		},
	})
}

type weComResult struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (w *WeCom) send(body weComType) error {
	r, err := req.R().
		SetHeader("Content-Type", "application/json").
		SetBodyJsonMarshal(&body).Post(w.uri)
	if err != nil {
		return err
	}
	var result weComResult
	if err = r.UnmarshalJson(&result); err != nil {
		return fmt.Errorf("解析请求结果异常：%v", err.Error())
	}
	if result.Errcode != 0 {
		return fmt.Errorf("code:%v:%v", result.Errcode, result.Errmsg)
	}
	return nil
}
