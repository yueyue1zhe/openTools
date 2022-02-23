import util from "@/common/util.js";
const install = (Vue, vm) => {
	let conf = util.init();
	Vue.prototype.$u.http.setConfig({
		baseUrl: util.apiBase(conf),
		loadingText: '努力加载中~',
		loadingTime: 2000,
		// ......
	});
	
	Vue.prototype.$u.http.interceptor.request = (config) => {
		config.header.ytoken = vm.vuex_token.data;
		return config;
	}
	
	Vue.prototype.$u.http.interceptor.response = (res) => {
		return util.apiPostResParse(res);
	}
}

export default {
	install
}