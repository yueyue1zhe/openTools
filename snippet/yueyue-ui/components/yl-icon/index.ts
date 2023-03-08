import type { App } from "vue";

import * as elIcons from "@element-plus/icons-vue";

export default {
  install: (app: App) => {
    for (const [key, component] of Object.entries(elIcons)) {
      app.component(`el-icon-${key}`, component);
    }
  },
};
