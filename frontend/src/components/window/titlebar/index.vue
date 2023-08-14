<template>
  <header id="header" style="--wails-draggable: drag" :style="{
    display: isFullScreen ? 'none' : 'flex',
  }">
    <div class="header-wrapper">
      <div class="header-left" @dblclick="maximiseWindow">
        <slot name="extend-left" />
      </div>
      <div class="header-center">
        <slot name="extend-center" />
      </div>
      <div class="header-right" style="--wails-draggable: none">
        <slot name="extend-right" />
        <window-controls :WindowOnTop="WindowOnTop" :WindowMinimise="WindowMinimise" :WindowMaximise="WindowMaximise"
          :WindowClose="WindowClose" :WindowIsMaximised="WindowIsMaximised" :WindowIsOnToped="WindowIsOnToped"
          :style="props.style?.control" />
        <!-- <old-window-controls :WindowOnTop="WindowOnTop" :WindowMinimise="WindowMinimise" :WindowMaximise="WindowMaximise"
          :WindowClose="WindowClose" :WindowIsMaximised="WindowIsMaximised" :WindowIsOnToped="WindowIsOnToped" /> -->
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { PropType, ref, } from "vue";
import { WindowIsMaximised } from "@wailsjs/runtime";
import WindowControls from "../control/index.vue";
// @ts-ignore
import OldWindowControls from "../control/old.vue";

const props = defineProps({
  WindowMinimise: {
    type: Function as PropType<() => Promise<void> | void>,
    required: true,
  },
  WindowMaximise: {
    type: Function as PropType<() => Promise<void> | void>,
    required: true,
  },
  WindowClose: {
    type: Function as PropType<(isHidden: boolean) => Promise<void> | void>,
    required: true,
  },
  WindowOnTop: {
    type: Function as PropType<() => Promise<void> | void>,
    required: true,
  },
  WindowIsMaximised: {
    type: Function as PropType<() => Promise<boolean> | boolean>,
    required: true,
  },
  WindowIsOnToped: {
    type: Function as PropType<() => Promise<boolean> | boolean>,
    required: true,
  },
  WindowIsFullScreen: {
    type: Function as PropType<() => Promise<boolean> | boolean>,
    required: true,
  },
  style: {
    type: Object as PropType<{
      height?: string;
      color?: string;
      backgroundColor?: string;
      borderBottomColor?: string;
      control?: {
        fill?: string;
        hoverColorExcludeClose?: string;
        hoverStyleOnlyClose?: {
          backgroundColor?: string;
          fill?: string;
        };
        onTopActiveColor?: string;
      }
    }>,
    required: false,
  }
});

async function maximiseWindow() {
  await props.WindowMaximise();
}

const isFullScreen = ref(await props.WindowIsFullScreen());

window.addEventListener("resize", async () => {
  isFullScreen.value = await props.WindowIsFullScreen();
});
</script>

<style scoped lang="less">
header {
  width: 100vw;
  overflow: hidden;
  height: v-bind("props.style?.height || '32px'");
  min-height: v-bind("props.style?.height || '32px'");
  //background-color: transparent;
  //background-color: @background-secondary;
  background-color: v-bind("props.style?.backgroundColor || 'transparent'");
  display: flex;
  align-items: center;
  user-select: none;
  position: relative;
  color: v-bind("props.style?.color || 'auto'");

  &::after {
    content: "";
    position: absolute;
    bottom: 0;
    height: 1px;
    width: 100%;
    background-color: v-bind("props.style?.borderBottomColor || 'transparent'");
  }

  .header-wrapper {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: row;

    .header-left {
      margin-right: auto;
      display: flex;
      align-items: center;
    }

    .header-right {
      margin-left: auto;
      display: flex;
      flex-direction: row;

      span {
        cursor: pointer;
        display: inline-block;
        height: 100%;
        //width: 47px;
        overflow: hidden;
        padding: 2px;
        position: relative;
        transition: background-color 300ms;

        &::before {
          position: absolute;
          top: 50%;
          left: 50%;
          transform: translate(-50%, -50%);
        }
      }
    }
  }
}
</style>
