<template>
  <view>
    <uni-popup ref="refUniPopup" type="bottom" @change="popupChange">
      <view class="color-picker-box" @touchmove="preventDefault">
        <view class="head y-flex y-col-center">
          <view class="side left">
            <uni-icons @click="state.popupShow = false" type="closeempty" size="30"></uni-icons>
          </view>
          <view class="side right">
            <uni-icons @click="state.popupShow = false" type="checkmarkempty" size="30"></uni-icons>
          </view>
          <view class="main y-flex-1">
            <view class="title">请选择颜色</view>
            <view class="y-font-xs y-flex y-row-center y-col-center y-tips-color">
              已选颜色: <view class="color-preview"></view> {{ state.colorRes }}
            </view>
          </view>
        </view>
        <movable-area id="target" class="target" :style="targetStyle">
          <view class="white"></view>
          <view class="black"></view>
          <movable-view
              direction="all"
              @change="changeSV"
              @touchend="onEnd"
              :x="state.x"
              :y="state.y"
              :animation="false"
              class="aimBox y-flex y-row-center y-col-center"
          >
            <view class="dot"></view>
            <view class="line left"></view>
            <view class="line top"></view>
            <view class="line right"></view>
            <view class="line bottom"></view>
          </movable-view>
        </movable-area>
        <slider activeColor="transparent" backgroundColor="transparent" class="ribbon"
                max="360" :value="state.hsv.h" :block-color="state.colorRes"
                @changing="changeHue"
                @change="changeHue"
                @touchend="onEnd"/>
      </view>
    </uni-popup>
  </view>
</template>

<script lang="ts" setup>
import {computed, getCurrentInstance, nextTick, onMounted, reactive, ref, watch} from "vue";
import UniPopup from "@/components/uni-ui/lib/uni-popup/uni-popup.vue";
import UniIcons from "@/components/uni-ui/lib/uni-icons/uni-icons.vue";

const props = withDefaults(defineProps<{
  initColor: string;
  modelValue: boolean;
}>(), {
  initColor: "rgb(255,0,0)",
  modelValue: false,
})
let state = reactive({
  hueColor: "",
  SV: {
    W: 0,
    H: 0,
    Step: 0,
  },
  hsv: {
    h: 0,
    s: 0,
    v: 0,
  },
  x: 0,
  y: 0,
  colorRes: "",

  popupShow:false,
})
watch(
    () => props.modelValue,
    () => {
        state.popupShow = props.modelValue;
    },
    {
      immediate:true,
      deep:true
    }
)
watch(()=>state.popupShow,()=>{
  if (state.popupShow) {
    refUniPopup.value?.open();
  } else {
    refUniPopup.value?.close();
  }
  emit("update:modelValue",state.popupShow)
})
const targetStyle = computed(() => {
  return `background-color:${state.hueColor}`;
})
const refUniPopup = ref();
const popupChange = (e: { show: boolean }) => {
  state.popupShow = e.show;
  if (e.show) {
    nextTick(() => {
      initPicker();
    })
  }
}
const thisApp = getCurrentInstance()
const initPicker = () => {
  const query = uni.createSelectorQuery().in(thisApp);
  query.select("#target").boundingClientRect(res => {
    const rect = res as NodeInfo;
    if (rect) {
      state.SV = {
        W: rect.width - 28, //block-size=28
        H: rect.height - 28,
        Step: (rect.width - 28) / 100
      }
      let {h, s, v} = rgb2hsv(props.initColor)
      // 初始化定位
      state.hsv.h = h
      state.hsv.s = s
      state.hsv.v = v
      state.x = Math.round(s * state.SV.Step)
      state.y = Math.round((100 - v) * state.SV.Step)
    }
  }).exec()
}
onMounted(() => {
  state.hueColor = hsv2rgb((rgb2hsv(props.initColor)).h, 100, 100)
})

const emit = defineEmits<{
  (e: "change", color: string): void
  (e: "update:modelValue", out: boolean): void;
}>()
const onEnd = () => {
  emit("change", state.colorRes);
}

const changeHue = (e: { detail: { value: number; }; }) => {
  let hue = e.detail.value;
  state.hsv.h = hue;
  state.hueColor = hsv2rgb(hue, 100, 100);
  state.colorRes = hsv2rgb(hue, state.hsv.s, state.hsv.v)
}

const changeSV = (e: { detail: { x: number; y: number; }; }) => {
  let {
    x,
    y
  } = e.detail;
  x = Math.round(x / state.SV.Step);
  y = 100 - Math.round(y / state.SV.Step);
  state.hsv.s = x;
  state.hsv.v = y;
  state.colorRes = hsv2rgb(state.hsv.h, x, y);
}

const preventDefault = () => {

}

