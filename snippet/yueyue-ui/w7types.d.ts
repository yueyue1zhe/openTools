//微前端
export declare global {
    interface Window {
        __MICRO_APP_BASE_ROUTE__: never;
        __MICRO_APP_ENVIRONMENT__: boolean;
        __MICRO_APP_PUBLIC_PATH__: string;
        __MICRO_APP_NAME__: string;
        __MICRO_APP_BASE_APPLICATION__: boolean;

        fullLoadingExist: boolean;

        microApp: W7UtilTypes.microApp;
    }

    namespace W7UtilTypes {
        //微擎基座应用结构
        interface microApp {
            getData: () => GetDataTypes;

            addDataListener(cb: (data: AddDataListenerCbData) => void);

            dispatch(opt: microAppDispatchOption): void;
        }

        interface AddDataListenerCbData {
            path: string
            type: string
            data:AnyObject
        }

        type DispatchOption = {
            refresh?: boolean;
            menu?: string;
        };


        interface GetDataTypes {
            path: string; // 路由 监听响应跳转
            w7: W7Types;
            baseURL: string; //请求baseurl
        }

        interface W7Types {
            setStorage({key: string, value: string}): void;

            getStorage(key: string): string;

            removeStorage(key: string): void;

            clearStorage(): void;

            login(): Promise.resolve<{ code: string }>; // 登陆code
            getModuleInfo(): { module_name: string; module_version: string };

            changeMenuCollapse(state: boolean): void

            pay(ticket: string, callback: VoidCallBack): void

            jsTicket(): string

            realname(): W7RealNameResult

            navigate(opt: W7NavigateOpt): void

            closeCard(): void
        }

        interface W7RealNameResult {
            is_certification: boolean//是否实名认证
            name: string//公司名或用户名
        }

        interface W7NavigateOpt {
            sitekey?: string
            modulename: string
            type: "micro" | "iframe"
            route: string
            style?: {
                width: number,
                height: number,
            }
        }
    }
}






