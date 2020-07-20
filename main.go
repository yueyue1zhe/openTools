package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	initDir()
}

func initDir() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	baseDir, pwdDir := baseDir(path)
	uniDir := fmt.Sprintf(`%v\%v_uni\unpackage\dist\build\h5`, baseDir, pwdDir)
	mobileDir := fmt.Sprintf(`%v\%v\template\mobile`, baseDir, pwdDir)
	uniMove(uniDir, mobileDir)
}

func uniMove(uniPath, mobileDir string) {
	uniFiles, err := ioutil.ReadDir(uniPath)
	if err != nil {
		fmt.Println("uni目录不存在")
		return
	}
	if len(uniFiles) > 0 {
		fmt.Println("开始移动uni文件")
		mobileFiles, err := ioutil.ReadDir(mobileDir)
		if err != nil {
			fmt.Println("mobile目录不存在")
			return
		}
		for _, v := range mobileFiles {
			if os.RemoveAll(mobileDir+`\`+v.Name()) != nil {
				fmt.Println(fmt.Sprintf("删除失败：%v", v.Name()))
			} else {
				fmt.Println(fmt.Sprintf("删除成功：%v", v.Name()))
			}
		}
		for _, v := range uniFiles {
			if os.Rename(uniPath+`\`+v.Name(), mobileDir+`\`+v.Name()) != nil {
				fmt.Println(fmt.Sprintf("移动失败：%v", v.Name()))
			} else {
				fmt.Println(fmt.Sprintf("移动成功：%v", v.Name()))
			}
		}
	}
	fmt.Println("uni文件移动操作结束")
}

func baseDir(path string) (baseDir string, pwdDir string) {
	s := strings.Split(path, `\`)
	pwdDir = s[len(s)-1]
	for k, v := range s {
		if v != s[len(s)-1] {
			if k != 0 {
				baseDir += fmt.Sprintf(`\%v`, v)
			} else {
				baseDir += v
			}
		}
	}

	fmt.Println("工作区路径: " + baseDir)
	fmt.Println("当前所在：" + pwdDir)
	return baseDir, pwdDir
}
