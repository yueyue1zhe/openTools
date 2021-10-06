package main

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/gookit/gcli/v3/interact"
	"openTools/unpackage"
	"os"
)

func main() {
	var (
		basePath  string
		err       error
		dirEntry  []os.DirEntry
		selectOpt []string
	)
	basePath, _ = os.Getwd()
	dirEntry, err = os.ReadDir(basePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, i := range dirEntry {
		if i.IsDir() {
			selectOpt = append(selectOpt, i.Name())
		}
	}
	ans := interact.SelectOne(
		"Your city name(use array)?",
		selectOpt,
		"",
	)
	color.Comment.Println("your select is: ", ans)
	unpackage.W7Module(basePath + "/" + ans)
}
