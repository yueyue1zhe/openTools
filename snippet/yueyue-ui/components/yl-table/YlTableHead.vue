<template>
  <transition name="el-fade-in">
    <div v-if="state.showComSearch" class="y-table-com-search">
      <el-row>
        <slot name="com-search"></slot>
        <y-table-com-search-item>
          <el-button @click="comSearchGo" type="primary">搜索</el-button>
          <el-button @click="comSearchReset">重置</el-button>
        </y-table-com-search-item>
      </el-row>
    </div>
  </transition>
  <div class="y-table-header">
    <y-tips-button
      v-if="props.refresh"
      @click="refreshAction"
      :tips="props.refreshTitle"
      icon="fa fa-refresh"
      type="info"
    ></y-tips-button>
    <slot name="left"></slot>
    <div class="y-table-search">
      <el-input
        v-if="fastSearch"
        @change="fastSearchGo"
        class="y-xs-hidden"
        v-model="state.searchKeyword"
        :placeholder="fastSearchPlaceholder"
      />
      <el-button-group v-if="comSearch" class="y-table-search-button-group">
        <el-button
          @click="state.showComSearch = !state.showComSearch"
          color="#dcdfe6"
          plain
        >
          <y-icon size="14" color="#303133" name="el-icon-Search" />
        </el-button>
      </el-button-group>
      <div class="y-m-l-15">
        <slot name="right"></slot>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from "vue";
import YIcon from "@/components/yueyue-ui/components/yl-icon/YlIcon.vue";
import YTipsButton from "@/components/yueyue-ui/components/yl-button/YlTipsButton.vue";
import YTableComSearchItem from "@/components/yueyue-ui/components/y-table/YTableComSearchItem.vue";

const props = withDefaults(
  defineProps<{
    refresh?: boolean;
    refreshTitle?: string;
    search?: boolean;
    fastSearch?: boolean;
    comSearch?: boolean;
    fastSearchPlaceholder?: string;
  }>(),
  {
    refresh: false,
    refreshTitle: "刷新",
    search: false,
    fastSearch: false,
    comSearch: false,
    fastSearchPlaceholder: "搜索",
  }
);

const state = reactive({
  showComSearch: false,
  searchKeyword: "",
});

const refreshAction = () => {
  emit("refresh");
};

const fastSearchGo = (val: string) => {
  emit("fast-search", val);
};

const comSearchReset = () => {
  emit("com-search-reset");
};
const comSearchGo = () => {
  emit("com-search-go");
};

const emit = defineEmits<{
  (e: "refresh"): void;
  (e: "com-search-go"): void;
  (e: "com-search-reset"): void;
  (e: "fast-search", value: string): void;
}>();

const FastSearchReset = () => {
  state.searchKeyword = "";
};

defineExpose({ FastSearchReset });
</script>

<style lang="scss" scoped>
.y-table-header {
  position: relative;
  overflow: hidden;
  box-sizing: border-box;
  display: flex;
  align-items: center;
  width: 100%;
  max-width: 100%;
  background-color: #ffffff;
  border: 1px solid #f6f6f6;
  border-bottom: none;
  padding: 13px 15px;
  font-size: 14px;
}

.y-table-com-search {
  overflow: hidden;
  box-sizing: border-box;
  width: 100%;
  max-width: 100%;
  background-color: #ffffff;
  border: 1px solid #f6f6f6;
  border-bottom: none;
  padding: 13px 15px;
  font-size: 14px;
}

.y-table-search {
  display: flex;
  margin-left: auto;
}

.y-table-search-button-group {
  display: flex;
  margin: 0 12px;

  button:focus,
  button:active {
    background-color: #ffffff;
  }

  button:hover {
    background-color: #dcdfe6;
  }
}
</style>
