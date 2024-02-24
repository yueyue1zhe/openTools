<template>
  <div class="layout-logo">
    <img
      v-if="logoShow && !layoutConfig.menuCollapse"
      class="logo-img"
      :src="logoShow"
      alt="logo"
    />
    <div
      v-if="!layoutConfig.menuCollapse"
      :style="{ color: layoutConfig.menuActiveColor }"
      class="website-name text-ellipsis"
    >
      {{ titleShow }}
    </div>
    <y-icon
      @click="onMenuCollapse"
      :name="layoutConfig.menuCollapse ? 'fa fa-indent' : 'fa fa-dedent'"
      :class="layoutConfig.menuCollapse ? 'unfold' : ''"
      :color="layoutConfig.menuActiveColor"
      size="18"
      class="fold"
    />
  </div>
</template>

<script setup lang="ts">
import useLayoutConfigStore from "@/stores/layoutConfig";

import YIcon from "@/components/yueyue-ui/components/yl-icon/YlIcon.vue";
import { useAsideTabsStore } from "@/stores/asideTabs";
import { useSettingStore } from "@/stores/setting";
import { computed } from "vue";

const asideTabs = useAsideTabsStore();
const settingStore = useSettingStore();
const titleShow = computed((): string => {
  let title = asideTabs.title || settingStore.title || "管理后台";
  let contact = title && settingStore.title ? "-" : "";
  title += contact + settingStore.append_title;
  return title;
});
const logoShow = computed((): string => {
  return asideTabs.logo || settingStore.logo;
});

const layoutConfig = useLayoutConfigStore();
function onMenuCollapse() {
  layoutConfig.menuCollapse = !layoutConfig.menuCollapse;
}
</script>

<style scoped lang="scss">
.layout-logo {
  width: 100%;
  height: 50px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-sizing: border-box;
  padding: 10px;
  background: transparent;
}

.logo-img {
  width: 28px;
}

.website-name {
  padding-left: 4px;
  font-size: var(--el-font-size-extra-large);
  font-weight: 600;
}

.fold {
  margin-left: auto;
}

.unfold {
  margin: 0 auto;
}
</style>
