const connect = (account: string, password: string): Promise<any> => {
    return new Promise((resolve, reject) => {
        // #ifdef MP-WEIXIN
        const getConnectWifi = () => {
            return new Promise<WechatMiniprogram.WifiInfo>((r, j) => {
                wx.getConnectedWifi({
                    success: (e) => {
                        if (e.wifi.SSID == "") {
                            j("未检测到已连接的Wi-Fi")
                            return
                        }
                        r(e.wifi);
                    },
                    fail: (e) => {
                        j(errorCodeParse(e))
                    }
                })
            })
        }
        const connectStart = () => {
            wx.onWifiConnectedWithPartialInfo((res) => {
                if (res.wifi.SSID == account || res.wifi.SSID.replaceAll('"', "") == account) {
                    resolve("连接成功")
                    return
                }
                reject("当前连接Wi-Fi为：" + res.wifi.SSID)
            })
            uni.showLoading({
                mask: true,
                title: "请稍候..."
            })
            wx.connectWifi({
                SSID: account,
                password: password,
                success: (e) => {
                    if (e.errMsg != "connectWifi:ok") {
                        reject(e.errMsg);
                    }
                },
                fail: (err) => {
                    reject(errorCodeParse(err));
                },
                complete: () => {
                    uni.hideLoading();
                }
            })
        }
        wx.startWifi({
            success: async () => {
                getConnectWifi().then(res => {
                    if (res.SSID == account) {
                        reject("已连接此Wi-Fi，请勿重复连接");
                        return
                    }
                    connectStart();
                }).catch(() => {
                    connectStart();
                })
            },
            fail: (err) => {
                reject(errorCodeParse(err));
            }
        })
        // #endif
    })
}


export default {
    connect
}

const errorCodeParse = (err: UniNamespace.WifiError): string => {
    const errMap: { [key: number]: string } = {
        12000: "未先调用 startWifi 接口",
        12001: "当前系统不支持相关能力",
        12002: "密码错误",
        12003: "Android 设备连接超时",
        12004: "重复连接 Wi-Fi",
        12005: "Android 设备未打开 Wi-Fi 开关",
        12006: "Android 设备，未打开 GPS 定位开关",
        12007: "用户拒绝授权链接 Wi-Fi",
        12008: "无效 SSID",
        12009: "系统运营商配置拒绝连接 Wi-Fi",
        12010: "系统其他错误",
        12011: "应用在后台无法配置 Wi-Fi",
        12013: "Android 设备,系统保存的 Wi-Fi 配置过期，建议忘记 Wi-Fi 后重试",
        12014: "iOS 设备，无效的 WEP / WPA 密码",
    }
    const errMsg = errMap[err.errCode];
    return errMsg ? errMsg : `code:${err.errCode};msg:${err.errMsg}`
}
