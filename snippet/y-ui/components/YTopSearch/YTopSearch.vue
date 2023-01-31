<template>
  <view>
    <view @click="clickThis" class="y-top-search-box uni-radius">
      <uni-icons :color="searchBoxColor" type="search" size="22px"></uni-icons>
      <view class="uni-ml-2">搜索</view>
    </view>
    <view class="top-box-bg"></view>
    <view v-if="fullBgShow" class="full-bg"></view>
  </view>
</template>

<script setup lang="ts">
import {computed, ref} from "vue";

const props = withDefaults(defineProps<{
  bgColor?: string,//仅小程序中有效 未激活背景色 默认透明
  pageScroll?: number,//宿主页面传入的scroll值用于改变背景色
  changeBgLimit?: number,//背景色改变阀值
  activeBgColor?: string,//激活背景色 h5 中始终显示
  fullBg?: boolean,
}>(), {
  bgColor: "rgba(0,0,0,0)",
  pageScroll: 0,
  changeBgLimit: 0,
  activeBgColor: "red",
  fullBg: true,
})

let fullBgShow = ref(props.fullBg);
//#ifdef H5
fullBgShow.value = true;
//#endif

const emit = defineEmits<{ (e: 'click'): void }>();

function clickThis() {
  emit('click')
}


//激活背景色透明度计算
let bgOp = computed(() => {
  let op: number = 1;
  //#ifdef MP-WEIXIN
  let o = 50
  let i = props.changeBgLimit - props.pageScroll;
  if (props.changeBgLimit == 0) return 1;
  op = props.changeBgLimit == 0 ? 1 : (i <= 0 && i >= -o) ? (o - (i + o)) / o : 1;
  if (props.fullBg) op = 1;
  //#endif
  return op;
})

//背景色选择计算
let bgColor = computed(() => {
  let out: string
  //#ifdef MP-WEIXIN
  out = (props.changeBgLimit == 0 || props.changeBgLimit > props.pageScroll) ? props.bgColor : props.activeBgColor;
  if (props.fullBg) out = props.activeBgColor
  //#endif
  //#ifdef H5
  out = props.activeBgColor
  //#endif
  return out
})

//搜索框背景色
let searchBoxColor = 'rgba(255, 255, 255, .6)';

//搜索框关键点定位
let boxPos = computed(() => {
  let out = {
    lr: '',
    top: '',
    width: '',
    height: '',
    bottom: '',
  }
  //#ifdef MP-WEIXIN
  let res = uni.getMenuButtonBoundingClientRect();
  let sysInfo = uni.getSystemInfoSync();
  let LR = sysInfo.windowWidth - res.right
  out.lr = `${LR}px`;
  out.top = `${res.top}px`;
  out.width = `${res.left - (LR * 2)}px`;
  out.height = `${res.height}px`;
  out.bottom = `calc(calc(${res.bottom - res.top}px - var(--status-bar-height)) + ${res.bottom}px)`
  //#endif
  //#ifdef H5
  out.lr = `12.5px`
  out.top = `7px`
  out.width = `calc(100% - 25px)`
  out.height = `30px`
  out.bottom = `44px`
  //#endif
  return out;
})

</script>
<style lang="scss" scoped>
$search-box-color: v-bind('searchBoxColor');
$search-index: 99;
.full-bg {
  height: v-bind('boxPos.bottom');
}

.top-box-bg {
  background-color: v-bind('bgColor');
  height: v-bind('boxPos.bottom');
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  opacity: v-bind('bgOp');
  z-index: $search-index - 1;
}

.y-top-search-box {
  display: flex;
  align-items: center;
  background-color: $search-box-color;
  color: $search-box-color;
  font-size: 14px;
  padding: 0 rpx(25);
  box-sizing: border-box;
  position: fixed;
  z-index: $search-index;
  top: v-bind('boxPos.top');
  left: v-bind('boxPos.lr');
  width: v-bind('boxPos.width');
  height: v-bind('boxPos.height');
}
</style>