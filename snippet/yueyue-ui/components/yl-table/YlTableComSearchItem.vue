<template>
  <el-col :span="showSpan">
    <div class="y-com-search-col" :class="{ 'y-p-l-20': !label }">
      <template v-if="label">
        <div class="y-com-search-col-label">{{ label }}</div>
        <div class="y-com-search-col-input">
          <slot></slot>
        </div>
      </template>
      <template v-else><slot></slot></template>
    </div>
  </el-col>
</template>

<script lang="ts" setup>
import { computed } from "vue";

const props = withDefaults(
  defineProps<{
    label?: string;
    span?: string;
  }>(),
  {
    label: "",
    span: "6",
  }
);

const showSpan = computed(() => {
  return parseInt(props.span);
});
</script>

<style lang="scss" scoped>
.y-com-search-col {
  display: flex;
  align-items: center;
  padding-top: 8px;
  color: var(--color-regular);
  font-size: 13px;
}

.y-com-search-col-label {
  width: 33.33%;
  padding: 0 15px;
  text-align: right;
  overflow: hidden;
  white-space: nowrap;
}

.y-com-search-col-input {
  padding: 0 15px;
  width: 66.66%;
}
</style>
