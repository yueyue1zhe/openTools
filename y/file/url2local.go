package file

import (
	"errors"
	"github.com/imroc/req"
	"strings"
)

func (y *File) UrlImg2File(url string) (path string, err error) {
	if err = y.UrlIsImage(url); err != nil {
		return "", err
	}
	r, err := req.Get(url)
	if err != nil {
		return "", errors.New("网络请求异常")
	}
	path, err = y.MakePath().Image(".png")
	if err != nil {
		return "", errors.New("目录权限异常")
	}
	if err = r.ToFile(y.AttachmentPath(path)); err != nil {
		return "", errors.New("图片保存失败")
	}
	return path, nil
}

func (y *File) UrlIsImage(url string) error {
	r, err := req.Head(url)
	if err != nil {
		return errors.New("网络请求异常")
	}
	if !strings.Contains(r.Response().Header.Get("Content-Type"), "image") {
		return errors.New("不是图片")
	}
	return nil
}
