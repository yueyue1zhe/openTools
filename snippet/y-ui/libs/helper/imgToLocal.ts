import {scopeWritePhotosAlbum} from "./scope";
import {toast} from "../function/toast";

export function imgTmpPathToToPhotosAlbum(tmpPath: string) {
    return new Promise<void>(async (resolve, reject) => {
        await scopeWritePhotosAlbum()
        uni.saveImageToPhotosAlbum({
            filePath: tmpPath,
            success: function () {
                toast("保存成功");
                resolve();
            },
            fail() {
                toast("图片保存失败");
                reject();
            }
        });
    })
}

export function imgUrlToLocal(url: string) {
    return new Promise(async (resolve, reject) => {
        if (url.indexOf("wxfile://") === 0) {
            console.log(url, "util.saveImg2Local");
            resolve(url);
        }
        uni.showLoading({
            title: "保存中..."
        })
        uni.downloadFile({
            url: url,
            success: (res) => {
                if (res.statusCode === 200) {
                    imgTmpPathToToPhotosAlbum(res.tempFilePath).then(resolve).catch(reject);
                } else {
                    toast("图片下载失败");
                    reject();
                }
            },
            fail() {
                toast("图片下载失败");
            }
        });
    })
}