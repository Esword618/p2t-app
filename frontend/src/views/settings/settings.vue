<template>
  <el-col :span="8">
    <el-card shadow="hover">
      <div v-for="o in 4" :key="o" class="text item">
        {{ "设置 开发中" }}
      </div>
    </el-card>
  </el-col>
</template>

<script lang="ts" setup></script>

<style scoped>
.text {
  font-size: 14px;
}

.item {
  padding: 18px 0;
}

.box-card {
  width: 480px;
}
</style>
