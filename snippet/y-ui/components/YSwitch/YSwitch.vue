<template>
  <view class="y-switch" :class="[outValue ? 'y-switch--on' : '', disabled ? 'y-switch--disabled' : '']"
        @click="onClick"
        :style="[switchStyle]">
    <view class="y-switch__node node-class" :style="switchNodeStyle">
      <y-loading :show="loading" class="y-switch__loading" :size="size * 0.6" :color="loadingColor"></y-loading>
    </view>
  </view>
</template>
<script lang="ts" setup>
import {computed} from "vue";
import YLoading from "@/components/y-ui/components/YLoading/YLoading.vue";

const props = withDefaults(defineProps<{
  modelValue: boolean;
  size?: number | string;
  loading?: boolean;
  disabled?: boolean;
  activeColor?: string;
  inactiveColor?: string;
  vibrateShort?: boolean;
  beforeChanged?:(value:boolean)=>Promise<boolean>
}>(), {
  modelValue: false,
  size: "50",
  loading: false,
  disabled: false,
  activeColor: "#07c160",
  inactiveColor: "#ffffff",
  vibrateShort: false,
})
const emit = defineEmits<{
  (e: "update:modelValue", val: boolean): void
}>()
const outValue = computed({
  get: () => props.modelValue,
  set: (val) => emit("update:modelValue", val)
})

const switchNodeStyle = computed(() => {
  return {
    width: uni.$y.addUnit(props.size),
    height: uni.$y.addUnit(props.size),
  }
})
const switchStyle = computed(() => {
  return {
    fontSize: uni.$y.addUnit(props.size),
    backgroundColor: outValue.value ? props.activeColor : props.inactiveColor
  }
})
const loadingColor = computed(() => {
  return outValue.value ? props.activeColor : null;
})
const onClick = () => {
  if (!props.disabled && !props.loading) {
    if (props.vibrateShort) uni.vibrateShort({});
    if (props.beforeChanged){
      props.beforeChanged(!outValue.value).then(res=>{
        outValue.value = res;
      })
    }else {
      outValue.value = !outValue.value;
    }
  }
}
</script>

<style lang="scss" scoped>
@import "../../libs/scss/style.components";

.y-switch {
  position: relative;
  /* #ifndef APP-NVUE */
  display: inline-block;
  /* #endif */
  box-sizing: initial;
  width: 2em;
  height: 1em;
  background-color: #fff;
  border: 1px solid rgba(0, 0, 0, 0.1);
  border-radius: 1em;
  transition: background-color 0.3s;
  font-size: 50rpx;
}

.y-switch__node {
  @include vue-flex;
  align-items: center;
  justify-content: center;
  position: absolute;
  top: 0;
  left: 0;
  border-radius: 100%;
  z-index: 1;
  background-color: #fff;
  background-color: #fff;
  box-shadow: 0 3px 1px 0 rgba(0, 0, 0, 0.05), 0 2px 2px 0 rgba(0, 0, 0, 0.1), 0 3px 3px 0 rgba(0, 0, 0, 0.05);
  box-shadow: 0 3px 1px 0 rgba(0, 0, 0, 0.05), 0 2px 2px 0 rgba(0, 0, 0, 0.1), 0 3px 3px 0 rgba(0, 0, 0, 0.05);
  transition: transform 0.3s cubic-bezier(0.3, 1.05, 0.4, 1.05);
  transition: transform 0.3s cubic-bezier(0.3, 1.05, 0.4, 1.05), -webkit-transform 0.3s cubic-bezier(0.3, 1.05, 0.4, 1.05);
  transition: transform cubic-bezier(0.3, 1.05, 0.4, 1.05);
  transition: transform 0.3s cubic-bezier(0.3, 1.05, 0.4, 1.05)
}

.y-switch__loading {
  @include vue-flex;
  align-items: center;
  justify-content: center;
}

.y-switch--on {
  background-color: #07c160;
}

.y-switch--on .y-switch__node {
  transform: translateX(100%);
}

.y-switch--disabled {
  opacity: 0.4;
}
</style>
