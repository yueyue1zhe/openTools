import {scopeWritePhotosAlbum} from "@/common/utils/function/scope";
import utils from "@/common/utils";

export function imgTmpPathToToPhotosAlbum(tmpPath :string){
    return new Promise<void>(async (resolve, reject)=>{
        await scopeWritePhotosAlbum()
        uni.saveImageToPhotosAlbum({
            filePath: tmpPath,
            success: function() {
                utils.toast("保存成功");
                resolve();
            },
            fail() {
                utils.toast("图片保存失败");
                reject();
            }
        });
    })
}

export function imgUrlToLocal(url:string){
    return new Promise(async (resolve,reject)=>{
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
                    utils.toast("图片下载失败");
                    reject();
                }
            },
            fail() {
                utils.toast("图片下载失败");
            }
        });
    })
}