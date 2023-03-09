import router from "@/router";
import type { RouteLocationRaw, Router } from "vue-router";
import { useUserStore } from "@/stores/user";
import NProgress from "nprogress";
import "nprogress/nprogress.css";
import { loading } from "@/components/yueyue-ui/libs/function/loading";
import W7helper from "@/components/yueyue-ui/libs/helper/W7helper";

export function UseRouterPush(path: string) {
  if (window.microApp) {
    window.microApp.dispatch({ menu: window.__MICRO_APP_BASE_ROUTE__ + path });
    return;
  }
  router.push({ path: path }).catch((e) => console.warn(e));
}

export function YUseRouterPush(to: RouteLocationRaw) {
  router.push(to).catch((e) => console.info(e));
}

export function routerSafe(
  router: Router,
  ignore: string[],
  loginPageName: string,
  indexPageName: string
) {
  router.beforeEach(async (to, from, next) => {
    //进度条启动
    NProgress.configure({ showSpinner: false });
    NProgress.start();
    if (!window.fullLoadingExist) {
      loading.show();
      window.fullLoadingExist = true;
    }

    //页面标题
    document.title = to.meta?.title ? to.meta?.title : "";

    //页面登录状态维护
    const userStore = useUserStore();
    if (W7helper.CanUse()) {
      await userStore.IsLogin();
    } else {
      if (!to.name || !ignore.includes(to.name.toString())) {
        if (!(await userStore.IsLogin()))
          return next({
            name: loginPageName,
            query: { redirect: to.fullPath },
          });
      }
      if (
        to.name &&
        ignore.includes(to.name.toString()) &&
        (await userStore.IsLogin())
      ) {
        return next({
          name: indexPageName,
        });
      }
    }
    next();
  });

  router.afterEach(() => {
    //进度条结束
    NProgress.done();
    if (window.fullLoadingExist) loading.hide();
  });
}
