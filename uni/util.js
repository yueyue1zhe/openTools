// #ifdef H5
import mpTools from "@/common/h5/tools.js";
// #endif
// #ifdef MP-WEIXIN
import mpTools from "@/common/mp-weixin/tools.js";
// #endif
import uniCopy from "@/common/lib/uni-copy.js";	

function openMap(lat,long,address=""){
	mpTools.openMap(lat,long,address);
}

function pay(conf){
	return mpTools.pay(conf);
}

function login(){
	mpTools.login();
}
function fetchUserInfo(modal=true){
	return new Promise((resolve,reject)=>{
		let vm = vmGet();
		if(!vm.vuex_user.id){
			uni.$once("vuex_user_register",()=>{
				fetchUserInfo(true).then(res=>{
					resolve(res);
				}).catch(err=>{
					reject(err);
				})
			})
			return;
		}
		if(vm.vuex_user.avatar != "" && vm.vuex_user.nickname != ""){
			resolve();
			return;
		};
		mpTools.fetchUserInfo(modal).then(res=>{
			resolve(res);
		}).catch(err=>{
			reject(err);
		});
	})
}
function init(){
	return mpTools.initGlobal();
}
function apiBase(conf=false){
	let g = {};
	if(!conf){
		g = vmGet().vuex_global_set;
	}else{
		g = conf
	}
	return g.site_root + "api/frontend/" + g.uni_acid;
}

function getCurPage(){
	let pages = getCurrentPages();
	return pages.pop();
}

/**
 * 获取当前页面路由信息
 * @returns {{fullPath: string, path: string, query: {}}}
 */
function getCurRoute(){
	return mpTools.getCurRoute();
}
function vmGet(){
	return mpTools.vmGet();
}
function apiPost(uri,params={}){
	let query = getCurRoute().query;
	let from_uid = query.from_uid
	let from = query.from;
	if(from_uid)params.from_uid = parseInt(from_uid);
	// #ifdef MP-WEIXIN
	params.from = "wxapp";
	// #endif
	// #ifdef H5
	params.from = "official";
	// #endif
	return new Promise( async (resolve,reject)=>{
		uni.$u.post(uri,params).then(res=>{
			resolve(res);
		}).catch(err=>{
			reject(err);
		})
	})
}
function apiPostResParse(res){
	if(res.errno === 0){
		return res.data;
	}else if(res.errno == 1){
		if(res.message){
			toast(res.message);
		}else{
			console.error(res);
		}
		return false;
	}else if(res.errno == 2){
		toast(res.message,()=>{
			if(res.data.indexOf("/pages/index/index") !== -1){
				uni.$u.vuex("vuex_active_id",false);
				uni.$u.vuex("vuex_active_tag",false);
			}
			rediect2link(res.data);
		});
		return false;
	}else if(res.errno == 40019){
		tokenSet("");
		login();
		return false;
	}else{
		console.log(res);
		return false;
	}
}

function rediect2link(link){
	if (!link)return;
	setTimeout(()=>{
		if(link.indexOf("http") === 0){
			// #ifdef H5
			location.replace(link);
			// #endif
			// #ifndef H5
			uni.redirectTo({
				url:"/pages/webview?uri=" + encodeURIComponent(link)
			})
			// #endif
			return;
		}
		if(link.indexOf("/") !== 0){
			link = "/" + link;
		}
		uni.redirectTo({
			url:link
		})
	},500)
}
function nav2link(link,query={}){
	if (!link)return;
	if(link.indexOf("http") === 0){
		// #ifdef H5
		location.href = link;
		// #endif
		// #ifndef H5
		uni.navigateTo({
			url:"/pages/webview?uri=" + encodeURIComponent(link)
		})
		// #endif
		return;
	}
	uni.$u.route(link,query);
}


function toast(title,callback){
	if(title){
		uni.$u.toast(title);
		setTimeout(()=>{
			callback && callback();
		},1500);
	}else{
		callback && callback();
	}
}

function toActive(active){
	let page = "";
	switch(parseInt(active.mode)){
		case 1:
			page = "/pages/active/index";
		break;
		case 2:
			page = "/pages/active/video";
		break;
		case 3:
			page = "/pages/active/classic";
		break;
		case 4:
			page = "/pages/active/group-buy";
		break;
		case 5:
			page = "/pages/active/free";
		break;
		default:
			console.warn("mode is undefined :",active);
		break;
	}
	uni.$u.route(page,{
		active_id:active.id
	})
}

