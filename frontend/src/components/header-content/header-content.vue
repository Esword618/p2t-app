<template>
  <header class="header">
    <div class="header-left">
      <div class="brand" @click="router.push('/')">
        <div class="logo">
          <img src="../../assets/logo.jpg" alt="Esword" />
<!--          <img src="../../assets/logo.png" alt="Esword" />-->
        </div>
        <div class="title">p2t-app</div>
      </div>
      <div class="website-count">
        <el-tooltip
          content="感觉好用，点 Star 支持作者(*￣︶￣)"
          placement="bottom"
        >
        </el-tooltip>
      </div>
      <!--      <div class="header-right">-->
      <!--        <div class="user-info" @click="onUserInfoClick">-->
      <!--          <div class="username">Esword</div>-->

      <!--          <div class="avatar">-->
      <!--            <el-icon :size="22"><UserFilled /></el-icon>-->
      <!--          </div>-->

      <!--          <el-dropdown-->
      <!--            trigger="click"-->
      <!--            @command="handleCommand"-->
      <!--            v-if="userConfigInfo?.avatarUrl"-->
      <!--          >-->
      <!--            <span class="el-dropdown-link">-->
      <!--              <span class="avatar">-->
      <!--                <img-->
      <!--                  :src="userConfigInfo?.avatarUrl"-->
      <!--                  :alt="userConfigInfo?.owner"-->
      <!--                />-->
      <!--              </span>-->
      <!--            </span>-->
      <!--            <template #dropdown>-->
      <!--              <el-dropdown-menu>-->
      <!--                <el-dropdown-item command="logout"> 退出登录 </el-dropdown-item>-->
      <!--              </el-dropdown-menu>-->
      <!--            </template>-->
      <!--          </el-dropdown>-->
      <!--        </div>-->
      <!--      </div>-->
    </div>
  </header>
</template>

<script lang="ts" setup>
import { useRouter } from "vue-router";
const router = useRouter();
</script>

<style scoped lang="scss">
.header {
  width: 100%;
  height: 100%;
  background: var(--background-color);
  padding: 0 20px;
  box-sizing: border-box;
  display: flex;
  justify-content: space-between;
  align-items: center;

  .header-left {
    height: 100%;
    display: flex;
    justify-content: flex-start;

    .brand {
      height: 100%;
      display: flex;
      justify-content: flex-start;
      align-items: center;
      cursor: pointer;
      .logo {
        width: 46px;
        height: 46px;
        margin-right: 10px;
        img {
          width: 100%;
          height: 100%;
          border-radius: 50%;
        }
      }
      .title {
        font-size: 36px;
        font-weight: bold;
      }
    }
    .website-count {
      box-sizing: border-box;
      display: flex;
      align-items: flex-end;
      font-size: 14px;
      margin-left: 10px;
      padding-bottom: 12px;
      cursor: pointer;
    }
  }
}
</style>
