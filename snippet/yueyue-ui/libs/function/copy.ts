import { useClipboard } from "@vueuse/core";
import { ElMessage } from "element-plus";

const { copy, isSupported } = useClipboard();

export default function (str: string) {
  if (!isSupported) {
    return ElMessage.error("您的浏览器不支持Clipboard API");
  }
  copy(str)
    .then(() => {
      ElMessage.success("复制成功");
    })
    .catch((err) => {
      console.warn(err);
    });
}
