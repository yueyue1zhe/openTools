<template>
  <view>
    <view
        v-for="(item,key) in props.opts" :key="key"
        class="opt-form-item y-flex y-col-center"
    >
      <view class="y-font-md label y-text-justify y-flex">
        <view class="y-w100">{{ showLabel(item) }}</view>
      </view>
      <view class="y-flex-1 y-flex y-col-center">
        <view
            v-if="item.type === YOptFormItemTypeState.text.value"
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
            ></uni-easyinput>
          </view>
          <slot :name="item.name"></slot>
        </view>
        <view
            v-if="item.type === YOptFormItemTypeState.uploadImage.value"
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
  </view>
</template>

<script lang="ts" setup>
import UniEasyinput from "@/components/uni-ui/lib/uni-easyinput/uni-easyinput.vue";
import {YOptFormItemTypeState} from "@/components/y-ui/components/YOptForm/state";
import {computed, ref} from "vue";
import YImageUploadPreviewBox from "@/components/y-ui/components/YImage/YImageUploadPreviewBox.vue";


interface PropsType {
  modelValue: AnyObject;
  opts: YOptForm.OptsItemType[];
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
        reject("请输入" + showLabel(item))
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
  if (typeof props.labelWidth == "number") {
    return props.labelWidth + "rpx"
  }
  if (!isNaN(parseInt(props.labelWidth))) {
    return props.labelWidth + "rpx";
  }
  return props.labelWidth;
})

const showLabel = (item: YOptForm.OptsItemType): string => {
  return item.label ? item.label : item.name;
}
const showPlaceholder = (item: YOptForm.OptsItemType) => {
  if (!item.placeholder) {
    let out = "请输入";
    switch (item.type) {
      case YOptFormItemTypeState.text.value:
        out += showLabel(item)
        break;
      case YOptFormItemTypeState.uploadImage.value:
        out = "点击上传" + showLabel(item)
        break;
    }
    return out
  }
  return item.placeholder;
}


const itemRefs = ref<Array<any>>([]);

const setItemRefs = (el: HTMLElement, name: string) => {
  if (el) {
    itemRefs.value.push({
      name: name,
      el
    })
  }
}

const uploadImageAction = (item: YOptForm.OptsItemType) => {
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
.opt-form-item {
  .label {
    height: 70rpx;
    width: v-bind(useLabelWidth);
  }
}

.opt-form-item + .opt-form-item {
  border-top: 1px solid $y-border-color;
  padding-top: 10rpx;
  margin-top: 10rpx;
}
</style>