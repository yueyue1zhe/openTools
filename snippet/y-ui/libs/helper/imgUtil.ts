import {scopeWritePhotosAlbum} from "./scope";
import {toast} from "../function/toast";

export function tmpPathSavePhotosAlbum(tmpPath: string) {
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

export function urlSaveToPhotosAlbum(url: string) {
    return new Promise(async (resolve, reject) => {
        if (url.indexOf("wxfile://") === 0) {
            console.log(url, "urlSaveToPhotosAlbum");
            resolve(url);
        }
        uni.showLoading({
            title: "保存中..."
        })
        uni.downloadFile({
            url: url,
            success: (res) => {
                if (res.statusCode === 200) {
                    tmpPathSavePhotosAlbum(res.tempFilePath).then(resolve).catch(reject);
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
const urlToTmpPath = (src:string):Promise<string> =>{
    return new Promise<string>((resolve, reject) => {
        uni.downloadFile({
            url:src,
            success:(res)=>{
                resolve(res.tempFilePath);
            },
            fail:(err=>{
                reject(err);
            })
        })
    })
}
export default {
    urlSaveToPhotosAlbum,
    tmpPathSavePhotosAlbum,
    urlToTmpPath,
}