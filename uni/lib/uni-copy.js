export default function uniCopy({content,success,error}) {
	
	content = typeof content === 'string' ? content : content.toString()
	
	//#ifndef H5
	uni.setClipboardData({
		data: content,
		success: function() {
			success("复制成功~")
			console.log('success');
		},
		fail:function(){
			success("复制失败~")
		}
	});
	//#endif
	
	/**
	 * H5端的复制逻辑
	 */
	// #ifdef H5
	if (!document.queryCommandSupported('copy')) {
		error('浏览器不支持')
	}
	let textarea = document.createElement("textarea")
	textarea.value = content
	textarea.readOnly = "readOnly"
	document.body.appendChild(textarea)
	textarea.select()
	textarea.setSelectionRange(0, content.length)
	let result = document.execCommand("copy")
	if(result){
		success("复制成功~")
	}else{
		error("h5安全拦截")
	}	
	textarea.remove()
	// #endif
}
