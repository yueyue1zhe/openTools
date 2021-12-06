package file

import (
	"fmt"
	"io"
	"openTools/y/global"
	"os"
	"strings"
)

type File struct {
	UploadFiled    string
	AttachRes      AttachRes `json:"attach_res"`
	attachmentRoot string
}

func NewFile() *File {
	return &File{
		UploadFiled:    "file",
		attachmentRoot: global.NewGlobal().AttachmentRoot,
	}
}

type AttachRes struct {
	Id   uint   `json:"id"`
	Path string `json:"path"`
	Url  string `json:"url"`
}

func (y *File) Remove(path string) error {
	if path == "" {
		return nil
	}
	path = y.AttachmentPath(path)
	if !y.FileExists(path) {
		return nil
	}
	return os.Remove(path)
}
func (y *File) MustOpen(path string) (*os.File, error) {
	path = y.AttachmentPath(path)
	return os.Open(path)
}

func (y *File) AttachmentPath(path string) string {
	if strings.Index(path, y.attachmentRoot) != 0 {
		path = y.attachmentRoot + path
	}
	return path
}

func (y *File) CopyFile(src, dst string) (int64, error) {
	if !y.FileExists(src) {
		return 0, fmt.Errorf("源文件不存在")
	}
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()
	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func pathExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return err
}

func (y *File) PathExists(path string) error {
	return pathExists(path)
}

func (y *File) FileExists(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func (y *File) Bytes2File(path string, data []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		y.Remove(path)
		return err
	}
	return nil
}
