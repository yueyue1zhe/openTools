import type { AxiosRequestConfig } from "axios";
import axios from "axios";
import { ElMessage } from "element-plus";
import { JsonResponseParse } from "./response";
import config from "@/config";
import { useUserStore } from "@/stores/user";
import { GetBaseUrl } from "./common";
import { NewAxiosInstance, YResponseDataType } from "./types";

const axiosInstance: NewAxiosInstance = axios.create({
  baseURL: GetBaseUrl(),
  timeout: 15000, //request timeout
  withCredentials: true, //设置跨域请求
});

// axios实例拦截请求
axiosInstance.interceptors.request.use((conf: AxiosRequestConfig) => {
  const userStore = useUserStore();
  if (conf?.headers) conf.headers[config.JWTTokenKey] = userStore.token.data;
  return conf;
});

// axios实例拦截响应
axiosInstance.interceptors.response.use(
  (response) => Promise.resolve(response),
  (error) => {
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

export function GetJson<T>(url: string) {
  return request<T>({
    url: url,
    method: "GET",
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
        result ? resolve(result) : reject(res.data);
      })
      .catch((error) => {
        reject(error);
      });
  });
}
