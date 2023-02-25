<template>
  <aside class="nav">
    <ul class="nav-list">
      <li
        class="nav-item flex-center"
        v-for="(navItem, index) in navList"
        :key="index"
        :class="{ active: navItem.isActive }"
        @click="navClick(navItem)"
        v-show="navItem.isShow"
      >
        <div class="nav-content">
          <el-icon :size="navIconSize">
            <component :is="navItem.icon"></component>
          </el-icon>
          <span class="nav-name">{{ navItem.name }}</span>
        </div>
      </li>
    </ul>
    <div
      class="nav-item quick-actions flex-center"
      :class="{ active: isShowQuickActions }"
    >
      <!--      快捷操作-->
      <!--      <div class="nav-content">-->
      <!--        <el-icon :size="navIconSize">-->
      <!--          <Operation />-->
      <!--        </el-icon>-->
      <!--        <span class="nav-name">快捷操作</span>-->
      <!--      </div>-->
      <div class="quick-actions-box" v-show="isShowQuickActions">
        <!--        <el-switch-->
        <!--          class="mb-2"-->
        <!--          active-text="暗夜模式"-->
        <!--          @change="themeModeChange"-->
        <!--        />-->
        <!--        <el-switch-->
        <!--          class="mb-2"-->
        <!--          active-text="压缩图片"-->
        <!--          @change="persistUserSettings"-->
        <!--        />-->
        <!--        <el-switch-->
        <!--          class="mb-2"-->
        <!--          active-text="转换 Markdown"-->
        <!--          @change="persistUserSettings"-->
        <!--        />-->
        <!--        <el-switch-->
        <!--          v-model="userSettings.isCompress"-->
        <!--          class="mb-2"-->
        <!--          active-text="压缩图片"-->
        <!--          @change="persistUserSettings"-->
        <!--        />-->
        <!--        <el-switch-->
        <!--            v-model="userSettings.defaultMarkdown"-->
        <!--            class="mb-2"-->
        <!--            active-text="转换 Markdown"-->
        <!--            @change="persistUserSettings"-->
        <!--        />-->
      </div>
    </div>
  </aside>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();

const navIconSize = computed(() => {
  // switch (userSettings.elementPlusSize) {
  //   case 'small':
  //     return 22
  //   case 'large':
  //     return 30
  //   default:
  //     return 26
  // }
  return 22;
});

const navList = ref([
  {
    name: "上传图片",
    icon: "upload",
    isActive: false,
    path: "/upload",
    isShow: true,
  },
  {
    name: "图片管理",
    icon: "box",
    isActive: false,
    path: "/management",
    isShow: true,
  },
  {
    name: "我的设置",
    icon: "setting",
    isActive: false,
    path: "/settings",
    isShow: true,
  },
  {
    name: "使用教程",
    icon: "magic-stick",
    isActive: false,
    path: "/tutorials",
    isShow: true,
  },
  {
    name: "关于",
    icon: "chat-dot-round",
    isActive: false,
    path: "/about",
    isShow: true,
  },
]);

const isShowQuickActions = ref<Boolean>(false);

const navClick = (e: any) => {
  const { path } = e;
  router.push(path);
};

const changeNavActive = (currentPath: string) => {
  navList.value.forEach((v) => {
    const temp = v;
    temp.isActive = v.path === currentPath;
    return temp;
  });
};

const persistUserSettings = () => {
  // console.log("persistUserSettings");
  // store.dispatch('USER_SETTINGS_PERSIST')
};

const themeModeChange = () => {
  // if (userSettings.themeMode === 'dark') {
  //   userSettings.themeMode = 'light'
  // } else {
  //   userSettings.themeMode = 'dark'
  // }
  // persistUserSettings()
};

watch(
  () => router.currentRoute.value,
  (_n) => {
    changeNavActive(_n.path);
  }
);

onMounted(() => {
  router.isReady().then(() => {
    changeNavActive(router.currentRoute.value.path);
  });

  document
    .querySelector(".quick-actions .quick-actions-box")
    ?.addEventListener("click", (e) => {
      isShowQuickActions.value = true;
      e.stopPropagation();
    });

  document.querySelector(".quick-actions")?.addEventListener("click", (e) => {
    isShowQuickActions.value = !isShowQuickActions.value;
    e.stopPropagation();
  });

  document.addEventListener("click", () => {
    if (isShowQuickActions.value) {
      isShowQuickActions.value = false;
    }
  });
});
</script>

<style scoped lang="scss">
$z-index-1: 1001;
.nav {
  position: relative;
  width: 100%;
  height: 100%;
  box-sizing: border-box;
  background: var(--background-color);
  display: flex;
  flex-direction: column;
  justify-content: space-between;

  .nav-list {
    padding: 0;
    margin: 0;
  }
  .nav-item {
    position: relative;
    width: 100%;
    height: 76px;
    cursor: pointer;
    user-select: none;
    &.active {
      font-weight: bold;
      background: var(--second-background-color);
    }
    .nav-content {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;

      .nav-name {
        margin-top: 5px;
        font-size: 12px;
      }
    }
  }
  .quick-actions {
    //position: fixed;
    //bottom: 0;
    .quick-actions-box {
      position: absolute;
      bottom: 0;
      left: 100%;
      box-sizing: border-box;
      box-shadow: 1px -1px 2.5px var(--shadow-color);
      width: 180px;
      z-index: $z-index-1;
      background: var(--background-color);
      padding: 6px 10px;
    }
  }
}
</style>
