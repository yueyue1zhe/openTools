import config from "@/config";

export function GetBaseUrl(): string {
  return window?.microApp?.getData
    ? window.microApp.getData().baseURL
    : config.dev_base;
}

/*
{
    "baseURL": "https://sh-dev-c.yueyuebang.cn",
    "path": "",
    "w7": {},
    "key": 347114,
    "site_name": "测试go_38",
    "sitename": "测试go_38",
    "family": "v",
    "version": "1.0.3",
    "can_license_above_three": false,
    "is_star": false,
    "is_installed": true,
    "is_install": true,
    "is_tcb": true,
    "is_not_app": true,
    "logo": "https://cdn.w7.cc/images/2020/03/14/NFyyBIOk25m23vPmDokDb9PyaAsNShJHDltieQ0V.jpeg",
    "family_txt": "普通版",
    "family_text": "普通版",
    "is_w7": false,
    "is_new_version": true,
    "is_offline_w7": false,
    "is_owner": true,
    "storage_key": "bee3f70936bae88be813139aa4cac0b6",
    "ip": "",
    "mark_token": "db859******83b1766",
    "is_sass": false,
    "username": "岳岳小组",
    "html": "https://cdn.w7.cc/web-app/test_go/t1.0.7/test_go/index.html",
    "webapp_url": "https://cdn.w7.cc/web-app/test_go/t1.0.7/test_go/index.html",
    "module_name": "test_go",
    "module_version": "1.0.7",
    "goods_url": "https://market.w7.cc/detail?id=0&type=0",
    "menu": [
        {
            "displayorder": "9",
            "is_default": 2,
            "do": "/avatar/theme/manage",
            "title": "头像专题管理",
            "icon": "caozuotai"
        },
        {
            "displayorder": "8",
            "is_default": 1,
            "do": "/avatar/client/wechat-mini",
            "title": "微信小程序对接",
            "icon": "goods"
        },
        {
            "displayorder": "7",
            "is_default": 1,
            "do": "/founder/system/attach",
            "title": "附件配置",
            "icon": "manage-menu"
        }
    ],
    "framework": "vue2",
    "has_web": true,
    "can_offline": false,
    "has_multi_role": true,
    "has_mini_program": false,
    "has_index": true,
    "plugins": [],
    "title": "测试go",
    "healthy_check_url": "",
    "role": "founder",
    "has_direct": false,
    "defaultdo": "/avatar/theme/manage"
}



 */
