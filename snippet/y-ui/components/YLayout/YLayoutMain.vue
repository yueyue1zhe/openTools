<template>
  <view class="layout-main-box">
    <slot name="header"></slot>
    <view class="y-flex-1">
      <slot></slot>
    </view>
    <slot name="footer"></slot>
  </view>
</template>

<script lang="ts" setup>

</script>

<style lang="scss" scoped>
.layout-main-box {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}
</style>