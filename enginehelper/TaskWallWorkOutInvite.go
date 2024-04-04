package integral

import (
	"d-photo/logic"
	"e.coding.net/zhechat/magic/taihao/bootstrap"
	"e.coding.net/zhechat/magic/taihao/function/commutil"
	"e.coding.net/zhechat/magic/taihao/library/dbutil"
	"e.coding.net/zhechat/magic/taihao/library/logutil"
	"fmt"
)

func TaskWallWorkOutInvite(uid, inviteUid uint) {
	var err error
	defer func() {
		logutil.Info("任务墙拉新任务完成情况检测结束", uid, err)
	}()
	logutil.Info("任务墙拉新任务完成情况开始检测", uid)
	var up logic.UserProfile
	if err = dbutil.Get().First(&up, uid).Error; err != nil {
		return
	}
	facade := GetByOther(bootstrap.FacadeGetByTenantID(up.TenantID))
	tcac := facade.TenantCreditAwardConf()
	if tcac.TaskWallUserInvitePoint <= 0 || tcac.TaskWallUserInviteMax <= 0 || tcac.TaskWallUserInviteExpireHour <= 0 {
		err = fmt.Errorf("站点未开启邀请好友任务")
		return
	}
	judgeQuery := dbutil.Get().Where(TaskWallUserLog{
		Openid: up.OpenID,
		Mode:   TaskWallModeInvite,
		Uid:    up.ID,
	}).Where("created_at BETWEEN ? AND ?", commutil.TimeTodayStart(), commutil.TimeTodayEnd())
	var judge TaskWallUserLog
	if err = facade.Spared(judgeQuery).First(&judge).Error; err != nil {
		err = fmt.Errorf("用户未领取今日拉新任务")
		return
	}
	judge.WorkOut(fmt.Sprintf("拉新 %v 今日任务完成", inviteUid))
}