const hsv2rgb = (h: number, s: number, v: number) => {
  let hsv_h = parseFloat((h / 360).toFixed(2));
  let hsv_s = parseFloat((s / 100).toFixed(2));
  let hsv_v = parseFloat((v / 100).toFixed(2));

  let i = Math.floor(hsv_h * 6);
  let f = hsv_h * 6 - i;
  let p = hsv_v * (1 - hsv_s);
  let q = hsv_v * (1 - f * hsv_s);
  let t = hsv_v * (1 - (1 - f) * hsv_s);

  let rgb_r = 0,
      rgb_g = 0,
      rgb_b = 0;
  switch (i % 6) {
    case 0:
      rgb_r = hsv_v;
      rgb_g = t;
      rgb_b = p;
      break;
    case 1:
      rgb_r = q;
      rgb_g = hsv_v;
      rgb_b = p;
      break;
    case 2:
      rgb_r = p;
      rgb_g = hsv_v;
      rgb_b = t;
      break;
    case 3:
      rgb_r = p;
      rgb_g = q;
      rgb_b = hsv_v;
      break;
    case 4:
      rgb_r = t;
      rgb_g = p;
      rgb_b = hsv_v;
      break;
    case 5:
      rgb_r = hsv_v;
      rgb_g = p;
      rgb_b = q;
      break;
  }

  return 'rgb(' + (Math.floor(rgb_r * 255) + "," + Math.floor(rgb_g * 255) + "," + Math.floor(rgb_b * 255)) + ')';
}

const rgb2hsv = (color: string) => {
  let rgb = color.split(',');
  let R = parseInt(rgb[0].split('(')[1]);
  let G = parseInt(rgb[1]);
  let B = parseInt(rgb[2].split(')')[0]);

  let hsv_red = R / 255, hsv_green = G / 255, hsv_blue = B / 255;
  let hsv_max = Math.max(hsv_red, hsv_green, hsv_blue),
      hsv_min = Math.min(hsv_red, hsv_green, hsv_blue);
  let hsv_h = hsv_max;
  let hsv_s = hsv_max;
  let hsv_v = hsv_max;
  let hsv_d = hsv_max - hsv_min;
  hsv_s = hsv_max == 0 ? 0 : hsv_d / hsv_max;

  if (hsv_max == hsv_min) hsv_h = 0;
  else {
    switch (hsv_max) {
      case hsv_red:
        hsv_h = (hsv_green - hsv_blue) / hsv_d + (hsv_green < hsv_blue ? 6 : 0);
        break;
      case hsv_green:
        hsv_h = (hsv_blue - hsv_red) / hsv_d + 2;
        break;
      case hsv_blue:
        hsv_h = (hsv_red - hsv_green) / hsv_d + 4;
        break;
    }
    hsv_h /= 6;
  }
  return {
    h: parseFloat((hsv_h * 360).toFixed()),
    s: parseFloat((hsv_s * 100).toFixed()),
    v: parseFloat((hsv_v * 100).toFixed())
  }
}
</script>

<style lang="scss" scoped>

.target {
  height: 600rpx;
  width: 600rpx;
  margin: 0 auto;
  overflow: hidden;
  border: 0.5px solid rgba(0, 0, 0, 0.5);
  position: relative;

  .white, .black {
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
  }

  .white {
    background-image: linear-gradient(to right, #fff, rgba(255, 255, 255, 0));
  }

  .black {
    background-image: linear-gradient(to top, #000, rgba(0, 0, 0, 0));
  }
}


.ribbon {
  background: -webkit-linear-gradient(left, #f00 0%, #ff0 17%, #0f0 33%, #0ff 50%, #00f 67%, #f0f 83%, #f00 100%);
  width: 600rpx;
  margin: 40rpx auto;
}

.color-picker-box {
  width: 100%;
  background-color: #fff;
  border-radius: 25rpx 25rpx 0 0;
  overflow: hidden;

  .head {
    font-size: $y-font-xs;
    height: 120rpx;
    position: relative;

    .side{
      position: absolute;
    }
    .side.left {
      left: $y-font-lg;
    }
    .side.right{
      right: $y-font-lg;
    }

    .main {
      text-align: center;

      .title {
        font-weight: bold;
        font-size: $y-font-xl;
      }
    }
  }
}


.aimBox {
  width: 50rpx;
  height: 50rpx;
  border-radius: 60rpx;
  border: 2px solid #FFFFFF;
  position: relative;

  .dot {
    background-color: #FFFFFF;
    height: 15rpx;
    width: 15rpx;
    border-radius: 40rpx;
  }

  .line {
    background-color: #FFFFFF;
    position: absolute;
  }

  .line.left {
    height: 2px;
    width: 20rpx;
    left: -10rpx;
  }

  .line.right {
    height: 2px;
    width: 20rpx;
    right: -10rpx;
  }

  .line.top, .line.bottom {
    width: 2px;
    height: 20rpx;
  }

  .line.top {
    top: -10rpx;
  }

  .line.bottom {
    bottom: -10rpx;
  }
}

.color-preview{
  height: $y-font-md;
  width: $y-font-md;
  background-color: v-bind("state.colorRes");
  margin: 0 10rpx;
}
</style>