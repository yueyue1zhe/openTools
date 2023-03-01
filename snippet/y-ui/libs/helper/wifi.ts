const connect = (account: string, password: string): Promise<any> => {
    return new Promise((resolve, reject) => {
        // #ifdef MP-WEIXIN
        const connectStart = () => {
            wx.onWifiConnectedWithPartialInfo((res) => {
                res.wifi.SSID = trimSSIDStr(res.wifi.SSID);
                if (res.wifi.SSID == account) {
                    resolve("连接成功 listener")
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
                        return;
                    }
                    let eachNum = 0;
                    const eachConnectedDose = ()=>{
                        if (eachNum < 4){
                            setTimeout(()=>{
                                eachConnected();
                                eachNum++;
                            },500)
                        }
                    }
                    const eachConnected = ()=>{
                        getConnectWifiSSID().then(res=>{
                            if (res == account){
                                resolve("连接成功 each");
                                return
                            }
                            eachConnectedDose();
                        }).catch(()=>{
                            eachConnectedDose();
                        })
                    }
                    eachConnected();
                },
                fail: (err) => {
                    reject(errorCodeParse(err));
                },
                complete: () => {
                    uni.hideLoading();
                }
            })
        }
        startWifi().then(()=>{
            getConnectWifi().then(res => {
                if (res.SSID == account) {
                    reject("已连接此Wi-Fi，请勿重复连接");
                    return
                }
                connectStart();
            }).catch(() => {
                connectStart();
            })
        }).catch(err=>{
            reject(err);
        })
        // #endif
    })
}

const getConnectWifiSSID = () => {
    return new Promise<string>((resolve, reject) => {
        startWifi().then(() => {
            getConnectWifi().then(res => {
                resolve(res.SSID);
            }).catch(err => {
                reject(err);
            })
        }).catch(err => {
            reject(err);
        })
    })
}
export default {
    connect,
    getConnectWifiSSID,
}

const startWifi = () => {
    return new Promise<void>((resolve, reject) => {
        // #ifdef MP-WEIXIN
        wx.startWifi({
            success: async () => {
                resolve();
            },
            fail: (err) => {
                reject(errorCodeParse(err));
            }
        })
        // #endif
    })
}
const trimSSIDStr = (val: string): string => {
    return val.replaceAll('"', '')
}
const getConnectWifi = () => {
    // #ifdef MP-WEIXIN
    return new Promise<UniNamespace.WifiInfo>((resolve, reject) => {
        wx.getConnectedWifi({
            partialInfo: true,
            success: (e) => {
                resolve({...e.wifi, SSID: trimSSIDStr(e.wifi.SSID)});
            },
            fail: (err) => {
                reject(errorCodeParse(err));
            }
        })
    })
    // #endif
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
        12010: err.errMsg,
        12011: "应用在后台无法配置 Wi-Fi",
        12013: "Android 设备,系统保存的 Wi-Fi 配置过期，建议忘记 Wi-Fi 后重试",
        12014: "iOS 设备，无效的 WEP / WPA 密码",
    }
    const errMsg = errMap[err.errCode];
    switch (err.errMsg) {
        case "getConnectedWifi:fail:netInfo is null":
        case "getConnectedWifi:fail currentWifi is null":
        case "getConnectedWifi:fail wifi is disabled":
            return "未打开Wi-Fi开关 或 当前未连接Wi-Fi"
    }
    console.warn(err);
    return errMsg ? errMsg : `code:${err.errCode};msg:${err.errMsg}`
}
