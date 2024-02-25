package enginehelper

import (
	"io/fs"
	"os"
	"path/filepath"
)

func removePwdMov() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("pwd fail", err.Error())
	}
	var waitPaths []string
	if err = filepath.Walk(pwd, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(info.Name()) == ".mov" {
			waitPaths = append(waitPaths, path)
		}
		return nil
	}); err != nil {
		log.Fatal("walk fail", err.Error())
	}
	for _, path := range waitPaths {
		if err = os.Remove(path); err != nil {
			log.Println("remove fail", path)
		}
	}
	log.Println("finished")
}
