package tests

import (
	"fmt"
	"github.com/xjieinfo/xjgo/xjcore/xjexcel"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"testing"
)

type Item struct {
	Name string `excel:"column:A;desc:名称;width:30"`
	Link string `excel:"column:B;desc:链接;width:50"`
}

func TestWord(t *testing.T) {
	contentRar, err := os.ReadFile("content")
	if err != nil {
		panic(err)
	}
	arr := strings.Split(string(contentRar), "\n")
	var list []Item
	var a = regexp.MustCompile("^[\u4e00-\u9fa5]$")

	for _, s := range arr {
		for i, v := range s {
			//golang中string的底层是byte类型，所以单纯的for输出中文会出现乱码，这里选择for-range来输出
			if a.MatchString(string(v)) {
				list = append(list, Item{
					Name: s[i:],
					Link: strings.Replace(s[:i], "---", "", 1),
				})
				break
			}
		}
	}

	f := xjexcel.ListToExcel(list, "测试数据", "表")
	f.SaveAs("out.xls")
}

func TestReadLink(t *testing.T) {
	uri := "https://pan.quark.cn/s/34d541fac6d0%3Ctel%3E573741#/list/share"
	resp, err := http.Get(uri)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}
	fmt.Println(string(body))
}