function voiceRedPacketShareUrl(param={}){
	let vm = vmGet();
	let base = "";
	// #ifdef H5
	base = vm.vuex_global_set.site_root + "app/" + vm.vuex_global_set.uni_acid + "/index";
	param["visit-p"] = "voice-red-packet-custom";
	// #endif
	// #ifndef H5
	base = "/pages/plugin/voice-red-packet/custom";
	// #endif
	if(!param.from_uid && vm.vuex_user.id > 0){
		param.from_uid = vm.vuex_user.id;
	}
	return base + uni.$u.queryParams(param);
}

function activeShareUrl(param={}){
	let url = ""
	switch(parseInt(param.mode)){
		case 1:
			url = officialShareUrl(param);
		break;
		case 2:
			url = officialShareVideoUrl(param);
		break;
		case 3:
			url = officialShareClassicUrl(param);
		break;
		case 4:
			url = officialShareGroupBuyUrl(param);
		break;
		case 5:
			url = officialShareFreeUrl(param);
		break;
		default:
		console.warn("mode is undefined:",param)
		break;
	}
	return url;
}
function officialShareFreeUrl(param={}){
	let vm = vmGet();
	let base = "";
	// #ifdef H5
	base = vm.vuex_global_set.site_root + "app/" + vm.vuex_global_set.uni_acid + "/index";
	// #endif
	// #ifndef H5
	base = "/pages/active/free";
	// #endif
	if(!param.from_uid && vm.vuex_user.id > 0){
		param.from_uid = vm.vuex_user.id;
	}
	return base + uni.$u.queryParams(param);
}
function officialShareGroupBuyUrl(param={}){
	let vm = vmGet();
	let base = "";
	// #ifdef H5
	base = vm.vuex_global_set.site_root + "app/" + vm.vuex_global_set.uni_acid + "/index";
	// #endif
	// #ifndef H5
	base = "/pages/active/group-buy";
	// #endif
	if(!param.from_uid && vm.vuex_user.id > 0){
		param.from_uid = vm.vuex_user.id;
	}
	return base + uni.$u.queryParams(param);
}
function officialShareClassicUrl(param={}){
	let vm = vmGet();
	let base = "";
	// #ifdef H5
	base = vm.vuex_global_set.site_root + "app/" + vm.vuex_global_set.uni_acid + "/index";
	// #endif
	// #ifndef H5
	base = "/pages/active/classic";
	// #endif
	if(!param.from_uid && vm.vuex_user.id > 0){
		param.from_uid = vm.vuex_user.id;
	}
	return base + uni.$u.queryParams(param);
}

function officialShareVideoUrl(param={}){
	let vm = vmGet();
	let base = "";
	// #ifdef H5
	base = vm.vuex_global_set.site_root + "app/" + vm.vuex_global_set.uni_acid + "/index";
	// #endif
	// #ifndef H5
	base = "/pages/active/video";
	// #endif
	if(!param.from_uid && vm.vuex_user.id > 0){
		param.from_uid = vm.vuex_user.id;
	}
	return base + uni.$u.queryParams(param);
}

function officialShareUrl(param={}){
	let vm = vmGet();
	let base = "";
	// #ifdef H5
	base = vm.vuex_global_set.site_root + "app/" + vm.vuex_global_set.uni_acid + "/index";
	// #endif
	// #ifndef H5
	base = "/pages/active/index";
	// #endif
	if(!param.from_uid && vm.vuex_user.id > 0){
		param.from_uid = vm.vuex_user.id;
	}
	return base + uni.$u.queryParams(param);
}



function tokenSet(token){
	let update = {data:"",time:0};
	if(token){
		update.data = token;
		update.time = new Date().getTime();
	}
	uni.$u.vuex('vuex_token', update);
}
function tokenCanUse(token){
	if(!token.data)return false;
	let now = new Date().getTime();
	if(now > token.time + (60*60*2*1000)){
		tokenSet("");
		return false;
	}
	return true;
}











function previewImage(val){
	let urls = [];
	if(Array.isArray(val)){
		val.forEach(item=>{
			urls.push(doseMedia(item))
		})
	}else{
		urls = [doseMedia(val)]
	}
	uni.previewImage({
		urls:urls
	})
}

function doseMedia(value=""){
  if(!value)return '';
  if (value.includes("http"))return value;
  if (value.indexOf("../") === 0) return value;
  if (value.indexOf('data:')===0 && value.indexOf('base64')!=-1) return value;
  return vmGet().vuex_global_set.attachment_url + value;
}

