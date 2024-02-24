<template>
  <div
    @click="chooseItem"
    class="flex-def flex-zCenter flex-cCenter list-item-body"
  >
    <el-image
      class="y-pointer list-item-image"
      fit="contain"
      :src="mediaPath"
    ></el-image>
    <div class="y-pointer flex-def flex-zCenter flex-cCenter list-item-title">
      <div class="text-ellipsis" style="width: 95%">
        {{ props.attach.filename }}
      </div>
    </div>
    <div @click.stop class="item-action-box">
      <el-dropdown
        placement="bottom-end"
        trigger="click"
        @command="actionClick"
      >
        <div class="y-pointer flex-def flex-zCenter flex-cCenter more-action">
          <y-icon name="el-icon-More" size="10"></y-icon>
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item command="rename">
              <y-icon name="el-icon-Edit" color="#4ca2ff" size="18"></y-icon>
              <span>修改名称</span>
            </el-dropdown-item>
            <el-dropdown-item command="move-group">
              <y-icon name="el-icon-Fold" color="#4ca2ff" size="18"></y-icon>
              <span>移动分组</span>
            </el-dropdown-item>
            <el-dropdown-item command="delete">
              <y-icon name="el-icon-Delete" color="red" size="17"></y-icon>
              <span>删除素材</span>
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
  <el-dialog
    class="custom-image-body-item-dialog"
    top="35vh"
    destroy-on-close
    v-model="state.moveConShow"
    :show-close="false"
    width="20rem"
  >
    <el-space>
      <el-select v-model="state.chooseGroupId" placeholder="请选择">
        <el-option label="未分组" :value="0"></el-option>
        <el-option
          v-for="v in props.group"
          :key="v.id"
          :label="v.title"
          :value="v.id"
        >
        </el-option>
      </el-select>
      <el-button @click="moveGroup" type="primary">确认提交</el-button>
    </el-space>
  </el-dialog>
</template>

<script setup lang="ts">
import { computed, reactive } from "vue";
import { ElMessageBox } from "element-plus";
import { PostJson } from "@/components/yueyue-ui/libs/request/client";
import YIcon from "@/components/yueyue-ui/components/yl-icon/YlIcon.vue";
import type {
  groupListItemType,
  imageListItemType,
} from "@/components/yueyue-ui/components/yl-file-upload/types";
import { ToMedia } from "@/components/yueyue-ui/libs/function/common";

const props = withDefaults(
  defineProps<{
    attach: imageListItemType;
    group: groupListItemType[];
  }>(),
  {
    attach() {
      return {
        id: 0,
        filename: "",
        attachment: "",
        group_id: 0,
      };
    },
  }
);

const mediaPath = computed(() => {
  return ToMedia(props.attach.attachment);
});

const moveGroup = () => {
  state.moveConShow = false;
  PostJson("/admin/attach/group/move", {
    group_id: state.chooseGroupId,
    ids: [props.attach.id],
  }).then(() => emit("reload"));
};
const del = () => {
  ElMessageBox.confirm("确认要删除此资源?", "系统提示", {
    confirmButtonText: "确认删除",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    PostJson("/admin/attach/remove", { id: props.attach.id }).then(() => {
      emit("reload");
    });
  });
};

const rename = () => {
  ElMessageBox.prompt("请输入新的名称", "修改名称", {
    confirmButtonText: "确认提交",
    cancelButtonText: "取消",
    inputValue: props.attach.filename,
  }).then(({ value }) => {
    PostJson("/admin/attach/rename", {
      title: value,
      id: props.attach.id,
    }).then(() => {
      emit("reload");
    });
  });
};

let state = reactive({
  moveConShow: false,
  chooseGroupId: 0,
});

const actionClick = (e: string) => {
  switch (e) {
    case "delete":
      del();
      break;
    case "move-group":
      state.moveConShow = true;
      break;
    case "rename":
      rename();
      break;
  }
  console.log(e);
};

const chooseItem = () => {
  emit("choose-item", props.attach);
};

const emit = defineEmits<{
  (e: "choose-item", item: imageListItemType): void;
  (e: "reload"): void;
}>();
</script>
<style>
.custom-image-body-item-dialog .el-dialog__header {
  padding: 0 !important;
}
.custom-image-body-item-dialog .el-dialog__body {
  padding: 10px !important;
}
</style>
<style lang="scss" scoped>
$img-size: 8.7rem;

.list-item-body {
  width: $img-size;
  position: relative;
}

.list-item-image {
  background-color: #f8f8f8;
  width: $img-size;
  height: $img-size;
}

.list-item-title {
  color: #ffffff;
  background-color: rgba(0, 0, 0, 0.4);
  font-size: 0.8rem;
  line-height: 1.5rem;
  position: absolute;
  bottom: 0;
  width: $img-size;
}

.item-action-box {
  position: absolute;
  right: 0;
  top: 0;
}

.sort-action {
  line-height: 2rem;
  margin-bottom: 0.5rem;
  text-align: center;

  span {
    margin-left: 0.5rem;
  }
}

.more-action {
  border-radius: 2rem;
  padding: 0.2rem;
  background-color: #4ca2ff;
  color: #ffffff;
  margin-right: 5px;
  margin-top: 5px;
  opacity: 0;
}

.del-action {
  line-height: 2rem;
  text-align: center;

  span {
    margin-left: 0.5rem;
  }
}

.list-item-body:hover .more-action {
  opacity: 1;
}
</style>
