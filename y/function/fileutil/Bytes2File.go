package fileutil

import "os"

func Bytes2File(path string, data []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	_, err = f.Write(data)
	if err != nil {
		_ = Remove(path)
		return err
	}
	return nil
}
