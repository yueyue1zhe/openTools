package rgapi

import (
	"crypto/sha1"
	"e.coding.net/zhechat/magic/taihao/function/commutil"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"sort"
	"strings"
)

type AppEncryptor struct {
	AppID  string
	Token  string
	AesKey string
}

type XData struct {
	Encrypt      string `xml:"Encrypt" y-required-label:"Encrypt"`
	MsgSignature string `xml:"MsgSignature" y-required-label:"Encrypt"`
	Nonce        string `xml:"Nonce" y-required-label:"Encrypt"`
	TimeStamp    string `xml:"TimeStamp" y-required-label:"Encrypt"`
}

//todo:: 向微擎官方要回调解密规则 以及测试xml

func (a *AppEncryptor) Decrypt(xmlRaw string) {
	var xData XData
	if err := xml.Unmarshal([]byte(xmlRaw), &xData); err != nil {
		fmt.Println(err.Error(), "unmarshal")
		return
	}
	if err := commutil.StructRequiredJudge(xData); err != nil {
		fmt.Println("缺失必要参数", err.Error())
		return
	}
	if xData.MsgSignature != a.createSignature(xData.TimeStamp, xData.Nonce, xData.Encrypt) {
		fmt.Println("无效的签名")
		return
	}

}

func (a *AppEncryptor) createSignature(timestamp, nonce, encrypt string) string {
	rawArr := []string{a.Token, timestamp, nonce, encrypt}
	sort.Strings(rawArr)
	o := sha1.New()
	o.Write([]byte(strings.Join(rawArr, "")))
	return hex.EncodeToString(o.Sum(nil))
}
