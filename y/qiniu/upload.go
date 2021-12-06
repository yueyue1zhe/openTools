package qiniu

import (
	"context"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/sirupsen/logrus"
	"openTools/y/file"
	"openTools/y/global"
	"strings"
)

// Upload 传入本地文件绝对路径上传至七牛云 并删除本地文件
//attachment 目录下文件可自动替换
//key 为删除凭证 生成规则为 文件保存路径
func (q *QiNiu) Upload(localFile string) (err error) {
	return q.upload(localFile, true)
}

func (q *QiNiu) UploadWithoutRemoveLocalFile(localFile string) (err error) {
	return q.upload(localFile, false)
}

func (q *QiNiu) upload(localFile string, removeLocalFile bool) (err error) {
	key := localFile
	attachmentRoot := global.NewGlobal().AttachmentRoot
	if strings.Index(key, attachmentRoot) == 0 {
		key = strings.Replace(localFile, attachmentRoot, "", 1)
	} else {
		localFile = attachmentRoot + localFile
	}
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err = formUploader.PutFile(context.Background(), &ret, q.upToken(), key, localFile, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err":  err.Error(),
			"file": localFile,
		}).Info("七牛云文件上传失败")
		return err
	}
	if removeLocalFile {
		file.NewFile().Remove(key)
	}
	return nil
}
