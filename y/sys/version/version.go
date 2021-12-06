package version

import (
	"fmt"
	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"openTools/y"
	"openTools/y/global"
	"openTools/y/sys/authorize"
	"openTools/y/sys/conf"
	"os"
	"os/exec"
)

func UpdateVersion() error {
	v, err := authorize.RemoteVersion()
	if err != nil {
		return err
	}
	if v.Version == conf.Version {
		return fmt.Errorf("当前为最新版本无需升级")
	}
	path, err := downLoadVersionFile(v)
	if err != nil {
		return err
	}
	err = os.Rename(path, fmt.Sprintf("%v/%v", global.NewGlobal().IaRoot, conf.AppName))
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	go ServiceUpdateRestart()
	return nil
}

func downLoadVersionFile(v authorize.RemoteVersionData) (path string, err error) {
	path = y.File().MakePath().MakeVersionPath(v.Version)
	r, err := req.Get(v.Url)
	if err != nil {
		return "", fmt.Errorf("下载更新文件网络异常")
	}
	err = r.ToFile(path + conf.AppName)
	if err != nil {
		return "", fmt.Errorf("更新文件保存失败")
	}
	return path + conf.AppName, nil
}

func ServiceUpdateRestart() {
	var (
		appCmdPath string
		err        error
		cmd        *exec.Cmd
		stdout     io.ReadCloser
		opBytes    []byte
	)
	appCmdPath = fmt.Sprintf("%v/%v", y.Global().IaRoot, conf.AppName)
	err = os.Chmod(appCmdPath, 0755)
	if err != nil {
		logrus.Errorf("重启命令授权执行失败：%v", err.Error())
		return
	}
	cmd = exec.Command(appCmdPath, "restart")
	if stdout, err = cmd.StdoutPipe(); err != nil { //获取输出对象，可以从该对象中读取输出结果
		logrus.Errorf("重启命令执行失败：%v", err.Error())
		return
	} else {
		defer stdout.Close()
		if err = cmd.Start(); err != nil { // 运行命令
			logrus.Errorf("重启命令执行失败：%v", err.Error())
			return
		}
		if opBytes, err = ioutil.ReadAll(stdout); err != nil { // 读取输出结果
			logrus.Errorf("重启命令执行输出读取失败：%v", err.Error())
		} else {
			logrus.Println(string(opBytes))
		}
	}
}
