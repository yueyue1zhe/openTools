package qiniu

import (
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type TranscodeParam struct {
	VideoPath string
	Vb        int64
	W         int64
}

func (q *QiNiu) Transcoding(param TranscodeParam) error {
	mac := qbox.NewMac(q.Conf.AccessKey, q.Conf.SecretKey)
	cfg := storage.Config{}
	operationManager := storage.NewOperationManager(mac, &cfg)
	// 处理指令集合
	//vb 码率 取值范围 10-60000  常用码率 128k 128 1.25m 1280  5m 5120
	//单位（kbit/s）
	//w 分辨率640x360 取值范围 20-3840  仅指定转码后视频分辨率宽 高等比例自动缩放
	if param.Vb < 10 {
		param.Vb = 10
	}
	if param.Vb > 60000 {
		param.Vb = 60000
	}
	cmdW := ""
	if param.W != 0 && param.W > 20 && param.W < 3840 {
		cmdW = fmt.Sprintf("/s/%vx", param.W)
	}
	cmdVb := fmt.Sprintf("/vb/%vk", param.Vb)
	cmdSaveOut := storage.EncodedEntry(q.Conf.Bucket, param.VideoPath+fmt.Sprintf("-avthumb_vb%v_s%v", param.Vb, param.W))
	fopAvthumb := fmt.Sprintf("avthumb/mp4%v%v|saveas/%s", cmdVb, cmdW, cmdSaveOut)
	fopBatch := []string{fopAvthumb}
	fops := strings.Join(fopBatch, ";")
	persistentId, err := operationManager.Pfop(q.Conf.Bucket, param.VideoPath, fops, "", "", false)
	if err != nil {
		return err
	}
	return q.transcodingStatus(persistentId)
}

func (q *QiNiu) transcodingStatus(persistentId string) error {
	mac := qbox.NewMac(q.Conf.AccessKey, q.Conf.SecretKey)
	cfg := storage.Config{}
	operationManager := storage.NewOperationManager(mac, &cfg)
	prefop, err := operationManager.Prefop(persistentId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err":          err.Error(),
			"persistentId": persistentId,
		}).Info("七牛云转码查询异常")
		return err
	}
	switch prefop.Code {
	case 0:
		if prefop.Items != nil {
			err = q.move(prefop.Items[0].Key, prefop.InputKey)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"err":          err.Error(),
					"persistentId": persistentId,
					"input":        prefop.InputKey,
					"output":       prefop.Items[0].Key,
				}).Info("七牛云转码文件替换异常")
				return fmt.Errorf("七牛云转码文件替换异常")
			}
			return nil
		}
		return fmt.Errorf("七牛云转码成功后文件未知异常")
	case 1, 2:
		time.Sleep(time.Second * 1)
		return q.transcodingStatus(persistentId)
	default:
		logrus.WithFields(logrus.Fields{
			"result":       prefop,
			"persistentId": persistentId,
		}).Info("七牛云转码异常")
		return fmt.Errorf("七牛云转码异常")
	}
}
