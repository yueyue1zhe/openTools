package file

import (
	"errors"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/jpeg"
	"os"
)

type QrCode struct {
	Url   string
	Size  int
	Ext   string
	Level qr.ErrorCorrectionLevel
	Mode  qr.Encoding
	File  *File
}

const ExtJpg = ".jpg"

func (y *File) NewQrcode(url string, size int) *QrCode {
	return &QrCode{
		Url:   url,
		Size:  size,
		Ext:   ExtJpg,
		Level: qr.M,
		Mode:  qr.Auto,
		File:  y,
	}
}

func (m *QrCode) GetQrCodeExt() string {
	return m.Ext
}

func (m *QrCode) EnCode(path string) (string, error) {
	code, err := qr.Encode(m.Url, m.Level, m.Mode)
	if err != nil {
		return "", errors.New("二维码生成失败")
	}
	code, err = barcode.Scale(code, m.Size, m.Size)
	if err != nil {
		return "", errors.New("二维码设置失败")
	}
	f, err := os.Create(m.File.AttachmentPath(path))
	if err != nil {
		return "", errors.New("创建文件失败")
	}
	defer f.Close()

	err = jpeg.Encode(f, code, nil)
	if err != nil {
		return "", err
	}
	return path, nil
}
