<template>
  <el-dropdown
    trigger="click"
    placement="bottom-end"
    @command="dropdownCommand"
  >
    <div class="btn-bg">
      <action-icon name="fa-solid fa-ellipsis"></action-icon>
    </div>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item
          v-for="item in props.command"
          :key="item.name"
          :command="item.name"
          :divided="item?.divided"
          :style="{ color: item?.color }"
        >
          {{ item.title }}
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script setup lang="ts">
import ActionIcon from "@/components/business/action-icon/ActionIcon.vue";

export interface actionMorePropsTypeItem {
  title: string;
  name: string;
  divided?: boolean;
  color?: string;
}

interface actionMorePropsType {
  command: actionMorePropsTypeItem[];
  index?: number;
}

const props = defineProps<actionMorePropsType>();

const emit = defineEmits<{
  (e: "click", name: string, index: number): void;
}>();

const dropdownCommand = (e: string) => {
  let outIndex = -1;
  if (typeof props?.index === "number") {
    outIndex = props.index;
  }
  emit("click", e, outIndex);
};
</script>

<style lang="scss" scoped>
.btn-bg {
  height: 24px;
  width: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.btn-bg:hover {
  background-color: var(--color-sub-2);
  border-radius: 3px;
}

.btn-bg:hover :deep(.icon-action) {
  color: var(--color-text-primary);
}
</style>
