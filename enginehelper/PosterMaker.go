package logic

import (
	"d-photo/addons/resources/static"
	"e.coding.net/zhechat/magic/taihao/function/fileutil"
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"image"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"unicode/utf8"
)

type PosterMaker struct {
	OutPath string

	BgPath string

	Actions []PosterMakerItem

	LocalPathCache []string //待删除path
	LocalFileCache []*os.File
}

type PosterMakerItemMode int

const (
	PosterMakerItemModeImage = iota
	PosterMakerItemModeText
)

type PosterMakerItem struct {
	Mode PosterMakerItemMode // 0 image 1 text

	Remove bool

	Source string
	Width  int
	Height int
	Left   int
	Top    int
}

func (m *PosterMaker) imgMustLocalOpen(path string) (openFile *os.File, err error) {
	openFile, err = fileutil.MustOpen(fileutil.AttachmentPath(path))
	if err != nil {
		return nil, err
	}
	m.LocalFileCache = append(m.LocalFileCache, openFile)
	return
}

func (m *PosterMaker) outClear() {
	for _, localFile := range m.LocalFileCache {
		_ = localFile.Close()
	}
	m.LocalFileCache = nil
	for _, path := range m.LocalPathCache {
		_ = fileutil.Remove(path)
	}
	m.LocalPathCache = nil
}

func (m *PosterMaker) Generate() (path string, err error) {
	outF, err := os.Create(fileutil.AttachmentPath(m.OutPath))
	if err != nil {
		err = fmt.Errorf("文件服务异常 %v", err.Error())
		return
	}

	m.LocalFileCache = append(m.LocalFileCache, outF)
	defer m.outClear()

	bgF, err := m.imgMustLocalOpen(m.BgPath)
	if err != nil {
		err = fmt.Errorf("背景异常:%v", err.Error())
		return
	}
	bgImage, _, err := image.Decode(bgF)
	if err != nil {
		err = fmt.Errorf("背景读取异常 %v", err.Error())
		return
	}
	jpg := image.NewRGBA(bgImage.Bounds())
	draw.Draw(jpg, jpg.Bounds(), bgImage, bgImage.Bounds().Min, draw.Over)

	//textBrush, err := newTextBrush("Alibaba-PuHuiTi-Regular.ttf", m.Poster.InviteCodeSize, image.Black)
	//if err != nil {
	//	//yFile.Remove(outPath)
	//	return "", errors.New("字体加载异常")
	//}
	//textBrush.FontColor = image.NewUniform(common.Hex2rgb(m.Poster.InviteCodeColor))
	//textBrush.DrawFontOnRGBA(jpg, image.Pt(m.Poster.InviteCodeLeft, m.Poster.InviteCodeTop), m.InviteCode)

	for i, action := range m.Actions {
		if action.Mode == PosterMakerItemModeImage {
			if action.Remove {
				m.LocalPathCache = append(m.LocalPathCache, action.Source)
			}
			tmpF, err := m.imgMustLocalOpen(action.Source)
			if err != nil {
				return "", fmt.Errorf("图片%v读取失败:%v", i, err.Error())
			}
			tmpImage, _, err := image.Decode(tmpF)
			if err != nil {
				return "", fmt.Errorf("图片%v解析异常 %v", i, err.Error())
			}
			if action.Width > 0 || action.Height > 0 {
				tmpImage = resize.Resize(uint(action.Width), uint(action.Height), tmpImage, resize.Lanczos3)
			}
			draw.Draw(jpg, jpg.Bounds(), tmpImage, tmpImage.Bounds().Min.Sub(image.Pt(action.Left, action.Top)), draw.Over)
		}
		if action.Mode == PosterMakerItemModeText {

		}
	}

	err = jpeg.Encode(outF, jpg, nil)
	if err != nil {
		err = fmt.Errorf("海报生成失败 %v", err.Error())
	}
	return
}

type textBrush struct {
	FontType  *truetype.Font
	FontSize  float64
	FontColor *image.Uniform
}

func NewTextBrush(FontFilePath string) (*textBrush, error) {
	fontFile, err := static.Embed.ReadFile(fmt.Sprintf("font/%v", FontFilePath))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fontType, err := truetype.Parse(fontFile)
	if err != nil {
		return nil, err
	}
	return &textBrush{FontType: fontType, FontSize: 20, FontColor: image.Black}, nil
}

