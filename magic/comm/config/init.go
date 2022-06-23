package config

import (
	"github.com/go-ini/ini"
	"openTools/magic/comm/encrypt"
	"openTools/magic/comm/log"
)

var WxCloudConf = &WxCloud{}
var MysqlConf = &Mysql{}
var ServerConf = &Server{}
var CommConf = &Comm{}
var WxApiConf = &WxApi{}

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("ini.toml")
	if err != nil {
		log.Errorf("load server.conf': %v", err)
		return
	}
	mapTo("mysql", MysqlConf)
	mapTo("wx-cloud", WxCloudConf)
	mapTo("server", ServerConf)
	mapTo("comm", CommConf)
	mapTo("wx-api", WxApiConf)
	if ServerConf.AesKey == "" {
		ServerConf.AesKey = encrypt.GenerateMd5(MysqlConf.Password)
	}
	if ServerConf.JwtSecret == "" {
		ServerConf.JwtSecret = encrypt.GenerateMd5(MysqlConf.Password)
	}
	log.Info(ServerConf)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Errorf("%s err: %v", section, err)
	}
}
