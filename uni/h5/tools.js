import conf from '@/common/h5/conf.js';
import util from "@/common/util.js";
import wxjssdk from "@/common/h5/jssdk.js";
export default{
	getUniacid(){
		let raw = util.getUrlParam(location.search,'i');
		return parseInt(raw.replace("/",""));
	},
	initGlobal(){
		uni.$u.vuex("vuex_global_set.attachment_url",conf.attachment_url);
		uni.$u.vuex("vuex_global_set.site_root",conf.site_root);
		uni.$u.vuex("vuex_global_set.waiter_qrcode",conf.waiter_qrcode);

		let apiBase = util.apiBase(conf);
		uni.request({
			url:apiBase + "/addons/chain/rule?i=" + this.getUniacid(),
			method:"POST",
			success(e) {
				if(e.statusCode === 200	&& e.data.errno === 0){
					uni.$u.vuex("addons_chain_rule",e.data.data);
					uni.$emit("addons_chain_rule_register")
				}
			}
		})
		uni.request({
			url:apiBase + "/system/page/common?i=" + this.getUniacid(),
			method:"POST",
			success(e) {
				if(e.statusCode === 200	&& e.data.errno === 0){
					uni.$u.vuex("vuex_common_set",e.data.data);
					uni.$emit("vuex_common_set_register")
				}
			}
		})
		return conf;
	},
	getCurRoute(){
		return this.vmGet().$route;
	},
	vmGet(){
		return getApp();
	},
	login(){
		return new Promise(resolve=>{
			// uni.$once("userInfoReload",()=>{
			// 	location.reload();
			// })
			uni.$u.api.memberInfoShow().then(res=>{
				uni.$u.vuex("vuex_user",res);
				uni.$emit("vuex_user_register");
				resolve();
			})
		}).catch(()=>{
			// uni.$emit("userInfoReload")
		})
	},
	fetchUserInfo(modal){
		return this.login()
	},
	openMap(lat,long,address=""){
		wxjssdk.openLocation(lat,long,address);
	},
	pay(res){
		return wxjssdk.pay(res);
	}
}