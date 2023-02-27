import {toast, toastFastModal} from "./libs/function/toast";
import route from "./libs/function/route";
import toMedia from "./libs/helper/toMedia";
import request from "./libs/request";
import getLaunchQuery from "./libs/helper/getLaunchQuery";
import queryParams from "./libs/function/queryParams";
import getCurRoute from "./libs/helper/getCurRoute";
import wifi from "./libs/helper/wifi";
import MpAd from "./libs/helper/MpAd";
import throttle from "./libs/function/throttle";
import colorGradient from "./libs/function/colorGradient";
import addUnit from "./libs/function/addUnit";
import canvas from "./libs/helper/canvas";
import lottie from "./libs/helper/lottie";
import imgUtil from "@/components/y-ui/libs/helper/imgUtil";

const $y = {
    addUnit,
    request,
    route,
    toast,
    toastFastModal,
    throttle,
    colorGradient,
    queryParams,


    imgUtil,
    toMedia,
    getLaunchQuery,
    getCurRoute,
    wifi,
    MpAd,
    canvas,

    lottie,
}

export type $YTypes = typeof $y;

uni.$y = $y;

export default {
    install: () => {

    }
}