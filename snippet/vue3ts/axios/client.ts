import axios, {
  AxiosError,
  AxiosInstance,
  AxiosInterceptorManager,
  AxiosPromise,
  AxiosRequestConfig,
  AxiosResponse,
} from "axios";
import { ElMessage } from "element-plus/es";
import { JsonResponseParse } from "@/common/req/response";

interface NewAxiosInstance extends AxiosInstance {
  //设置泛型T，将请求后的结果返回变成AxiosPromise<T>
  <T>(config: AxiosRequestConfig): AxiosPromise<T>;

  interceptors: {
    request: AxiosInterceptorManager<AxiosRequestConfig>;
    response: AxiosInterceptorManager<AxiosResponse>;
  };
}

const axiosInstance: NewAxiosInstance = axios.create({
  baseURL:
    window.microApp?.getData().baseURL || "http://127.0.0.1:4523/mock/1231005", //TODO::非微前端环境下 || 如何获取请求基础地址
  timeout: 15000, //request timeout
  withCredentials: true, //设置跨域请求
});

// axios实例拦截请求
axiosInstance.interceptors.request.use((config: AxiosRequestConfig) => {
  if (config?.headers) config.headers["y-token"] = "123token"; //TODO::读取store中的token
  return config;
});

// axios实例拦截响应
axiosInstance.interceptors.response.use(
  (response) => Promise.resolve(response),
  (error: AxiosError) => {
    if (!error.response?.data) ElMessage.warning("网络连接异常,请稍后再试!");
    return Promise.reject(error.response?.data);
  }
);

export function PostJson<T>(url: string, data?: AnyObject) {
  return request<T>({
    url: url,
    data: data,
    method: "POST",
    responseType: "json",
    headers: {
      "Content-Type": "application/json",
    },
  });
}

function request<T>(obj: AxiosRequestConfig & { url: string }) {
  return new Promise<T>((resolve, reject) => {
    axiosInstance<YResponseDataType<T>>({
      url: obj.url,
      data: obj.data || {},
      method: obj.method || "GET",
      responseType: obj.responseType || "json",
    })
      .then((res) => {
        const result = JsonResponseParse<T>(res.data);
        result ? resolve(result) : reject(res);
      })
      .catch((error) => {
        reject(error);
      });
  });
}
