export const imgUrlToTmpPath = (src:string):Promise<string> =>{
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