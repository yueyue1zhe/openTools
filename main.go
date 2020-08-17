package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const DefaultIgnore = `.gitignore`

func main() {
	baseDir, pwdDir := initDir()
	var do string
	var preEnd string
	var merge string
	flag.StringVar(&do, "do", "", "操作类型 默认为空")
	flag.StringVar(&preEnd, "p", "v1", "操作类型 默认为空")
	flag.StringVar(&merge, "m", "false", "操作类型 默认为空")
	flag.Parse()
	switch do {
	case "build":
		outPut(baseDir, pwdDir, preEnd, merge)
		fmt.Println("-do build 发布版本 操作结束")
	case "mobile":
		MoveMobileBuild(baseDir, pwdDir)
		fmt.Println("-do mobile 迁移移动端文件 操作结束")
	case "admin":
		MoveAdminBuild(baseDir, pwdDir)
		fmt.Println("-do admin 迁移后台文件 操作结束")
	default:
		fmt.Println("-do 指令缺失")
		fmt.Println("-do build 打包发行版本")
		fmt.Println("-do mobile 迁移移动端文件")
		fmt.Println("-do admin 迁移后台文件")
	}
	fmt.Println(time.Now().String())
}
func outPut(baseDir, pwdDir, preEnd, merge string) {
	codePath := fmt.Sprintf(`%v\%v`, baseDir, pwdDir)
	outPath := fmt.Sprintf(`%v\%v`, baseDir, pwdDir+preEnd)
	_ = os.Mkdir("../"+pwdDir+preEnd, os.ModePerm)
	copyDir(codePath, outPath)
	os.RemoveAll(outPath + `\.git`)
	os.RemoveAll(outPath + `\.idea`)
	os.RemoveAll(outPath + `\openTools.exe`)
	files, err := ioutil.ReadDir(outPath + `\lib`)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, file := range files {
			if merge == "true" {
				if !file.IsDir() && file.Name() != "enum.php" {
					tmpNew := otherPhpCode(outPath + `\lib\` + file.Name())
					tmpNew = strings.Replace(tmpNew, "<?php", "", -1)
					tmpNew = strings.Replace(tmpNew, "defined('IN_IA') or exit('Access Denied');", "", -1)
					tmpOld := fmt.Sprintf(`require_once IA_ROOT . "/addons/%v/lib/%v";`, pwdDir, file.Name())
					replaceFile(outPath+`\lib\enum.php`, tmpOld, tmpNew)
					os.Remove(outPath + `\lib\` + file.Name())
				}
			} else {
				if !file.IsDir() {
					clearFileExpNote(outPath + `\lib\` + file.Name())
				} else {
					replaceFileExpNoteEach(outPath + `\lib\` + file.Name())
				}
			}
		}
		if merge == "true" {
			clearFileExpNote(outPath + `\lib\enum.php`)
		}
	}
}
func replaceFileExpNoteEach(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, file := range files {
			if !file.IsDir() {
				clearFileExpNote(path + `\` + file.Name())
			} else {
				replaceFileExpNoteEach(path + `\` + file.Name())
			}
		}
	}
}

func initDir() (base, pwd string) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return "", ""
	}
	return baseDir(path)
}
func MoveMobileBuild(baseDir, pwdDir string) {
	uniDir := fmt.Sprintf(`%v\%v_uni\unpackage\dist\build\h5`, baseDir, pwdDir)
	mobileDir := fmt.Sprintf(`%v\%v\template\mobile`, baseDir, pwdDir)
	FilesMove(uniDir, mobileDir, "")
}
func MoveAdminBuild(baseDir, pwdDir string) {
	adminDir := fmt.Sprintf(`%v\%v_admin\dist`, baseDir, pwdDir)
	webDir := fmt.Sprintf(`%v\%v\template`, baseDir, pwdDir)
	FilesMove(adminDir, webDir, "mobile")
}

func FilesOut(path, toPath string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("目录不存在:" + path)
		return
	}
	if len(files) > 0 {
		fmt.Println("输出文件")
		for _, v := range files {
			_, err = CopyFile(path+`\`+v.Name(), toPath+`\`+v.Name())
			if err != nil {
				fmt.Println(fmt.Sprintf("复制失败：%v", v.Name()))
			} else {
				fmt.Println(fmt.Sprintf("移动成功：%v", v.Name()))
			}
		}
	} else {
		fmt.Println("没有可操作的文件")
	}
}

func CopyFile(dstFilePath string, srcFilePath string) (written int64, err error) {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		fmt.Printf("打开源文件错误，错误信息=%v\n", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFilePath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Printf("打开目标文件错误，错误信息=%v\n", err)
		return
	}
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)
}

func FilesMove(path, toPath, ignore string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("目录不存在:" + path)
		return
	}
	if len(files) > 0 {
		fmt.Println("准备移动文件")
		mobileFiles, err := ioutil.ReadDir(toPath)
		if err != nil {
			fmt.Println("目录不存在:" + toPath)
			return
		}
		for _, v := range mobileFiles {
			if v.Name() != ignore && v.Name() != DefaultIgnore {
				if os.RemoveAll(toPath+`\`+v.Name()) != nil {
					fmt.Println(fmt.Sprintf("删除失败：%v", v.Name()))
				} else {
					fmt.Println(fmt.Sprintf("删除成功：%v", v.Name()))
				}
			}
		}
		for _, v := range files {
			if os.Rename(path+`\`+v.Name(), toPath+`\`+v.Name()) != nil {
				fmt.Println(fmt.Sprintf("移动失败：%v", v.Name()))
			} else {
				fmt.Println(fmt.Sprintf("移动成功：%v", v.Name()))
			}
		}
	} else {
		fmt.Println("没有可操作的文件")
	}
}

func baseDir(path string) (baseDir string, pwdDir string) {
	s := strings.Split(path, `\`)
	pwdDir = s[len(s)-1]
	for k, v := range s {
		if v != s[len(s)-1] && k != len(s)-1 {
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

func isReplaceIgnoreDir(name string) bool {
	ignoreDir := []string{
		".idea",
		".DS_Store",
		"node_modules",
		"dist",
		"tests",
	}
	for _, v := range ignoreDir {
		if v == name {
			return true
		}
	}
	return false
}
func replaceEach(pwdPath string, files []os.FileInfo, old, new string) {
	for _, v := range files {
		if v.IsDir() && !isReplaceIgnoreDir(v.Name()) {
			newFiles, newErr := ioutil.ReadDir(pwdPath + `\` + v.Name())
			if newErr != nil {
				fmt.Println(newErr.Error())
			} else {
				replaceEach(pwdPath+`\`+v.Name(), newFiles, old, new)
			}
		} else {
			replaceFile(pwdPath+`\`+v.Name(), old, new)
		}
	}
}
func replaceFile(path, old, new string) {
	in, err := os.Open(path)
	if err != nil {
		fmt.Println("open file fail:", err)
		os.Exit(-1)
	}
	defer in.Close()

	out, err := os.OpenFile(path+".tmp", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("Open write file fail:", err)
		os.Exit(-1)
	}
	defer out.Close()

	br := bufio.NewReader(in)
	index := 1
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read err:", err)
			os.Exit(-1)
		}
		newLine := strings.Replace(string(line), old, new, -1)
		_, err = out.WriteString(newLine + "\n")
		if err != nil {
			fmt.Println("write to file fail:", err)
			os.Exit(-1)
		}
		index++
	}
	in.Close()
	out.Close()
	err = os.Rename(path+".tmp", path)
}

func otherPhpCode(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("error : %s", err)
		return ""
	}
	return string(bytes)
}

func clearFileExpNote(path string) {
	in, err := os.Open(path)
	if err != nil {
		fmt.Println("open file fail:", err)
		os.Exit(-1)
	}
	defer in.Close()

	out, err := os.OpenFile(path+".tmp", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("Open write file fail:", err)
		os.Exit(-1)
	}
	defer out.Close()

	br := bufio.NewReader(in)
	index := 1
	bigGoods := false
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read err:", err)
			os.Exit(-1)
		}
		newLine := ""
		isBidGoods := strings.Index(strings.Trim(string(line), " "), "/**")

		if isBidGoods == int(0) {
			bigGoods = true
		}
		if bigGoods {
			//多行注释 判断是否最后一行
			newLine = ""
			isBidGoodsEnd := strings.Index(strings.Trim(string(line), " "), "*/")
			if isBidGoodsEnd == int(0) {
				bigGoods = false
			}
		} else {
			//不是多行注释
			newLine = clearPhpSmall(line)
		}
		preEnd := "\n"
		if newLine == "" {
			preEnd = ""
		}
		_, err = out.WriteString(newLine + preEnd)
		if err != nil {
			fmt.Println("write to file fail:", err)
			os.Exit(-1)
		}
		index++
	}
	in.Close()
	out.Close()
	err = os.Rename(path+".tmp", path)
}

func clearPhpSmall(line []byte) string {
	isClearGoods := strings.Index(strings.Trim(string(line), " "), "//")
	newLine := ""
	//不是单行注释 直接返回
	if isClearGoods != int(0) {
		newLine = string(line)
	}
	return newLine
}

func copyDir(src string, dest string) {
	src = FormatPath(src)
	dest = FormatPath(dest)
	log.Println(src)
	log.Println(dest)

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("xcopy", src, dest, "/I", "/E")
	case "darwin", "linux":
		cmd = exec.Command("cp", "-R", src, dest)
	}

	outPut, e := cmd.Output()
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println(string(outPut))
}
func FormatPath(s string) string {
	switch runtime.GOOS {
	case "windows":
		return strings.Replace(s, "/", "\\", -1)
	case "darwin", "linux":
		return strings.Replace(s, "\\", "/", -1)
	default:
		fmt.Println("only support linux,windows,darwin, but os is " + runtime.GOOS)
		return s
	}
}
