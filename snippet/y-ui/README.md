# main.ts 中注册
    import YUi from "@/components/y-ui";

    app.use(YUi);
# uni.scss 中引入
    @import 'src/components/y-ui/libs/scss/theme.scss';
# App.vue 中引入
    @import 'src/components/y-ui/libs/scss/index.scss';
    page{
        background-color: $y-bg-color;
    }