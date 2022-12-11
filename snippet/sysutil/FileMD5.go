package fileutil

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func FileMD5(filePath string) (string, error) {
	file, err := os.Open(AttachmentPath(filePath))
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func FileReaderMd5(reader io.Reader) string {
	hash := md5.New()
	_, _ = io.Copy(hash, reader)
	return hex.EncodeToString(hash.Sum(nil))
}
