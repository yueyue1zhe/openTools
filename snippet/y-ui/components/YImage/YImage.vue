<template>
  <image
      :src="props.src"
      :style="props.yStyle"
      :class="yClass"
      :mode="props.mode"
      show-menu-by-longpress
      :draggable="false"
      style="display: flex"
  ></image>
</template>

<script lang="ts" setup>

const props = withDefaults(defineProps<{
  src: string,
  yStyle: string,
  yClass: string,
  mode?: "scaleToFill" | "aspectFit" | "aspectFill" | "widthFix" | "heightFix",
}>(), {
  src: "",
  yStyle: "max-width:100%;",
  yClass: "",
  mode: "scaleToFill",
})
</script>

<style lang="scss" scoped>

</style>