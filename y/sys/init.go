package sys

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	"github.com/sirupsen/logrus"
	"net/http"
	"openTools/y/sys/conf"
	"openTools/y/sys/env"
	"openTools/y/wechat"
	"os"
	"path/filepath"
)

type Services struct {
	Log service.Logger
	Srv *http.Server
	Cfg *service.Config
}

// 获取可执行文件的绝对路径
func ExecPath() string {
	file, e := os.Executable()
	if e != nil {
		logrus.Printf("Executable file path error : %s\n", e.Error())
	}
	path := filepath.Dir(file)
	return path
}

func GetSrv() service.Service {
	s := &Services{
		Cfg: &service.Config{
			Name:        conf.ServiceName,
			DisplayName: conf.ServiceName,
			Description: conf.ServiceDesc,
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
		srv.Log.Info("Start run http server")
	}
	gin.SetMode(gin.ReleaseMode)
	if err := env.Register(); err == nil {
		if env.Conf.Http.Debug {
			gin.SetMode(gin.DebugMode)
		}
	}
	go srv.StarServer()
	return nil
}

func (srv *Services) Stop(s service.Service) error {
	if srv.Log != nil {
		srv.Log.Info("Start stop http server")
	}
	logrus.Println("Server exiting")
	return srv.Srv.Shutdown(context.Background())
}

func (srv *Services) StarServer() {
	var err error
	gin.DisableConsoleColor()
	routes := gin.Default()
	//router.Register(routes)
	gin.SetMode(gin.ReleaseMode)
	if err = env.Register(); err == nil {
		wechat.Register()
		//db.Register()
		if env.Conf.Http.Debug {
			gin.SetMode(gin.DebugMode)
		}
	}
	srv.Srv = &http.Server{
		Addr:    ":" + env.Conf.Http.Port,
		Handler: routes,
	}
	srv.Srv.ListenAndServe()
}
