<template>
  <div class="y-table-box">
    <slot></slot>
  </div>
</template>

<script lang="ts" setup></script>

<style lang="scss">
.y-table-box {
  border-radius: var(--el-border-radius-round);

  .el-table .el-button + .el-button,
  .el-table .y-button + .y-button {
    margin-left: 6px;
  }

  .el-table {
    --el-table-border-color: #f6f6f6;
  }
}
</style>
