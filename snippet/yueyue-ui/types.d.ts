interface AnyObject {
    [key: string]: any; // eslint-disable-line
}

/*
引用json
tsconfig.json 中 resolveJsonModule 应配置为 true
 */
declare module "*.json" {
    const value: any; // eslint-disable-line
  export default value;
}

interface PageResult<T> {
  list: T[];
  page: number;
  total: number;
  size: number;
}

type VoidCallBack<T = any> = (res?: T) => void;

interface ApiToken {
  data: string;
  time: number;
}

interface BusinessEnum {
  [key: string]: {
    value: number;
    label: string;
  };
}

interface HTMLLinkElement {
  rel: string;
}
