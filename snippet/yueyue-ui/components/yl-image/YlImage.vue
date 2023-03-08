<template>
  <el-image
    :style="style"
    :src="showSrc"
    :zoom-rate="1.2"
    :preview-src-list="props.preview ? [showSrc] : []"
    fit="cover"
    preview-teleported
    hide-on-click-modal
    :class="props.class"
  />
</template>

<script lang="ts" setup>
import { ToMedia } from "@/components/yueyue-ui/libs/function/common";
import { computed } from "vue";

const props = withDefaults(
  defineProps<{
    src: string;
    style?: string;
    class?: string;
    preview?: boolean;
  }>(),
  {
    src: "",
    style: "width: 5rem; height: 5rem;border-radius:5px",
    class: "",
    preview: true,
  }
);

const showSrc = computed(() => {
  return ToMedia(props.src);
});
</script>

<style lang="scss" scoped></style>
