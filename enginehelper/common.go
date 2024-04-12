package tests

import (
	"d-photo/app"
	"d-photo/env"
	"e.coding.net/zhechat/magic/taihao/core"
)

func withDb() {
	logutil
	conf := env.GetConf()
	appConf := core.Conf{
		Name:                  env.AppName,
		Version:               "0.1.10",
		AttachmentPath:        "/attachment",
		JWTTokenKey:           env.JWTTokenKey,
		StructBasic:           env.StructBasic,
		MultiTenantIDTokenKey: env.MultiTenantIDTokenKey,
		MultiTenantIDQueryKey: env.MultiTenantIDQueryKey,
		Debug:                 conf.Debug,
		DevProxySiteDomain:    conf.DevProxySiteDomain,
		MultiTenant:           true,
	}
	core.Init(appConf)
	if err := app.RegisterByConf(); err != nil {
		panic("启动异常：" + err.Error())
	}
}
