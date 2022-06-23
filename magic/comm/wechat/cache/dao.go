package cache

import (
	"openTools/magic/comm/log"
	"time"
)

func (mem *WechatCache) Get(key string) interface{} {
	if value, found := mem.goCache.Get(key); found {
		return value
	}
	if kv, err := mem.dbGet(key); err != nil {
		return nil
	} else {
		if kv.ExpireTime.Before(time.Now()) {
			//数据库缓存过期删除
			_ = mem.dbDel(key)
			return nil
		} else {
			//数据库可用更新内存缓存
			mem.goCache.Set(key, kv.Value, time.Until(kv.ExpireTime))
			return kv.Value
		}
	}
}
func (mem *WechatCache) Set(key string, val interface{}, timeout time.Duration) error {
	mem.Lock()
	defer mem.Unlock()
	if err := mem.dbSet(key, val, timeout); err != nil {
		log.Error(err.Error())
		return err
	}
	mem.goCache.Set(key, val, timeout)
	return nil
}
func (mem *WechatCache) IsExist(key string) bool {
	if _, found := mem.goCache.Get(key); found {
		return true
	}
	if kv, err := mem.dbGet(key); err != nil {
		return false
	} else {
		if kv.ExpireTime.Before(time.Now()) {
			//数据库缓存过期删除
			_ = mem.dbDel(key)
			return false
		} else {
			//数据库可用更新内存缓存
			mem.goCache.Set(key, kv.Value, time.Until(kv.ExpireTime))
			return true
		}
	}
}

func (mem *WechatCache) Delete(key string) error {
	mem.Lock()
	defer mem.Unlock()
	if err := mem.dbDel(key); err != nil {
		log.Error(err.Error())
		return err
	}
	mem.goCache.Delete(key)
	return nil
}

func (mem *WechatCache) dbGet(key string) (kv wechatCommKv, err error) {
	err = mem.DB.Where(&wechatCommKv{Key: key}).First(&kv).Error
	return
}

func (mem *WechatCache) dbDel(key string) error {
	return mem.DB.Where(wechatCommKv{Key: key}).Unscoped().Delete(&wechatCommKv{}).Error
}

func (mem *WechatCache) dbSet(key string, value interface{}, d time.Duration) error {
	var kv wechatCommKv
	return mem.DB.Where(&wechatCommKv{Key: key}).Assign(wechatCommKv{
		Value:      value.(string),
		ExpireTime: time.Now().Add(d),
	}).FirstOrCreate(&kv).Error
}
