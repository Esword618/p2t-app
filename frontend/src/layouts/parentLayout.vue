<template>
  <div v-if="LayoutMode == 'column'" style="height: 100%;display: flex;flex-direction: column;">
    <!-- 异步组件 -->
    <suspense>
      <!-- 自定义titlebar -->
      <title-bar ref="header" :WindowMinimise="WindowMinimise" :WindowMaximise="WindowMaximise" :WindowClose="WindowClose"
        :WindowOnTop="WindowOnTop" :WindowIsMaximised="WindowIsMaximised" :WindowIsOnToped="WindowIsOnToped"
        :WindowIsFullScreen="WindowIsFullScreen" :style="{
          height: '32px',
          color: 'var(--text-primary)',
          backgroundColor: 'var(--background-secondary)',
          borderBottomColor: 'var(--border-primary)',
          control: {
            fill: 'var(--text-primary)',
            hoverColorExcludeClose: 'var(--button-hover)',
            hoverStyleOnlyClose: {
              backgroundColor: 'rgb(228, 79, 79)',
              fill: '#fff',
            },
          }
        }">
        <template #extend-left>
          <div style="display:flex;height: 100%;justify-content: center;align-items: center;">
            <img :src="appConfig.icon" style="height: 100%;padding: 5px;">
            <span>{{ appConfig.name }}</span>
          </div>
        </template>
        <template #extend-center>
          <div style="height: 100%;display: flex;justify-content: center;align-items: center;">
            <!-- <span>双击Loding界面即可撤去Loading</span> -->
          </div>
        </template>
        <template #extend-right>
          <toggle-theme class="toggle-theme" :size="18" :duration="200" />
        </template>
      </title-bar>
    </suspense>
    <section class="page-container" :class="{ 'loading': isLoading }" @dblclick="cancelLoading"
      style="--color-link: var(--text-link);--color-bg: var(--text-invert);height: calc(100% - 32px);">
      <a-side-bar-new>
        <toggle-theme style="width: 40px; height: 40px;  margin-left: auto; margin-right: auto; margin-top: 8.75px;"
          :size="20" :duration="200" />
      </a-side-bar-new>
      <main ref="smoothDom" class="page-content ">
        <router-view />
      </main>
    </section>
  </div>
  <div v-else style="height: 100%;display: flex;flex-direction: row;">
    <a-side-bar-new>
      <template #extend-bottom>
        <toggle-theme style="width: 40px; height: 40px;  margin-left: auto; margin-right: auto; margin-top: 8.75px;"
          :size="20" :duration="200" />
      </template>
    </a-side-bar-new>
    <div style="width: 100%;display: flex;flex-direction: column;">
      <suspense>
        <!-- 自定义titlebar -->
        <title-bar ref="header" :WindowMinimise="WindowMinimise" :WindowMaximise="WindowMaximise"
          :WindowClose="WindowClose" :WindowOnTop="WindowOnTop" :WindowIsMaximised="WindowIsMaximised"
          :WindowIsOnToped="WindowIsOnToped" :WindowIsFullScreen="WindowIsFullScreen" :style="{
            height: '32px',
            color: 'var(--text-primary)',
            backgroundColor: 'var(--background-secondary)',
            borderBottomColor: 'var(--border-primary)',
            control: {
              fill: 'var(--text-primary)',
              hoverColorExcludeClose: 'var(--button-hover)',
              hoverStyleOnlyClose: {
                backgroundColor: 'rgb(228, 79, 79)',
                fill: '#fff',
              },
            }
          }">
          <!-- <template #extend-left>
            <div style="display:flex;height: 100%;justify-content: center;align-items: center;">
              <img :src="appConfig.icon" style="height: 100%;padding: 5px;">
              <span>{{ appConfig.name }}</span>
            </div>
          </template> -->
          <template #extend-center>
            <div style="height: 100%;display: flex;justify-content: center;align-items: center;">
              <!-- <span>双击Loding界面即可撤去Loading</span> -->
            </div>
          </template>
          <template #extend-right>
            <toggle-theme class="toggle-theme" :size="18" :duration="200" />
          </template>
        </title-bar>
      </suspense>
      <section class="page-container" :class="{ 'loading': isLoading }" @dblclick="cancelLoading"
        style="--color-link: var(--text-link);--color-bg: var(--text-invert);height: calc(100% - 32px);">
        <main ref="smoothDom" class="page-content ">
          <router-view />
        </main>
      </section>

    </div>
  </div>
</template>

<script lang="ts" setup>
// @ts-nocheck
import TitleBar from "@/components/window/titlebar/index.vue";
import ASideBarNew from "@/components/window/asidebar/index.vue";

import ToggleTheme from "@/components/button/ToggleTheme.vue";
import { WindowMinimise, WindowMaximise, WindowClose, WindowOnTop, WindowIsMaximised, WindowIsOnToped, WindowIsFullScreen } from "@wailsjs/go/main/App";

import { appConfig } from "@/configs/app.ts";
import { onMounted, ref } from "vue";
// @ts-ignore
import { SmoothScroller, initSmoothScrolling } from "@/utils/page/smoothScroll";

defineProps({
  LayoutMode: {
    type: String as () => ("column" | "row"),
    default: "column",
  },
})

const isLoading = ref(false);
//@ts-ignore
const switchLoading = () => {
  isLoading.value = !isLoading.value;
};
const cancelLoading = () => {
  isLoading.value = false;
};

const smoothDom = ref<HTMLElement>();

// @ts-ignore 防抖
function debounce(fn: Function, delay: number) {
  let timer: any = null;
  return function () {
    if (timer) {
      clearTimeout(timer);
    }
    timer = setTimeout(fn, delay);
  };
}
// @ts-ignore 节流
function throttle(fn: Function, delay: number) {
  let timer: any = null;
  return function () {
    if (!timer) {
      timer = setTimeout(() => {
        fn();
        timer = null;
      }, delay);
    }
  };
}

onMounted(() => {
  // if (smoothDom.value) {
  //   initSmoothScrolling(smoothDom.value);
  // }
});

</script>


<style>
#app {
  height: 100vh;
  width: 100vw;
  display: flex;
  flex-direction: column;
}
</style>
<style lang="less" scoped>
.toggle-theme {
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0 11.5px;
  box-sizing: content-box;
}

.page-container {
  flex-grow: 1;
  display: flex;
  position: relative;
  width: 100%;

  .page-content {
    position: relative;
    flex-grow: 1;
    height: 100%;
    overflow-y: auto;
    // 主要是针对路由切换动画时出现的横向滚动条
    overflow-x: hidden;
  }
}
</style>


