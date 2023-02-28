import {$YTypes} from "@/components/y-ui/index";


declare global {
    interface Uni {
        $y: $YTypes
    }

    interface NodeInfo {
        width: number;
        height: number;
    }

    type GlobalToken = {
        data: string,
        time: number
    }

    type YCallBack<T = any> = (res?: T) => void;


    interface RequestResponse {
        data: any,
        errno: number,
        message: string
    }

    interface PageResult<T> {
        list: T[];
        page: number;
        total: number;
        size: number;
    }

    interface AttachResult {
        id: number
        filename: string
        attachment: string
    }

    type ToMediaFunc = (val: string) => string;

    namespace YEasyFormTypes {
        interface OptsItemType {
            name: string;
            type: number;
            label?: string;
            placeholder?: string;
            required?: boolean;

            defaultHide?: boolean;//默认隐藏
            showCond?: (editForm: AnyObject) => boolean;

            uploadImageOption?: UploadImageOption;
            radioOpts?: RadioOptsItem[];
        }

        interface UploadImageOption {
            toMediaFunc: ToMediaFunc
            actionFunc: () => Promise<AttachResult>
        }

        interface RadioOptsItem {
            text: string,
            value: number
        }
    }
}