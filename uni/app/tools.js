import conf from '@/common/app/conf.js';
import util from "@/common/util.js";

function getUniacid(){
	return conf.uni_acid;
}

function initGlobal(){
	uni.$u.vuex("vuex_global_set.site_root",conf.site_root);
	
	console.log(conf);
	let apiBase = util.apiBase(conf);
	uni.request({
		url:apiBase + "/system/page/global?i=" + getUniacid(),
		method:"POST",
		success(e) {
			if(e.statusCode === 200	&& e.data.errno === 0){
				uni.$u.vuex("vuex_global_set.attachment_url",e.data.data.attachment_url);
				uni.$u.vuex("vuex_global_set.waiter_qrcode",e.data.data.waiter_qrcode);
				uni.$emit("vuex_global_set_register")
			}
		}
	})
	uni.request({
		url:apiBase + "/addons/chain/rule?i=" + getUniacid(),
		method:"POST",
		success(e) {
			if(e.statusCode === 200	&& e.data.errno === 0){
				uni.$u.vuex("addons_chain_rule",e.data.data);
			}
		}
	})
	uni.request({
		url:apiBase + "/system/page/common?i=" + getUniacid(),
		method:"POST",
		success(e) {
			if(e.statusCode === 200	&& e.data.errno === 0){
				uni.$u.vuex("vuex_common_set",e.data.data);
				uni.$emit("vuex_common_set_register")
			}
		}
	})
	return conf;
}
function getCurRoute(){
	let route = {
		query:{},
		path:"",
		fullPath:""
	};
	let pages = getCurrentPages();
	let page = pages.pop();
	route.path = "/" + page.route;
	route.fullPath = "/" + page.route + uni.$u.queryParams(page.options);
	route.query = page.options;
	return route;
}

function localTokenCanUse(){
	return new Promise((resolve,reject)=>{
		let lifeData = uni.getStorageSync("lifeData");
		if(!lifeData){
			reject();
			return;
		}
		if(!lifeData.vuex_token){
			reject();
			return;
		}
		if(!lifeData.vuex_token.data){
			reject();
			return;
		}
		if(!util.tokenCanUse(lifeData.vuex_token)){
			reject();
			return;
		}
		uni.checkSession({
			success() {
				resolve();
			},
			fail() {
				reject();
			}
		})
	})
}

function login(){
	return new Promise(resolve=>{
		localTokenCanUse().then(()=>{
			uni.$u.api.memberInfoShow().then(res=>{
				uni.$u.vuex("vuex_user",res);
				uni.$emit("vuex_user_register");
				resolve();
			})
		}).catch(async ()=>{
			// util.tokenSet("");
			let code = await getLoginCode();
			uni.request({
				url:util.apiBase(conf)+"/wxapp/code2openid",
				data:{code},
				method:"POST",
				success(e) {
					if(e.statusCode !== 200){
						util.toast("网络异常",()=>{
							uni.reLaunch({
								url:"/pages/index/index"
							})
						});
						return
					}
					let res = util.apiPostResParse(e.data);
					if(!res)return;
					util.tokenSet(res.token);
					uni.$u.vuex("vuex_user",res.member);
					uni.$emit("vuex_user_register");
					resolve();
				}
			})
		})
	})
}

function fetchUserInfo(modal){
	return new Promise((resolve,reject)=>{
		getUserProfile(modal).then(res=>{
			uni.$u.api.memberInfoUpdate(res.userInfo).then(res=>{
				uni.$u.vuex("vuex_user",res);
				resolve(res);
			})
		}).catch(err=>{
			reject(err);
		})
	})
}

function vmGet(){
	return getApp().$vm;
}
function openMap(lat,long,address=""){
	uni.openLocation({
		address:address,
		latitude:parseFloat(lat),
		longitude:parseFloat(long),
		scale:12
	});
}
function pay(res){
	return new Promise((resolve,reject)=>{
		wx.requestPayment({
		    'timeStamp': res.timestamp,
		    'nonceStr': res.nonceStr,
		    'package': res.package,
		    'signType': res.signType,
		    'paySign': res.paySign,
		    'success': function (res) {
				resolve(res);
		    },
		    'fail': function (err) {
		        reject({errno:2,message:"支付失败：" + err.errMsg})
		    },
			'cancel':function(){
				reject({errno:3,message:"用户取消支付"})
			}
		})
	})
}
export default{
	initGlobal,
	getCurRoute,
	vmGet,
	login,
	fetchUserInfo,
	openMap,
	pay,
	getUniacid,
}

function getUserProfile(modal){
    return new Promise((resolve, reject) => {
		if(!modal){
			wx.getUserProfile({
			    lang: 'zh_CN',
			    desc: '用于完善会员资料',
			    success: function(wxInfo){
			        resolve(wxInfo);
			    },
			    fail: function(res) {
			        reject(res);
			    }
			})
		}else{
			uni.showModal({
			    title: '获取用户信息',
			    content: '请允许授权用以完善会员资料',
			    cancelColor:"#d4cdcd",
			    success: function (res) {
			        if (res.confirm) {
			            wx.getUserProfile({
			                lang: 'zh_CN',
			                desc: '用于完善会员资料',
			                success: function(wxInfo){
			                    resolve(wxInfo);
			                },
			                fail: function(res) {
			                    reject(res);
			                }
			            })
			        } else if (res.cancel) {
						reject(res);
			        }
			    }
			})
		}
    })
}

function getLoginCode(){
    return new Promise((resolve, reject) => {
        uni.login({
            provider: "weixin",
            success: function (loginRes) {
                if(loginRes.errMsg === "login:ok"){
                    resolve(loginRes.code);
                }
                reject(loginRes.errMsg)
            },
            fail(err) {
                reject(err);
            }
        });
    })
}


