<template>
  <div class="nav-menus default">
    <router-link class="h100" to="/">
      <div class="nav-menu-item">
        <y-icon
          :color="layoutConfig.headerBarTabColor"
          class="nav-menu-icon icon"
          name="el-icon-Odometer"
          size="18"
        />
      </div>
    </router-link>
    <router-link class="h100" to="/founder">
      <div class="nav-menu-item">
        <y-icon
          :color="layoutConfig.headerBarTabColor"
          class="nav-menu-icon icon"
          name="el-icon-Monitor"
          size="18"
        />
      </div>
    </router-link>
    <el-popover
      @show="onCurrentNavMenu(true, 'adminInfo')"
      @hide="onCurrentNavMenu(false, 'adminInfo')"
      placement="bottom-end"
      :hide-after="0"
      trigger="click"
    >
      <template #reference>
        <div
          class="admin-info"
          :class="state.currentNavMenu === 'adminInfo' ? 'hover' : ''"
        >
          <div class="admin-name">admin</div>
        </div>
      </template>
      <div
        class="admin-right-menu flex-def flex-zTopBottom flex-zCenter flex-cCenter"
      >
        <div>
          <el-button
            @click="router.push({ name: 'user-info' })"
            type="primary"
            plain
            >资料设置
          </el-button>
        </div>
        <div class="y-m-t-10">
          <el-button type="danger" plain>退出账号</el-button>
        </div>
      </div>
    </el-popover>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from "vue";
import YIcon from "@/components/yueyue-ui/components/yl-icon/YlIcon.vue";
import useLayoutConfigStore from "@/stores/layoutConfig";
import { useRouter } from "vue-router";

const state = reactive({
  isFullScreen: false,
  currentNavMenu: "",
  showLayoutDrawer: false,
});

const layoutConfig = useLayoutConfigStore();

const onCurrentNavMenu = (status: boolean, name: string) => {
  state.currentNavMenu = status ? name : "";
};

const router = useRouter();
</script>

<style scoped lang="scss">
.nav-menus.default {
  border-radius: var(--el-border-radius-base);
  box-shadow: var(--el-box-shadow-light);
}

.nav-menus {
  display: flex;
  align-items: center;
  height: 100%;
  //margin-left: auto;
  background-color: v-bind("layoutConfig.headerBarBackground");

  .nav-menu-item {
    height: 100%;
    width: 40px;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;

    .nav-menu-icon {
      box-sizing: content-box;
    }

    &:hover {
      .icon {
        animation: twinkle 0.3s ease-in-out;
      }
    }
  }

  .admin-info {
    display: flex;
    height: 100%;
    padding: 0 10px;
    align-items: center;
    cursor: pointer;
    user-select: none;
    color: v-bind("layoutConfig.headerBarTabColor");
  }

  .admin-name {
    padding-left: 6px;
  }

  .nav-menu-item:hover,
  .admin-info:hover,
  .nav-menu-item.hover,
  .admin-info.hover {
    background: v-bind("layoutConfig.headerBarHoverBackground");
  }
}

.chang-lang :deep(.el-dropdown-menu__item) {
  justify-content: center;
}

.admin-info-base {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  padding-top: 10px;

  .admin-info-other {
    display: block;
    width: 100%;
    text-align: center;
    padding: 10px 0;

    .admin-info-name {
      font-size: var(--el-font-size-large);
    }
  }
}

.admin-info-footer {
  padding: 10px 0;
  margin: 0 -12px -12px -12px;
  display: flex;
  justify-content: space-around;
  background: var(--color-bg-2);
}

@keyframes twinkle {
  0% {
    transform: scale(0);
  }
  80% {
    transform: scale(1.2);
  }
  100% {
    transform: scale(1);
  }
}
</style>
