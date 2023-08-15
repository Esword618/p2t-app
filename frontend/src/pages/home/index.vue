<template>
  <div class="container">
    <!-- 上传 -->
    <div class="upload-container"
    >
      <el-upload
          :before-upload="beforeUpload"
          :show-file-list="false"
          class="upload-box"
          drag
      >
        <div>
          <el-icon class="el-icon--upload">
            <upload-filled/>
          </el-icon>
          <div>
            Click, copy or drag images onto the page for recognition.
          </div>
        </div>
      </el-upload>
      <div v-if="imgBs64.length > 0"
           class="show-image-box"
           v-loading="loading"
           element-loading-text="Loading..."
           :element-loading-spinner="svg"
           element-loading-svg-view-box="-10, -10, 50, 50"
           element-loading-background="rgba(122, 122, 122, 0.8)"
      >
        <img :src="imgBs64" />
      </div>
    </div>
    <!-- 按钮 -->
    <div class="btn-container">
      <el-divider border-style="dashed">
        <el-icon>
          <star-filled/>
        </el-icon>
      </el-divider>
      <div class="btn-box">
        <el-button @click="onReset" type="warning">
          <el-icon>
            <Refresh/>
          </el-icon>
          reset
        </el-button>
        <el-button @click="onParse" type="primary">
          <el-icon>
            <ZoomIn/>
          </el-icon>
          parse
        </el-button>
      </div>
      <el-divider border-style="dashed"/>
    </div>
    <!-- 结果 -->
    <div class="result-container">
      <div class="result-box">
        <div class="result-container-inner">
          <div class="result-left-box">
            <div class="dropdown-container">
              <el-dropdown @command="handleCommand">
                <el-button plain >
                  <el-icon ><arrow-down /></el-icon>
                </el-button>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item command="Latex">Latex</el-dropdown-item>
                    <el-dropdown-item command="Tex">Tex</el-dropdown-item>
                    <el-dropdown-item command="MathML">MathML</el-dropdown-item>
                    <el-dropdown-item divided command="PNG">PNG</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </div>
            <textarea v-model="textarea" class="text-textarea"
                      placeholder="Please upload an image or enter your content."/>
          </div>
          <div class="hidden-space"></div>
          <div class="result-right-box">
            <div>
              <el-tooltip
                  content="view!"
                  placement="top-end"
              >
                <el-button plain size="small" @click="onView">
                  <el-icon>
                    <View/>
                  </el-icon>
                </el-button>
              </el-tooltip>
              <el-dialog width="80%" v-model="centerDialogVisible" center>
                <div id="home-dialog-preview"></div>
              </el-dialog>
            </div>
            <div class="latex-box">
              <div id="home-preview" class="rendered-content"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

// https://github.com/vuejs/core/issues/7799
// 解决了setup组件中没有显式定义组件name的问题  路由缓存的include、exclude属性匹配的是组件的name属性
<script lang="ts">
import {PageEnum} from "@/enums/pageEnum";

export default {
  name: PageEnum.BASE_HOME_NAME,
};
</script>
<script lang="ts" setup>
// @ts-nocheck
import {
  UploadFilled,
  StarFilled,
  Refresh,
  View,
  ZoomIn,
  ArrowDown,
} from '@element-plus/icons-vue'
import {ref, watch, nextTick, onMounted} from "vue";
import {ElMessage, ElNotification, UploadRawFile} from "element-plus";
import {ClipboardSetText, EventsOn} from "@wailsjs/runtime";
import {convertTex2Mml, convertTex2SvgThenPng, renderLatex} from "@/utils/util/tex.ts";
import {ClipboardSetImage, GetTex, ParseImg} from "@wailsjs/go/utils/Utils";
import {useRoute, useRouter} from "vue-router";

const router = useRouter();
const route = useRoute();
const centerDialogVisible = ref(false);
const textarea = ref(``);
const imgBs64 = ref('');
const loading = ref(false);
const svg = `
        <path class="path" d="
          M 30 15
          L 28 17
          M 25.61 25.61
          A 15 15, 0, 0, 1, 15 30
          A 15 15, 0, 1, 1, 27.99 7.5
          L 15 15
        " style="stroke-width: 4px; fill: rgba(0, 0, 0, 0)"/>
      `;