// 图片插入文字
func (fb *textBrush) DrawFontOnRGBA(rgba *image.RGBA, pt image.Point, content string) {

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(fb.FontType)
	c.SetHinting(font.HintingFull)
	c.SetFontSize(fb.FontSize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fb.FontColor)

	_, _ = c.DrawString(content, freetype.Pt(pt.X, pt.Y+int(fb.FontSize)))
}
func (fb *textBrush) TextWidth(text string) int {
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(fb.FontType)
	c.SetHinting(font.HintingFull)
	c.SetFontSize(fb.FontSize)
	c.SetSrc(fb.FontColor)
	f, _ := c.DrawString(text, freetype.Pt(0, 0))
	return f.X.Ceil()
}

func (fb *textBrush) TruncateStringWithFreetype(s string, maxWidth int) (string, error) {
	// 创建一个新的FreeType字体上下文
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(fb.FontType)
	c.SetFontSize(fb.FontSize)
	c.SetSrc(fb.FontColor)

	// 如果字符串为空，则直接返回
	if s == "" {
		return "", nil
	}

	// 初始化截断字符串和宽度
	var truncated string
	var currentWidth int

	// 逐个rune处理字符串，以避免切断多字节字符
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		runeStr := string(r)

		// 测量当前rune的宽度
		w, err := c.DrawString(runeStr, freetype.Pt(0, 0))
		if err != nil {
			return "", err
		}

		// 如果加上当前rune的宽度超过了最大宽度，则截断字符串并返回
		if currentWidth+w.X.Ceil() > maxWidth && truncated != "" {
			return truncated + "...", nil
		}
		// 更新截断字符串和宽度
		truncated += runeStr
		currentWidth += w.X.Ceil()

		// 移除已处理的rune
		s = s[size:]
	}

	// 如果字符串宽度没有超过最大宽度，则返回原始字符串
	return truncated, nil
}

/*
package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/font/gofont/gofont"
	"golang.org/x/image/math/fixed"
)

// TruncateStringWithFreetype 使用freetype库测量并截断字符串
func TruncateStringWithFreetype(s string, maxWidth fixed.Int26_6, fontSize float64, fontData []byte) (string, error) {
	// 解析字体
	ttf, err := freetype.ParseFont(fontData)
	if err != nil {
		return "", err
	}

	// 创建一个新的FreeType字体上下文
	c := freetype.NewContext()
	defer c.Dispose()

	// 设置DPI
	c.SetDPI(72)

	// 设置字体和大小
	c.SetFont(ttf, fontSize)

	// 如果字符串为空，则直接返回
	if s == "" {
		return "", nil
	}

	// 初始化截断字符串和宽度
	var truncated string
	var currentWidth fixed.Int26_6

	// 逐个rune处理字符串，以避免切断多字节字符
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		runeStr := string(r)

		// 测量当前rune的宽度
		_, _, w, err := c.Bounds(runeStr, freetype.Pt(0, 0))
		if err != nil {
			return "", err
		}

		// 如果加上当前rune的宽度超过了最大宽度，则截断字符串并返回
		if currentWidth+w > maxWidth && truncated != "" {
			return truncated + "...", nil
		}

		// 更新截断字符串和宽度
		truncated += runeStr
		currentWidth += w

		// 移除已处理的rune
		s = s[size:]
	}

	// 如果字符串宽度没有超过最大宽度，则返回原始字符串
	return s, nil
}

func main() {
	// 加载字体数据（这里使用Go的内置字体作为示例）
	fontData, err := truetype.GOBytes(gofont.Regular, nil)
	if err != nil {
		panic(err)
	}

	// 设置最大宽度（以1/64像素为单位）和字体大小
	maxWidth := fixed.Int26_6(200 * 64) // 200像素
	fontSize := 12.0                    // 字体大小，单位点数（pt）

	// 要截断的字符串
	text := "这是一个非常长的字符串，需要被截断以适应特定的宽度限制。"

	// 截断字符串并输出结果
	truncatedText, err := TruncateStringWithFreetype(text, maxWidth, fontSize, fontData)
	if err != nil {
		panic(err)
	}
	fmt.Println("原始字符串:", text)
	fmt.Println("截断后的字符串:", truncatedText)
}
*/
