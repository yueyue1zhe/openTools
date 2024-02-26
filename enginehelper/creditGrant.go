package _integral

import (
	"d-photo/logic"
	"e.coding.net/zhechat/magic/taihao/bootstrap"
	"e.coding.net/zhechat/magic/taihao/library/dbutil"
	"fmt"
	"time"
)

type CreditGrant struct {
	dbutil.Model
	//来源
	Source int `json:"source" gorm:"type:int;not null"`
	//代理商的是代理
	SourceID uint  `json:"source_id" gorm:"type:int;not null"`
	Count    int64 `json:"count" gorm:"type:int;not null"`
	//使用的
	Used int64 `json:"used" gorm:"type:int;not null"`

	//归档后 解锁 count-used
	PlaceOnFile       bool `json:"place_on_file" gorm:"type:bool;not null"`
	PlaceOnFileReason int  `json:"place_on_file_reason" gorm:"type:int;not null"`

	Reason int `json:"reason" gorm:"type:int;not null"`

	//过期时间 用户的积分有过期时间，到期扫描恢复
	Expire int64 `json:"expire" gorm:"type:int;not null"`

	ReceiverUid uint `json:"receiver_uid" gorm:"type:int;not null"`

	Content string `json:"content" gorm:"type:string;not null"`
}

const (
	CreditReasonUniFollow = iota + 1 //关注平台二维码
	CreditReasonTenantUserRegister
)
const (
	CreditPlaceOnFileReasonLost  = iota + 1 //过期归档
	CreditPlaceOnFileReasonEmpty            //消耗光归档
)
const (
	CreditSourceUni = iota + 1 //平台
	CreditSourceTenant
)

func CreditUniGrantTotal() (total, used, lock int64) {
	//积分总量 发放出去锁定的积分总量
	//自己的积分 划扣过来 要从账户扣除掉
	//var (
	//	waitUseList []CreditGrant
	//	allCount    int64
	//	allUsed     int64
	//)
	//dbutil.Get().Model(&CreditGrant{}).Where(CreditGrant{
	//	Source: CreditSourceUni,
	//}).Where("place_on_file = ? AND count > used AND expire > ?", false, time.Now().Unix()).
	//	Order("expire asc").Find(&waitUseList)
	//if len(waitUseList) <= 0 {
	//	return
	//}
	//for _, record := range waitUseList {
	//	allCount += record.Count
	//	allUsed += record.Used
	//}
	//return allCount - allUsed
	return
}

func GrantByFollowUni(uid uint) error {
	count := int64(50)
	expire := time.Now().AddDate(0, 0, 1).Unix()

	var up logic.UserProfile
	if err := dbutil.Get().First(&up, uid).Error; err != nil {
		return fmt.Errorf("指定用户数据未找到 %v", err.Error())
	}
	grantData := CreditGrant{
		Source:      CreditSourceUni,
		Count:       count,
		Reason:      CreditReasonUniFollow,
		Expire:      expire,
		ReceiverUid: uid,
		Content:     fmt.Sprintf("关注指定平台公众号获取 %v  有效期至 %v", count, formatTimestamp(expire)),
	}
	if err := dbutil.Get().Create(&grantData).Error; err != nil {
		return fmt.Errorf("分发流程创建异常 %v", err.Error())
	}
	facade := GetByOther(bootstrap.FacadeGetByTenantID(up.TenantID))
	uc := facade.UserCredit(uid)
	if err := uc.creditGet(grantData); err != nil {
		return err
	}
	return nil
}
