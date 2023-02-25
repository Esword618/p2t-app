<template>
  <main class="main-container">
    <div class="top-container">
      <header-content />
    </div>

    <div class="bottom-container">
      <div class="bottom-content">
        <div class="left">
          <nav-content />
        </div>

        <div class="right">
          <div class="content">
            <router-view />
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script lang="ts" setup>
import NavContent from "@/components/nav-content/nav-content.vue";
import HeaderContent from "@/components/header-content/header-content.vue";
</script>

<style lang="scss" scoped>
$top-container-height: 60px;
$left-side-width: 80px;
$component-interval: 16px;
.main-container {
  position: absolute;
  box-sizing: border-box;
  width: 100%;
  height: 100%;
  background: var(--second-background-color);
  font-size: 15px;

  .top-container {
    background: #8693b2;
    position: relative;
    box-sizing: border-box;
    width: 100%;
    height: $top-container-height;
  }

  .bottom-container {
    position: relative;
    box-sizing: border-box;
    width: 100%;
    height: calc(100% - $top-container-height);
    padding-top: $component-interval;

    .bottom-content {
      position: relative;
      box-sizing: border-box;
      width: 100%;
      height: 100%;
      padding-left: $left-side-width;

      .left {
        position: absolute;
        box-sizing: border-box;
        width: $left-side-width;
        height: 100%;
        top: 0;
        left: 0;
      }

      .right {
        position: relative;
        width: 100%;
        height: 100%;
        box-sizing: border-box;
        padding: 0 $component-interval 0 $component-interval;
      }

      .content {
        background-color: #19afc9;
        position: relative;
        box-sizing: border-box;
        width: 100%;
        height: 100%;
      }
    }
  }
}
</style>
