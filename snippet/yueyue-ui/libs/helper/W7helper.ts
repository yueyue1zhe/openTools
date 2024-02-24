const W7Helper = {
    CanUse: (): boolean => {
        return typeof window.microApp != "undefined"
    },
    microApp: window.microApp,
    w7: window.microApp?.getData().w7,

    UploadMiniPackage: (packageUri: string, appid: string, supportType: number) => {
        const route = `/upload?url=${packageUri}&app_id=${appid}&support_type[]=${supportType}`
        W7Helper.w7.navigate({
            modulename: "w7_rangineapi",
            type: "micro",
            route: route,
        })
    },
    BindApp: (url: string, token: string, aes_key: string, cb: VoidCallBack<{
        appID: string,
        appSecret: string,
    }>) => {
        W7Helper.w7.navigate({
            modulename: 'w7_rangineapi',
            type: 'micro',
            route: `/card?url=${url}&token=${token}&aes_key=${aes_key}`,
        });
        W7Helper.microApp.addDataListener((e) => {
            console.log(e);
            if (e.type == "returnData") {
                cb({
                    appID: e.data.app_id,
                    appSecret: e.data.app_secret,
                })
            }
        })
    }
}
export default W7Helper;