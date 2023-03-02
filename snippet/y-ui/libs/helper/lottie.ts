import lottie from "lottie-miniprogram";
import {getCurrentInstance} from "vue";


function load(el: string,url:string, opt: LottieLoadOpts = {
    loop: true,
    autoplay: true
}): Promise<LoadAnimationReturnType> {
    // LoadAnimationReturnType
    return new Promise<LoadAnimationReturnType>(resolve => {
        uni.createSelectorQuery().in(getCurrentInstance()).select(el).node(res => {
            const canvas = res.node;
            const context = canvas.getContext('2d');
            lottie.setup(canvas);
            resolve(lottie.loadAnimation({
                loop: opt.loop,
                autoplay: opt.autoplay,
                path: url,
                rendererSettings: {
                    context: context,
                },
            }));
        }).exec();
    })

}

export default {
    load,
}