<template>
  <div
    class="upload-area active-upload"
    :class="{ focus: uploadAreaActive }"
    element-loading-text="图片上传中..."
    element-loading-background="rgba(0, 0, 0, 0.5)"
  >
    <label for="uploader" class="active-upload" v-if="uploadAreaActive"></label>
    <input id="uploader" @change="onSelect" type="file" multiple autofocus />
    <div class="tips active-upload" v-if="!imgBase64Url">
      <el-icon class="icon active-upload">
        <UploadFilled />
      </el-icon>
      <div class="text active-upload">粘贴、或点击此处上传</div>
    </div>
    <el-image
      class="active-upload-image"
      v-if="imgBase64Url"
      :src="imgBase64Url"
      alt="Pictures to be uploaded"
      fit="contain"
    />
  </div>
</template>

<script lang="ts" setup>
import imgStore from "@/store/imgStore";
import { ref, computed, onMounted } from "vue";
import { EventsOn, EventsEmit } from "../../../wailsjs/runtime";
// import { UploadFilled } from "@element-plus/icons-vue";
// import { Search,FolderOpened } from '@element-plus/icons-vue';

const useImgeStore = imgStore();
const uploadAreaActive = ref(true);

// const imgBase64Url = ref("");

// 动态监听imgBase64Url
const imgBase64Url = computed(() => {
  return useImgeStore.getImgBase64Url();
});

const onSelect = async (e: any) => {
  // console.log("imgBase64Url:", imgBase64Url);
  for (const file of e.target.files) {
    // console.log(file);
    const blob = await file.type;
    // console.log(blob);
    if (blob === "image/png" || blob === "image/jpeg") {
      const bs64img = await blobToBase64(file);
      useImgeStore.imgBase64Url = bs64img;
      EventsEmit("set_base64_img", bs64img);
    }
  }
};

onMounted(() => {
  // 获取 go 发送过来的图片
  EventsOn("clipboard_image", (bs64img: string) => {
    if (useImgeStore.imgBase64Url != bs64img) {
      useImgeStore.imgBase64Url = bs64img;
      // imgBase64Url.value = bs64img;
    }
  });
});

onMounted(() => {
  // ctrl + v 快捷键检测，并接收 go 发送接的图片
  document.onkeydown = async (e) => {
    if (((e.ctrlKey || e.metaKey) && e.key === "v") || e.key === "V") {
      const clipboardItems = await navigator.clipboard.read();
      for (const clipboardItem of clipboardItems) {
        for (const type of clipboardItem.types) {
          const blob = await clipboardItem.getType(type);
          // console.log(blob);
          if (blob.type === "image/png" || blob.type === "image/jpeg") {
            const bs64img = await blobToBase64(blob);
            // imgBase64Url.value = bs64img;
            useImgeStore.imgBase64Url = bs64img;
            EventsEmit("set_base64_img", bs64img);
          }
        }
      }
      e.preventDefault();
    }
  };
});

// file转base64
const blobToBase64 = (blob: Blob): Promise<string> => {
  return new Promise((resolve, _) => {
    const reader = new FileReader();
    reader.onloadend = () => resolve(reader.result as string);
    reader.readAsDataURL(blob);
  });
};
</script>

<style scoped lang="scss">
.upload-area {
  position: relative;
  width: 100%;
  height: 300px;
  border: 4px dashed var(--third-text-color);
  box-sizing: border-box;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 999;
  .active-upload{
    width: 100%;
    margin-bottom: 20px;
  }
  .active-upload-image {
    width: 100%;
    margin-bottom: 100px;
  }
  &.focus {
    border-color: var(--upload-area-focus-color);
  }

  &:hover {
    border-color: var(--upload-area-focus-color);
  }

  label {
    display: block;
    position: absolute;
    width: 100%;
    height: 100%;
    z-index: 1000;
    cursor: pointer;
  }

  input[type="file"] {
    position: absolute;
    left: -9999px;
    top: -9999px;
  }

  .tips {
    text-align: center;
    color: #aaa;

    .icon {
      font-size: 100px;
    }

    .text {
      cursor: default;
      font-size: 20px;
      margin-bottom: 100px;
    }
  }
}
</style>
