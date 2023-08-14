import { RouteRecordRaw } from "vue-router";
import { PageEnum } from "@/enums/pageEnum";

const importPath = {
    "PageEnum.BASE_ABOUT_NAME": () => import("@/pages/about/index.vue"),
};

export const aboutRouters: RouteRecordRaw = {
    path: PageEnum.BASE_ABOUT,
    name: PageEnum.BASE_ABOUT_NAME,
    component: importPath["PageEnum.BASE_ABOUT_NAME"],
    meta: {
        title: "关于",
        keepAlive: true,
    },
    children: [],
};