<template>
  <div v-if="showInput" class="w100">
    <div class="flex-def">
      <el-input
          :placeholder="props.placeholder"
          :value="modelValue"
          @change="(e:string)=>!e?out = '':''"
      ></el-input>
      <el-button class="y-m-l-10" @click="chooseStart">选择图片</el-button>
    </div>
    <el-image
        v-if="modelValue"
        :src="toMediaModelValue"
        :preview-src-list="[toMediaModelValue]"
        :hide-on-click-modal="true"
        fit="contain"
        class="image-slot"
    ></el-image>
    <div v-else class="image-slot">
      <y-icon name="el-icon-Picture" size="40"></y-icon>
    </div>
  </div>
  <el-dialog
      v-model="state.popupShow"
      destroy-on-close
      :show-close="false"
      width="60rem"
      class="custom-dialog"
  >
    <template #header>
      <div class="custom-dialog-header">
        <div class="y-section-title">选择图片</div>
        <el-upload
            ref="refElUploader"
            :action="elUploadActionUrl"
            :headers="elUploadHeaders"
            :data="elUploadData"
            name="file"
            :show-file-list="false"
            accept="image/jpg,image/jpeg,image/png,image/ico,image/svg,image/bmp,image/gif"
            :on-success="elUploadSuccess"
            :on-error="elUploadError"
            :on-progress="elUploadProgress"
            multiple
        >
          <el-button type="primary">点击上传</el-button>
        </el-upload>
      </div>
      <div v-show="state.progressShow" class="y-upload-progress">
        <el-progress
            :stroke-width="3"
            indeterminate
            stroke-linecap="square"
            :percentage="state.progressPercent"
            :show-text="false"
        />
      </div>
    </template>
    <el-container class="y-body">
      <image-body-aside
          ref="refAside"
          @tab-change="groupChange"
          @group-load="groupLoad"
      ></image-body-aside>
      <image-body-list
          ref="refBody"
          :group="state.group"
          :group_id="state.groupId"
          @choose="imageBodyListChoose"
      ></image-body-list>
    </el-container>
  </el-dialog>
</template>

<script lang="ts" setup>
import {GetBaseUrl} from "../../libs/request/common";
import {useUserStore} from "@/stores/user";
import config from "@/config";
import {computed, nextTick, reactive, ref} from "vue";
import YIcon from "@/components/yueyue-ui/components/yl-icon/YlIcon.vue";
import ImageBodyAside from "@/components/yueyue-ui/components/yl-file-upload/image-upload/imageBodyAside.vue";
import ImageBodyList from "@/components/yueyue-ui/components/yl-file-upload/image-upload/imageBodyList.vue";
import type {groupListItemType, imageListItemType,} from "@/components/yueyue-ui/components/yl-file-upload/types";
import type {UploadProps} from "element-plus";
import {ElMessage, ElUpload} from "element-plus";
import {ToMedia} from "@/components/yueyue-ui/libs/function/common";
import {YResponseDataType} from "@/components/yueyue-ui/libs/request/types";

const elUploadActionUrl = GetBaseUrl() + "/admin/attach/upload";
let userStore = useUserStore();
let elUploadHeaders: Record<string, string> = {};
elUploadHeaders[config.JWTTokenKey] = userStore.token.data;

let state = reactive<{
  popupShow: boolean;
  callback?: VoidCallBack<string>;
  groupId: number;
  group: groupListItemType[];
  progressShow: boolean;
  progressPercent: number;
}>({
  popupShow: false,
  groupId: -1,
  group: [],
  progressShow: false,
  progressPercent: 0,
});

const elUploadData = computed((): Record<string, number> => {
  return {
    group_id: state.groupId >= 0 ? state.groupId : 0,
    project_id: 0,
  };
});

const toMediaModelValue = computed((): string => {
  return ToMedia(props.modelValue);
});

const imageBodyListChoose = (item: imageListItemType) => {
  let outVal = item.attachment;
  if (props.full) {
    outVal = ToMedia(item.attachment);
  }
  state.popupShow = false;

  out.value = outVal;
  if (typeof state?.callback == "function") {
    state?.callback(outVal);
  }
};

const chooseStart = (cb?: VoidCallBack<string>) => {
  state.popupShow = true;
  if (cb) state.callback = cb;
  nextTick(() => {
    refAside.value?.openAside(-1);
    refBody.value?.openBody();
  });
};

defineExpose({chooseStart});

interface YUploadImageProps {
  full?: boolean;
  modelValue?: string;
  showInput?: boolean;
  placeholder?: string; //仅showInput为true时有效
}

const props = withDefaults(defineProps<YUploadImageProps>(), {
  full: false,
  modelValue: "",
  showInput: true,
});

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

const refElUploader = ref<InstanceType<typeof ElUpload> | null>(null);
const elUploadError = (error: Error) => {
  console.warn(error);
};
const elUploadSuccess = (res: YResponseDataType<imageListItemType>) => {
  if (res.message) ElMessage.error(res.message);
  refElUploader.value?.clearFiles();
  state.progressShow = false;
  state.progressPercent = 0;
  refBody.value?.attachLoad();
};
const elUploadProgress: UploadProps["onProgress"] = (evt) => {
  state.progressShow = true;
  state.progressPercent = evt.percent;
};

const refAside = ref<InstanceType<typeof ImageBodyAside> | null>(null);
const refBody = ref<InstanceType<typeof ImageBodyList> | null>(null);

const groupChange = (index: number) => {
  state.groupId = index;
};
const groupLoad = (group: groupListItemType[]) => (state.group = group);
</script>

<style>
.custom-dialog .el-dialog__body {
  padding: 0;
}

.custom-dialog .el-dialog__header {
  padding: 10px 10px 10px 20px;
  margin-right: 0;
}
</style>
<style lang="scss" scoped>
.y-upload-progress {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
}

.y-body {
  height: 33rem;
  border: 1px solid var(--color-sub-3);
}

.y-section-title {
  margin: 0;
}

.custom-dialog-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: relative;
}

.image-slot {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 8rem;
  height: 8rem;
  background: #f5f7fa;
  color: #909399;
  font-size: 30px;
  max-width: 8rem;
  margin-top: 0.5rem;
  border-radius: 5px;
  border: 1px solid var(--color-sub-3);
}
</style>
