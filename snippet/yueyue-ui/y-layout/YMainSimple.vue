<template>
  <el-main class="layout-main">
    <el-scrollbar
      class="layout-main-scrollbar"
      :style="layoutMainScrollbarStyle()"
      ref="mainScrollbarRef"
    >
      <slot></slot>
      <router-view v-slot="{ Component }">
        <transition :name="layoutMainAnimation" mode="out-in">
          <component :is="Component" :key="state.componentKey" />
        </transition>
      </router-view>
    </el-scrollbar>
  </el-main>
</template>

<script setup lang="ts">
import { computed, reactive, watch } from "vue";
import { useRoute } from "vue-router";
import { mainHeight as layoutMainScrollbarStyle } from "@/common/util/layout";
import useLayoutConfigStore from "@/stores/layoutConfig";

const route = useRoute();

const layoutConfig = useLayoutConfigStore();
const layoutMainAnimation = computed(() => layoutConfig.mainAnimation);

const state: {
  componentKey: string;
} = reactive({
  componentKey: route.path,
});

watch(
  () => route.path,
  () => {
    state.componentKey = route.path;
  }
);
</script>

<style scoped lang="scss">
.layout-container .layout-main {
  padding: 0 !important;
  overflow: hidden;
  width: 100%;
  height: 100%;
}

.layout-main-scrollbar {
  width: 100%;
  position: relative;
  overflow: hidden;
}
</style>
