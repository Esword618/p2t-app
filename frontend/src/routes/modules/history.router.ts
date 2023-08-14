import { RouteRecordRaw } from "vue-router";
import { PageEnum } from "@/enums/pageEnum";

const importPath = {
    "PageEnum.BASE_HISTORY_NAME": () => import("@/pages/history/index.vue"),
};

export const historyRouters: RouteRecordRaw = {
    path: PageEnum.BASE_HISTORY,
    name: PageEnum.BASE_HISTORY_NAME,
    component: importPath["PageEnum.BASE_HISTORY_NAME"],
    meta: {
        title: "历史记录",
        keepAlive: true,
    },
    children: [],
};