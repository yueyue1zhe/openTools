package excelutil

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"reflect"
	"strconv"
)

type Map2ExcelHeaderItem struct {
	Key   string
	Label string
	Show  bool

	Image bool

	ImageHeight float64
}

// Map2ExcelPrint
//
//	网页输出
//	fileName := fmt.Sprintf("%v%v%v.xlsx", time.Now().Format("20060102150405"), "-", title)
//	c.Header("Content-Type", "application/octet-stream")
//	c.Header("Content-Disposition", "attachment; filename="+fileName)
//	c.Header("Content-Transfer-Encoding", "binary")
//	_ = f.Write(c.Writer)
//
//	保存文件
//	f.SaveAs(path)
func Map2ExcelPrint(list interface{}, headers []Map2ExcelHeaderItem) (*excelize.File, error) {
	var err error
	sheet1 := "Sheet1"
	f := excelize.NewFile()
	getValue := reflect.ValueOf(list)
	if getValue.Kind() != reflect.Slice {
		return nil, fmt.Errorf("数据异常")
	}
	headerIndex := 0
	for _, header := range headers {
		if header.Show {
			err = f.SetCellValue(sheet1, string(rune(65+headerIndex))+"1", header.Label)
			if err != nil {
				return nil, fmt.Errorf("制表异常:" + err.Error())
			}
			headerIndex++
		}
	}
	length := getValue.Len()
	if length > 0 {
		line := 2
		for i := 0; i < length; i++ {
			value := getValue.Index(i)
			typeL := value.Type()
			if typeL.Kind() != reflect.Map {
				return nil, fmt.Errorf("制表异常:list must be slice of map")
			}
			lineChr := strconv.Itoa(line)
			i2 := 0
			for _, header := range headers {
				if header.Show {
					if header.Image {
						imgHeight := 50.00
						if header.ImageHeight > 0 {
							imgHeight = header.ImageHeight
						}
						imgPath := value.MapIndex(reflect.ValueOf(header.Key)).String()
						imgFormat := GetPhotoFormat(imgPath, imgHeight)
						err = f.AddPicture(sheet1, string(rune(65+i2))+lineChr, imgPath, &excelize.GraphicOptions{
							AutoFit: true,
							ScaleX:  imgFormat,
							ScaleY:  imgFormat,
						})
					} else {
						err = f.SetCellValue(sheet1, string(rune(65+i2))+lineChr, value.MapIndex(reflect.ValueOf(header.Key)))
					}
					if err != nil {
						return nil, fmt.Errorf("制表异常:" + err.Error())
					}
					i2++
				}
			}
			line++
		}
	}
	return f, nil
}

func GetPhotoFormat(photo string, height float64) float64 {
	file, _ := os.Open(photo)
	img, _, _ := image.Decode(file)
	_ = file.Close()
	b := img.Bounds()
	return height / float64(b.Max.Y)
}
