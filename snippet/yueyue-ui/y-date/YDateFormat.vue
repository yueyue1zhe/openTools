<template>
  <span>{{ show.value }}</span>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import { useDateFormat } from "@vueuse/core";

interface dateObjType {
  Time: string;
  Valid: boolean;
}
const props = withDefaults(
  defineProps<{
    date: Date | number | string | dateObjType;
    format?: string;
  }>(),
  {
    date: "",
    format: "YYYY-MM-DD HH:mm:ss",
  }
);
const show = computed(() => {
  let useDateRaw: string | Date | number;
  if (typeof props.date === "object") {
    const sqlNullTime = props.date as dateObjType;
    useDateRaw = sqlNullTime.Valid ? sqlNullTime.Time : "";
  } else {
    useDateRaw = props.date;
  }
  if (!useDateRaw) return "";
  return useDateFormat(useDateRaw, props.format);
});
</script>

<style lang="scss" scoped></style>
