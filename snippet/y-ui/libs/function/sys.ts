function os() {
	return uni.getSystemInfoSync().platform;
}

function sys() {
	return uni.getSystemInfoSync();
}

export default {
	os,
	sys
}
