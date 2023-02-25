<template>
  <el-config-provider>
    <main-container />
  </el-config-provider>
  <!--  <MainPage />-->
</template>

<script setup lang="ts">
import MainContainer from "@/components/main-container/main-container.vue";
// import MainPage from "@/components/main-page/main-page.vue";
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  box-sizing: border-box;
  position: relative;
  width: 100%;
  height: 100%;
}
</style>
