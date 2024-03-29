package main

// 弹幕消息类型
type PackMsgType int

const (
	PackMsgTypeNone PackMsgType = iota
	PackMsgTypeDanmuMsg
	PackMsgTypeLikeMsg
	PackMsgTypeEnterLive
	PackMsgTypeFollowMsg
	PackMsgTypeGiftMsg
	PackMsgTypeLiveStats
	PackMsgTypeFansclubMsg
	PackMsgTypeShareLive
	PackMsgTypeOffline
)

var packMsgTypeDescs = map[PackMsgType]string{
	PackMsgTypeNone:        "无",
	PackMsgTypeDanmuMsg:    "弹幕消息",
	PackMsgTypeLikeMsg:     "点赞",
	PackMsgTypeEnterLive:   "进房",
	PackMsgTypeFollowMsg:   "关注",
	PackMsgTypeGiftMsg:     "礼物",
	PackMsgTypeLiveStats:   "直播间统计",
	PackMsgTypeFansclubMsg: "粉团",
	PackMsgTypeShareLive:   "直播间分享",
	PackMsgTypeOffline:     "下播",
}

func (t PackMsgType) String() string {
	desc, ok := packMsgTypeDescs[t]
	if !ok {
		return "未知类型"
	}
	return desc
}

// 粉丝团消息类型
type FansclubType int

const (
	FansclubTypeNone FansclubType = iota
	FansclubTypeUpgrade
	FansclubTypeJoin
)

// 直播间分享目标
type ShareType int

const (
	ShareTypeUnknown ShareType = iota
	ShareTypeWeChat
	ShareTypeMoments
	ShareTypeWeibo
	ShareTypeQzone
	ShareTypeQQ
	ShareTypeDouyinFriends
)

// 数据包装器
type BarrageMsgPack struct {
	Type        PackMsgType `json:"Type"`
	ProcessName string      `json:"ProcessName"`
	Data        string      `json:"Data"`
}

// 消息
type Msg struct {
	MsgId     int64   `json:"MsgId"`
	User      MsgUser `json:"User"`
	Content   string  `json:"Content"`
	RoomId    int64   `json:"RoomId"`
	WebRoomId int64   `json:"WebRoomId"`
}

// 粉丝团信息
type FansClubInfo struct {
	ClubName string `json:"ClubName"`
	Level    int    `json:"Level"`
}

// 用户弹幕信息
type MsgUser struct {
	Id             int64         `json:"Id"`
	ShortId        int64         `json:"ShortId"`
	DisplayId      string        `json:"DisplayId"`
	Nickname       string        `json:"Nickname"`
	Level          int           `json:"Level"`
	PayLevel       int           `json:"PayLevel"`
	Gender         int           `json:"Gender"`
	HeadImgUrl     string        `json:"HeadImgUrl"`
	SecUid         string        `json:"SecUid"`
	FansClub       *FansClubInfo `json:"FansClub"`
	FollowerCount  int64         `json:"FollowerCount"`
	FollowStatus   int64         `json:"FollowStatus"`
	FollowingCount int64         `json:"FollowingCount"`
}

func (u *MsgUser) GenderString() string {
	switch u.Gender {
	case 1:
		return "男"
	case 2:
		return "女"
	default:
		return "妖"
	}
}

// 礼物消息
type GiftMsg struct {
	Msg
	GiftId       int64    `json:"GiftId"`
	GiftName     string   `json:"GiftName"`
	GroupId      int64    `json:"GroupId"`
	GiftCount    int64    `json:"GiftCount"`
	RepeatCount  int64    `json:"RepeatCount"`
	DiamondCount int      `json:"DiamondCount"`
	Combo        bool     `json:"Combo"`
	ImgUrl       string   `json:"ImgUrl"`
	ToUser       *MsgUser `json:"ToUser,omitempty"`
}

// 点赞消息
type LikeMsg struct {
	Msg
	Count int64 `json:"Count"`
	Total int64 `json:"Total"`
}

// 直播间统计消息
type UserSeqMsg struct {
	Msg
	OnlineUserCount    int64  `json:"OnlineUserCount"`
	TotalUserCount     int64  `json:"TotalUserCount"`
	TotalUserCountStr  string `json:"TotalUserCountStr"`
	OnlineUserCountStr string `json:"OnlineUserCountStr"`
}

// 粉丝团消息
type FansclubMsg struct {
	Msg
	Type  int `json:"Type"`
	Level int `json:"Level"`
}

// 来了消息
type MemberMessage struct {
	Msg
	CurrentCount int64 `json:"CurrentCount"`
}

// 直播间分享
type ShareMessage struct {
	Msg
	ShareType ShareType `json:"ShareType"`
}
