import getCurRoute from "./getCurRoute";

export default function () {
    //    #ifdef H5
    location.reload();
    //    #endif
    //    #ifdef MP-WEIXIN
    let page = getCurRoute().fullPath
    uni.reLaunch({
        url: page,
    })
    // #endif
}