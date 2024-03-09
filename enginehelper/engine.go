package bootstrap

import (
	"e.coding.net/zhechat/magic/taihao/library/logutil"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func RestartUsePath(appCmdPath string) {
	var (
		err     error
		cmd     *exec.Cmd
		stdout  io.ReadCloser
		opBytes []byte
	)
	err = os.Chmod(appCmdPath, 0755)
	if err != nil {
		logutil.ErrorF("重启命令授权执行失败：%v", err.Error())
		return
	}
	cmd = exec.Command(appCmdPath, "restart")
	if stdout, err = cmd.StdoutPipe(); err != nil { //获取输出对象，可以从该对象中读取输出结果
		logutil.ErrorF("重启命令执行失败：%v", err.Error())
		return
	} else {
		defer func(stdout io.ReadCloser) {
			_ = stdout.Close()
		}(stdout)
		if err = cmd.Start(); err != nil { // 运行命令
			logutil.ErrorF("重启命令执行失败：%v", err.Error())
			return
		}
		if opBytes, err = ioutil.ReadAll(stdout); err != nil { // 读取输出结果
			logutil.ErrorF("重启命令执行输出读取失败：%v", err.Error())
		} else {
			logutil.Info(string(opBytes))
		}
	}
}