const beforeUpload = (rawFile: UploadRawFile) => {
  const a = new FileReader()
  a.readAsDataURL(rawFile)
  a.onload = (e) => {
    if (typeof e.target?.result === 'string') {
      imgBs64.value = e.target.result
      ElNotification({
        title: '图片上传！',
        message: `图片上传成功！`,
        type: 'success',
        duration: 1500,
      })
    } else {
      ElNotification({
        title: '图片上传！',
        message: `图片上传失败！`,
        type: 'error',
        duration: 1500,
      })
    }
  }
  return false;
};

const onReset = () => {
  imgBs64.value = "";
  loading.value = false;
}

const onParse = async () => {
  // console.log(imgBs64.value)
  loading.value = true;
  textarea.value = await ParseImg(imgBs64.value);
  // textarea.value = await SimplexApi(imgBs64.value)
  loading.value = false;
  ElNotification({
    title: '图片解析！',
    message: `图片解析成功！`,
    type: 'success',
    duration: 1500,
  });
}

const onView = () => {
  if (!textarea.value) {
    ElMessage({
      message: '无latex公式',
      duration: 1500,
      grouping: true,
      type: 'warning',
    });
    return;
  }
  centerDialogVisible.value = true
  nextTick(async () => {
    await renderLatex('home-dialog-preview',textarea.value);
  })
};

const handleCommand = async (command: string | number | object) => {
  const texStr = await GetTex(textarea.value);
  switch (command) {
    case "Latex":
      await ClipboardSetText(textarea.value);
      break
    case "Tex":
      await ClipboardSetText(texStr);
      break
    case "MathML":
      console.log(await convertTex2Mml(texStr))
      await ClipboardSetText(await convertTex2Mml(texStr));
      break;
    case "PNG":
      const pngBase64 = await convertTex2SvgThenPng(texStr);
      await ClipboardSetImage(pngBase64)
      break;
  }
  ElNotification({
    title: '复制通知!',
    message: `${command} 复制成功!`,
    type: 'success',
    duration: 1500,
  })
}

// @ts-ignore
watch(textarea, async (newValue, oldValue) => {
  await renderLatex('home-preview',newValue);
}, {immediate: true})

onMounted(() => {
  // 获取 go 发送过来的解析结果
  EventsOn("watch_clipboard_img", (message: string) => {
    // 获取并打印当前路由的路径
    if (router.name != "Home") {
      router.push("/home")
      imgBs64.value = message;
    }
  });
})

</script>

