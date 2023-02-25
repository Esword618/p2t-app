import {
  createRouter,
  createWebHashHistory,
  type RouteRecordRaw,
} from "vue-router";

import upload from "@/views/upload/upload.vue";
import management from "@/views/management/management.vue";
import tutorials from "@/views/tutorials/tutorials.vue";
import settings from "@/views/settings/settings.vue";

// 2. 配置路由
const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "图片上传",
    redirect: {
      name: "upload",
    },
  },
  {
    path: "/upload",
    name: "upload",
    component: upload,
    meta: {
      title: `图片上传`,
    },
  },
  {
    path: "/management",
    name: "management",
    component: management,
    meta: {
      title: `图片管理`,
    },
  },
  {
    path: "/tutorials",
    name: "tutorials",
    component: tutorials,
    meta: {
      title: `使用教程`,
    },
  },
  {
    path: "/settings",
    name: "settings",
    component: settings,
    meta: {
      title: `我的设置`,
    },
  },
  {
    path: "/about",
    name: "about",
    // route level code-splitting
    // this generates a separate chunk (About.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import("@/views/about/about.vue"),
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
});

export default router;
