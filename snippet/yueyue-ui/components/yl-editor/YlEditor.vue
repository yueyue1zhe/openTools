<template>
  <div style="border: 1px solid #ccc" class="w100">
    <Toolbar
      style="border-bottom: 1px solid #ccc"
      :editor="editorRef"
      :defaultConfig="toolbarConfig"
      :mode="mode"
    />
    <Editor
      :style="editorStyle"
      v-model="valueHtml"
      :defaultConfig="editorConfig"
      :mode="mode"
      @onCreated="handleCreated"
    />
    <y-upload-image :show-input="false" ref="refYUploadImage"></y-upload-image>
  </div>
</template>

<script lang="ts" setup>
import "@wangeditor/editor/dist/css/style.css";
import { Editor, Toolbar } from "@wangeditor/editor-for-vue";
import { computed, onBeforeUnmount, ref, shallowRef } from "vue";
import { IDomEditor, IEditorConfig, IToolbarConfig } from "@wangeditor/editor";
import YUploadImage from "@/components/yueyue-ui/components/yl-file-upload/YlUploadImage.vue";
import { ToMedia } from "@/components/yueyue-ui/libs/function/common";

const mode = "default";

const editorRef = shallowRef();
const refYUploadImage = ref<InstanceType<typeof YUploadImage>>();

const props = withDefaults(
  defineProps<{
    height?: string;
    modelValue: string;
    placeholder?: string;
  }>(),
  {
    modelValue: "",
    height: "23rem",
    placeholder: "请输入内容...",
  }
);
const emit = defineEmits<{
  (e: "update:modelValue", out: string): void;
}>();

const editorStyle = computed(() => {
  return `height:${props.height};overflow-y: hidden;`;
});

const valueHtml = computed({
  get(): string {
    return props.modelValue;
  },
  set(value: string) {
    emit("update:modelValue", value);
  },
});

const toolbarConfig: Partial<IToolbarConfig> = {
  insertKeys: {
    index: 25,
    keys: "imageUploader",
  },
  excludeKeys: [
    "headerSelect",
    "blockquote",
    "code",
    "insertLink",
    "fontFamily",
    "todo",
    "codeBlock",
    "group-video",
    "insertTable",
    "fullScreen",
    "group-image",
    "editImage",
    "bulletedList",
    "numberedList",
  ],
};
const editorConfig: Partial<IEditorConfig> = {
  placeholder: props.placeholder,
  hoverbarKeys: {
    text: {
      menuKeys: [
        "headerSelect",
        "bold",
        "through",
        "color",
        "bgColor",
        "clearStyle",
      ],
    },
    image: {
      menuKeys: [
        "imageWidth30",
        "imageWidth50",
        "imageWidth100",
        "deleteImage",
      ],
    },
  },
};

// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
  const editor = editorRef.value;
  if (editor == null) return;
  editor.destroy();
});

const handleCreated = (editor: IDomEditor) => {
  editor.on("want-upload-image", function () {
    refYUploadImage.value?.chooseStart((e: string) => {
      const imgUrl = ToMedia(e);
      editor.dangerouslyInsertHtml(`<img src="${imgUrl}"  alt=""/>`);
    });
  });
  editorRef.value = editor; // 记录 editor 实例，重要！
};
</script>

<style lang="scss" scoped></style>
