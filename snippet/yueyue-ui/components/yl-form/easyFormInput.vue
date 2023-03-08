<template>
  <el-input v-model="useOut" :placeholder="placeholder" autocomplete="off" />
</template>

<script lang="ts" setup>
import { computed } from "vue";
import {
  easyFormGetProperty,
  easyFormSetProperty,
} from "@/components/yueyue-ui/components/yl-form/easyFormTools";

interface propsType {
  modelValue: AnyObject;
  opt: YlEasyFormTypes.OptsItemType;
  placeholder: string;
}

const props = withDefaults(defineProps<propsType>(), {
  modelValue: () => {
    return {};
  },
  opt: () => {
    return {} as YlEasyFormTypes.OptsItemType;
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
