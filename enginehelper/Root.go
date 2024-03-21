package core

import (
	"os"
	"path/filepath"
)

// IaRoot 获取当前执行文件所在目录
func IaRoot() string {
	if len(os.Args) > 1 && os.Args[1] == "-test.v" {
		p, _ := os.Getwd()
		return p
	}
	ePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	//20230523 文件路径应优先使用 filepath处理路径
	return filepath.Dir(ePath)
}

func AttachmentRoot() string {
	return IaRoot() + appRegisterConf.AttachmentPath + "/"
}
