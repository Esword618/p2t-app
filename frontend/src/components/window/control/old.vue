<template>
  <!-- 置顶——状态切换 -->
  <span :title="isOntoped ? '取消置顶' : '置顶'" class="sticky iconfont icon-zhiding" @click="onTopWindow"></span>
  <span title="最小化" class="minimize iconfont icon-zuixiaohua" @click="minimizeWindow"></span>
  <!-- 最大化应用程序——状态切换 -->
  <span :title="isMaximised ? '取消最大化' : '最大化'" class="maximize iconfont"
    :class="isMaximised ? 'icon-shouqi' : 'icon-quanping'" @click="maximiseWindow"></span>
  <span title="关闭" class="close iconfont icon-guanbi" @click="closeWindow"></span>
</template>

<script setup lang="ts">
import { ref } from "vue";

const props = defineProps({
  WindowMinimise: {
    type: Object as () => () => Promise<void> | void,
    required: true,
  },
  WindowMaximise: {
    type: Object as () => () => Promise<void> | void,
    required: true,
  },
  WindowClose: {
    type: Object as () => (isHidden: boolean) => Promise<void> | void,
    required: true,
  },
  WindowOnTop: {
    type: Object as () => () => Promise<void> | void,
    required: true,
  },
  WindowIsMaximised: {
    type: Object as () => () => Promise<boolean> | boolean,
    required: true,
  },
  WindowIsOnToped: {
    type: Object as () => () => Promise<boolean> | boolean,
    required: true,
  },
});

async function minimizeWindow() {
  await props.WindowMinimise();
}

async function maximiseWindow() {
  await props.WindowMaximise();
  isMaximised.value = !isMaximised.value;
}

async function closeWindow() {
  // wails 运行时关闭程序 传参true表示隐藏窗口而不关闭
  await props.WindowClose(false);
}


async function onTopWindow() {
  // OnTopWindow
  await props.WindowOnTop();
  isOntoped.value = await props.WindowIsOnToped();
  if (isOntoped.value) {
    style.value.stickyRotate = "-45deg";
    style.value.stickyColor = "limegreen";
  } else {
    style.value.stickyRotate = "0deg";
    style.value.stickyColor = "var(--text-primary)";
  }
}

let isMaximised = ref(await props.WindowIsMaximised());
let isOntoped = ref(await props.WindowIsOnToped());
const style = ref({
  stickyRotate: isOntoped.value ? "-45deg" : "0deg",
  stickyColor: isOntoped.value ? "limegreen" : "var(--text-primary)",
});

window.onresize = async (_event) => {
  isMaximised.value = await props.WindowIsMaximised();
};
</script>

<style scoped lang="less">
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

.sticky {
  //color: @sticky-color;
  color: v-bind("style.stickyColor");
  width: 43px;

  &::before {
    transition: transform 300ms;
    transform: translate(-50%, -50%) rotate(v-bind("style.stickyRotate"));
  }

  &:hover {
    background-color: var(--button-hover);
    //color: @background-primary;
  }
}

.minimize {
  color: var(--text-primary);
  width: 43px;

  &:hover {
    background-color: var(--button-hover);
    //color: @background-primary;
  }
}

.maximize {
  color: var(--text-primary);
  width: 43px;

  &:hover {
    background-color: var(--button-hover);
    //color: @background-primary;
  }
}

.close {
  color: var(--text-primary);
  width: 41px;

  &:hover {
    background-color: @close-hover-color;
    color: @background-primary;
  }
}
</style>
