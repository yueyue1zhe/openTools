package wechat

import (
	"fmt"
	"regexp"
)

func parseWx110() {
	content := `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta id="viewport" name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0, viewport-fit=cover" />
    <title></title>
    <link rel="stylesheet" type="text/css" href="https://res.wx.qq.com/t/wx_fed/cdn_libs/res/weuicss/2.5.0/weui.min.css" />
    <style>body{background-color: var(--weui-BG-2); color: var(--weui-FG-0);}</style>
  <link href="//res.wx.qq.com/t/wx_fed/wx110/wx110/res/css/banurl.c9674390.css" rel="preload" as="style"><link href="//res.wx.qq.com/t/wx_fed/wx110/wx110/res/js/banurl.29c14bb9.js" rel="preload" as="script"><link href="//res.wx.qq.com/t/wx_fed/wx110/wx110/res/js/chunk-common.b4d9f5a7.js" rel="preload" as="script"><link href="//res.wx.qq.com/t/wx_fed/wx110/wx110/res/js/chunk-vendors.1684286d.js" rel="preload" as="script"><link href="//res.wx.qq.com/t/wx_fed/wx110/wx110/res/css/banurl.c9674390.css" rel="stylesheet"></head>
  <body ontouchstart="">
    <div id="app"></div>
    <script>
      var cgiData = {"retcode":0,"type":"gray","title":"将要访问","desc":"该地址为IP地址，请使用域名访问网站。","url":"http:&#x2f;&#x2fns":[{"name":"继续访问","url":"http:&#x2f;&#x2f;192.144.227.223&#x2f;","type":"plain-primary"}]};
    </script>
    <script src="https://res.wx.qq.com/a/wx_fed/cdn_libs/res/vue/2.6.11/vue.min.js"></script>
  <script type="text/javascript" src="//res.wx.qq.com/t/wx_fed/wx110/wx110/res/js/chunk-common.b4d9f5a7.js"></script><script type="text/javascript" src="//res.wx.qq.com/t/wx_fed/wx110/wx110/res/js/chunk-vendors.1684286d.js"></script><script type="text/javascript" src="//res.wx.qq.com/t/wx_fed/wx110/wx110/res/js/banurl.29c14bb9.js"></script></body>
</html>
`
	//ql, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
	str, err := RegexpStr(content, `\s+cgiData\s+=\s+{(.*?)}`)
	if err != nil {
		fmt.Println("err", err.Error())
		return
	}
	fmt.Println(str)
}

func RegexpStr(str, rule string) (res string, err error) {
	reg := regexp.MustCompile(rule)
	strReg := reg.FindStringSubmatch(str)
	if strReg == nil {
		return "", fmt.Errorf("未匹配")
	}
	return strReg[1], nil
}
