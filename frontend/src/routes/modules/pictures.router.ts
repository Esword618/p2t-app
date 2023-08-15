import { RouteRecordRaw } from "vue-router";
import { PageEnum } from "@/enums/pageEnum";

const importPath = {
    "PageEnum.BASE_PICTURES_NAME": () => import("@/pages/pictures/index.vue"),
};

export const picturesRouters: RouteRecordRaw = {
    path: PageEnum.BASE_PICTURES,
    name: PageEnum.BASE_PICTURES_NAME,
    component: importPath["PageEnum.BASE_PICTURES_NAME"],
    meta: {
        title: "图库",
        keepAlive: true,
    },
    children: [],
};