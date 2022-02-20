package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"yueyue-sell/cmd"
	"yueyue-sell/config"
	"yueyue-sell/config/mysql"
)

func main() {
	mysql.Register()
	s := cmd.GetSrv()
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "version":
			fmt.Println(config.Version)
			return
		case "install":
			err := s.Install()
			if err != nil {
				logrus.Fatalf("Install service error:%s\n", err.Error())
			}
			fmt.Printf("服务已安装\n")
		case "uninstall":
			err := s.Uninstall()
			if err != nil {
				logrus.Fatalf("Uninstall service error:%s\n", err.Error())
			}
			fmt.Printf("服务已卸载\n")
		case "start":
			err := s.Start()
			if err != nil {
				logrus.Fatalf("Start service error:%s\n", err.Error())
			}
			fmt.Printf("服务已启动\n")
		case "restart":
			err := s.Restart()
			if err != nil {
				logrus.Fatalf("restart service error:%s\n", err.Error())
			}
			fmt.Printf("服务已重启\n")
		case "stop":
			err := s.Stop()
			if err != nil {
				logrus.Fatalf("top service error:%s\n", err.Error())
			}
			fmt.Printf("服务已关闭\n")
		}
		return
	}
	err := s.Run()
	if err != nil {
		logrus.Fatalf("Run programe error:%s\n", err.Error())
	}
}
