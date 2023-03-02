import { App } from "vue";
import YIcon from "@/components/yueyue-ui/y-icon";
import YEditor from "@/components/yueyue-ui/y-editor";

export default {
  install: (app: App) => {
    app.use(YIcon);
    app.use(YEditor);
  },
};
