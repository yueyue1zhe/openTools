import {toast, toastFastModal} from "./libs/function/toast";
import route from "./libs/function/route";
import toMedia from "./libs/helper/toMedia";
import request from "./libs/request";
import {getLaunchOpt,parseQueryScene} from "./libs/helper/getLaunchOpt";
import queryParams from "./libs/function/queryParams";
import getCurRoute from "./libs/helper/getCurRoute";
import wifi from "./libs/helper/wifi";
import MpAd from "./libs/helper/MpAd";
import throttle from "./libs/function/throttle";
import colorGradient from "./libs/function/colorGradient";
import addUnit from "./libs/function/addUnit";
import canvas from "./libs/helper/canvas";
import lottie from "./libs/helper/lottie";
import imgUtil from "./libs/helper/imgUtil";
import uniCopy from "@/components/y-ui/libs/helper/uni-copy";

const $y = {
    addUnit,
    request,
    route,
    toast,
    toastFastModal,
    throttle,
    colorGradient,
    queryParams,
    copy:uniCopy,


    imgUtil,
    toMedia,
    getLaunchOpt,
    parseQueryScene,
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