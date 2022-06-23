package inits

import (
	"openTools/magic/comm/log"
	"openTools/magic/comm/wechat/cache"
	"openTools/magic/comm/wechat/openplatform"
)

type AppOption func() error

var appOpts []AppOption

func include(opts ...AppOption) {
	appOpts = append(appOpts, opts...)
}

// 执行所有功能板块初始化操作

func Init() error {

	// db.Init must be the first
	include(db.Init, cache.Init, openplatform.Init, wxcallback.Init)
	//include(db.Init, dao.Init, admin.Init, proxy.Init)

	for i, opt := range appOpts {
		log.Infof("[%d]--begin init--", i)
		if err := opt(); err != nil {
			log.Errorf("inits failed, err:%v\n", err)
			return err
		} else {
			log.Infof("[%d]--init succ--", i)
		}
	}
	return nil
}
