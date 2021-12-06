package qiniu

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/sirupsen/logrus"
)

// Remove 传入文件保存路径 删除七牛云空间中的文件
func (q *QiNiu) Remove(path string) (err error) {
	cfg := storage.Config{}
	mac := qbox.NewMac(q.Conf.AccessKey, q.Conf.SecretKey)
	bucketManager := storage.NewBucketManager(mac, &cfg)
	err = bucketManager.Delete(q.Conf.Bucket, path)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err":  err.Error(),
			"file": path,
		}).Info("七牛云文件删除失败")
		return
	}
	return nil
}
