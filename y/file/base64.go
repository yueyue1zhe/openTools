package file

import (
	"encoding/base64"
	"errors"
	"io"
	"os"
	"strings"
)

func (y *File) Base64Img2File(row string) (path string, err error) {
	i := strings.Split(row, ",")
	if len(i) != 2 {
		return "", errors.New("参数异常")
	}
	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(i[1]))
	path, err = y.MakePath().Image(".png")
	if err != nil {
		return "", errors.New("目录权限异常")
	}

	f, err := os.Create(y.AttachmentPath(path))
	if err != nil {
		return "", errors.New("创建文件失败")
	}
	defer f.Close()

	_, err = io.Copy(f, dec)
	if err != nil {
		return "", errors.New("复制内容失败")
	}
	return path, nil
}
