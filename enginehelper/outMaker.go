package logic

import (
	"bytes"
	"e.coding.net/zhechat/magic/taihao/function/commutil"
	"e.coding.net/zhechat/magic/taihao/function/fileutil"
	"e.coding.net/zhechat/magic/taihao/library/qrcodeutil"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"os"
	"time"
)

type OutMaker struct {
	bottomRaw *image.RGBA
}

func OutMakerGet() *OutMaker {
	return &OutMaker{}
}

func (o *OutMaker) Draw(sourcePath string) (out *image.RGBA, err error) {
	sourceRaw, err := os.Open(fileutil.AttachmentPath(sourcePath))
	if err != nil {
		err = fmt.Errorf("输出文件异常 %v", err.Error())
		return
	}
	defer sourceRaw.Close()
	sourceImg, _, err := image.Decode(sourceRaw)
	if err != nil {
		err = fmt.Errorf("解码输出文件异常 %v", err.Error())
		return
	}
	sWidth := sourceImg.Bounds().Max.X
	sHeight := sourceImg.Bounds().Max.Y
	outWidth := sWidth
	bottomHeight := int(float64(sWidth) / float64(750) * float64(130))
	outHeight := sHeight + bottomHeight
	out = image.NewRGBA(image.Rectangle{
		Max: image.Point{X: outWidth, Y: outHeight},
	})
	draw.Draw(out, out.Bounds(), sourceImg, sourceImg.Bounds().Min, draw.Over)
	br := resize.Resize(uint(sWidth), uint(bottomHeight), o.bottomRaw, resize.Lanczos3)
	draw.Draw(out, out.Bounds(), br, image.Pt(0, -sHeight), draw.Over)
	return
}

func (o *OutMaker) BottomDraw(txt, qrcode string) error {
	const (
		width        = 750
		height       = 130
		margin       = 10
		fontSize     = 36
		descFontSize = 28
	)
	//txt = strings.ReplaceAll(txt, " ", ",")
	//content := "你是无敌的无敌的小可爱你是无敌的无敌的小可爱你是无敌的无敌的小可爱你是无敌的无敌的小可爱"
	//qrcode := "你是无敌的无敌的小可爱你是无敌的无敌的小可爱你是无敌的无敌的小可爱你是无敌的无敌的小可爱"

	desc := time.Now().Format("2006.01.02 15:04:05")
	bg := image.NewRGBA(image.Rectangle{
		Max: image.Point{X: width, Y: height},
	})
	for x := 0; x < bg.Bounds().Max.X; x++ {
		for y := 0; y < bg.Bounds().Max.Y; y++ {
			bg.Set(x, y, commutil.Hex2rgb("#ffffff"))
		}
	}
	buf := new(bytes.Buffer)
	if err := qrcodeutil.Make(buf, qrcode, height-margin*2); err != nil {
		return fmt.Errorf("二维码生成异常 %v", err.Error())
	}
	qrcodeImage, _, err := image.Decode(buf)
	if err != nil {
		return fmt.Errorf("二维码解码异常 %v", err.Error())
	}
	draw.Draw(bg, bg.Bounds(), qrcodeImage, image.Pt(-(width-height), -margin), draw.Over)

	tBrush, err := NewTextBrush("Alibaba-PuHuiTi-Regular.ttf")
	if err != nil {
		return fmt.Errorf("字体加载异常 %v", err.Error())
	}
	tBrush.FontSize = fontSize

	//useContent, err := tBrush.TruncateStringWithFreetype(txt, width-height-margin-fontSize)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	tBrush.DrawFontOnRGBA(bg, image.Pt(margin*2, margin*2), txt)

	tBrush.FontSize = descFontSize
	tBrush.FontColor = image.NewUniform(commutil.Hex2rgb("#909399"))
	tBrush.DrawFontOnRGBA(bg, image.Pt(margin*2, fontSize*2), desc)

	o.bottomRaw = bg
	return nil
}