<style lang="less" scoped>
/* styles here */
.container {
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  // 上传容器
  .upload-container {
    max-height: 23vh;
    display: flex;
    justify-content: center;
    box-sizing: border-box;
    margin: 16px;
    flex-direction: row; /* 使子元素按行排列 */
    align-items: stretch; /* 使子元素填充至最大高度 */
    // 劫持elemnet-plus的主题颜色，变为自定义的颜色
    --el-fill-color-blank: var(--background-primary);

    .upload-box {
      flex: 1; /* 划分相等的空间 */
      box-sizing: border-box;
      width: 100%;
      max-width: 80%;
      margin: 0 auto;
    }
    .show-image-box {
      flex: 1; /* 划分相等的空间 */
      box-sizing: border-box;
      display: flex;
      justify-content: center; // 主轴对齐
      align-items: center;

      img {
        height: 100%;
        width:100%;
        max-height: 18vh;      // 高度不超过其容器的高度
        object-fit: contain;   // 图片将会按照其原本的宽高比缩放，并自动在空白区域添加白色背景。
      }
    }
  }

  // 按钮容器
  .btn-container {
    width: 100%;
    max-width: 80%;
    margin: 0 auto;
    --el-bg-color: var(--background-primary);
    --el-text-color-primary: var(--icon-information);

    .btn-box {
      text-align: right; /* 将按钮文本对齐到右侧 */

      button {
        margin-right: 10px; /* 或者你想要的间距值 */
      }
    }
  }

  // 结果容器
  .result-container {
    display: flex;
    justify-content: center;
    box-sizing: border-box;
    margin: 16px;
    --el-fill-color-blank: var(--background-primary);
    --el-text-color-primary: var(--text-primary);

    .result-box {
      box-sizing: border-box;
      width: 100%;
      max-width: 90%;
      margin: 0 auto;

      .result-container-inner {
        display: flex;
        justify-content: space-between;
        align-items: center; /* 垂直居中两个盒子和分隔符 */

        .result-left-box,
        .result-right-box {
          height: 42vh; /* 设置高度为页面高度的 50% */
          flex: 1; /* 让两个盒子占据相同的宽度，包括分隔符的宽度 */
          border: 0.5px dashed #d9d9d9 !important;
          border-radius: 10px;
          padding: 10px;
          /* 新增样式属性 */
          position: relative;
          width: 50vh;
        }

        .result-left-box {
          .dropdown-container {
            position: absolute;
            top: 0;
            right: 0;

            button {
              color: #626aef;
              margin: 5px; /* 可以调整按钮与容器边缘的距离 */
              z-index: 1;
              border: none;
              outline: none; /* 去除选中时的边框 */
            }
          }
          .text-textarea {
            background-color: transparent;
            width: 100%;
            height: 100%;
            border: none;
            outline: none; /* 去除选中时的边框 */
            resize: none; /* 去除右下角拉取调整边框 */
            overflow: auto; /* 自动显示垂直滚动条 */

            // 定义 WebKit 浏览器（如 Chrome 和 Safari）的滚动条样式。
            &::-webkit-scrollbar {
              width: 5px;
              height: 8px;
              background-color: var(--text-primary);
            }

            // 定义滚动条的轨道部分的背景颜色。
            &::-webkit-scrollbar-track {
              background: var(--background-primary);
            }

            // 定义滚动条的滑块部分的背景颜色，以及滑块的边框半径。
            &::-webkit-scrollbar-thumb {
              background: var(--text-inactive);
              border-radius: 80px;
            }

            // 定义滑块在鼠标悬停时的背景颜色。
            &::-webkit-scrollbar-thumb:hover {
              background: var(--text-inactive);
            }

            // 定义滚动条的角落部分的背景颜色。
            &::-webkit-scrollbar-corner {
              background: var(--background-primary);
            }
          }
        }

        .result-right-box {
          --el-bg-color: var(--background-primary);
          --el-dialog-bg-color: var(--background-primary);
          --el-text-color-regular: var(--text-primary);

          button {
            color: #626aef;
            position: absolute;
            top: 0;
            right: 0;
            margin: 5px; /* 可以调整按钮与容器边缘的距离 */
            z-index: 1;
            border: none;
            outline: none; /* 去除选中时的边框 */
          }

          .latex-box {
            // 设置内容不超出容器，使用绝对定位
            position: relative;
            overflow: hidden;
            width: 100%;
            height: 100%;
            //禁止右键菜单
            //pointer-events: none;
            .rendered-content {
              width: 100%;
              height: 100%;
              overflow: auto; /* 自动显示垂直滚动条 */
              // 确保容器内的内容不会撑开容器

              // 定义 WebKit 浏览器（如 Chrome 和 Safari）的滚动条样式。
              &::-webkit-scrollbar {
                width: 5px;
                height: 8px;
                background-color: var(--text-primary);
              }

              // 定义滚动条的轨道部分的背景颜色。
              &::-webkit-scrollbar-track {
                background: var(--background-primary);
              }

              // 定义滚动条的滑块部分的背景颜色，以及滑块的边框半径。
              &::-webkit-scrollbar-thumb {
                background: var(--text-inactive);
                border-radius: 80px;
              }

              // 定义滑块在鼠标悬停时的背景颜色。
              &::-webkit-scrollbar-thumb:hover {
                background: var(--text-inactive);
              }

              // 定义滚动条的角落部分的背景颜色。
              &::-webkit-scrollbar-corner {
                background: var(--background-primary);
              }
            }

          }

        }

        .hidden-space {
          width: 5px; /* 调整间距的宽度 */
        }
      }
    }
  }
}
</style>
