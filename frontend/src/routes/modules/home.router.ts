import { RouteRecordRaw } from "vue-router";
import { PageEnum } from "@/enums/pageEnum";

const importPath = {
  "PageEnum.BASE_HOME_NAME": () => import("@/pages/home/index.vue"),
};

export const homeRouters: RouteRecordRaw = {
  path: PageEnum.BASE_HOME,
  name: PageEnum.BASE_HOME_NAME,
  component: importPath["PageEnum.BASE_HOME_NAME"],
  meta: {
    title: "主页",
    keepAlive: true,
  },
  children: [],
};
