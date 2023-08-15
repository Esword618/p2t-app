<template>
  <on-top :WindowOnTop="WindowOnTop" :WindowIsOnToped="WindowIsOnToped" :style="{
    fill: props.style?.fill || '#000',
    activeColor: props.style?.onTopActiveColor,
    hoverColor: props.style?.hoverColorExcludeClose || 'rgb(255, 255, 255)',
  }" />
  <div title="最小化" class="frameless-titlebar-button frameless-titlebar-minimize" @click="minimizeWindow">
    <svg width="10" height="10">
      <path v-for="(d, _index) in minimizePath" :d="d" />
    </svg>
  </div>
  <div title="最大化" class="frameless-titlebar-button frameless-titlebar-toggle" @click="maximiseWindow">
    <svg width="10" height="10">
      <path v-for="(d, _index) in (isMaximised ? restorePath_win11 : maximizePath_win11)" :d="d" />
    </svg>
  </div>
  <div title="关闭" class="frameless-titlebar-button frameless-titlebar-close" @click="closeWindow">
    <svg width="10" height="10">
      <path v-for="(d, _index) in closePath" :d="d" />
    </svg>
  </div>
</template>

<script setup lang="ts">
// @ts-ignore
import { minimizePath, maximizePath, closePath, restorePath, restorePath_win11, maximizePath_win11 } from "./assets";
import { PropType, ref } from "vue";
import OnTop from "./items/OnTop.vue";

const props = defineProps({
  WindowOnTop: {
    type: Function as PropType<() => Promise<void> | void>,
    required: true
  },
  WindowMinimise: {
    type: Function as PropType<() => Promise<void> | void>,
    required: true
  },
  WindowMaximise: {
    type: Function as PropType<() => Promise<void> | void>,
    required: true
  },
  WindowClose: {
    type: Function as PropType<(isHidden: boolean) => Promise<void> | void>,
    required: true
  },
  WindowIsMaximised: {
    type: Function as PropType<() => Promise<boolean> | boolean>,
    required: true
  },
  WindowIsOnToped: {
    type: Function as PropType<() => Promise<boolean> | boolean>,
    required: true
  },
  style: {
    type: Object as PropType<{
      fill?: string;
      hoverColorExcludeClose?: string;
      hoverStyleOnlyClose?: {
        backgroundColor?: string;
        fill?: string;
      };
      onTopActiveColor?: string;
    }>,
    required: false,
  }
})

async function minimizeWindow() {
  await props.WindowMinimise()
}

async function closeWindow() {
  // wails 运行时关闭程序 传参true表示隐藏窗口而不关闭
  await props.WindowClose(false)
}

async function maximiseWindow() {
  await props.WindowMaximise()
  isMaximised.value = !isMaximised.value
}

const isMaximised = ref(await props.WindowIsMaximised())

window.onresize = async (_event) => {
  isMaximised.value = await props.WindowIsMaximised()
}
</script>

<style scoped lang="less">
.frameless-titlebar-button {
  position: relative;
  display: block;
  width: 46px;
  height: 100%;
  fill: v-bind("props.style?.fill || '#000'");
  cursor: pointer;

  svg {
    position: absolute;
    display: inline-flex;
    top: 50%;
    left: 50%;
    transform: translateX(-50%) translateY(-50%);
  }

  &.frameless-titlebar-close:hover {
    background-color: v-bind("props.style?.hoverStyleOnlyClose?.backgroundColor || 'rgb(228, 79, 79)'");

    svg {
      fill: v-bind("props.style?.hoverStyleOnlyClose?.fill || '#fff'");
    }
  }

  &.frameless-titlebar-minimize:hover,
  &.frameless-titlebar-toggle:hover {
    background-color: v-bind("props.style?.hoverColorExcludeClose || 'rgb(255, 255, 255)'");
  }
}
</style>