import { createApp } from "vue";
import { createPinia } from "pinia";
import App from "./App.vue";
import router from "./router";
import i18n from "./i18n";
import * as ElementPlusIconsVue from "@element-plus/icons-vue";
import ElementPlus from "element-plus";
import VueLatex from "vatex";

import "element-plus/dist/index.css";

// import "./style.scss";

// import "./assets/main.css";

const app = createApp(App);

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component);
}
// 循环注册图标
// Object.keys(ElementPlusIconsVue).forEach(key=>{
//   //elementPlusIcons as any 这里ts学的不咋地 只能先any大法了
//   app.component(key,(ElementPlusIconsVue as any)[key])
//   // 也可这样写
//   //app.component(key,elementPlusIcons[key as keyof typeof elementPlusIcons])
// })


app.use(createPinia());
app.use(router);
app.use(i18n);
app.use(ElementPlus);
app.use(VueLatex);
app.mount("#app");
//https://github.com/Shimada666/VaTex
