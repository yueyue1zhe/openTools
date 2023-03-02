<template>
  <div class="nav-bar">
    <NavMenus></NavMenus>
  </div>
</template>

<script setup lang="ts">
import useLayoutConfigStore from "@/stores/layoutConfig";

import NavMenus from "@/components/yueyue-ui/y-layout/navMenus.vue";
import { computed } from "vue";

const layoutConfig = useLayoutConfigStore();

const headerBarTabColor = computed(() => layoutConfig.headerBarTabColor);
const headerBarTabActiveColor = computed(
  () => layoutConfig.headerBarTabActiveColor
);
const headerBarTabActiveBackground = computed(
  () => layoutConfig.headerBarTabActiveBackground
);
</script>

<style lang="scss" scoped>
.nav-bar {
  display: flex;
  position: absolute;
  z-index: 1000;
  left: var(--main-space);
  width: calc(100% - var(--main-space) * 3);
  background-color: rgba(0, 0, 0, 0);
  justify-content: right;
  height: 50px;
  box-sizing: content-box;
  padding: 20px var(--main-space) calc(var(--main-space) / 2);
  margin-bottom: calc(var(--main-space) / 2);
  background-image: radial-gradient(transparent 1px, var(--color-bg-1) 1px);
  backdrop-filter: saturate(50%) blur(4px);
  background-size: 4px 4px;

  :deep(.nav-tabs) {
    display: flex;
    height: 100%;
    position: relative;

    .bd-nav-tab {
      display: flex;
      align-items: center;
      justify-content: center;
      padding: 0 20px;
      cursor: pointer;
      z-index: 1;
      user-select: none;
      opacity: 0.7;
      color: v-bind(headerBarTabColor);

      .close-icon {
        padding: 2px;
        margin: 2px 0 0 4px;
      }

      .close-icon:hover {
        background: var(--color-primary-sub-0);
        color: var(--color-sub-1) !important;
        border-radius: 50%;
      }

      &.active {
        color: v-bind(headerBarTabActiveColor);
      }

      &:hover {
        opacity: 1;
      }
    }

    .nav-tabs-active-box {
      position: absolute;
      height: 40px;
      border-radius: var(--el-border-radius-base);
      background-color: v-bind(headerBarTabActiveBackground);
      box-shadow: var(--el-box-shadow-light);
      transition: all 0.2s;
      -webkit-transition: all 0.2s;
    }
  }
}
</style>
