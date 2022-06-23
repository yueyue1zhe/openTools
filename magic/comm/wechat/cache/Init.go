package cache

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"sync"
	"time"
)

/*
使用微信相关模块应先加载此项

微信token go-cache + mysql 缓存持久化
*/

type WechatCache struct {
	sync.Mutex
	goCache *cache.Cache
	DB      *gorm.DB
}

type wechatCommKv struct {
	db.Model
	Key        string    `json:"key" gorm:"unique;type:string;size:255;not null"`
	Value      string    `json:"value" gorm:"type:text;not null"`
	ExpireTime time.Time `json:"expire_time" gorm:"type:time;not null"`
}

func Init() error {
	mysql := db.Get()
	if !mysql.Migrator().HasTable(&wechatCommKv{}) {
		if err := mysql.Migrator().CreateTable(&wechatCommKv{}); err != nil {
			fmt.Println("table create fail wechatCommKv", err.Error())
			return err
		}
	}
	wechatCacheInstance = &WechatCache{
		goCache: db.GetCache(),
		DB:      db.Get(),
	}
	return nil
}

var wechatCacheInstance *WechatCache

func Get() *WechatCache {
	return wechatCacheInstance
}
