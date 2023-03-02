<template>
  <view @click="emit('click')" class="static-bg" :class="props.yClass" :style="useStyle">
    <slot></slot>
  </view>
</template>

<script lang="ts" setup>
import {computed} from "vue";
import widthHeightAppendStyle from "@/components/y-ui/components/YImage/widthHeightAppendStyle";

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
  minHeight: "",
  yStyle: "",
  yClass: "",
})

const useSize = computed(() => {
  let outStyle = ";";
  const useMinHeight = uni.$y.addUnit(props.minHeight);
  if (useMinHeight) outStyle += `min-height:${useMinHeight};`
  return outStyle + widthHeightAppendStyle(props.width, props.height);
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