import { ElMessage } from "element-plus";
import { useUserStore } from "@/stores/user";
import { YResponseDataType } from "./types";

export function JsonResponseParse<T>(res: YResponseDataType<T>): T | false {
  if (res.message) ElMessage.warning(res.message);
  const userStore = useUserStore();
  switch (res.errno) {
    case 0:
      return res.data;
    case 1:
      return false;
    case 2:
      // utils.toast(res.message, () => {
      //   linkDo(res.data, "nav");
      // });
      return false;
    case 4:
      // utils.toast(res.message, () => {
      //   linkDo(res.data, "redirect");
      // });
      return false;
    case 40019:
      userStore.setToken("");
      userStore.IsLogin().finally();
      return false;
    default:
      console.warn("request fail:", res);
      return false;
  }
}
