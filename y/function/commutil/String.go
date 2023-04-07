package commutil

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// StringLenLimit 检测字符串最大及最小长度
func StringLenLimit(judge string, min, max int) error {
	if len(judge) < min {
		return fmt.Errorf("长度最小%v位", min)
	}
	if len(judge) > max {
		return fmt.Errorf("长度最大%v位", max)
	}
	return nil
}

// StringRandom 生成随机字符串
func StringRandom(len int) string {
	var str = "abcdefghijklmnopqrstuvwxyABCDEFGHIJKLMNOPQRSTUVWXY"
	return StringRandomUseRaw(len, str)
}

// StringRandomUseRaw 使用指定字符生成随机字符串
func StringRandomUseRaw(len int, raw string) string {
	var container string
	b := bytes.NewBufferString(raw)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(raw[randomInt.Int64()])
	}
	return container
}

// StringSplitUint 逗号分割字符串后转为uint
func StringSplitUint(str string) (res []uint) {
	raw := strings.Split(str, ",")
	if len(raw) > 0 {
		for _, s := range raw {
			if uR, err := StringToUint(s); err == nil {
				res = append(res, uR)
			}
		}
	}
	return
}

// StringCut 截取指定下标范围内字符串
func StringCut(str string, start, end int) string {
	r := []rune(str)
	realEnd := len(r)
	if realEnd > end {
		realEnd = end
	}
	realStart := start
	if realStart < 0 {
		realStart = 0
	}
	if realEnd <= 0 {
		realEnd = 0
		return string(r[realStart:])
	}
	return string(r[realStart:realEnd])
}

// StringToUint 字符串转uint
func StringToUint(str string) (uint, error) {
	out, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(out), nil
}

// StringToInt64 字符串转int 64
func StringToInt64(str string) int64 {
	out, _ := strconv.ParseInt(str, 10, 64)
	return out
}

// StringSplitKeywordArr 将字符串根据关键字拆分成字符串切片,未命中时返回原字符串
func StringSplitKeywordArr(str, keyword string) []string {
	if keyword == "" || !strings.Contains(str, keyword) {
		return []string{str}
	}
	//输入关键词后拆分处理
	splitKey := "+|+split-key+|+"
	useRow := strings.ReplaceAll(str, keyword, keyword+splitKey)
	waitArr := strings.Split(useRow, keyword)
	var tmpArr []string
	for _, s := range waitArr {
		if strings.Contains(s, splitKey) {
			cArr := strings.Split(strings.ReplaceAll(s, splitKey, keyword+","), ",")
			tmpArr = append(tmpArr, cArr...)
		} else {
			tmpArr = append(tmpArr, s)
		}
	}
	return tmpArr
}
