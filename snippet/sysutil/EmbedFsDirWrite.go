package fileutil

import (
	"embed"
	"fmt"
	"io/fs"
	"strings"
)

type EmbedFsDirWriteConf struct {
	DirFs     embed.FS                       //处理的源素材所在embed
	Dir       []fs.DirEntry                  //迭代的文件目录
	FsBase    string                         //处理的素材 位于 源素材embed的目录
	OutBase   string                         //输出文件跟目录
	ClearLeft string                         //清理fBase中的文件夹名称
	EnvCall   map[string]func([]byte) []byte //根据指定路径重写文件
}
type EmbedFsDirWriteConfEnvCall map[string]func([]byte) []byte

// EmbedFsDirWrite 嵌入二进制文件中的文件夹 输出到指定目录
// 可清理左侧问降价名称 ｜ 根据指定路径重写文件
func EmbedFsDirWrite(conf EmbedFsDirWriteConf) error {
	var (
		readDir []fs.DirEntry
		err     error
	)
	for _, entry := range conf.Dir {
		if entry.IsDir() {
			readDir, err = conf.DirFs.ReadDir(conf.FsBase + entry.Name())
			if err != nil {
				return fmt.Errorf("源文件子目录获取异常：%v", err.Error())
			}
			if err := EmbedFsDirWrite(EmbedFsDirWriteConf{
				DirFs:     conf.DirFs,
				Dir:       readDir,
				FsBase:    conf.FsBase + entry.Name() + "/",
				OutBase:   conf.OutBase,
				ClearLeft: conf.ClearLeft,
				EnvCall:   conf.EnvCall,
			}); err != nil {
				return err
			}
		} else {
			itemPath := conf.OutBase + "/" + outFsBaseClearLeft(conf.FsBase, conf.ClearLeft)
			if err := PathMustExists(itemPath); err != nil {
				return fmt.Errorf("源文件处理异常：%v", err.Error())
			}
			useContent, err := conf.DirFs.ReadFile(conf.FsBase + entry.Name())
			if err != nil {
				return fmt.Errorf("源文件目标文件获取异常：%v", err.Error())
			}
			seeEnvCall := conf.EnvCall[outFsBaseClearLeft(conf.FsBase+entry.Name(), conf.ClearLeft)]
			if seeEnvCall != nil {
				useContent = seeEnvCall(useContent)
			}
			if err := Bytes2File(itemPath+entry.Name(), useContent); err != nil {
				return fmt.Errorf("源文件写入异常：%v", err.Error())
			}
		}
	}
	return nil
}

func outFsBaseClearLeft(fsBase, left string) string {
	if strings.Index(fsBase, left) == 0 {
		return strings.Replace(fsBase, left, "", 1)
	}
	return fsBase
}
