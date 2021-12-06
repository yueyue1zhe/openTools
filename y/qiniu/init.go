package qiniu

import (
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"net/url"
)

type QiNiu struct {
	Conf Conf
}
type Conf struct {
	Bucket    string
	AccessKey string
	SecretKey string
	Url       string
}

func NewQiNiu(conf Conf) *QiNiu {
	return &QiNiu{
		Conf: Conf{
			Bucket:    conf.Bucket,
			AccessKey: conf.AccessKey,
			SecretKey: conf.SecretKey,
			Url:       conf.Url,
		},
	}
}

func (q *QiNiu) upToken() string {
	putPolicy := storage.PutPolicy{
		Scope: q.Conf.Bucket,
	}
	mac := qbox.NewMac(q.Conf.AccessKey, q.Conf.SecretKey)
	return putPolicy.UploadToken(mac)
}

func (q *QiNiu) CanUse() error {
	if q.Conf.Url == "" {
		return fmt.Errorf("配置异常:url不能为空")
	}
	parse, err := url.Parse(q.Conf.Url)
	if err != nil {
		return fmt.Errorf("配置异常:%v", err.Error())
	}
	mac := qbox.NewMac(q.Conf.AccessKey, q.Conf.SecretKey)
	cfg := storage.Config{}
	domains, err := storage.NewBucketManager(mac, &cfg).ListBucketDomains(q.Conf.Bucket)
	if err != nil {
		return fmt.Errorf("配置异常:%v", err.Error())
	}
	hostOk := false
	for _, domain := range domains {
		if domain.Tbl == q.Conf.Bucket && domain.Domain == parse.Host {
			hostOk = true
		}
	}
	if !hostOk {
		return fmt.Errorf("配置异常:url不在当前储存空间")
	}
	return nil
}
