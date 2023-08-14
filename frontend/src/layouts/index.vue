<template>
  <!-- 缓存处理 -->
  <!-- 这种方式当从缓存路由切换到非缓存路由之后 之前的缓存路由会丢失 -->
  <!-- <router-view>
    <template #default="{ Component, route }">
      <component
        v-if="!route.meta.keepAlive"
        :is="Component"
        :key="route.fullPath"
      ></component>
      <keep-alive v-else>
        <component :is="Component" :key="route.fullPath"></component>
      </keep-alive>
    </template>
  </router-view> -->
  <router-view v-slot="{ Component, route }" class="child-view">
    <!--
      这里的include\exclude匹配的是组件的name属性 而不是路由的name属性
      setup的组件是没有显式定义name属性的 所以需要在script标签中显式定义name属性
      详见：https://github.com/vuejs/core/issues/7799
    -->
    <keep-alive v-if="!isTransition" :include="cachedViews">
      <component :is="Component" :key="route.fullPath" />
    </keep-alive>
  <!--  路由动画  -->
    <transition v-else :name="route.meta.transition">
      <keep-alive :include="cachedViews">
        <component :is="Component" :key="route.fullPath" />
      </keep-alive>
    </transition>
  </router-view>
</template>

<script setup lang="ts">
import { ref, watch } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();

type MatchPattern = string | RegExp;
// 缓存的页面
const cachedViews = ref<MatchPattern[]>([]);
// 是否开启路由过渡动画
const isTransition = ref(true);

// 监听路由改变
watch(
  () => route.path,
  () => {
    // 缓存页面
    const routeName = route.name as string;
    if (route.meta.keepAlive && cachedViews.value.indexOf(routeName) == -1) {
      cachedViews.value.push(routeName);
    }
  },
  {
    immediate: true,
  }
);
</script>

<style lang="less" scoped>
.child-view {
  position: absolute;
  width: 100%;
  transition: all 0.3s cubic-bezier(0.55, 0, 0.1, 1);
  color: var(--text-primary);
}
// 路由动画对应css
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.75s ease;
}

.fade-enter-active,
.fade-leave-active {
  opacity: 0;
}

.slide-left-enter-active,
.slide-right-leave-active {
  opacity: 0;
  transform: translate(30px, 0);
}

.slide-left-leave-active,
.slide-right-enter-active {
  opacity: 0;
  transform: translate(-30px, 0);
}
</style>
