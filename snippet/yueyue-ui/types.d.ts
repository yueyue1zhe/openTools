namespace YPopupEasyFormTypes {
  interface OptsItemType {
    name: string;
    nameDeep?: boolean; //适用于嵌套对象
    type: number;
    label?: string;
    placeholder?: string;
    required?: boolean;

    powerSortOpts?: {
      powerName: string;
      sortName: string;
    };

    defaultHide?: boolean; //默认隐藏
    showCond?: (editForm: AnyObject) => boolean;

    radioOpts?: {
      [key: string]: {
        value: number | string;
        label: string;
      };
    };
  }
}
