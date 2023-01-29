import {$YTypes} from "@/components/y-ui/index";


declare global {
    interface Uni{
        $y:$YTypes
    }

    type GlobalToken = {
        data: string,
        time: number
    }

    type VoidCallBack = (res?: any) => void;


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
}