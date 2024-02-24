/* 加载网络css文件 */
import { useSettingStore } from "@/stores/setting";

export function loadCss(url: string): void {
  const link = document.createElement("link");
  link.rel = "stylesheet";
  link.href = url;
  link.crossOrigin = "anonymous";
  document.getElementsByTagName("head")[0].appendChild(link);
}

/* 加载网络js文件 */
export function loadJs(url: string): void {
  const link = document.createElement("script");
  link.src = url;
  document.body.appendChild(link);
}

//判断是否站外网络路径
export function IsExternal(path: string) {
  return /^(https?|ftp|mailto|tel):/.test(path);
}

export function EnumValue2Label(val: number, raw: BusinessEnum): string {
  const keys = Object.keys(raw);
  let out = "未知";
  keys.forEach((item) => {
    if (raw[item].value === val) {
      out = raw[item].label;
    }
  });
  return out;
}

export function ToMedia(value: string): string {
  if (!value) return "";
  if (value.includes("http")) return value;
  const settingStore = useSettingStore();
  return settingStore.attach_prefix + value;
}

interface linkIcon extends Element {
  rel: string;
  type: string;
  href: string;
}

export function ICOChange(url: string) {
  const link =
    <linkIcon>document.querySelector("link[rel*='icon']") ||
    document.createElement("link");
  link.type = "image/x-icon";
  link.rel = "shortcut icon";
  link.href = url;
  document.getElementsByTagName("head")[0].appendChild(link);
}
