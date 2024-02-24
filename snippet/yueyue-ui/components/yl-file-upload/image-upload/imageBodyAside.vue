<template>
  <el-aside width="150px" style="background-color: #ededed">
    <el-scrollbar class="list-scroll-body">
      <template v-for="(item, key) in groupListShow" :key="key">
        <div
          class="flex-def flex-cCenter aside-item-box"
          style="position: relative"
        >
          <div
            style="width: 100%"
            @click="changeCurTab(item.id)"
            :class="curTab === item.id ? 'active' : ''"
            class="text-ellipsis y-pointer aside-item"
          >
            {{ item.title }}
          </div>
          <div
            v-if="item.id > 0"
            style="position: absolute; right: 0.7rem"
            class="y-pointer delete-con"
          >
            <el-popconfirm
              @confirm="groupDel(item.id)"
              title="确定删除这个分组吗？"
              width="200"
              confirm-button-text="确认"
              cancel-button-text="取消"
            >
              <template #reference>
                <y-icon name="el-icon-Delete" size="12" color="red"></y-icon>
              </template>
            </el-popconfirm>
          </div>
        </div>
      </template>
    </el-scrollbar>
    <div
      @click="groupAdd"
      class="aside-item create-con y-pointer flex-def flex-cCenter"
    >
      <y-icon name="el-icon-Plus" size="15"></y-icon>
      <span>新建分组</span>
    </div>
  </el-aside>
</template>

<script setup lang="ts">
import { computed, ref } from "vue";
import { ElMessageBox } from "element-plus";
import { PostJson } from "@/components/yueyue-ui/libs/request/client";
import YIcon from "@/components/yueyue-ui/components/yl-icon/YlIcon.vue";
import type { groupListItemType } from "@/components/yueyue-ui/components/yl-file-upload/types";

let groupList = ref<groupListItemType[]>([]);
let curTab = ref<number>(-1);

const groupListShow = computed(() => {
  let arr = [
    { id: -1, title: "全部" },
    { id: 0, title: "未分组" },
  ];
  return [...arr, ...groupList.value];
});

const groupLoad = () => {
  PostJson<groupListItemType[]>("/admin/attach/group/all").then((res) => {
    groupList.value = res;
    emit("group-load", res);
  });
};

const emit = defineEmits<{
  (e: "tab-change", index: number): void;
  (e: "group-load", group: groupListItemType[]): void;
}>();

const changeCurTab = (index: number) => {
  curTab.value = index;
  emit("tab-change", index);
};

const openAside = (index: number) => {
  curTab.value = index;
  groupLoad();
};

const groupDel = (id: number) => {
  PostJson("/admin/attach/group/del", { id }).then(() => {
    changeCurTab(-1);
    groupLoad();
  });
};

const groupAdd = () => {
  ElMessageBox.prompt("请输入分组名称", "创建新的分组", {
    confirmButtonText: "确认提交",
    cancelButtonText: "取消",
  }).then(({ value }) => {
    PostJson("/admin/attach/group/add", { title: value }).then(() => {
      groupLoad();
    });
  });
};

defineExpose({ changeCurTab, openAside });
</script>

<style lang="scss" scoped>
.aside-item.active {
  background-color: var(--color-primary-sub-9);
  color: var(--color-primary);
}

.aside-item {
  line-height: 3rem;
  padding: 0 2rem;
}

.delete-con {
  display: none;
}

.aside-item-box:hover > .delete-con {
  display: block;
}

.aside-item.create-con {
  background-color: var(--color-basic-white);
  color: var(--color-primary);
  border-right: 1px solid var(--color-sub-3);

  span {
    margin-left: 0.5rem;
    font-weight: 600;
    font-size: var(--el-font-size-extra-small);
  }
}

.list-box {
  display: flex;
  justify-content: center;
  position: relative;
}

.list-scroll-body {
  height: calc(100% - 3rem);
}
</style>
