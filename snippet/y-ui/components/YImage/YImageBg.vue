<template>
  <view @click="emit('click')" class="static-bg" :class="props.yClass" :style="useStyle">
    <slot></slot>
  </view>
</template>

<script lang="ts" setup>
import {computed} from "vue";

const props = withDefaults(defineProps<{
  src: string;
  width?: string | number;
  height?: string | number;
  minHeight?: string | number;
  yStyle?: string;
  yClass?: string;
}>(), {
  src: "",
  width: "100%",
  height: "",
  minHeight: "200rpx",
  yStyle: "",
  yClass: "",
})
const safeSize = (val: number | string) => {
  if (typeof val === "number") {
    return val + "rpx";
  }
  if (!isNaN(parseInt(val))) {
    val += "rpx";
  }
  return val;
}

const useSize = computed(() => {
  let outStyle = ";";
  let useWidth = safeSize(props.width);
  let useHeight = safeSize(props.height);
  if (!useWidth && !useHeight) return outStyle;
  outStyle += `height:${useHeight ? useHeight : useWidth};`;
  outStyle += `width:${useWidth ? useWidth : useHeight};`;
  const useMinHeight = safeSize(props.minHeight);
  if (useMinHeight) outStyle += `min-height:${useMinHeight};`
  return outStyle;
})

const useStyle = computed((): string => {
  let outStyle = props.yStyle + useSize.value;
  return outStyle + `background-image:url(${props.src})`;
})

const emit = defineEmits<{
  (e: "click"): void
}>()
</script>

<style lang="scss" scoped>
.static-bg {
  background-size: 100% 100%;
  background-repeat: no-repeat;
}
</style>