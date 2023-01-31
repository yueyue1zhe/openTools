<template>
  <view>
    <uni-swiper-dot :current="swiperCurrent" mode="round" :info="showList" :dotsStyles="dotStyle">
      <swiper @change="swiperCurrentChange" :style="firstImgHeight" autoplay :interval="interval"
              :duration="500" circular>
        <swiper-item v-for="(item,key) in showList" :key="key">
          <image @click="clickThis(key)" :style="firstImgHeight" :src="item" class="image"></image>
        </swiper-item>
      </swiper>
    </uni-swiper-dot>
  </view>
</template>

<script setup lang="ts">
import {computed, ref} from "vue";
import UniSwiperDot from "@/components/uni-ui/lib/uni-swiper-dot/uni-swiper-dot.vue";


const props = withDefaults(defineProps<{
  list: string[],
  interval?: number,
  dot?: boolean,
  dotBottom?: number,
}>(), {
  list: () => [],
  interval: 3000,
  dot: true,
  dotBottom: 10,
})
let firstImgHeight = ref('height:0');
let swiperCurrent = ref(0);
let dotStyle = {
  bottom: props.dotBottom,
  selectedBackgroundColor: "#ffffff",
  border: 'none',
  selectedBorder: 'none'
}

function swiperCurrentChange(e: CustomEvent) {
  swiperCurrent.value = e.detail.current
}

const showList = computed(() => {
  if (!props.list.length) return [];
  let sysInfo = uni.getSystemInfoSync();
  uni.getImageInfo({
    src: props.list[0],
    success: (e) => {
      console.log(e);
      firstImgHeight.value = `height:${((sysInfo.windowWidth * 2 / e.width) * e.height)}rpx`;
    }
  })
  return props.list
})

const emit = defineEmits<{
  (e: "click", key: number): void
}>();

const clickThis = (key :number)=>{
  emit("click",key);
}

</script>

<style lang="scss" scoped>
.image {
  width: 100%;
}
</style>