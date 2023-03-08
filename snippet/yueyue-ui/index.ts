import { App } from "vue";
import YIcon from "@/components/yueyue-ui/components/yl-icon";
import YEditor from "@/components/yueyue-ui/components/yl-editor";

export default {
  install: (app: App) => {
    app.use(YIcon);
    app.use(YEditor);
  },
};
