<template>
  <view>
    <view class="y-bottom-button-safe"></view>
    <view class="y-bottom-button-box y-flex y-row-center y-col-center">
      <button @click="clickThis" class="y-reset-button custom-button">
        <slot></slot>
      </button>
    </view>
  </view>
</template>

<script lang="ts" setup>
import {reactive} from "vue";
const emit = defineEmits<{
  (e:"click"):void
}>()
const props = withDefaults(defineProps<{
  safeArea: boolean
}>(), {
  safeArea: true,
})
const state = reactive({
  safeAreaHeight: "0px",
})
const useSafeArea = () => {
  const {
    safeArea,
    screenHeight,
    safeAreaInsets
  } = uni.getSystemInfoSync()
  if (safeArea && props.safeArea) {
    // #ifdef MP-WEIXIN
    state.safeAreaHeight = screenHeight - safeArea.bottom + "px"
    // #endif
    // #ifndef MP-WEIXIN
    state.safeAreaHeight = (safeAreaInsets ? safeAreaInsets.bottom : 0) + "px"
    // #endif
  }
}
useSafeArea()

const clickThis = ()=>{
  uni.$y.throttle(()=>{
    emit('click')
  })
}

</script>

<style lang="scss" scoped>
$button-height:120rpx;
.y-bottom-button-safe{
  height: $button-height;
  padding-bottom: v-bind("state.safeAreaHeight");
}
.y-bottom-button-box {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  height: $button-height;
  color: #FFFFFF;
  background-color: #FAFAFAFF;
  padding-bottom: v-bind("state.safeAreaHeight");
  border-top: 1px solid $y-border-color;
  box-sizing: content-box;

  .custom-button{
    background-color: $y-main-theme-color;
    width: 700rpx;
    line-height: 80rpx;
  }
}
</style>