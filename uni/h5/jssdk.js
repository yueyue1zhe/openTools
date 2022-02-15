import jweixin from 'jweixin-module';
import util from "@/common/util.js";
export default {
	isWxClient() {
		let ua = navigator.userAgent.toLowerCase();
		return ua.match(/MicroMessenger/i) == "micromessenger";
	},
	initJssdk: function(callback) {
		uni.$u.api.officialJssdk({
			uri: location.href.split('#')[0]
		}).then(res => {
			jweixin.config({
				debug: false,
				appId: res.app_id,
				timestamp: res.timestamp,
				nonceStr: res.nonce_str,
				signature: res.signature,
				jsApiList: jsApiList
			});
			jweixin.ready(() => {
				jweixin.showAllNonBaseMenuItem();
				jweixin.showOptionMenu();
				jweixin.hideMenuItems({
					menuList: [
						"menuItem:copyUrl",
						"menuItem:originPage",
						"menuItem:readMode",
						"menuItem:openWithQQBrowser",
						"menuItem:openWithSafari",
						"menuItem:share:email",
						"menuItem:share:qq",
						"menuItem:share:weiboApp",
						"menuItem:favorite",
						"menuItem:share:facebook",
						"menuItem:share:QZone",
					]
				})
				callback && callback()
			})
		})
	},
	pay: function(res) {
		return new Promise((resolve, reject) => {
			jweixin.chooseWXPay({
				timestamp: res.timestamp,
				nonceStr: res.nonceStr,
				package: res.package,
				signType: res.signType,
				paySign: res.paySign,
				success: function(res) {
					resolve(res);
				},
				cancel: function() {
					reject({
						errno: 3,
						message: "用户取消支付"
					})
				},
				fail: function(err) {
					reject({
						errno: 2,
						message: "支付失败：" + err.errMsg
					})
				}
			})
		})
		// return new Promise((resolve,reject)=>{
		// 	if(!this.isWxClient()){
		// 		uni.$u.toast("非微信浏览器调用")
		// 		reject({errno:4,message:"非微信浏览器调用"});
		// 		return;
		// 	}
		// 	let params = {fee,tid,title}
		// 	uni.$u.api.WxJsPay(params).then(res=>{

		// 	}).catch(err=>{
		// 		console.warn(err);
		// 		uni.$u.toast("订单生成失败")
		// 		reject({errno:1,message:"订单生成失败"});
		// 	})
		// })
	},
	closeWindow: function() {
		if (!this.isWxClient()) {
			return;
		}
		jweixin.closeWindow();
	},
	noMenus: function() {
		if (!this.isWxClient()) {
			return;
		}
		this.initJssdk(() => {
			jweixin.hideAllNonBaseMenuItem();
			jweixin.hideOptionMenu();
		})
	},
	ready: function(callback) {
		jweixin.ready(() => {
			callback && callback();
		})
	},
	getLocation: function(callback, errCallBack) {
		this.ready(() => {
			jweixin.getLocation({
				type: 'gcj02',
				success: function(res) {
					callback && callback({
						lat: e.latitude,
						long: e.longitude
					});
				},
				fail: function(e) {
					errCallBack && errCallBack(e);
				}
			})
		});
	},
	openLocation: function(lat, long, address = "") {
		jweixin.openLocation({
			address: address,
			latitude: parseFloat(lat),
			longitude: parseFloat(long),
			scale: 12
		})
	},
	shareIndex: function(share_data) {
		if (!this.isWxClient()) {
			return;
		}
		this.initJssdk(() => {
			let vm = util.vmGet();
			let shareData = {
				title: share_data.title,
				desc: share_data.desc,
				link: vm.vuex_global_set.site_root + "app/" + vm.vuex_global_set.uni_acid + "/index",
				imgUrl: util.doseMedia(share_data.pic)
			};
			jweixin.updateAppMessageShareData(shareData);
			jweixin.updateTimelineShareData(shareData);
		})
	},
	shareVoiceRedPacket: function(share_data, query = {}) {
		if (!this.isWxClient()) {
			return;
		}
		let vm = util.vmGet();
		if (!vm.vuex_user.id) {
			uni.$once("vuex_user_register", () => {
				this.shareVoiceRedPacket(share_data, query);
			})
			return;
		}
		this.initJssdk(() => {
			let shareData = {
				title: util.doseShareVariable(share_data.title),
				desc: util.doseShareVariable(share_data.desc),
				link: util.voiceRedPacketShareUrl(query),
				imgUrl: util.doseMedia(share_data.pic)
			};
			jweixin.updateAppMessageShareData(shareData);
			jweixin.updateTimelineShareData(shareData);
		})
	},
	share: function(share_data, query = {}) {
		if (!this.isWxClient()) {
			return;
		}
		let vm = util.vmGet();
		if (!vm.vuex_user.id) {
			uni.$once("vuex_user_register", () => {
				this.share(share_data, query);
			})
			return;
		}
		this.initJssdk(() => {
			let shareData = {
				title: util.doseShareVariable(share_data.title),
				desc: util.doseShareVariable(share_data.desc),
				link: util.activeShareUrl(query),
				imgUrl: util.doseMedia(share_data.pic)
			};
			jweixin.updateAppMessageShareData(shareData);
			jweixin.updateTimelineShareData(shareData);
		})
	},
	startRecord(callback,statusBack) {
		if (!this.isWxClient()) {
			return;
		}
		this.initJssdk(() => {
			this.callback = callback;
			let t = this;
			jweixin.startRecord();
			jweixin.onVoiceRecordEnd({
				complete: function(res) {
					t.uploadVoice(res.localId)
				}
			});
			statusBack && statusBack();
		})
		
	},
	callback:false,
	minWx:jweixin,
	stopRecord() {
		if (!this.isWxClient()) {
			return;
		}
		let t = this;
		jweixin.stopRecord({
			success: function(res) {
				t.uploadVoice(res.localId)
				t.callback && t.callback(res.localId);
			}
		});
	},
	uploadVoice(localId,callback=false){
		let t = this;
		jweixin.uploadVoice({
		  localId: localId, // 需要上传的音频的本地ID，由stopRecord接口获得
		  isShowProgressTips: 1, // 默认为1，显示进度提示
		  success: function (res) {
			t.callback && t.callback(res.serverId);
		  }
		});
	},
}

const jsApiList = [
	'checkJsApi',
	'updateAppMessageShareData',
	'updateTimelineShareData',
	'onMenuShareWeibo',
	'onMenuShareQZone',
	'startRecord',
	'stopRecord',
	'onVoiceRecordEnd',
	'playVoice',
	'pauseVoice',
	'stopVoice',
	'onVoicePlayEnd',
	'uploadVoice',
	'downloadVoice',
	'chooseImage',
	'previewImage',
	'uploadImage',
	'downloadImage',
	'translateVoice',
	'getNetworkType',
	'openLocation',
	'getLocation',
	'hideOptionMenu',
	'showOptionMenu',
	'hideMenuItems',
	'showMenuItems',
	'hideAllNonBaseMenuItem',
	'showAllNonBaseMenuItem',
	'closeWindow',
	'scanQRCode',
	'chooseWXPay',
	'openProductSpecificView',
	'addCard',
	'chooseCard',
	'openCard'
]
