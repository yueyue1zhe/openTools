//微前端
export declare global {
  interface Window {
    __MICRO_APP_BASE_ROUTE__: never;
    __MICRO_APP_ENVIRONMENT__: boolean;
    __MICRO_APP_PUBLIC_PATH__: string;
    __MICRO_APP_NAME__: string;
    __MICRO_APP_BASE_APPLICATION__: boolean;

    fullLoadingExist: boolean;

    microApp?: microApp;
  }
}

//微擎基座应用结构
interface microApp {
  getData?: () => microAppData;

  addDataListener(cb: (data: { path: string }) => void);

  dispatch(opt: microAppDispatchOption): void;
}

type microAppDispatchOption = {
  refresh?: boolean;
  menu?: string;
};

interface microAppData {
  path: string; // 路由 监听响应跳转
  w7: microAppDataW7;
  baseURL: string; //请求baseurl
}

interface microAppDataW7 {
  setStorage({ key: string, value: string }): void;

  getStorage(key: string): string;

  removeStorage(key: string): void;

  clearStorage(): void;

  login(): Promise.resolve<{ code: string }>; // 登陆code
  getModuleInfo(): { module_name: string; module_version: string };
}
