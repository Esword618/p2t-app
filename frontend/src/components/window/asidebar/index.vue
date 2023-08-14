<template>
  <aside class="page-asidebar" style="--wails-draggable: drag;">
    <router-link v-for="(item, index) in asideList.top" :key="index" class="asider-item icon" :class="{
            active: $route.path === item.path,
        }" :to="item.path">
      <el-tooltip
          :content="item.text"
          placement="right"
      >
        <svg viewBox="0 0 1024 1024">
          <path v-for="(d, i) in item.svg" :key="i" :d="d"></path>
        </svg>
      </el-tooltip>
    </router-link>
    <slot name="extend-top" />
    <div style="margin-top: auto;display: flex;flex-direction: column;">
      <slot name="extend-bottom" />
      <router-link v-for="(item, index) in asideList.bottom" :key="index" class="asider-item icon" :class="{
                active: $route.path === item.path,
            }" :to="item.path">
        <el-tooltip
            :content="item.text"
            placement="right"
        >
          <svg viewBox="0 0 1024 1024">
            <path v-for="(d, i) in item.svg" :key="i" :d="d"></path>
          </svg>
        </el-tooltip>
      </router-link>
    </div>

  </aside>
</template>

<script lang="ts" setup>
import {asideList} from "./list"
</script>

<style lang="less" scoped>
.page-asidebar {
  display: flex;
  flex-direction: column;
  width: 60px;
  height: 100%;
  background-color: var(--background-secondary);
  padding: 10px 0;
  box-sizing: border-box;
  overflow: hidden;
  transition: width 0.3s ease-in-out;

  // &:hover {
  //     width: 200px;
  // }

  .asider-item {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 40px;
    width: 100%;
    color: var(--text-secondary);
    text-decoration: none;
    margin-top: 8.75px;
    overflow: hidden;

    &:hover {
      background-color: var(--background-tertiary);
    }

    &.active {
      background-color: var(--background-tertiary);
    }

    &.icon {
      width: 40px;
      height: 40px;
      margin-left: auto;
      margin-right: auto;
      display: flex;
      justify-content: center;
      align-items: center;
      border-radius: 10px;
      overflow: hidden;
      padding: 10px;
      box-sizing: border-box;

      svg {
        width: 100%;
        height: 100%;
        fill: var(--text-secondary);
      }
    }

    &.text {
      font-size: 12px;
    }

    &.icon-text {
      flex-direction: column;
    }
  }


}
</style>