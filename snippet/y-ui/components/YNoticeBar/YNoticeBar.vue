<template>
  <swiper :style="swiperHeight" class="w100" vertical circular autoplay>
    <template v-for="(item,key) in props.list" :key="key">
      <swiper-item>
        <view class="y-line-1">{{ item }}</view>
      </swiper-item>
    </template>
  </swiper>
</template>

<script lang="ts" setup>
import {computed} from "vue";

const props = withDefaults(defineProps<{
  list: string[];
  fontSize?: number;
  bgColor?: string;
}>(), {
  list: () => {
    return []
  },
  fontSize: 28,
  bgColor: "#ffffff",
})
const swiperHeight = computed(() => {
  return `background-color:${props.bgColor};line-height:${props.fontSize * 1.5}rpx;height:${props.fontSize * 1.5}rpx;font-size:${props.fontSize}rpx`;
})
</script>

<style lang="scss" scoped>

</style>