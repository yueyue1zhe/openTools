package tests

import (
	"d-photo/addons/super"
	"e.coding.net/zhechat/magic/taihao/library/dbutil"
	"testing"
)

func TestNavFirstAiPhoto(t *testing.T) {
	withDb()
	dbutil.Get().Create(&super.HomeNav{
		Model:         dbutil.Model{ID: 1},
		Enable:        true,
		GroupId:       6,
		Title:         "AI梦幻写真馆",
		TopTips:       "快照3秒出片",
		BottomTip:     "在家轻松拍梦幻写真",
		Preview:       "",
		Mode:          super.NavModeLink,
		Tag:           "热门",
		ActionContent: "/pages/index/index",
		ActionTips:    "",
		ActionBtn:     "点击体验",
	})
	dbutil.Get().Create(&super.HomeNav{
		Model:         dbutil.Model{ID: 2},
		Enable:        true,
		GroupId:       5,
		Title:         "玩梗头像",
		TopTips:       "直播伴侣",
		BottomTip:     "姓氏匹配谐音成语",
		Preview:       "div-avatar/preview/family/1.jpg",
		Mode:          super.NavModeLink,
		Tag:           "热门",
		ActionContent: "/pages/diy-avatar/index",
		ActionTips:    "扫码体验",
		ActionBtn:     "点击体验",
	})
	dbutil.Get().Create(&super.HomeNav{
		Model:         dbutil.Model{ID: 3},
		Enable:        true,
		GroupId:       6,
		Title:         "艺术引流码",
		TopTips:       "全域引流神器",
		BottomTip:     "美的不像二维码",
		Preview:       "div-avatar/preview/family/1.jpg",
		Mode:          super.NavModeLink,
		Tag:           "私域必备",
		ActionContent: "/pages/art-qrcode/index",
		ActionTips:    "",
		ActionBtn:     "点击体验",
	})
	dbutil.Get().Create(&super.HomeNav{
		Model:         dbutil.Model{ID: 4},
		Enable:        true,
		GroupId:       7,
		Title:         "AI艺术写真",
		TopTips:       "高端定制操作手册",
		BottomTip:     "AI写真馆绘梦师必备",
		Preview:       "div-avatar/preview/family/1.jpg",
		Mode:          super.NavModeToast,
		Tag:           "待上线",
		ActionContent: "请耐心等待",
		ActionTips:    "",
		ActionBtn:     "",
	})
}
