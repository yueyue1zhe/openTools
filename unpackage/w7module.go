package unpackage

import (
	"bufio"
	"fmt"
	"io"
	"openTools/utils"
	"os"
	"strings"
	"time"
)

const ignore = "enum.php"

var content = ""

func W7Module(basePath string) {
	var (
		err      error
		destPath string
	)
	destPath = basePath + "_out_" + time.Now().Format("20060102150405")
	err = utils.CopyDir(basePath, destPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dir(destPath + "/lib")
	if content == "" {
		fmt.Println("未找到可操作文件")
		return
	}
	err = os.WriteFile(destPath+"/enum.php", []byte(content), 0755)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	clearFileExpNote(destPath + "/enum.php")
	_ = os.RemoveAll(destPath + "/.git")
	err = os.RemoveAll(destPath + "/lib")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("执行完成")
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
			if index != 1 {
				newLine = clearOther([]byte(newLine))
			} else {
				newLine = "<?php\ndefined('IN_IA') or exit('Access Denied');"
			}
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

func clearOther(line []byte) string {
	if strings.Index(strings.Trim(string(line), " "), "<?php") == 0 {
		return ""
	}
	if strings.Index(strings.Trim(string(line), " "), "defined('IN_IA') or exit('Access Denied');") == 0 {
		return ""
	}
	if strings.Index(strings.Trim(string(line), " "), "#") != -1 {
		eIndex := strings.Index(strings.Trim(string(line), " "), "#")
		return string(line)[:eIndex]
	}
	return string(line)
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

func dir(readPath string) {
	var (
		err     error
		readDir []os.DirEntry
	)
	readDir, err = os.ReadDir(readPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, entry := range readDir {
		if entry.IsDir() {
			dir(readPath + "/" + entry.Name())
		} else {
			if entry.Name() == ignore {
				continue
			}
			content += getFileContent(readPath + "/" + entry.Name())
		}
	}
	return
}

func getFileContent(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(file) + "\n"
}
