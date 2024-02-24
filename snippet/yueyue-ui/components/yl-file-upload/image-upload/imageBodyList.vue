<template>
  <el-main style="padding: 0">
    <div class="flex-def flex-warp list-panel">
      <div
        class="flex-def flex-zCenter flex-cCenter list-item-box"
        v-for="(item, key) in attach.list"
        :key="key"
      >
        <image-body-item
          @choose-item="chooseItem"
          @reload="attachLoad"
          :attach="item"
          :group="props.group"
        ></image-body-item>
      </div>
      <div
        v-if="attach.list.length === 0"
        class="flex-def flex-zCenter flex-cCenter no-attach"
      >
        还没有资源...
      </div>
    </div>
    <div class="flex-def flex-zCenter flex-cCenter bottom-pagination">
      <el-pagination
        :hide-on-single-page="true"
        background
        small
        layout="total,prev, pager,next"
        :total="attach.total"
        :page-size="attach.size"
        :current-page="attach.page"
        @current-change="pageChange"
      >
      </el-pagination>
    </div>
  </el-main>
</template>

<script setup lang="ts">
import { reactive, watch } from "vue";
import { PostJson } from "@/components/yueyue-ui/libs/request/client";
import type {
  imageListItemType,
  groupListItemType,
} from "@/components/yueyue-ui/components/yl-file-upload/types";
import ImageBodyItem from "@/components/yueyue-ui/components/yl-file-upload/image-upload/imageBodyItem.vue";

const attach = reactive<PageResult<imageListItemType>>({
  list: [],
  page: 1,
  total: 0,
  size: 15,
});

const props = withDefaults(
  defineProps<{
    group_id: number;
    group: groupListItemType[];
  }>(),
  {
    group_id: -1,
  }
);
watch(
  () => props.group_id,
  () => {
    attach.page = 1;
    attach.list.length = 0;
    attachLoad();
  }
);

const emit = defineEmits<{
  (e: "choose", item: imageListItemType): void;
}>();

const chooseItem = (item: imageListItemType) => {
  emit("choose", item);
};

const attachLoad = () => {
  PostJson<PageResult<imageListItemType>>("/admin/attach/search", {
    group_id: props.group_id >= 0 ? props.group_id : 0,
    all: props.group_id == -1,
    page: attach.page,
  }).then((res) => {
    attach.list = res.list;
    attach.total = res.total;
    attach.size = res.size;
  });
};

const pageChange = (index: number) => {
  attach.page = index;
  attachLoad();
};
const openBody = () => {
  attachLoad();
};
defineExpose({ openBody, attachLoad });
</script>

<style lang="scss" scoped>
.list-panel {
  height: calc(100% - 3rem);
  padding: 5px 10px;
}

$img-size: 8.7rem;

.list-item-box {
  margin: 0.5rem;
  position: relative;
  width: calc(20% - 1rem);
  height: $img-size;
}

.no-attach {
  width: 100%;
  font-size: 1rem;
  color: #ededed;
  font-weight: 600;
}

.bottom-pagination {
  height: 3rem;
  border-top: 1px solid var(--color-sub-3);
}
</style>
