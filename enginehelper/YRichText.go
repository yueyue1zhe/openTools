package dto

import (
	"database/sql/driver"
	"encoding/json"
)

type YRichText []YRichTextItem

func (p YRichText) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *YRichText) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &p)
}

type YRichTextItem struct {
	Mode    YRichTextItemMode `json:"mode"`
	Content string            `json:"content"`

	Style string `json:"style"`

	ImageOpts YRichTextItemImageOpts `json:"image_opts" y-ts-types-spacer:"?:"`
}
type YRichTextItemImageOpts struct {
	ShowMenuByLongpress bool   `json:"show_menu_by_longpress" y-ts-types-spacer:"?:"`
	Mode                string `json:"mode" y-ts-types:"YImageMode" y-ts-types-spacer:"?:"`
	Width               string `json:"width" y-ts-types-spacer:"?:"`
	Height              string `json:"height" y-ts-types-spacer:"?:"`
	BorderRadius        string `json:"border_radius" y-ts-types-spacer:"?:"`
}

type YRichTextItemMode = int

const (
	yRichTextItemModeText YRichTextItemMode = iota + 1
	yRichTextItemModeImage
)
