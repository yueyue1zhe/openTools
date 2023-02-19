import {$YTypes} from "@/components/y-ui/index";


declare global {
    interface Uni {
        $y: $YTypes
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
        id :number
        filename:string
        attachment:string
    }

    type ToMediaFunc = (val:string)=>string;
}