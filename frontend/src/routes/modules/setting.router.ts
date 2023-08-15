import { RouteRecordRaw } from "vue-router";
import { PageEnum } from "@/enums/pageEnum";

const importPath = {
  "PageEnum.BASE_SETTING_NAME": () => import("@/pages/setting/index.vue"),
};

export const settingRouters: RouteRecordRaw = {
  path: PageEnum.BASE_SETTING,
  name: PageEnum.BASE_SETTING_NAME,
  component: importPath["PageEnum.BASE_SETTING_NAME"],
  meta: {
    title: "设置",
    keepAlive: true,
  },
  children: [],
};
