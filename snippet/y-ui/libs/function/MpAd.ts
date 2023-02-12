import {onHide} from "@dcloudio/uni-app";

//传入广告ID 延迟时间单位秒 指定时间后拉起插屏广告
//2001，2002错误时将计时5秒后重新尝试拉起广告
const InterstitialAd = function (adID: string, timeout: number = 0) {
    let interstitialAd: WechatMiniprogram.InterstitialAd | null = null;
    onHide(() => {
        interstitialAd?.destroy();
        interstitialAd = null;
        console.info("interstitialAd is destroy onHide")
    })
    let timeOutID: number;
    let timeShow: boolean = true;
    const mustShow = () => {
        if (!timeShow) return;
        const adShow = () => {
            interstitialAd?.show().catch(err => {
                console.log(err);
                if ([2001, 2002].includes(err.errCode)) {
                    timeOutID = setTimeout(() => {
                        mustShow()
                    }, 5000)
                }
            });
        }
        if (timeout <= 0) {
            adShow()
        } else {
            setTimeout(() => {
                adShow()
            }, timeout * 1000)
        }
    }
    if (wx.createInterstitialAd) {
        interstitialAd = wx.createInterstitialAd({
            adUnitId: adID
        })
        interstitialAd.onLoad(() => {
            console.log('interstitialAd onLoad event emit');
            mustShow();
        })
        interstitialAd.onError((err) => {

            console.log('interstitialAd onError event emit', err)
        })
        interstitialAd.onClose(() => {
            console.log('interstitialAd onClose event emit')
            clearTimeout(timeOutID);
            timeShow = false;
        })
    }
    return
}

// 传入广告id拉起激励视频广告
const RewardedVideoAd = (adID: string) => {
    let rewardedVideoAd: WechatMiniprogram.RewardedVideoAd | null = null;
    return new Promise<void>((resolve, reject) => {
        if (!wx.createRewardedVideoAd) {
            return reject("视频加载失败：请升级微信到最新版本");
        }
        if (!rewardedVideoAd) {
            uni.showLoading({
                title: "赞助视频加载中...",
                mask: true,
            })
            rewardedVideoAd = wx.createRewardedVideoAd({adUnitId: adID})
            rewardedVideoAd.onLoad(() => {
                uni.hideLoading();
                console.log('onLoad event emit')
            })
            rewardedVideoAd.onError((err) => {
                let useErr = errorCodeParse(err.errCode);
                if (useErr == "") useErr = "未知异常，请联系站点管理员：" + err.errMsg;
                reject(useErr);
            })
            rewardedVideoAd.onClose((res) => {
                if (res.isEnded) {
                    resolve()
                } else {
                    reject("播放中途退出");
                }
            })
        }
        rewardedVideoAd?.load().then(() => {
            rewardedVideoAd?.show();
        })
    })
}


export default {
    InterstitialAd,
    RewardedVideoAd,
}

const errorCodeParse = (code: number): string => {
    const map = {
        1000: "赞助广告异常，请稍候再试",
        1001: "使用方法错误，请联系网站管理员",
        1002: "广告单元无效，可能是拼写错误、或者误用了其他APP的广告ID",
        1003: "赞助广告内部错误，请稍候再试",
        1004: "无适合的广告，请稍候再试",
        1005: "广告组件审核中，请稍候再试",
        1006: "广告组件被驳回06，请联系网站管理员",
        1007: "广告组件被驳回07，请联系网站管理员",
        1008: "广告单元已关闭08，请联系网站管理员",
    }
    type mapType = keyof typeof map
    return map[code as mapType] || "";
}