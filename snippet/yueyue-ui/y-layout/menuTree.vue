<template>
  <template v-for="menu in props.menus">
    <template v-if="menu.children && menu.children.length > 0">
      <el-sub-menu :index="menu.path" :key="menu.path">
        <template #title>
          <y-icon
            :color="menuColor"
            :name="menu.icon ? menu.icon : defaultIcon"
          />
          <span>{{ menu.title ? menu.title : noTitle }}</span>
        </template>
        <menu-tree :menus="menu.children"></menu-tree>
      </el-sub-menu>
    </template>
    <template v-else>
      <el-menu-item
        v-if="menu.type === 'tab'"
        :index="menu.path"
        :key="menu.path"
      >
        <y-icon
          :color="menuColor"
          :name="menu.icon ? menu.icon : defaultIcon"
        />
        <span>{{ menu.title ? menu.title : noTitle }}</span>
      </el-menu-item>
      <el-menu-item
        v-if="menu.type === 'link'"
        index=""
        :key="menu.path"
        @click="onLink(menu.path)"
      >
        <y-icon
          :color="menuColor"
          :name="menu.icon ? menu.icon : defaultIcon"
        />
        <span>{{ menu.title ? menu.title : noTitle }}</span>
      </el-menu-item>
    </template>
  </template>
</template>
<script setup lang="ts">
import { computed, withDefaults } from "vue";
import type { viewMenu } from "@/stores/asideTabs";
import useLayoutConfigStore from "@/stores/layoutConfig";
import YIcon from "@/components/yueyue-ui/y-icon/y-icon.vue";

const layoutConfig = useLayoutConfigStore();

interface Props {
  menus: viewMenu[];
}

const props = withDefaults(defineProps<Props>(), {
  menus: () => [],
});

const noTitle = "未设置菜单名";

const defaultIcon = computed(() => layoutConfig.menuDefaultIcon);
const menuColor = computed(() => layoutConfig.menuColor);
const menuActiveBackground = computed(() => layoutConfig.menuActiveBackground);

const onLink = (url: string) => {
  window.open(url, "_blank");
};
</script>

<style scoped lang="scss">
.el-sub-menu .icon,
.el-menu-item .icon {
  vertical-align: middle;
  margin-right: 5px;
  width: 24px;
  text-align: center;
}

.is-active .icon {
  color: var(--el-menu-active-color) !important;
}

.el-menu-item.is-active {
  background-color: v-bind(menuActiveBackground);
}
</style>
