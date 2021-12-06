package sys

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"openTools/y/global"
	"openTools/y/sys/conf"
	"time"
)

func registerLog() {
	path := fmt.Sprintf("%v/log/%v.log", global.NewGlobal().IaRoot, conf.AppName)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	writer, _ := rotatelogs.New(
		path+".%y%m%d%h%m",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(720)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)
	logrus.SetOutput(writer)
}
