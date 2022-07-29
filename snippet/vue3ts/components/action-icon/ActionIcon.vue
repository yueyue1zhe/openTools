<template>
  <app-icon
    class="icon-action y-pointer"
    :size="prop.size"
    :name="prop.name"
    :color="prop.color"
    @click="emit('click')"
  ></app-icon>
</template>

<script setup lang="ts">
import AppIcon from "@/components/icon/AppIcon.vue";

interface actionIconPropsType {
  size?: string;
  name: string;
  color?: string;
}

const prop = withDefaults(defineProps<actionIconPropsType>(), {
  size: "16",
  color: "",
});

const emit = defineEmits<{
  (e: "click"): void;
}>();
</script>

<style lang="scss" scoped>
.icon-action {
  color: var(--color-placeholder);
}

.icon-action:hover {
  color: var(--color-text-primary);
}
</style>
