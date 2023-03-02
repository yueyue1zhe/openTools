<template>
  <el-scrollbar ref="verticalMenusRef" class="vertical-menus-scrollbar">
    <el-menu
      class="layouts-menu-vertical"
      router
      :collapse-transition="false"
      :unique-opened="layoutConfig.menuUniqueOpened"
      :default-active="state.defaultActive"
      :collapse="layoutConfig.menuCollapse"
      :background-color="layoutConfig.menuBackground"
      :text-color="layoutConfig.menuColor"
      :active-text-color="layoutConfig.menuActiveColor"
    >
      <MenuTree :menus="navTabs.viewRoutes"></MenuTree>
    </el-menu>
  </el-scrollbar>
</template>

<script setup lang="ts">
import { computed, nextTick, onMounted, reactive, ref } from "vue";
import { onBeforeRouteUpdate, useRoute } from "vue-router";
import type { RouteLocationNormalized } from "vue-router";
import type { ElScrollbar } from "element-plus";

import useLayoutConfigStore from "@/stores/layoutConfig";
import MenuTree from "@/components/yueyue-ui/y-layout/menuTree.vue";
import { useAsideTabsStore } from "@/stores/asideTabs";

const route = useRoute();
const layoutConfig = useLayoutConfigStore();
const navTabs = useAsideTabsStore();

const verticalMenusRef = ref<InstanceType<typeof ElScrollbar>>();

const state = reactive({
  defaultActive: "",
});

const verticalMenusScrollbarHeight = computed(() => {
  let menuTopBarHeight = 0;
  if (layoutConfig.menuShowTopBar) {
    menuTopBarHeight = 50;
  }
  return "calc(100vh - " + (32 + menuTopBarHeight) + "px)";
});

// 激活当前路由的菜单
//RouteLocationNormalizedLoaded
const currentRouteActive = (currentRoute: RouteLocationNormalized) => {
  let useActiveRoute = currentRoute.path;
  const constPathSplitArr = currentRoute.path.split("/");
  if (constPathSplitArr.length > 4)
    useActiveRoute = constPathSplitArr.slice(0, 4).join("/");

  state.defaultActive = useActiveRoute;
};

// 滚动条滚动到激活菜单所在位置
const verticalMenusScroll = () => {
  nextTick(() => {
    let activeMenu: HTMLElement | null = document.querySelector(
      ".el-menu.layouts-menu-vertical li.is-active"
    );
    if (!activeMenu) return false;
    verticalMenusRef.value?.setScrollTop(activeMenu.offsetTop);
  });
};

onMounted(() => {
  currentRouteActive(route);
  verticalMenusScroll();
});

onBeforeRouteUpdate((to) => {
  currentRouteActive(to);
});
</script>
<style>
.vertical-menus-scrollbar {
  height: v-bind(verticalMenusScrollbarHeight);
}

.layouts-menu-vertical {
  border: 0;
}
</style>
