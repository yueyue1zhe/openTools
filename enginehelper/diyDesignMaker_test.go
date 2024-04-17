package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestName00(t *testing.T) {
	const firstReplaceTag = "[first]"
	const firstContentReplaceTag = "[first-content]"
	const basePath = "diy-design"
	const firstFont = basePath + "/fonts/【灵瞳】一醉方休.TTF"
	const firstContentFontPath = basePath + "/fonts/千图小兔体.ttf"
	const bgPath = basePath + "/tmpl/00/bg.png"
	const gapPath = basePath + "/tmpl/00/gap.png"
	const maskPath = basePath + "/tmpl/00/mask.png"
	const useWidth = 1000
	const useHeight = 1000
	maker := DiyDesignMaker{
		Width:           useWidth,
		Height:          useHeight,
		UseFirstContent: true,
		ContentUseIdiom: true,
		OutPath:         "diy-design-test.png",
		Actions: []DiyDesignActionItem{
			{
				Mode:   DiyDesignActionItemModeImage,
				Source: bgPath,
				Width:  useWidth,
				Height: useHeight,
			},
			{
				Mode:     DiyDesignActionItemModeText,
				Source:   firstReplaceTag,
				Left:     155,
				Top:      155,
				Color:    "#0000004D",
				FontSize: 700,
				FontPath: firstFont,
			},
			{
				Mode:     DiyDesignActionItemModeText,
				Source:   firstReplaceTag,
				Left:     150,
				Top:      150,
				Color:    "#00000026",
				FontSize: 700,
				FontPath: firstFont,
			},
			{
				Mode:     DiyDesignActionItemModeText,
				Source:   firstReplaceTag,
				Left:     160,
				Top:      160,
				Color:    "#7A1C12",
				FontSize: 700,
				FontPath: firstFont,
			},
			{
				Mode:   DiyDesignActionItemModeImage,
				Remove: false,
				Source: maskPath,
				Width:  useWidth,
				Height: useHeight,
			},
			{
				Mode:   DiyDesignActionItemModeImage,
				Source: gapPath,
				Width:  useWidth,
				Height: useHeight,
			},
			{
				Mode:     DiyDesignActionItemModeText,
				Source:   firstContentReplaceTag,
				Width:    0,
				Height:   0,
				Left:     250,
				Top:      500,
				Color:    "#7A1C12",
				FontSize: 130,
				FontPath: firstContentFontPath,
			},
		},
		LocalPathCache: nil,
		LocalFileCache: nil,
	}
	fmt.Println(maker.Generate())
	js, _ := json.Marshal(&maker)
	println(string(js))
	f, _ := os.Create("diy-design.json")
	f.Write(js)
	f.Close()
}

func TestName01(t *testing.T) {
	const firstReplaceTag = "[first]"
	const firstContentReplaceTag = "[first-content]"

	const useWidth = 1000
	const useHeight = 1000
	const basePath = "diy-design"

	const firstFont = basePath + "/fonts/造字工房凌毅常规体.ttf"
	const firstContentFontPath = basePath + "/fonts/【灵瞳】一醉方休.TTF"
	const bgPath = basePath + "/tmpl/00/bg.png"
	const gapPath = basePath + "/tmpl/00/gap.png"
	const maskPath = basePath + "/tmpl/01/mask.png"

	maker := DiyDesignMaker{
		Width:           useWidth,
		Height:          useHeight,
		UseFirstContent: true,
		ContentUseIdiom: true,
		OutPath:         "diy-design-test.png",
		Actions: []DiyDesignActionItem{
			{
				Mode:   DiyDesignActionItemModeImage,
				Source: bgPath,
				Width:  useWidth,
				Height: useHeight,
			},
			{
				Mode:     DiyDesignActionItemModeText,
				Source:   firstReplaceTag,
				Left:     95,
				Top:      195,
				Color:    "#0000004D",
				FontSize: 800,
				FontPath: firstFont,
			},
			{
				Mode:     DiyDesignActionItemModeText,
				Source:   firstReplaceTag,
				Left:     90,
				Top:      190,
				Color:    "#00000026",
				FontSize: 800,
				FontPath: firstFont,
			},
			{
				Mode:     DiyDesignActionItemModeText,
				Source:   firstReplaceTag,
				Left:     100,
				Top:      200,
				Color:    "#00275E",
				FontSize: 800,
				FontPath: firstFont,
			},
			{
				Mode:   DiyDesignActionItemModeImage,
				Source: maskPath,
				Width:  useWidth,
				Height: useHeight,
			},
			{
				Mode:   DiyDesignActionItemModeImage,
				Source: gapPath,
				Width:  useWidth,
				Height: useHeight,
			},
			{
				Mode:     DiyDesignActionItemModeText,
				Source:   firstContentReplaceTag,
				Left:     240,
				Top:      500,
				Color:    "#00275E",
				FontSize: 120,
				FontPath: firstContentFontPath,
			},
		},
	}
	fmt.Println(maker.Generate())
	js, _ := json.Marshal(&maker)
	println(string(js))
	f, _ := os.Create("diy-design.json")
	f.Write(js)
	f.Close()
}

func TestName02(t *testing.T) {
	const firstReplaceTag = "[first]"
	const firstContentReplaceTag = "[first-content]"

	const useWidth = 1000
	const useHeight = 1000
	const basePath = "diy-design"

	const firstFont = basePath + "/fonts/【灵瞳】一醉方休.TTF"
	const firstContentFontPath = basePath + "/fonts/义启中秋体.ttf"
	const bgPath = basePath + "/tmpl/02/bg.png"
	const gapPath = basePath + "/tmpl/02/gap.png"

	maker := DiyDesignMaker{
		Width:           useWidth,
		Height:          useHeight,
		UseFirstContent: true,
		ContentUseIdiom: true,
		OutPath:         "diy-design-test.png",
		Actions: []DiyDesignActionItem{
			{
				Mode:   DiyDesignActionItemModeImage,
				Source: bgPath,
				Width:  useWidth,
				Height: useHeight,
			},
			{
				Mode:     DiyDesignActionItemModeText,
				Source:   firstReplaceTag,
				Left:     195,
				Top:      195,
				Color:    "#0000004D",
				FontSize: 600,
				FontPath: firstFont,
			},
			{
				Mode:     DiyDesignActionItemModeText,
				Source:   firstReplaceTag,
				Left:     190,
				Top:      190,
				Color:    "#00000026",
				FontSize: 600,
				FontPath: firstFont,
			},
			{
				Mode:     DiyDesignActionItemModeText,
				Source:   firstReplaceTag,
				Left:     200,
				Top:      200,
				Color:    "#CCB373",
				FontSize: 600,
				FontPath: firstFont,
			},
			{
				Mode:   DiyDesignActionItemModeImage,
				Source: gapPath,
				Width:  useWidth,
				Height: useHeight,
			},
			{
				Mode:     DiyDesignActionItemModeText,
				Source:   firstContentReplaceTag,
				Left:     360,
				Top:      460,
				Color:    "#CCB373",
				FontSize: 90,
				FontPath: firstContentFontPath,
			},
		},
		LocalPathCache: nil,
		LocalFileCache: nil,
	}
	fmt.Println(maker.Generate())
	js, _ := json.Marshal(&maker)
	println(string(js))
	f, _ := os.Create("diy-design.json")
	f.Write(js)
	f.Close()
}
