<template>
  <image
      @click="clickThis"
      :src="props.src"
      :style="useStyle"
      :class="yClass"
      :mode="props.mode"
      show-menu-by-longpress
      :draggable="false"
      style="display: flex"
  ></image>
</template>

<script lang="ts" setup>

import {computed} from "vue";
import widthHeightAppendStyle from "@/components/y-ui/components/YImage/widthHeightAppendStyle";

const props = withDefaults(defineProps<{
  src: string,
  yStyle?: string,
  yClass?: string,
  mode?: "scaleToFill" | "aspectFit" | "aspectFill" | "widthFix" | "heightFix",
  width?: string | number,
  height?: string | number,
}>(), {
  src: "",
  yStyle: "max-width:100%;",
  yClass: "",
  mode: "scaleToFill",
  width: "",
  height: "",
})

const useStyle = computed(()=>{
  return props.yStyle + widthHeightAppendStyle(props.width,props.height);
})
const emit = defineEmits<{
  (e:"click"):void
}>()
const clickThis = ()=>{
  uni.$y.throttle(()=>{
    emit("click")
  })
}
</script>

<style lang="scss" scoped>

</style>