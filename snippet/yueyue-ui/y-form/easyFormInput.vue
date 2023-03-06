<template>
  <el-input v-model="useOut" :placeholder="placeholder" autocomplete="off" />
</template>

<script lang="ts" setup>
import { computed } from "vue";
import {
  easyFormGetProperty,
  easyFormSetProperty,
} from "@/components/yueyue-ui/y-form/easyFormTools";

interface propsType {
  modelValue: AnyObject;
  opt: YPopupEasyFormTypes.OptsItemType;
  placeholder: string;
}

const props = withDefaults(defineProps<propsType>(), {
  modelValue: () => {
    return {};
  },
  opt: () => {
    return {} as YPopupEasyFormTypes.OptsItemType;
  },
  placeholder: "",
});
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
const useOut = computed({
  get(): AnyObject {
    return easyFormGetProperty(out.value, props.opt.name);
  },
  set(value: AnyObject) {
    easyFormSetProperty(out.value, props.opt.name, value);
  },
});
</script>

<style lang="scss" scoped></style>
