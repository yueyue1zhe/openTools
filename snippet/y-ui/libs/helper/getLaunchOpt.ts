/*
启动参与 注意事项
已启动在后台 重新扫码时 启动参数扔为初始启动的参数

当前认为仅适用于用户来源认定
 */
export function getLaunchOpt() {
    const launchOpt = uni.getLaunchOptionsSync();
    let out = {
        path: launchOpt.path,
        query: launchOpt.query,
    }
    if (launchOpt.query.scene || [1047, 1048, 1049].includes(launchOpt.scene)) {
        out.query = parseQueryScene(out.query)
    }
    return out
}

// 20230305 废弃 重新整合为 getLaunchOpt parseQueryScene
// export function getLaunchQuery() {
//     let glos = uni.getLaunchOptionsSync();
//     if ([1047, 1048, 1049].includes(glos.scene)) {
//         let queryRaw = decodeURIComponent(glos.query.scene);
//         if (!queryRaw) return {};
//         let query: AnyObject = {};
//         let queryArr = queryRaw.split("&");
//         queryArr.forEach(item => {
//             let kv = item.split("=");
//             query[kv[0]] = decodeURIComponent(kv[1]);
//         })
//         return query;
//     }
//     return glos.query;
// }

export function parseQueryScene(query: AnyObject | undefined): AnyObject {
    if (!query) return {};
    if (!query.scene) return {};
    const scene = decodeURIComponent(query.scene);
    const queryArr = scene.split("&");
    let out: AnyObject = {};
    queryArr.forEach(item => {
        const kv = item.split("=");
        out[kv[0]] = decodeURIComponent(kv[1])
    })
    return out;
}