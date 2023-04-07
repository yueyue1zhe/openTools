package fileutil

import (
	"e.coding.net/zhechat/magic/taihao/core"
	"fmt"
	"io"
	"os"
	"strings"
)

func Remove(path string) error {
	if path == "" {
		return nil
	}
	path = AttachmentPath(path)
	if !FileExists(path) {
		return nil
	}
	return os.Remove(path)
}
func RemoveAll(path string) error {
	if path == "" {
		return nil
	}
	path = AttachmentPath(path)
	if !FileExists(path) {
		return nil
	}
	return os.RemoveAll(path)
}

func MustOpen(path string) (*os.File, error) {
	path = AttachmentPath(path)
	return os.Open(path)
}

// AttachmentPath 补全附件路径
func AttachmentPath(path string) string {
	attachRoot := core.AttachmentRoot()
	if strings.Index(path, attachRoot) != 0 && strings.Index(path, core.IaRoot()) != 0 {
		path = attachRoot + path
	}
	return path
}

// FileExists 检测文件是否存在
func FileExists(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func CopyFile(src, dst string) (int64, error) {
	if !FileExists(src) {
		return 0, fmt.Errorf("源文件不存在")
	}
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer func(source *os.File) {
		_ = source.Close()
	}(source)
	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer func(destination *os.File) {
		_ = destination.Close()
	}(destination)
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// PathMustExists 传入指定路径若不存在 则创建
func PathMustExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return err
}
