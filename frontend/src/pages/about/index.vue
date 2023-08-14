<template>
  <el-container>
    <el-main class="main-container">
        <el-card class="box-card" shadow="never">
          <!--开头信息-->
          <el-descriptions
              class="des-box"
              :title="title"
              :column="3"
              size="large"
              border
          >
            <!--作者-->
            <el-descriptions-item>
              <template #label>
                <div class="cell-item">
                  <el-icon style="margin-right: 8px">
                    <User />
                  </el-icon>
                  Author
                </div>
              </template>
              Esword
            </el-descriptions-item>
            <!--邮箱-->
            <el-descriptions-item>
              <template #label>
                <div class="cell-item">
                  <el-icon style="margin-right: 8px">
                    <Message />
                  </el-icon>
                  Mail
                </div>
              </template>
              <el-button type="warning" link @click="copy('jyj349559953@gmail.com')">
                jyj349559953@gmail.com
              </el-button>
            </el-descriptions-item>
            <!--github-->
            <el-descriptions-item>
              <template #label>
                <div class="cell-item">
                  <el-icon style="margin-right: 8px">
                    <Link />
                  </el-icon>
                  github
                </div>
              </template>
              <el-button
                  type="primary"
                  link
                  @click="openLink('github.com/Esword618')"
              >
                github.com/Esword618
              </el-button>
            </el-descriptions-item>
            <!--公众号-->
            <el-descriptions-item>
              <template #label>
                <div class="cell-item">
                  <el-icon style="margin-right: 8px">
                    <Orange />
                  </el-icon>
                  公众号
                </div>
              </template>
              <el-popover
                  placement="top-start"
                  title="公众号二维码"
                  :width="200"
                  trigger="hover"
                  content="this is content, this is content, this is content"
              >
                <template #reference>
                  <el-button class="m-2" type="success" round plain>
                    Spiders and AI
                  </el-button>
                </template>
                <!--公众号图片展示-->
                <template #default>
                  <div
                      class="demo-rich-conent"
                      style="display: flex; gap: 16px; flex-direction: column"
                  >
                    <el-image
                        class="active-upload"
                        :src="gzhImg"
                        alt="Pictures to be uploaded"
                        fit="contain"
                    />
                  </div>
                </template>
              </el-popover>
            </el-descriptions-item>
            <!--感谢-->
            <el-descriptions-item>
              <template #label>
                <div class="cell-item">
                  <el-icon style="margin-right: 8px">
                    <Bell />
                  </el-icon>
                  Thanks
                </div>
              </template>
              <el-button
                  type="danger"
                  class="m-2"
                  round
                  plain
                  style="margin-left: 24px"
                  @click="drawer = true"
              >
                click me
              </el-button>
              <el-drawer
                  v-model="drawer"
                  title="I am the title"
                  :with-header="false"
                  size="40%"
              >
                <el-timeline>
                  <el-timeline-item
                      center
                      v-for="(activity, index) in thankActivities"
                      :key="index"
                      :color="activity.color"
                  >
                    <el-card>
                      <h4>{{ activity.title }}</h4>
                      <el-button
                          v-for="(a, index) in activity.links"
                          :key="index"
                          type="danger"
                          link
                          @click="openLink(a.link)"
                      >
                        {{ a.name }}
                      </el-button>
                    </el-card>
                  </el-timeline-item>
                </el-timeline>
              </el-drawer>
            </el-descriptions-item>
            <!--code-->
            <el-descriptions-item>
              <template #label>
                <div class="cell-item">
                  <el-icon style="margin-right: 8px">
                    <Document />
                  </el-icon>
                  Code
                </div>
              </template>
              python、go、javascript、vue、uniapp、爬虫、js逆向
            </el-descriptions-item>
          </el-descriptions>
          <!--分隔条-->
          <el-divider border-style="dashed" />
          <!--任务安排-->
          <div class="todo-list">
            <h3>Todo List</h3>
            <el-table height="18vh" :data="tasks" style="width: 100%">
              <el-table-column prop="state" label="状态" width="75">
                <template #default="scope">
                  <el-tag
                      :type="scope.row.state === '√' ? 'success' : 'danger'"
                      disable-transitions
                  >{{ scope.row.state }}
                  </el-tag>
                </template>
              </el-table-column>
              <el-table-column prop="description" label="任务名称" />
            </el-table>
          </div>
          <!--分隔条-->
          <el-divider border-style="dashed" />
          <!--时间线-->
          <div>
            <h3>Timeline</h3>
            <el-scrollbar height="18vh">
              <el-timeline>
                <el-timeline-item
                    v-for="(activity, index) in timeActivities"
                    :key="index"
                    :color="activity.color"
                    :timestamp="activity.timestamp"
                >
                  {{ activity.content }}
                </el-timeline-item>
              </el-timeline>
            </el-scrollbar>
          </div>
        </el-card>
    </el-main>

  </el-container>

