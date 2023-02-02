// 判断arr是否为一个数组，返回一个bool值
function isArray(arr: any) {
    return Object.prototype.toString.call(arr) === '[object Array]';
}

// 深度克隆
function deepClone(obj: any): object | [] {
    // 对常见的“非”值，直接返回原来值
    if ([null, undefined, NaN, false].includes(obj)) return obj;
    if (typeof obj !== "object" && typeof obj !== 'function') {
        //原始类型直接返回
        return obj;
    }
    let o
    if (isArray(obj)) {
        o = <any[]>[]
        for (let i in obj as any[]) {
            o[i] = obj[i]
        }
        return o
    } else {
        let keys = Object.keys(obj)
        o = <Record<string, any>>{}
        for (let i in obj as object) {
            o[i] = obj[i]
        }
    }
    return o
}

export default deepClone;
