<template>
  <view
      v-for="(item,key) in useOpts" :key="key"
      class="y-easy-form-item y-flex y-col-center"
  >
    <view class="y-font-md label y-text-justify y-flex">
      <view class="y-w100">{{ showLabel(item) }}</view>
    </view>
    <view class="y-flex-1 y-flex y-col-center">
      <view
          v-if="item.type === YEasyFormItemTypeState.switch.value"
          class="y-w100 y-padding-left-20 y-flex y-row-between y-col-center"
      >
        <view class="y-tips-color">{{ showPlaceholder(item) }}</view>
        <y-switch v-model="out[item.name]"></y-switch>
      </view>
      <view
          v-if="item.type === YEasyFormItemTypeState.radio.value"
          class="y-w100 y-padding-left-20"
      >
        <uni-data-checklist
            mode="button"
            v-model="out[item.name]"
            :localdata="item.radioOpts"
            selected-color="#07c160"
        ></uni-data-checklist>
      </view>
      <view
          v-if="item.type === YEasyFormItemTypeState.text.value"
          class="y-flex y-col-center y-w100"
      >
        <view class="y-flex-1">
          <uni-easyinput
              v-model="out[item.name]"
              :input-border="false"
              primary-color="#909399"
              :placeholder="showPlaceholder(item)"
              placeholder-style="font-size: 28rpx;padding-top: 3rpx;"
              :trim="item.required"
              :styles="{
                color: '#333',
                backgroundColor: 'rgba(0,0,0,0)',
                disableColor: '#F7F6F6',
                borderColor: '#e5e5e5'
              }"
          ></uni-easyinput>
        </view>
        <slot :name="item.name"></slot>
      </view>
      <view
          v-if="item.type === YEasyFormItemTypeState.uploadImage.value"
          class="y-flex y-row-between y-w100 y-padding-left-20"
          @click="uploadImageAction(item)"
      >
        <view class="y-tips-color">{{ showPlaceholder(item) }}</view>
        <y-image-upload-preview-box
            :src="out[item.name]"
            :to-media-func="item.uploadImageOption.toMediaFunc"
            :size="80"
        ></y-image-upload-preview-box>
      </view>
    </view>
  </view>
</template>
<script lang="ts" setup>
import {computed} from "vue";
import {YEasyFormItemTypeState} from "@/components/y-ui/components/YEasyForm/state";
import UniEasyinput from "@/components/uni-ui/lib/uni-easyinput/uni-easyinput.vue";
import YImageUploadPreviewBox from "@/components/y-ui/components/YImage/YImageUploadPreviewBox.vue";
import UniDataChecklist from "@/components/uni-ui/lib/uni-data-checkbox/uni-data-checkbox.vue";
import YSwitch from "@/components/y-ui/components/YSwitch/YSwitch.vue";

interface PropsType {
  modelValue: AnyObject;
  opts: YEasyFormTypes.OptsItemType[];
  labelWidth?: string | number;
}

const props = withDefaults(defineProps<PropsType>(), {
  modelValue: () => {
    return {}
  },
  opts: () => {
    return []
  },
  labelWidth: "100rpx"
})
const useOpts = computed(() => {
  let useOpts :YEasyFormTypes.OptsItemType[] = [];
  props.opts.forEach(item=>{
    if (!item.defaultHide || (item?.showCond && item?.showCond(out.value))){
      useOpts.push(item)
    }
  })
  return useOpts;
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
const ActionCheck = () => {
  return new Promise<void>((resolve, reject: (msg: string) => void) => {
    props.opts.forEach(item => {
      if (item.required && !out.value[item.name]) {
        reject(showPlaceholder(item))
        return
      }
    })
    resolve();
  })
}
defineExpose({
  ActionCheck
})
const useLabelWidth = computed((): string => {
  return uni.$y.addUnit(props.labelWidth);
})
const showLabel = (item: YEasyFormTypes.OptsItemType): string => {
  return item.label ? item.label : item.name;
}
const showPlaceholder = (item: YEasyFormTypes.OptsItemType) => {
  if (!item.placeholder) {
    let out = "请输入";
    switch (item.type) {
      case YEasyFormItemTypeState.text.value:
        out += showLabel(item)
        break;
      case YEasyFormItemTypeState.uploadImage.value:
        out = "点击上传" + showLabel(item)
        break;
      case YEasyFormItemTypeState.switch.value:
        out = "开关" + showLabel(item)
        break;
    }
    return out
  }
  return item.placeholder;
}
const uploadImageAction = (item: YEasyFormTypes.OptsItemType) => {
  uni.showLoading({
    title: "请稍候...",
    mask: true
  })
  item.uploadImageOption?.actionFunc().then((res) => {
    out.value[item.name] = res.attachment
  }).finally(() => {
    uni.hideLoading();
  })
}
</script>
<style lang="scss" scoped>

.y-easy-form-item {
  .label {
    height: 80rpx;
    width: v-bind(useLabelWidth);
    color: $y-content-color;
  }
}

.y-easy-form-item + .y-easy-form-item {
  border-top: 1px solid $y-border-color;
  padding-top: 10rpx;
  margin-top: 10rpx;
}
</style>