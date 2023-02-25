<template>
  <div class="upload-page-container">
    <!--    <div-->
    <!--      class="upload-page-left page-container"-->
    <!--      :style="{-->
    <!--        width: '280rem',-->
    <!--      }"-->
    <!--    >-->
    <!--      <div-->
    <!--        class="uploaded-item"-->
    <!--        v-for="(item, index) in uploadedImageList"-->
    <!--        :key="index"-->
    <!--      >-->
    <!--        <image-card :image-obj="item" :is-uploaded="true" />-->
    <!--      </div>-->
    <!--    </div>-->

    <div class="upload-page-right page-container">
      <!-- 上传区域 -->
      <!--      <div class="row-item">-->
      <!--        <div class="content-box">-->
      <!--          <upload-area-->
      <!--            :image-loading="imageLoading"-->
      <!--            ref="uploadAreaDom"-->
      <!--          ></upload-area>-->
      <!--        </div>-->
      <!--      </div>-->

      <!-- 待上传的图片列表 -->
      <!--      <div class="row-item">-->
      <!--        <div class="content-box">-->
      <!--          <to-upload-image-card-->
      <!--            ref="toUploadImageCardDom"-->
      <!--            :loading-all-image="imageLoading"-->
      <!--          />-->
      <!--        </div>-->
      <!--      </div>-->
      <!-- 重置 & 上传 -->
      <div class="row-item">
        <div class="content-box" style="text-align: right">
          <el-button plain type="warning" @click="resetUploadImage"
            >重置
          </el-button>
          <el-button plain type="primary" @click="uploadImage">
            上传
          </el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
// import ImageCard from '@/components/image-card/image-card.vue'
// import ToUploadImageCard from '@/components/to-upload-image-card/to-upload-image-card.vue'
// import UploadArea from '@/components/upload-area/upload-area.vue'
import { ElNotification } from "element-plus";

const resetUploadImage = () => {
  // console.log("重置成功");
  notification("success", "重置成功");
};
const uploadImage = () => {
  // console.log("上传成功");
  notification("success", "上传成功");
};

const notification = (type: string, message: string) => {
  ElNotification({
    title: type,
    message: message,
    position: "bottom-right",
    type: "success",
    duration: 1000,
  });
};
</script>

<style scoped></style>
