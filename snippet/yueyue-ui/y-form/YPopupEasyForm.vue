<template>
  <div>
    <el-dialog
      v-model="state.editBoxShow"
      :width="props.width"
      :title="props.title"
      @close="editBoxClosed"
    >
      <el-form :label-width="props.labelWidth" label-position="left">
        <template v-for="(item, key) in useOpts" :key="key">
          <el-form-item :label="showLabel(item)">
            <div
              v-if="item.type === YPopupEasyFormItemTypeState.powerSort.value"
              class="flex-def flex-zBetween w100"
            >
              <el-switch
                v-model="out[item.powerSortOpts.powerName]"
              ></el-switch>
              <div>
                <span style="color: #606266; margin-right: 12px">排序</span>
                <el-input-number
                  v-model="out[item.powerSortOpts.sortName]"
                ></el-input-number>
              </div>
            </div>
            <template
              v-if="item.type === YPopupEasyFormItemTypeState.text.value"
            >
              <easy-form-input
                :opt="item"
                v-model="out"
                :placeholder="showPlaceholder(item)"
              ></easy-form-input>
            </template>
            <template
              v-if="item.type === YPopupEasyFormItemTypeState.number.value"
            >
              <el-input-number v-model="out[item.name]"></el-input-number>
            </template>
            <template
              v-if="item.type === YPopupEasyFormItemTypeState.uploadImage.value"
            >
              <y-upload-image v-model="out[item.name]"></y-upload-image>
            </template>
            <template
              v-if="item.type === YPopupEasyFormItemTypeState.switch.value"
            >
              <el-switch v-model="out[item.name]"></el-switch>
              <span v-if="item.placeholder" class="y-desc y-m-l-5">{{
                item.placeholder
              }}</span>
            </template>
            <template
              v-if="item.type === YPopupEasyFormItemTypeState.radio.value"
            >
              <el-radio-group v-model="out[item.name]">
                <el-radio
                  v-for="(v, k) in item.radioOpts"
                  :key="k"
                  :label="v.value"
                  >{{ v.label }}</el-radio
                >
              </el-radio-group>
            </template>
          </el-form-item>
        </template>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="state.editBoxShow = false">取消</el-button>
          <el-button type="primary" @click="editFormSubmit">提交</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { computed, reactive } from "vue";
import { YPopupEasyFormItemTypeState } from "@/components/yueyue-ui/y-form/yPopupEasyFormState";
import YUploadImage from "@/components/yueyue-ui/y-file-upload/YUploadImage.vue";
import { ElMessage } from "element-plus";
import EasyFormInput from "@/components/yueyue-ui/y-form/easyFormInput.vue";
import { easyFormGetProperty } from "@/components/yueyue-ui/y-form/easyFormTools";

interface propsTypes {
  title?: string;
  width?: string;
  labelWidth?: string;
  modelValue: AnyObject;
  opts: YPopupEasyFormTypes.OptsItemType[];
}

const useOpts = computed(() => {
  let useOpts: YPopupEasyFormTypes.OptsItemType[] = [];
  props.opts.forEach((item) => {
    if (!item.defaultHide || (item?.showCond && item?.showCond(out.value))) {
      useOpts.push(item);
    }
  });
  return useOpts;
});
const props = withDefaults(defineProps<propsTypes>(), {
  title: "编辑",
  width: "",
  labelWidth: "",
  modelValue: () => {
    return {};
  },
  opts: () => {
    return [];
  },
});
const emit = defineEmits<{
  (e: "close"): void;
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

const state = reactive<{
  editBoxShow: boolean;
  callback?: ActionOpenCallback;
}>({
  editBoxShow: false,
});
const editBoxClosed = () => {
  emit("close");
};
const editFormSubmit = () => {
  optsCheck().then(() => {
    state.callback &&
      state.callback(() => {
        state.editBoxShow = false;
      });
  });
};
type ActionOpenCallback = (close: () => void) => void;
const ActionOpen = (cb?: ActionOpenCallback) => {
  state.editBoxShow = true;
  if (cb) state.callback = cb;
};
const optsCheck = () => {
  return new Promise<void>((resolve, reject: (msg: string) => void) => {
    props.opts.forEach((item) => {
      const isShow =
        !item.showCond || (item.showCond && item.showCond(out.value));
      if (
        isShow &&
        item.required &&
        !easyFormGetProperty(out.value, item.name)
      ) {
        const msg = showPlaceholder(item);
        ElMessage.error(msg);
        reject(msg);
        return;
      }
    });
    resolve();
  });
};
defineExpose({
  ActionOpen,
});

const showLabel = (item: YPopupEasyFormTypes.OptsItemType): string => {
  return item.label ? item.label : item.name;
};
const showPlaceholder = (item: YPopupEasyFormTypes.OptsItemType) => {
  if (!item.placeholder) {
    let outStr = "请输入";
    switch (item.type) {
      case YPopupEasyFormItemTypeState.text.value:
        outStr += showLabel(item);
        break;
      case YPopupEasyFormItemTypeState.uploadImage.value:
        outStr = "请上传" + showLabel(item);
        break;
      case YPopupEasyFormItemTypeState.switch.value:
        outStr = "开关" + showLabel(item);
        break;
    }
    return outStr;
  }
  return item.placeholder;
};
</script>

<style lang="scss" scoped></style>
