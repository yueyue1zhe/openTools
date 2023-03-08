<template>
  <div>
    <input
      class="y-file-input"
      :id="props.id"
      @change="fileInputChange"
      type="file"
    />
    <el-button class="w100" @click="actionFileClick">{{ btnTxt }}</el-button>
    <textarea v-if="out !== ''" v-model="out" rows="4" class="w100"></textarea>
  </div>
</template>

<script lang="ts" setup>
import { computed, reactive } from "vue";
import { ElMessage } from "element-plus";

interface propType {
  id?: string;
  modelValue?: string;
}

const props = withDefaults(defineProps<propType>(), {
  id: "files",
  modelValue: "",
});
const emit = defineEmits<{
  (e: "update:modelValue", out: string): void;
}>();
const out = computed({
  get: (): string => {
    return props.modelValue;
  },
  set: (value: string) => {
    emit("update:modelValue", value);
  },
});

const actionFileClick = () => {
  document.getElementById(props.id)?.click();
};

let state = reactive({
  fileName: "",
  content: "",
});

const btnTxt = computed(() => {
  return state.fileName || (props.modelValue ? "已上传" : "") || "选择文件";
});

const fileInputChange = (e: Event) => {
  let eventFiles = (e.target as HTMLInputElement).files;
  if (!eventFiles || eventFiles.length <= 0)
    return ElMessage.info("未选择文件");
  const chooseFile = eventFiles[0];
  const reader = new FileReader();
  reader.readAsText(chooseFile, "UTF-8");
  state.fileName = chooseFile.name;
  reader.onload = function (e) {
    if (typeof e.target?.result == "string") {
      out.value = state.content = e.target.result;
    } else {
      ElMessage.error("请选择正确的文件");
    }
  };
};
</script>

<style lang="scss" scoped>
.y-file-input {
  display: none;
}
</style>
