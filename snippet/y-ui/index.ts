import toast from "./libs/function/toast";
import route from "./libs/function/route";
import toMedia from "./libs/function/toMedia";
import request from "./libs/request";
import getLaunchQuery from "./libs/function/getLaunchQuery";
import queryParams from "./libs/function/queryParams";
import getCurRoute from "./libs/function/getCurRoute";
import wifi from "./libs/function/wifi";
import MpAd from "./libs/function/MpAd";

const $y = {
    request,
    route,
    toast,
    toMedia,

    getLaunchQuery,

    queryParams,
    getCurRoute,
    wifi,
    MpAd,
}

export type $YTypes = typeof $y;

uni.$y = $y;

export default {
    install: () => {

    }
}