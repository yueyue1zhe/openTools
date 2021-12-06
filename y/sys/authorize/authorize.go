package authorize

import (
	"fmt"
	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
	"openTools/y/sys/conf"
	"os"
	"strconv"
)

const (
	baseUri    = ""
	authUri    = ""
	versionUri = ""
)

type RemoteRes struct {
	Errno   int64  `json:"errno"`
	Message string `json:"message"`
}
type RemoteAuthRes struct {
	RemoteRes
	Data RemoteAuthData `json:"data"`
}
type RemoteAuthData struct {
	Id               string `json:"_id"`
	Ip               string `json:"ip"`
	Status           bool   `json:"status"`
	StartTime        int64  `json:"start_time"`
	UniTotal         int64  `json:"uni_total"`
	LastVisitTime    int64  `json:"last_visit_time"`
	LastVisitVersion string `json:"last_visit_version"`
	CreateTime       int64  `json:"createtime"`
}

func RemoteAuth() (res RemoteAuthData, err error) {
	var result RemoteAuthRes
	var r *req.Resp
	r, err = req.Get(fmt.Sprintf("%v?v=%v", baseUri+authUri, conf.Version))
	if err != nil {
		return RemoteAuthData{}, fmt.Errorf("服务请求初始化失败")
	}
	if err = r.ToJSON(&result); err != nil {
		return RemoteAuthData{}, fmt.Errorf("服务请求异常:%v", err.Error())
	}
	if result.Errno != 0 {
		return RemoteAuthData{}, fmt.Errorf(result.Message)
	}
	result.Data.StartTime = doseTimestamp(result.Data.StartTime)
	result.Data.LastVisitTime = doseTimestamp(result.Data.LastVisitTime)
	result.Data.CreateTime = doseTimestamp(result.Data.CreateTime)
	return result.Data, nil
}

type RemoteVersionData struct {
	Version string `json:"version"`
	Url     string `json:"url"`
}

func RemoteVersion() (res RemoteVersionData, err error) {
	var r *req.Resp
	r, err = req.Get(baseUri + versionUri)
	if err != nil {
		return RemoteVersionData{}, fmt.Errorf("服务请求初始化失败")
	}
	if err = r.ToJSON(&res); err != nil {
		return RemoteVersionData{}, fmt.Errorf("服务请求异常:%v", err.Error())
	}
	if res.Url == "" || res.Version == "" {
		return RemoteVersionData{}, fmt.Errorf("版本服务异常")
	}
	return res, nil
}

func doseTimestamp(timestamp int64) int64 {
	if len(strconv.Itoa(int(timestamp))) > 10 {
		return timestamp / 1000
	}
	return timestamp
}

func RemoteStatusMustOk() {
	auth, err := RemoteAuth()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("授权服务异常中止运行")
		return
	}
	if !auth.Status {
		logrus.Error("站点未授权中止运行")
		os.Exit(0)
	}
}
