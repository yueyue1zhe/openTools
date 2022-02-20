package cmd

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"yueyue-sell/app"
	"yueyue-sell/config"
	"yueyue-sell/y"
)

type Services struct {
	Log service.Logger
	Srv *http.Server
	Cfg *service.Config
}

func GetSrv() service.Service {
	s := &Services{
		Cfg: &service.Config{
			Name:        config.ServiceName,
			DisplayName: config.ServiceName,
			Description: config.ServiceDesc,
		}}
	serv, er := service.New(s, s.Cfg)
	if er != nil {
		logrus.Printf("Set logger error:%s\n", er.Error())
	}
	s.Log, er = serv.SystemLogger(nil)
	return serv
}

func (srv *Services) Start(s service.Service) error {
	if srv.Log != nil {
		_ = srv.Log.Info("Start run http server")
	}
	go srv.StarServer()
	return nil
}

func (srv *Services) Stop(s service.Service) error {
	if srv.Log != nil {
		_ = srv.Log.Info("Start stop http server")
	}
	logrus.Println("Server exiting")
	return srv.Srv.Shutdown(context.Background())
}

func (srv *Services) StarServer() {
	var err error
	gin.DisableConsoleColor()
	routes := gin.Default()
	app.Register(routes)
	srv.Srv = &http.Server{
		Addr:    ":" + config.Env.Http.Port,
		Handler: routes,
	}
	err = srv.Srv.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func ServiceRestart() {
	err := os.Chmod(y.Global().IaRoot+"/"+config.AppName, 0755)
	if err != nil {
		logrus.Errorf("重启命令授权执行失败：%v", err.Error())
		return
	}
	cmd := exec.Command(y.Global().IaRoot+"/"+config.AppName, "restart")
	if stdout, err := cmd.StdoutPipe(); err != nil {
		logrus.Errorf("重启命令执行失败：%v", err.Error())
		return
	} else {
		defer stdout.Close()
		if err = cmd.Start(); err != nil {
			logrus.Errorf("重启命令执行失败：%v", err.Error())
			return
		}
		if opBytes, err := ioutil.ReadAll(stdout); err != nil {
			logrus.Errorf("重启命令执行输出读取失败：%v", err.Error())
		} else {
			logrus.Println(string(opBytes))
		}
	}
}
