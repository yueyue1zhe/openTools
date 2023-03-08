<template>
  <el-aside class="app-aside-default">
    <y-logo v-if="layoutConfig.menuShowTopBar"></y-logo>
    <menu-vertical></menu-vertical>
  </el-aside>
</template>

<script setup lang="ts">
import useLayoutConfigStore from "@/stores/layoutConfig";
import YLogo from "@/components/yueyue-ui/components/yl-layout/YlLogo.vue";
import MenuVertical from "@/components/yueyue-ui/components/yl-layout/menuVertical.vue";

const layoutConfig = useLayoutConfigStore();
</script>

<style lang="scss" scoped>
.app-aside-default {
  background: var(--color-basic-white);
  margin: 16px 0 16px 16px;
  height: calc(100vh - 32px);
  box-shadow: var(--el-box-shadow-light);
  border-radius: var(--el-border-radius-base);
  overflow: hidden;
  transition: width 0.3s ease;
  width: v-bind("layoutConfig.GetMenuWidth");
  cursor: pointer;
}
</style>