</template>

<script lang="ts">
import { PageEnum } from "@/enums/pageEnum";
export default {
  name: PageEnum.BASE_ABOUT_NAME,
};
</script>
<script lang="ts" setup>
import {onMounted, ref} from "vue";
import { BrowserOpenURL,ClipboardSetText } from "@wailsjs/runtime";
import { Message,User, Link, Orange, Bell, Document } from '@element-plus/icons-vue'
import { ElNotification } from "element-plus";
import gzhImg from "/public/vx.png";
import {GetVersion} from "@wailsjs/go/utils/Utils";

const title = ref('');

const drawer = ref(false);

// 时间线
const timeActivities = [
  {
    content: "v0.2.0 发布（与官方api的返回结果适配，加入新的api，新的ui界面）",
    color: "#f11b1b",
    timestamp: "2023.08.14",
  },
  {
    content: "v0.1.2 发布（与官方api的返回结果适配）",
    color: "#19afc9",
    timestamp: "2023.07.04",
  },
  {
    content:
        "v0.1.1 发布（修复文字与公式无法一起渲染bug，加入对公式编辑立即渲染功能）",
    color: "#434de7",
    timestamp: "2023.02.26",
  },
  {
    content: "v0.1.0 发布（基础功能完成）",
    color: "#0bbd87",
    timestamp: "2023.02.25",
  },

];
// 感谢
const thankActivities = [
  {
    title: "感谢",
    color: "#0bbd87",
    links: [
      {
        name: "Vanisper",
        link: "https://github.com/Vanisper",
      },
      {
        name: "breezedeus",
        link: "github.com/breezedeus",
      },
    ],
  },
  {
    title: "感谢项目",
    color: "#ef2562",
    links: [
      {
        name: "wails",
        link: "github.com/wailsapp",
      },
      {
        name: "Pix2Text",
        link: "github.com/breezedeus/Pix2Text",
      },
    ],
  },
  {
    title: "参考项目",
    color: "#4d5e27",
    links: [
      {
        name: "MoocDownload",
        link: "https://github.com/Esword618/MoocDownload",
      },
    ],
  },
];
// todo
const tasks = ref([
  { state: "√", description: "v0.2.0 完成其它页面的编写,软件重构" },
  { state: "×", description: "v0.3.0 性能优化" },
  { state: "×", description: "v0.4.0 模型本地化部署" },
  { state: "×", description: "...待加入" },
]);
// 打开链接
const openLink = (link: string) => {
  BrowserOpenURL("https://" + link);
  // console.log(link);
};
// 复制功能
const copy = async (text: string) => {
  try {
    await ClipboardSetText(text);
    notificationSuccess("success", "邮箱复制成功");
  } catch (e) {
    notificationError("error", "邮箱复制失败");
  }
};
// 通知
const notificationSuccess = (title: string, message: string) => {
  ElNotification({
    title: title,
    message: message,
    position: "top-right",
    type: "success",
    duration: 1000,
  });
};
const notificationError = (title: string, message: string) => {
  ElNotification({
    title: title,
    message: message,
    position: "top-right",
    type: "error",
    duration: 1000,
  });
};

onMounted(async () => {
  const version = await GetVersion();
  title.value = "P2t-app " + version;
})
</script>

<style scoped lang="less">

.main-container {
  --el-fill-color-blank: var(--background-primary);
  --el-text-color-primary: var(--text-primary);
  --el-bg-color: var(--background-primary);
  --el-dialog-bg-color: var(--background-primary);
  --el-text-color-regular: var(--text-primary);
  --el-fill-color-light:var(--fill-color-light);
  --el-border-color-light:var(--background-toc-active);
  display: flex;
  justify-content: center;
  align-items: center;
  .box-card {
    height: 100%;
    //width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    .des-box {
      margin-top: 20px;
      .cell-item {
        display: flex;
        align-items: center;
      }
    }
  }
}
</style>