<template>
  <div class="y-table-pagination">
    <el-pagination
      v-model:currentPage="state.pageIndex"
      v-model:page-size="state.pageSize"
      :page-sizes="[10, 20, 50]"
      background
      layout="sizes,total, ->, prev, pager, next, jumper"
      :total="state.total"
      hide-on-single-page
    />
  </div>
</template>

<script lang="ts" setup>
import { computed, reactive, watch } from "vue";

const prop = withDefaults(
  defineProps<{
    opt: Omit<PageResult<AnyObject>, "list">;
  }>(),
  {}
);
let state = reactive({
  pageIndex: 0,
  pageSize: 0,
  total: 0,
});
watch(
  () => prop.opt,
  (item) => {
    if (item.page != state.pageIndex) state.pageIndex = item.page;
    if (item.size != state.pageSize) state.pageSize = item.size;
    if (item.total != state.total) state.total = item.total;
  },
  {
    immediate: true,
    deep: true,
  }
);
watch(
  () => state,
  (item) => {
    emit("change", { page: item.pageIndex, size: item.pageSize });
  },
  {
    immediate: true,
    deep: true,
  }
);

const emit = defineEmits<{
  (e: "change", res: PaginationResult): void;
}>();

const padding = computed(() => {
  return prop.opt.total > prop.opt.size ? "13px 15px" : "";
});
</script>

<style lang="scss" scoped>
.y-table-pagination {
  background-color: #ffffff;
  box-sizing: border-box;
  width: 100%;
  max-width: 100%;
  padding: v-bind(padding);
}
</style>
