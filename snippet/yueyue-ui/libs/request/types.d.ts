import {
  AxiosInstance,
  AxiosInterceptorManager,
  AxiosPromise,
  AxiosRequestConfig,
  AxiosResponse,
} from "axios";

interface YResponseDataType<T> {
  errno: number;
  message: string;
  data: T;
}

interface NewAxiosInstance extends AxiosInstance {
  //设置泛型T，将请求后的结果返回变成AxiosPromise<T>
  <T>(config: AxiosRequestConfig): AxiosPromise<T>;

  interceptors: {
    request: AxiosInterceptorManager<AxiosRequestConfig>;
    response: AxiosInterceptorManager<AxiosResponse>;
  };
}
