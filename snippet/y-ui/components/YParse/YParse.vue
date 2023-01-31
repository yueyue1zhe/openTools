<template>
  <rich-text :nodes="props.content"></rich-text>
</template>

<script lang="ts" setup>
const props = withDefaults(defineProps<{
  content:string;
}>(),{
  content:"",
})
</script>

<style lang="scss" scoped>

</style>