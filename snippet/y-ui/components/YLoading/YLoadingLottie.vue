<template>
  <view class="page-loading-box y-flex y-flex-d-column">
    <canvas class="page-loading" canvas-id="page-loading" id="page-loading" type="2d"></canvas>
  </view>
</template>

<script lang="ts" setup>
import {onMounted} from "vue";
import service from "@/common/service";

onMounted(() => {
  uni.$y.lottie.load("#page-loading", service.toStatic("/lottie/page-loading.json"), {
    loop: true,
    autoplay: true,
  }).then(loader => {
    loader.play()
  })
})
</script>

<style lang="scss" scoped>
.page-loading-box {
  position: fixed;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  background-color: $y-bg-color;
  z-index: $y-page-loading-index;

  .page-loading {
    position: relative;
    top: 30vh;
  }
}
</style>