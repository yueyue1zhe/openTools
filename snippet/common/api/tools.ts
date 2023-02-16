import {useGlobalStore} from "@/stores/global";
import conf from "@/conf";
import service from "@/common/service";
import {useUserStore} from "@/stores/user";

function apiPost<T>(uri: string, params: AnyObject = {}) {
    //请求前拦截
    return new Promise((resolve: (res: T) => void, reject) => {
        uni.$y.request.config.baseUrl = apiBase();
        uni.$y.request.config.header[conf.JWTTokenKey] = useUserStore().token.data;
        params = {
            ...params,
            ...getRequestOtherParam()
        }
        uni.$y.request.post(uri, params).then(resp => {
            let res = apiPostResParse<T>(resp)
            if (typeof res != "boolean") {
                resolve(res)
            } else {
                reject()
            }
        })
    })
}

interface requestOtherParamTypes {
    launch_from_uid?: number;
    from_channel: string;
    from_client_version: string;
}

function getRequestOtherParam(): requestOtherParamTypes {
    const systemInfo = uni.getSystemInfoSync()
    const globalStore = useGlobalStore();
    let obj: requestOtherParamTypes = {
        from_channel: systemInfo.uniPlatform,
        from_client_version: globalStore.version,
    }
    const fromUID = parseInt(uni.$y.getLaunchQuery()?.from_uid);
    if (fromUID) obj.launch_from_uid = fromUID;
    return obj;
}

function uploadFile<T>(uri: string, tmpPath: string, name: string = "file") {
    return new Promise<T>((resolve, reject) => {
        const userStore = useUserStore();
        let formData: AnyObject = getRequestOtherParam()
        let header: AnyObject = {
            "Content-Type": "multipart/form-data",
        };
        if (userStore.token.data) header[conf.JWTTokenKey] = userStore.token.data;
        uni.uploadFile({
            url: apiBase() + uri + uni.$y.queryParams(formData),
            name: name,
            header,
            formData,
            filePath: tmpPath,
            success(result) {
                console.log(result);
                if (result.statusCode != 200) {
                    return reject({
                        result: result,
                        msg: "请求响应异常"
                    })
                }
                let res = apiPostResParse<T>(JSON.parse(result.data));
                typeof res != "boolean" ? resolve(res) : reject({
                    result: result,
                    msg: "请求解析异常"
                })
            },
            fail(e) {
                reject({result: e, msg: "请求发起异常：" + e.errMsg})
            }
        })
    })
}

function apiPostResParse<T>(res: RequestResponse): T | boolean {
    if (res.message) uni.$y.toast(res.message);
    switch (res.errno) {
        case 0:
            return res.data as T
        case 1:
            return false;
        case 2:
            uni.$y.toast(res.message, () => {
                service.linkGoTo.nav(res.data)
            })
            return false;
        case 4:
            uni.$y.toast(res.message, () => {
                service.linkGoTo.redirect(res.data);
            })
            return false;
        case 40019:
            return false;
        default:
            console.warn("request fail:", res);
            return false;
    }
}


function apiBase(): string {
    return useGlobalStore().site_root + uni.getSystemInfoSync().uniPlatform
}

export default {
    apiBase,
    apiPost,
    uploadFile,
    apiPostResParse,
}