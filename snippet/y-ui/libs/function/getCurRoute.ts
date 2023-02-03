import queryParams from "./queryParams";

const routeTpl = {
    query: {},
    path: "",
    fullPath: ""
}
const getCurRoute = (): typeof routeTpl => {
    let route = {...routeTpl};

    // #ifdef MP-WEIXIN
    let pages = getCurrentPages();
    if (pages.length > 0) {
        let page = pages.pop();
        if (page) {
            route.path = "/" + page.route;
            let out = <AnyObject>page;
            route.fullPath = "/" + page.route + queryParams(out.options);
            route.query = out.options ? out.options : {};
        }
    }
    // #endif

    // #ifdef H5
    let vmR = getApp()?.$route
    route = {
        fullPath: vmR.fullPath,
        path: vmR.path,
        query: vmR.query,
    };
    // #endif

    return route;
}

export default getCurRoute