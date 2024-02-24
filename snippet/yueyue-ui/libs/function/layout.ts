import type { CSSProperties } from "vue";

/**
 * main高度
 * @param extra main高度额外减去的px数,可以实现隐藏原有的滚动条
 * @returns CSSProperties
 */
export function mainHeight(extra = 0): CSSProperties {
  return {
    height: "calc(100vh - " + extra.toString() + "px)",
  };
}
