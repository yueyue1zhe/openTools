function toast(title: string | number, callback: VoidCallBack | false = false, duration = 1500) {
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

export default toast
