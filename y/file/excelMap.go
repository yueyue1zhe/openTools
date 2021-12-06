package file

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strconv"
)

//excel 全部以服务器本地数据保存下载

type Map2ExcelHeaderItem struct {
	Key   string
	Label string
	Show  bool
}

func (y *File) Map2Excel(list interface{}, headers []Map2ExcelHeaderItem, title string) (path string, err error) {
	sheet1 := "Sheet1"
	f := excelize.NewFile()
	getValue := reflect.ValueOf(list)
	if getValue.Kind() != reflect.Slice {
		return "", fmt.Errorf("list must be slice")
	}
	headerIndex := 0
	for _, header := range headers {
		if header.Show {
			err = f.SetCellValue(sheet1, string(rune(65+headerIndex))+"1", header.Label)
			if err != nil {
				return "", err
			}
			headerIndex++
		}
	}
	length := getValue.Len()
	if length > 0 {
		line := 2
		for i := 0; i < length; i++ {
			value := getValue.Index(i)
			typel := value.Type()
			if typel.Kind() != reflect.Map {
				return "", fmt.Errorf("list must be slice of map")
			}
			lineChr := strconv.Itoa(line)
			i2 := 0
			for _, header := range headers {
				if header.Show {
					err = f.SetCellValue(sheet1, string(rune(65+i2))+lineChr, value.MapIndex(reflect.ValueOf(header.Key)))
					if err != nil {
						return "", err
					}
					i2++
				}
			}
			line++
		}
	}
	path, _ = y.MakePath().Excel(title)
	if err = f.SaveAs(y.AttachmentPath(path)); err != nil {
		return "", err
	}
	return path, nil
}
