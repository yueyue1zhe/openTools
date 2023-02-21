export function toast(title: string | number, callback: YCallBack | false = false, duration = 1500) {
    uni.showToast({
        title: String(title),
        icon: 'none',
        duration: duration
    })
    if (typeof callback != "boolean") {
        setTimeout(() => {
            callback()
        }, duration)
    }
}


export function toastFastModal(content: string, opt ?: { title?: string, page?: string }) {
    uni.showModal({
        title: opt?.title || "系统提示",
        content: content,
        showCancel: false,
        success: () => {
            uni.reLaunch({
                url: opt?.page || "/pages/index/index"
            })
        }
    })
}