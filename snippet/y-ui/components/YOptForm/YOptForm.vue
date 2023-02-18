<template>
  <view>
    <view
        v-for="(item,key) in state.opts" :key="key"
        class="opt-form-item y-flex y-col-center"
    >
      <view class="y-font-md y-flex y-col-center label">
        {{ showLabel(item) }}
      </view>
      <view class="y-flex-1 y-flex y-col-center">
        <uni-easyinput
            v-if="item.type === YOptFormItemTypeState.text.value"
            v-model="out[item.name]"
            :input-border="false"
            primary-color="#909399"
            :placeholder="showPlaceholder(item)"
            placeholder-style="font-size: 28rpx;padding-top: 3rpx"
            :trim="item.required"
        ></uni-easyinput>
      </view>
    </view>
  </view>
</template>

<script lang="ts" setup>
import UniEasyinput from "@/components/uni-ui/lib/uni-easyinput/uni-easyinput.vue";
import {YOptFormItemTypeState} from "@/components/y-ui/components/YOptForm/state";
import {computed, reactive} from "vue";

interface AnyObject {
  [key: string]: any;
}

const props = withDefaults(defineProps<{
  modelValue: AnyObject
}>(), {
  modelValue: () => {
    return {}
  }
})
const emit = defineEmits<{
  (e: "update:modelValue", out: AnyObject): void;
}>();
const out = computed({
  get(): AnyObject {
    return props.modelValue;
  },
  set(value: AnyObject) {
    emit("update:modelValue", value);
  },
});

let state = reactive<{
  opts: YOptFormItemType[]
}>({
  opts: [],
})
const ActionSetOpts = (opts: YOptFormItemType[]) => {
  state.opts = opts;
}
const ActionCheck = ()=>{
  return new Promise<void>((resolve, reject)=>{
    state.opts.forEach(item=>{
      if (item.required && !out.value[item.name]){
        reject("请设置" + showLabel(item))
        return
      }
    })
    resolve();
  })
}

defineExpose({
  ActionSetOpts,
  ActionCheck
})

const showLabel = (item: YOptFormItemType): string => {
  return item.label ? item.label : item.name;
}
const showPlaceholder = (item: YOptFormItemType) => {
  if (!item.placeholder) return `请输入${item.label}`;
  return item.placeholder;
}
</script>

<style lang="scss" scoped>
.opt-form-item {
  .label {
    height: 70rpx;
  }
}

.opt-form-item + .opt-form-item {
  border-top: 1px solid $y-border-color;
  padding-top: 10rpx;
  margin-top: 10rpx;
}
</style>