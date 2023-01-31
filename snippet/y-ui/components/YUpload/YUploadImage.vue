<template>
  <view class="y-flex y-row-center y-col-center">
    <view @click="uploadImage" class="image-slot y-flex y-row-center y-col-center">
      <image v-if="out" :src="toMediaModelValue" class="image-slot"></image>
      <uni-icons v-else type="folder-add-filled" color="#909399" size="70rpx"></uni-icons>
    </view>
  </view>
</template>

<script lang="ts" setup>
import {computed} from "vue";
import UniIcons from "@/components/uni-ui/lib/uni-icons/uni-icons.vue";
import tools from "@/common/api/tools";
import utils from "@/common/utils";

const props = withDefaults(defineProps<{
  modelValue: string;
  size?: number | string //图片预览大小 默认200rpx 默认单位 rpx
}>(), {
  modelValue: "",
  size: 200
})
const emit = defineEmits<{
  (e: "update:modelValue", out: string): void;
}>();
const out = computed({
  get(): string {
    return props.modelValue;
  },
  set(value: string) {
    emit("update:modelValue", value);
  },
});

const toMediaModelValue = computed(():string => {
  return utils.toMedia(out.value);
});


const uploadImage = () => {
  uni.chooseImage({
    count: 1,
    sizeType: ["compressed"],
    success: (res) => {
      uni.uploadFile({
        url: tools.apiBase() + "/attach/upload-image",
        filePath: res.tempFilePaths[0],
        success: (uploadRes) => {
          let result = tools.apiPostResParse<AttachResult>(JSON.parse(uploadRes.data))
          if (typeof result != "boolean") {
            out.value = result.attachment
          }
        }
      })
    }
  })
}

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