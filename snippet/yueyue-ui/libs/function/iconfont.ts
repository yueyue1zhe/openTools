import { nextTick } from "vue";
import { loadCss, loadJs } from "./common";
import * as elIcons from "@element-plus/icons-vue";

const cssUrls: Array<string> = [
  "//netdna.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css",
];
const jsUrls: Array<string> = [];

export default function init() {
  if (cssUrls.length > 0) {
    cssUrls.map((v) => {
      loadCss(v);
    });
  }

  if (jsUrls.length > 0) {
    jsUrls.map((v) => {
      loadJs(v);
    });
  }
}

/*
 * 获取当前页面中从指定域名加载到的样式表，样式表未载入前无法获取
 */
function getStylesFromDomain(domain: string) {
  const sheets = [];
  const styles: StyleSheetList = document.styleSheets;
  for (const key in styles) {
    if (styles[key].href && (styles[key].href as string).indexOf(domain) > -1) {
      sheets.push(styles[key]);
    }
  }
  return sheets;
}

/*
 * 获取本地自带的图标
 * /src/assets/icons文件夹内的svg文件
 */
export function getLocalIconfontNames() {
  return new Promise<string[]>((resolve, reject) => {
    nextTick(() => {
      let iconfontArr: string[] = [];

      const svgEl = document.getElementById("local-icon");
      if (svgEl?.dataset.iconName) {
        iconfontArr = (svgEl?.dataset.iconName as string).split(",");
      }

      if (iconfontArr.length > 0) {
        resolve(iconfontArr);
      } else {
        reject("No ElementPlus Icons");
      }
    }).finally();
  });
}

/*
 * 获取 Awesome-Iconfont 的 name 列表
 */
export function getAwesomeIconfontNames() {
  return new Promise<string[]>((resolve, reject) => {
    nextTick(() => {
      const iconfontArr = [];
      const sheets = getStylesFromDomain(
        "cdn.bootcdn.net/ajax/libs/font-awesome/"
      );
      for (const key in sheets) {
        const rules: any = sheets[key].cssRules; // eslint-disable-line
        for (const k in rules) {
          if (
            rules[k].selectorText &&
            /^\.fa-(.*)::before$/g.test(rules[k].selectorText)
          ) {
            if (rules[k].selectorText.indexOf(", ") > -1) {
              const iconNames = rules[k].selectorText.split(", ");
              /*
              // 含图标别名
              for (const i_k in iconNames) {
                  iconfonts.push(`${iconNames[i_k].substring(1, iconNames[i_k].length).replace(/\:\:before/gi, '')}`)
              } */
              iconfontArr.push(
                `${iconNames[0]
                  .substring(1, iconNames[0].length)
                  .replace(/::before/gi, "")}` // eslint-disable-line
              );
            } else {
              iconfontArr.push(
                `${rules[k].selectorText
                  .substring(1, rules[k].selectorText.length)
                  .replace(/::before/gi, "")}` // eslint-disable-line
              );
            }
          }
        }
      }

      if (iconfontArr.length > 0) {
        resolve(iconfontArr);
      } else {
        reject("No AwesomeIcon style sheet");
      }
    }).finally();
  });
}

/*
 * 获取 Iconfont 的 name 列表
 */
export function getIconfontNames() {
  return new Promise<string[]>((resolve, reject) => {
    nextTick(() => {
      const iconfonts = [];
      const sheets = getStylesFromDomain("at.alicdn.com");
      for (const key in sheets) {
        const rules: any = sheets[key].cssRules; // eslint-disable-line
        for (const k in rules) {
          if (
            rules[k].selectorText &&
            /^\.icon-(.*)::before$/g.test(rules[k].selectorText)
          ) {
            iconfonts.push(
              `${rules[k].selectorText
                .substring(1, rules[k].selectorText.length)
                .replace(/::before/gi, "")}`
            );
          }
        }
      }

      if (iconfonts.length > 0) {
        resolve(iconfonts);
      } else {
        reject("No Iconfont style sheet");
      }
    }).finally();
  });
}

/*
 * 获取element plus 自带的图标
 */
export function getElementPlusIconfontNames() {
  return new Promise<string[]>((resolve, reject) => {
    nextTick(() => {
      const iconfontArr = [];
      for (const [key] of Object.entries(elIcons)) {
        iconfontArr.push(`el-icon-${key}`);
      }
      if (iconfontArr.length > 0) {
        resolve(iconfontArr);
      } else {
        reject("No ElementPlus Icons");
      }
    }).finally();
  });
}
