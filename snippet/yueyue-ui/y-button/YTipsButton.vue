<template>
  <YTipsButton></YTipsButton>
</template>
<script lang="tsx" setup>
import { computed, h, resolveComponent } from "vue";
import YIcon from "@/components/yueyue-ui/y-icon/y-icon.vue";
import { ElPopconfirm } from "element-plus";

interface YTipsButtonPropsType {
  type:
    | ""
    | "default"
    | "primary"
    | "success"
    | "warning"
    | "info"
    | "danger"
    | "text";
  tips?: string;
  text?: string;
  icon?: string;
  small?: boolean;
  mode?: "tips" | "confirm" | "normal";
  size?: "large" | "default" | "small";
  confirmTips?: string; //确认弹窗提示文本 为空且tips存在时 给tips 追加 确认 。。 ?
}

const props = withDefaults(defineProps<YTipsButtonPropsType>(), {
  tips: "",
  text: "",
  icon: "",
  type: "",
  small: false,
  mode: "tips",
  size: "default",
});

const modeShow = computed(() => {
  let mode = props.mode;
  if (mode === "tips" && props.tips == "") {
    mode = "normal";
  }
  return mode;
});

const emit = defineEmits<{
  (e: "click"): void;
}>();
const clickThis = () => {
  emit("click");
};

const YTipsButton = () => {
  switch (modeShow.value) {
    case "confirm":
      return confirmRender;
    case "normal":
      return normalRender;
    case "tips":
      return tipsRender;
    default:
      return normalRender;
  }
};

const iconRender = h(YIcon, { name: props.icon });
const spanRender = h("span", { class: ["y-button-text"] }, props.text);
const normalRender = h(
  resolveComponent("el-button"),
  {
    onClick: modeShow.value !== "confirm" ? clickThis : "",
    class: ["y-button", { small: props.small }],
    type: props.type,
    size: props.size,
  },
  () => {
    let arr = [];
    if (props.icon) arr.push(iconRender);
    if (props.text) arr.push(spanRender);
    return arr;
  }
);
const tipsRender = h(
  resolveComponent("el-tooltip"),
  {
    content: props.tips,
    placement: "top",
  },
  () => {
    return normalRender;
  }
);

const confirmTips = computed(() => {
  let str = props.confirmTips;
  if (!props.confirmTips && props.tips) str = "确认要" + props.tips + "?";
  return str || "确认？";
});
const confirmRender = h(
  ElPopconfirm,
  {
    onConfirm: clickThis,
    title: confirmTips.value,
  },
  {
    reference: () => {
      return h(
        "div",
        {
          style: "display: inline-flex;vertical-align: middle",
          class: "y-button",
        },
        props.tips ? tipsRender : normalRender
      );
    },
  }
);
</script>

<style lang="scss">
.y-button-text {
  margin-left: 6px;
}

.y-button.small .icon {
  font-size: 18px !important;
}

.y-button.small {
  padding: 4px 5px;
  height: auto;
}

.y-button .icon {
  font-size: 14px !important;
  color: #ffffff !important;
}
</style>
