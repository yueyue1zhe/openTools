export function scopeAuthorize(scope: string) {
    return new Promise(resolve => {
        uni.authorize({
            scope: scope,
            success() {
                resolve(true)
            },
            fail() {
                resolve(false);
            }
        })
    })
}

export function scopeWritePhotosAlbum() {
    const scopeKey = "scope.writePhotosAlbum"
    return new Promise<void>(async (resolve, reject) => {
        if (!await scopeAuthorize(scopeKey)) {
            uni.showModal({
                title: "授权提醒",
                content: "请授权保存到相册、用以保存图片",
                success(e) {
                    if (e.confirm) {
                        uni.openSetting({
                            success(res) {
                                if (!res.authSetting[scopeKey]) {
                                    uni.showModal({
                                        title: "授权提醒",
                                        content: "你拒绝了授权，图片保存失败",
                                        showCancel: false
                                    });
                                    reject("未授权");
                                } else {
                                    resolve();
                                }
                            }
                        });
                    } else {
                        reject("未授权");
                    }
                }
            })
        } else {
            resolve()
        }
    })
}
