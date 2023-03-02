<template>
  <YLayoutDefault></YLayoutDefault>
</template>

<script setup lang="ts">
import { h, resolveComponent } from "vue";
import YAsideSimple from "@/components/yueyue-ui/y-layout/YAsideSimple.vue";
import YHeaderSimple from "@/components/yueyue-ui/y-layout/YHeaderSimple.vue";
import YMainSimple from "@/components/yueyue-ui/y-layout/YMainSimple.vue";
import config from "@/config";

const YLayoutDefault = () => {
  if (config.mode === "self-console") {
    return h(
      resolveComponent("el-container"),
      {
        class: ["layout-container"],
      },
      () => {
        return [
          h(YAsideSimple),
          h(
            resolveComponent("el-container"),
            { class: ["content-wrapper"] },
            () => {
              return [
                h(YHeaderSimple),
                h(YMainSimple, () => h("div", { class: "help-height" })),
              ];
            }
          ),
        ];
      }
    );
  }
  return h(
    "div",
    { class: "help-height-padding" },
    h(resolveComponent("router-view"))
  );
};
</script>

<style lang="scss">
$help-height: calc(50px + 20px + var(--main-space));
.help-height {
  height: $help-height;
}

.help-height-padding {
  padding-top: var(--main-space);
}

.layout-container {
  height: 100%;
  width: 100%;
}

.content-wrapper {
  flex-direction: column;
  width: 100%;
  height: 100%;
  position: relative;
}
</style>
