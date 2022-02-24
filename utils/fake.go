package utils

import (
	"fmt"
	"openTools/y"
)

var NickNames = []string{
	"醉挽清风", "素衣清颜淡若尘", "夜雨微澜", "墨玲珑", "天生傲骨岂能输",
	"凉薄少女夢", "若只如初见", "秦凝雪月", "镜湖月", "古雪娃",
	"你最倾城", "落叶飘零", "浅夏丿初晴", "糖果控", "瑶风韵蕊",
	"北岸初晴i", "倾弦", "柚子味儿的西瓜", "暮雨浔茶", "一只酷宝贝",
	"紫陌寒", "一米阳光", "花月夜", "一只酷宝贝", "婧婧的旋转",
	"别吵醒寂寞", "当年华褪去生涩", "柒墨姬", "时有幽花", "森花",
	"雪落纷纷", "一曲一场叹", "夕阳下边", "固执", "别怂",
	"无言情话", "失她失心", "两人的开始", "青树柠檬", "寂寞的眼泪",
	"魅雪灵", "酷与孤独", "方蕊", "森花", "时有幽花",
	"柒墨姬", "当年华褪去生涩", "别吵醒寂寞", "忘川秋叶", "踏花游湖",
}

func MakeMember() (nickname, avatar string) {
	yRandom := y.Compute().NewRandom()
	nameIndex := yRandom.Int64(0, 49)
	nickname = NickNames[nameIndex]
	avatarIndex := yRandom.Int64(0, 49)
	avatar = fmt.Sprintf("fake/%v.jpg", avatarIndex)
	return nickname, avatar
}
