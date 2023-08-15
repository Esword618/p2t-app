import { createApp } from "vue";
import ElementPlus from 'element-plus'
// @ts-ignore
import App from "./App.vue";
import { setupRouter } from "./routes";

import 'element-plus/dist/index.css'
import "./assets/styles/style.less";

function init() {
  const app = createApp(App);
  setupRouter(app);
  app.use(ElementPlus);
  app.mount("#app");
};

void init();
