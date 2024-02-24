<template>
  <div>
    <el-input disabled placeholder="请选择链接" :model-value="out">
      <template #append>
        <div @click="state.show = true" class="y-pointer">选择链接</div>
      </template>
    </el-input>
    <el-dialog
      append-to-body
      title="选择链接"
      :show-close="false"
      v-model="state.show"
      width="50rem"
      class="custom"
    >
      <el-tabs v-model="state.activeName" style="height: 20rem">
        <template v-for="(item, key) in showLinkList" :key="key">
          <el-tab-pane :label="item.title" :name="item.name">
            <template v-for="(v, k) in item.list" :key="k">
              <el-button style="margin: 0.5rem" @click="chooseLink(v.link)">{{
                v.title
              }}</el-button>
            </template>
          </el-tab-pane>
        </template>
        <el-tab-pane v-if="props.useCustom" label="自定义链接" name="custom">
          <div class="flex-def flex-zCenter flex-cCenter" style="height: 10rem">
            <el-input
              style="width: 30rem"
              placeholder="请输入"
              v-model="state.custom"
            >
              <template #append>
                <el-button
                  @click="sureCustom"
                  style="padding: 0 3rem"
                  size="small"
                  >确认</el-button
                >
              </template>
            </el-input>
          </div>
        </el-tab-pane>
      </el-tabs>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { computed, reactive } from "vue";
import { useSettingStore } from "@/stores/setting";
const settingStore = useSettingStore();

const showLinkList = computed(() => {
  return settingStore.link_choose_list;
});

const props = withDefaults(
  defineProps<{
    modelValue: string;
    useCustom?: boolean;
  }>(),
  {
    modelValue: "",
    useCustom: true,
  }
);
const emit = defineEmits<{
  (e: "update:modelValue", out: string): void;
}>();

const out = computed({
  get(): string {
    return props.modelValue;
  },
  set(value: string) {
    emit("update:modelValue", value);
    state.show = false;
  },
});

let state = reactive({
  show: false,
  custom: "",
  activeName: "",
});

if (settingStore.link_choose_list.length > 0) {
  state.activeName = settingStore.link_choose_list[0].name;
} else {
  state.activeName = "custom";
}

const chooseLink = (link: string) => {
  out.value = link;
};
const sureCustom = () => {
  out.value = state.custom;
};
</script>

<style lang="scss">
.custom > .el-dialog__body {
  padding: 0 20px 30px !important;
  margin: 0;
}
</style>
