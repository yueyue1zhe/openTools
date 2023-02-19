<template>
  <view class="y-flex y-row-center y-col-center">
    <view class="image-slot y-flex y-row-center y-col-center">
      <image v-if="props.src" :src="showMediaSrc" class="image-slot"></image>
      <uni-icons v-else type="folder-add-filled" color="#909399" size="70rpx"></uni-icons>
    </view>
  </view>
</template>

<script lang="ts" setup>
import {computed} from "vue";
import UniIcons from "@/components/uni-ui/lib/uni-icons/uni-icons.vue";

const props = withDefaults(defineProps<{
  src: string;
  size?: number | string; //图片预览大小 默认200rpx 默认单位 rpx
  toMediaFunc?: ToMediaFunc;
}>(), {
  src: "",
  size: 200
})

const showMediaSrc = computed((): string => {
  return props.toMediaFunc ? props.toMediaFunc(props.src) : props.src;
})

const imgSize = computed(() => {
  return props.size + "rpx"
})
</script>

<style lang="scss" scoped>
.image-slot {
  background-color: #f8f8f8;
  width: v-bind(imgSize);
  height: v-bind(imgSize);
  border-radius: 10rpx;
  display: flex;
}
</style>