package openplatform

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/openplatform"
	"github.com/silenceper/wechat/v2/openplatform/config"
	"yueyue-magic/comm/wechat/cache"
)

/*
使用站点单个开放平台 应先于开放平台应用加载此项
*/

var openPlatform *openplatform.OpenPlatform

func Init() error {
	wc := wechat.NewWechat()
	wechatCache := cache.Get()
	cfg := &config.Config{
		AppID:          "wx5a68e98c25b87d43",
		AppSecret:      "ab0bf0b492a7665def3a38ca9af2e05f",
		Token:          "l5spjlm0dgo0govtsc1gjs2y5om1gchm",
		EncodingAESKey: "cTzBXv34izGt42gGlPKjwCy2UAvH65Gb0N1VX8ghnWl",
		Cache:          wechatCache,
	}
	openPlatform = wc.GetOpenPlatform(cfg)
	return nil
}

func Get() *openplatform.OpenPlatform {
	return openPlatform
}
