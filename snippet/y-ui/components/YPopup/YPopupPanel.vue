<template>
  <uni-popup ref="refPanelPopup" @change="popupChange" type="center">
    <view class="panel-box y-flex-d-column y-flex y-row-between">
      <view @click="wantClose" class="close-box">
        <uni-icons type="closeempty" size="20" color="#909399"></uni-icons>
      </view>
      <view class="y-w100">
        <view v-if="props.title" class="title">{{ props.title }}</view>
        <slot></slot>
      </view>
      <view
          v-if="props.btnShow"
          @click="clickThis"
          class="copy-btn y-default-shadow y-w100"
      >{{ props.btnText }}</view>
    </view>
  </uni-popup>
</template>

<script lang="ts" setup>

import UniPopup from "@/components/uni-ui/lib/uni-popup/uni-popup.vue";
import {computed, ref, watch} from "vue";
import UniIcons from "@/components/uni-ui/lib/uni-icons/uni-icons.vue";

const props = withDefaults(defineProps<{
  modelValue: boolean;
  title?: string;
  btnShow?: boolean;
  btnText?: string;
  width?: string | number;
}>(), {
  modelValue: false,
  title: "",
  btnShow: false,
  btnText: "",
  width: "700",
})
const useWidth = computed(()=>{
  return uni.$y.addUnit(props.width);
})
const refPanelPopup = ref();
const emit = defineEmits<{
  (e: "btn-click",close:YCallBack): void;
  (e: "update:modelValue", show: boolean): void;
}>()
watch(
    () => props.modelValue,
    (val) => {
      if (val) {
        refPanelPopup.value?.open();
      } else {
        refPanelPopup.value?.close();
      }
    },
    {
      immediate: true,
      deep: true
    }
)
const popupChange = (e: { show: boolean }) => {
  emit("update:modelValue", e.show);
}
const wantClose = ()=>{
  refPanelPopup.value?.close();
}
const clickThis = () => {
  uni.$y.throttle(() => {
    emit("btn-click",wantClose);
  })
}
</script>

<style lang="scss" scoped>
.panel-box {
  background-color: #FFFFFF;
  width: v-bind(useWidth);
  min-height: 300rpx;
  position: relative;
  border-radius: 25rpx;
  padding: 25rpx;

  .close-box {
    position: absolute;
    right: 15rpx;
    top: 15rpx;
  }

  .title {
    font-size: $y-font-xl;
    color: $y-content-color;
    font-weight: bold;
    text-align: center;
    padding-bottom: 25rpx;
    border-bottom: 1px solid $y-border-color;
  }

  .copy-btn {
    text-align: center;
    line-height: 80rpx;
    border-radius: 80rpx;
    background-color: $y-main-theme-color;
    font-weight: bold;
    color: #FFFFFF;
    margin-top: 25rpx;
  }
}
</style>