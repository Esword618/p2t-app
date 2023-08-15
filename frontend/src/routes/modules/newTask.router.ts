import { RouteRecordRaw } from "vue-router";
import { PageEnum } from "@/enums/pageEnum";

const importPath = {
    "PageEnum.BASE_NEW_TASK_NAME": () => import("@/pages/newTask/index.vue"),
};

export const newTaskRouters: RouteRecordRaw = {
    path: PageEnum.BASE_NEW_TASK,
    name: PageEnum.BASE_NEW_TASK_NAME,
    component: importPath["PageEnum.BASE_NEW_TASK_NAME"],
    meta: {
        title: "添加任务",
        keepAlive: true,
    },
    children: [],
};
