package qiniu

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func (q *QiNiu) move(old, now string) error {
	cfg := storage.Config{}
	mac := qbox.NewMac(q.Conf.AccessKey, q.Conf.SecretKey)
	bucketManager := storage.NewBucketManager(mac, &cfg)
	return bucketManager.Move(q.Conf.Bucket, old, q.Conf.Bucket, now, true)
}
