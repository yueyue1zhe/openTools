
interface WriteWordsOptions {
    word:string;
    maxWidth:number;
    fontSize:number;
    ctx:Uni.CanvasContext;
    maxLine:number;
    x:number;
    y:number;
}

interface LottieLoadOpts {
    loop?:boolean,
    autoplay?:boolean,
}



interface LoadAnimationReturnType {
    play(): void;
    stop(): void;
    pause(): void;
    setSpeed(speed: number): void;
    goToAndPlay(value: number, isFrame?: boolean): void;
    goToAndStop(value: number, isFrame?: boolean): void;
    setDirection(direction: AnimationDirection): void;
    playSegments(segments: AnimationSegment | AnimationSegment[], forceFlag?: boolean): void;
    setSubframe(useSubFrames: boolean): void;
    destroy(): void;
    getDuration(inFrames?: boolean): number;
    triggerEvent<T = any>(name: AnimationEventName, args: T): void;
    addEventListener<T = any>(name: AnimationEventName, callback: AnimationEventCallback<T>): void;
    removeEventListener<T = any>(name: AnimationEventName, callback: AnimationEventCallback<T>): void;
}