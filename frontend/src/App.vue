<template>
  <parent-layout/>
  <!-- 更新弹窗 -->
  <el-dialog
      v-model="updateVisible"
      title="新版本更新提示"
      :before-close="handleClose"
      center
  >
    <span> 发现新版本，请前往更新新版本！ </span>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="updateVisible = false">Cancel</el-button>
        <el-button type="primary" @click="onConfirm">
          Confirm
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
// @ts-nocheck
import parentLayout from "./layouts/parentLayout.vue";
import {GetLatestReleases} from "@wailsjs/go/utils/Utils";
import {onMounted} from "vue";
import {BrowserOpenURL, EventsOn} from "@wailsjs/runtime";
import {ElMessageBox, ElNotification} from "element-plus";
import {ref} from "vue";
import {useRouter} from "vue-router";

const router = useRouter();
const updateVisible = ref(false);
const updateUrl = ref("");

const handleClose = (done: () => void) => {
  ElMessageBox.confirm('是否确定暂时不更新?')
      .then(() => {
        done()
      })
      .catch(() => {
        // catch error
      })
}

const onConfirm = () => {
  BrowserOpenURL(updateUrl.value)
  updateVisible.value = false;
}

onMounted(() => {
  // 获取 go 发送过来的解析结果
  EventsOn("try_msg", (message: string) => {
    router.push(message);
  });
})

onMounted(async () => {
  [updateVisible.value,updateUrl.value] = await GetLatestReleases()
})

onMounted(() => {
  // 获取 go 发送过来的解析结果
  EventsOn("watch_notification", (message: Notification) => {
    ElNotification({
      title: message.title,
      message: message.message,
      type: message.type,
      duration: 3000,
      zIndex: 10086,
    })
  });
})

interface Notification {
  title: string,
  message: string,
  type: string,
}
</script>

<style lang="less">
html,
body {
  overflow: hidden;
  background-color: var(--background-primary);
}
</style>
