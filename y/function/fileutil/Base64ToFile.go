package fileutil

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

func Base64ToFile(row, path string) error {
	i := strings.Split(row, ",")
	if len(i) != 2 {
		return fmt.Errorf("参数异常")
	}
	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(i[1]))
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("创建文件失败:%v", err.Error())
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	_, err = io.Copy(f, dec)
	if err != nil {
		return fmt.Errorf("复制内容失败%v", err.Error())
	}
	return nil
}
