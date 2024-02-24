import { nextTick } from "vue";
import "@/components/yueyue-ui/libs/scss/loading.scss";

export const loading = {
  show: () => {
    const bodyGet: Element = document.body;
    const div = document.createElement("div");
    div.className = "block-loading";
    div.innerHTML = `
            <div class="block-loading-box">
                <div class="block-loading-box-warp">
                    <div class="block-loading-box-item"></div>
                    <div class="block-loading-box-item"></div>
                    <div class="block-loading-box-item"></div>
                    <div class="block-loading-box-item"></div>
                    <div class="block-loading-box-item"></div>
                    <div class="block-loading-box-item"></div>
                    <div class="block-loading-box-item"></div>
                    <div class="block-loading-box-item"></div>
                    <div class="block-loading-box-item"></div>
                </div>
            </div>
        `;
    bodyGet.insertBefore(div, bodyGet.childNodes[0]);
  },
  hide: () => {
    nextTick(() => {
      setTimeout(() => {
        const el = document.querySelector(".block-loading");
        el && el.parentNode?.removeChild(el);
      }, 1000);
    }).catch();
  },
};
