import validate from "../function/test";

interface RequestConfigHeader {
    [key: string]: string
}
type RequestConfig = {
    baseUrl: string, // 请求的根域名
    // 默认的请求头
    header: RequestConfigHeader ,
    method: "POST" | "GET",
    // 设置为json，返回后uni.request会对数据进行一次JSON.parse
    dataType: 'json',
    // 此参数无需处理，因为5+和支付宝小程序不支持，默认为text即可
    responseType: 'text',
    showLoading: boolean, // 是否显示请求中的loading
    loadingText: string,
    loadingTime: number, // 在此时间内，请求还没回来的话，就显示加载中动画，单位ms
    timer: number | null | any, // 定时器
    loadingMask: boolean, // 展示loading的时候，是否给一个透明的蒙层，防止触摸穿透
}

class Request {
    request(options: UniNamespace.RequestOptions) {
        options.dataType = options.dataType || this.config.dataType;
        options.responseType = options.responseType || this.config.responseType;
        options.url = options.url || '';
        options.data = options.data || {};
        options.header = Object.assign({}, this.config.header, options.header);
        options.method = options.method || this.config.method;
        uni.showNavigationBarLoading();
        return new Promise((resolve: (res: any) => void, reject) => {
            options.success = (res) => {
                if (res.statusCode == 200) {
                    resolve(res.data)
                } else {
                    reject(res)
                }
            }
            options.fail = (resp) => {
                reject(resp)
            }
            options.complete = (response) => {
                // 请求返回后，隐藏loading(如果请求返回快的话，可能会没有loading)
                uni.hideNavigationBarLoading()
                uni.hideLoading();
                // 清除定时器，如果请求回来了，就无需loading
                if (typeof this.config.timer == "number") {
                    clearTimeout(this.config.timer);
                }
                this.config.timer = null;
            }

            // 判断用户传递的URL是否/开头,如果不是,加上/，这里使用了uView的test.js验证库的url()方法
            options.url = validate.url(options.url) ? options.url : (this.config.baseUrl + (options.url.indexOf('/') == 0 ?
                options.url : '/' + options.url));

            // 是否显示loading
            // 加一个是否已有timer定时器的判断，否则有两个同时请求的时候，后者会清除前者的定时器id
            // 而没有清除前者的定时器，导致前者超时，一直显示loading
            if (this.config.showLoading && !this.config.timer) {
                this.config.timer = setTimeout(() => {
                    uni.showLoading({
                        title: this.config.loadingText,
                        mask: this.config.loadingMask
                    })
                    this.config.timer = null;
                }, this.config.loadingTime);
            }
            uni.request(options);
        })

    }

    config: RequestConfig
    'get'
    post

    constructor() {
        this.config = {
            baseUrl: '',
            header: {},
            method: 'POST',
            dataType: 'json',
            responseType: 'text',
            showLoading: true,
            loadingText: '请求中...',
            loadingTime: 800,
            timer: null,
            loadingMask: true,
        }
        // get请求
        this.get = (url: string, data = {}, header = {}) => {
            return this.request({
                method: 'GET',
                url,
                header,
                data
            })
        }

        // post请求
        this.post = (url: string, data = {}, header = {}) => {
            return this.request({
                url,
                method: 'POST',
                header,
                data
            })
        }
    }

}

export default new Request