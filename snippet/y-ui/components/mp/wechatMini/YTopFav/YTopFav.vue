<template>
  <view>
    <!--   #ifdef MP-WEIXIN     -->
    <view v-show="show" class="tips-box">
      <view class="triangle"></view>
      <view class="u-flex u-col-top">
        <slot></slot>
        <view v-if="props.closeBtn" class="close-btn">
          <uni-icons @click="show=false" type="clear" color="#a8a1a2"></uni-icons>
        </view>
      </view>
    </view>
    <!--   #endif     -->
  </view>
</template>

<script setup lang="ts">
//#ifdef MP-WEIXIN
import {computed, ref} from "vue";
import UniIcons from "@/components/uni-ui/lib/uni-icons/uni-icons.vue";

const props = withDefaults(defineProps<{
  show?: boolean,//显示控制
  closeBtn?: boolean,//关闭按钮控制
  custom?: boolean,//自定义页面顶部导航定位
}>(), {
  show: true,
  closeBtn: true,
  custom: false,
})

let show = ref(props.show);

let triangleHeight = '6px';

//计算定位容器元素关键属性
let boxPos = computed(() => {
  let res = uni.getMenuButtonBoundingClientRect();
  let sysInfo = uni.getSystemInfoSync();
  return {
    top: props.custom ? `calc(${res.bottom}px + ${triangleHeight})` : triangleHeight,
    right: `${sysInfo.windowWidth - res.right}px`,
    triangleRight: `calc(${res.width}px - ${triangleHeight} * 4)`,
    minWidth: `${res.width}px`
  };
})
//#endif
</script>

<style lang="scss" scoped>
$box-bg-color: #333333;
$box-font-size: 10px;


.tips-box {
  background-color: $box-bg-color;
  color: #ffffff;
  font-size: $box-font-size;
  position: fixed;
  top: v-bind('boxPos.top');
  right: v-bind('boxPos.right');
  padding: 10rpx;
  border-radius: 10rpx;
  max-width: 375rpx;
  min-width: v-bind('boxPos.minWidth');
  min-height: 42rpx;
  z-index: 99;
}

.triangle {
  width: 0;
  height: 0;
  border: v-bind('triangleHeight') solid transparent;
  border-bottom-color: $box-bg-color;
  position: absolute;
  bottom: 100%;
  right: v-bind('boxPos.triangleRight');
}

.close-btn {
  position: absolute;
  right: 5rpx;
  top: 5rpx;
}
</style>