function copy(content){
	return new Promise((resolve,reject)=>{
		uniCopy({
			content: content,
			success:()=>{
				console.log("clipboard success");
				resolve();
			},
			error:(e)=>{
				console.warn("clipboard fail",e);
				reject(e);
			}
		})
	});
}


function pageLoadingShow(){
	uni.$u.vuex("vuex_page_loading",true);
}
function pageLoadingHide(){
	uni.$u.vuex("vuex_page_loading",false);
}

function pageReload(){
	// #ifdef H5
	location.reload();
	// #endif
	// #ifdef MP-WEIXIN
	let page = getCurRoute().fullPath
	uni.reLaunch({
		url:page
	})
	// #endif
}


function mustGetActiveId(e,tag=""){
	let active_id = parseInt(e.active_id);
	if(!active_id){
		let vmActiveId = vmGet().vuex_active_id;
		if(vmActiveId){
			return vmActiveId;
		}
		toast("异常访问",()=>{
			uni.reLaunch({
				url:"/pages/index/index"
			})
		})
		return;
	}
	uni.$u.vuex("vuex_active_id",active_id);
	if(tag){
		uni.$u.vuex("vuex_active_tag",tag);
	}
	return active_id;
}

function getRect(selector){
	return new Promise((resolve) => {
		let view = uni.createSelectorQuery().select(selector);
		view.fields({
			size: true,
			rect: true,
			scrollOffset:true
		}, (res) => {
			resolve(res);
		}).exec();
	})
}

function doseShareVariable(str){
	str = str.replace("[用户昵称]",vmGet().vuex_user.nickname)
	str =  str.replace("[活动名称]",vmGet().vuex_active_title)
	return str;
}


function authorizeScope(scope){
	return new Promise((resolve,reject)=>{
		uni.authorize({
		    scope: scope,
		    success() {
				resolve(true)
			},
			fail() {
				resolve(false);
			}
		})
	})
}

function saveImg2LocalScope(){
	return new Promise( async (resolve,reject)=>{
		let scope = await authorizeScope('scope.writePhotosAlbum');
		if(!scope){
			uni.showModal({
				title:"授权提醒",
				content:"请授权保存到相册、用以保存图片",
				success(e) {
					if(e.confirm){
						uni.openSetting({
							success(res) {
								if(!res.authSetting['scope.writePhotosAlbum']){
									uni.showModal({
										title:"授权提醒",
										content:"你拒绝了授权，图片保存失败",
										showCancel:false
									});
									reject({errno:1,message:"未授权"});
								}else{
									resolve();
								}
							}
						});
					}else{
						reject({errno:1,message:"未授权"});
					}
				}
			})
		}else{
			resolve()
		}
	})
}

function saveImg2Local(url){
	url = doseMedia(url);
	return new Promise(async (resolve,reject)=>{
		await saveImg2LocalScope();
		if(url.indexOf("wxfile://") === 0){
			console.log(url,"util.saveImg2Local");
			resolve(url);
		}
		uni.showLoading({
			title:"保存中..."
		})
		uni.downloadFile({
			url: url,
			success: (res) => {
				if (res.statusCode === 200) {
					uni.saveImageToPhotosAlbum({
						filePath: res.tempFilePath,
						success: function () {
							uni.showToast({
								icon:"none",
								title:"保存成功"
							})
							resolve(res.tempFilePath);
						},
						fail() {
							uni.showToast({
								icon:"none",
								title:"图片保存失败"
							})
							reject({errno:1,message:"图片保存失败"});
						}
					});
				}else{
					uni.showToast({
						icon:"none",
						title:"图片下载失败"
					})
					reject({errno:2,message:"图片下载失败"});
				}
			},
			fail() {
				uni.showToast({
					icon:"none",
					title:"图片下载失败"
				})
				reject({errno:2,message:"图片下载失败"});
			}
		});
	})
}


function getUrlParam(url, name) {
	var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
	var r = url.split('?')[1].match(reg);  //匹配目标参数
	if (r != null) return unescape(r[2]); return null; //返回参数值
}

export default {
	init,
	tokenSet,
	tokenCanUse,
	vmGet,
	apiPost,
	apiBase,
	toast,
	rediect2link,
	nav2link,
	activeShareUrl,
	doseMedia,
	pageLoadingShow,
	pageLoadingHide,
	copy,
	mustGetActiveId,
	login,
	apiPostResParse,
	fetchUserInfo,
	pageReload,
	openMap,
	pay,
	getRect,
	doseShareVariable,
	saveImg2Local,
	getUrlParam,
	authorizeScope,
	toActive,
	previewImage,
	voiceRedPacketShareUrl,
};