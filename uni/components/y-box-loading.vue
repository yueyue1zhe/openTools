<template>
  <view v-if="vuex_page_loading" class="y-box-loading">
    <view class="loading-21">
      <view class="loading-21-item loading-21-item-1"></view>
      <view class="loading-21-item loading-21-item-2"></view>
      <view class="loading-21-item loading-21-item-3"></view>
      <view class="loading-21-item loading-21-item-4"></view>
      <view class="loading-21-item loading-21-item-5"></view>
      <view class="loading-21-item loading-21-item-6"></view>
    </view>
  </view>
</template>

<script>
export default {
  name: "y-box-loading",
  data() {
    return {};
  }
};
</script>

<style scoped="true">
.y-box-loading{
  background-color: #f8f8f8;
  position: fixed;
  left: 0;
  top: 0;
  bottom: 0;
  right: 0;
  z-index: 10000;
}
.loading-21-item {
  width: 40rpx;
  height: 40rpx;
  border-radius: 20rpx;
  backface-visibility: hidden;
  position: absolute;
  animation-name: move;
  animation-timing-function: cubic-bezier(0.4, 0, 1, 0.8);
  animation-iteration-count: infinite;
  animation-duration: 3s;
  top: calc(50% - 200rpx);
  left: 50%;
  transform-origin: -20rpx center;
}
.loading-21-item-1{
  background: #226b80;
  animation-delay: -0.5s;
  opacity: 0;
}
.loading-21-item-2 {
  background: #226b80;
  animation-delay: -1s;
  opacity: 0;
}
.loading-21-item-3 {
  background: #35b0ab;
  animation-delay: -1.5s;
  opacity: 0;
}
.loading-21-item-4 {
  background: #c5f0a4;
  animation-delay: -2s;
  opacity: 0;
}
.loading-21-item-5 {
  background: #226b80;
  animation-delay: -2.5s;
  opacity: 0;
}
.loading-21-item-6 {
  background: #c5f0a4;
  animation-delay: -3s;
  opacity: 0;
}

@keyframes move {
  0% {
    transform: scale(1) rotate(0deg) translate3d(0, 0, 1px);
  }
  30% {
    opacity: 1;
  }
  100% {
    z-index: 10;
    transform: scale(0) rotate(360deg) translate3d(0, 0, 1px);
  }
}
</style>
