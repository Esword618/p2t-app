import { RouteRecordRaw } from "vue-router";
import { PageEnum } from "@/enums/pageEnum";

const importPath = {
  "PageEnum.BASE_TEST_NAME": () => import("@/pages/test/index.vue"),
};

export const testRouters: RouteRecordRaw = {
  path: PageEnum.BASE_TEST,
  name: PageEnum.BASE_TEST_NAME,
  component: importPath["PageEnum.BASE_TEST_NAME"],
  meta: {
    title: "测试",
    keepAlive: false,
  },
  children: [],
};
