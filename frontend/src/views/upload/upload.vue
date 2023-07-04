<template>
  <div class="upload-page-container">
    <div class="upload-page-right page-container">
      <!-- 上传区域 -->
      <el-card class="box-card" shadow="hover" v-loading="loading">
        <div class="row-item">
          <div class="content-box">
            <upload-area ref="uploadAreaDom"></upload-area>
          </div>
        </div>
      </el-card>
      <!--分隔条-->
      <el-divider border-style="dashed" />
      <!-- 重置 & 上传 -->
      <div class="row-item">
        <div class="content-box" style="text-align: right">
          <el-button plain type="warning" @click="resetUploadImage"
            >重置
          </el-button>
          <el-button plain type="primary" @click="parseUploadImage">
            解析
          </el-button>
        </div>
      </div>
      <!--分隔条-->
      <el-divider border-style="dashed" />
      <!--解析结果展示-->
      <div class="show-container">
        <el-row :gutter="12">
          <el-col :span="12">
            <el-card shadow="hover" class="textarea-formula-card">
              <el-input
                class="my-input"
                v-model="formula"
                :autosize="{ minRows: 8, maxRows: 8 }"
                type="textarea"
                resize="none"
                placeholder="formula 公式:"
              />
            </el-card>
          </el-col>
          <el-col :span="12">
            <el-card shadow="hover" class="textarea-formula-card">
              <div v-html="formulaResults"></div>
            </el-card>
          </el-col>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref, watch } from "vue";
import UploadArea from "@/components/upload-area/upload-area.vue";
import { ElNotification } from "element-plus";
import { EventsEmit, EventsOn } from "../../../wailsjs/runtime";
import imgStore from "@/store/imgStore";
import userStore from "@/store/userStore";
import mathjax3 from "markdown-it-mathjax3";
import md from "markdown-it";

const Md = md({
  breaks: true,
}).use(mathjax3, {
  loader: { load: ["input/tex", "output/chtml"] },
  tex: {
    inlineMath: [
      ["$", "$"],
      ["\\(", "\\)"],
    ],
  },
});

const useImgeStore = imgStore();
const IUserStore = userStore();
const loading = ref(false);
const formula = ref("");
const formulaResults = ref("");
// const formula = ref("");

// 动态监听 formula 公式
// formulaResults.value: = computed(() => {
//   return Md.render(formula.value);
// });
// 动态监听 formula 公式改变
watch(
  formula,
  () => {
    formulaResults.value = Md.render(formula.value);
  },
  { deep: true, immediate: true }
);

onMounted(() => {
  // 获取 go 发送过来的解析结果
  EventsOn("parse_result", (message: string) => {
    console.log(message);
    const info = JSON.parse(message);
    if (info["status_code"] == 200) {
      // formula.value = v.substring(2, v.length - 2);
      IUserStore.formula = info["results"];
      formula.value = info["results"];
      formulaResults.value = Md.render(info["results"]);
      notificationSuccess("success", "解析成功");
      // const v = infoHandler(message);
    } else {
      notificationError("error", "识别失败，请稍后重试");
    }
    loading.value = false;
  });
});


// 图片重置
const resetUploadImage = () => {
  useImgeStore.imgBase64Url = "";
  // console.log("----->", useImgeStore.getImgBase64Url());
  notificationSuccess("success", "重置成功");
};
// 图片解析
const parseUploadImage = () => {
  loading.value = true;
  EventsEmit("parse_img", "parse");
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
</script>

<style scoped lang="scss">
$content-max-width: 88px;
.upload-page-container {
  background: #ec2e2e;
  width: 100%;
  height: 50%;
  display: flex;
  justify-content: space-between;
  .upload-page-right {
    width: 100%;
    height: 100%;
    box-sizing: border-box;
    .box-card {
      height: 250px;
    }
  }
  .show-container {
    justify-content: space-between;
    border: 1px dashed black;
    padding: 10px;
    .textarea-formula-card {
      min-height: 220px;
      max-height: 220px;
    }
  }
}
</style>
