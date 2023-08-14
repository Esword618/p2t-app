<template>
  <div>
    <div v-if="cardsDate.length ==0 " class="container">
      <el-empty description="description" />
    </div>
    <div v-else class="container">
      <el-container>
        <el-header class="header-container">
          <div class="block">
            <el-date-picker
                v-model="selectData"
                :disabled-date="disabledDate"
                placeholder="Pick a day"
                type="date"
                format="YYYY/MM/DD"
                value-format="YYYY-MM-DD"
            />
          </div>
          <el-input
              v-model="latexInput"
              class="latex-input"
              placeholder="Please input latex text"
          >
            <template #append>
              <el-button @click="onSerch">
                <el-icon>
                  <Search/>
                </el-icon>
              </el-button>
            </template>
          </el-input>
        </el-header>
        <el-main class="main-container">
          <div
              class="card-container"
              @mouseenter="showScrollbar = true"
              @mouseleave="showScrollbar = false"
          >
            <el-card
                v-for="(cardDate) in cardsDate"
                :key="cardDate"
                class="card"
            >
              <div class="card-content">
                <img
                    :src="baseImgUrl+cardDate.Image"
                    class="card-image"
                    @click="onDetail($event,cardDate)"
                />
                <div style="padding: 14px">
                  <div class="card-bottom">
                    <time class="card-time">{{ cardDate.CreatedAt }}</time>
                    <el-popconfirm
                        class="card-btn"
                        :icon="InfoFilled"
                        cancel-button-text="No"
                        confirm-button-text="Yes"
                        confirm-button-type="danger"
                        title="Are you sure to delete this?"
                        trigger="click"
                        @cancel="cancelEvent"
                        @confirm="confirmEvent(cardDate.Uuid)"
                    >
                      <template #reference>
                        <el-button
                            :icon="Delete"
                            circle
                            class="detail-button"
                            type="danger"
                        >
                        </el-button>
                      </template>
                    </el-popconfirm>
                  </div>
                </div>
              </div>
              <el-tag class="card-api-type">ApiType：{{ cardDate.ApiType }}</el-tag>
            </el-card>
            <!-- 仿小红书弹窗 参考：https://github.com/xishandong/Vue3_web_redbook -->
            <transition
                name="fade"
                @before-enter="onBeforeEnter"
                @after-enter="onAfterEnter"
                @before-leave="onBeforeLeave"
                @after-leave="onAfterLeave"
            >
              <div v-if="showOverlay" class="overlay-box" @click="showOverlay = false">
                <button class="backPage" style="display:none;" @click="showOverlay = false">
                  <el-icon>
                    <Back/>
                  </el-icon>
                </button>
                <el-card
                    :body-style="{ padding: '0px' }"
                    class="overlay-card"
                    @click.stop
                >
                  <div class="overlay-card-upper-container">

                    <div class="center-image">
                      <el-image
                          :preview-src-list="[detailOriginImg]"
                          :src="detailOriginImg"
                          class="origin-image"
                          fit="cover"
                      />
                    </div>

                    <el-scrollbar height="40vh">
                      <div class="latex-box">
                        <div id="history-preview" class="rendered-content"></div>
                      </div>
                    </el-scrollbar>

                  </div>
                  <div class="overlay-card-lower-container"
                       style="padding: 14px"
                  >
                    <el-radio-group v-model="formulaType" label="label position" @click="isCopy = false">
                      <el-radio-button label="Latex">copy Latex</el-radio-button>
                      <el-radio-button label="Tex">copy Tex</el-radio-button>
                      <el-radio-button label="MathML">Copy MathML(Word)</el-radio-button>
                      <!--                    <el-radio-button label="PNG">Copy PNG</el-radio-button>-->
                    </el-radio-group>
                    <div class="spacer"></div>
                    <div class="textarea-box">
                      <el-button plain @click="onCopy">
                        <el-icon v-if="!isCopy">
                          <CopyDocument/>
                        </el-icon>
                        <el-icon v-if="isCopy">
                          <Check/>
                        </el-icon>
                      </el-button>
                      <textarea v-if="formulaType == 'Latex'" v-model="copyFormula.Latex" class="text-textarea"/>
                      <textarea v-else-if="formulaType == 'Tex'" v-model="copyFormula.Tex" class="text-textarea"/>
                      <textarea v-else-if="formulaType == 'MathML'" v-model="copyFormula.MathML" class="text-textarea"/>
                      <textarea v-else class="text-textarea">
                          图片如上图所示！
                      下面是图片的base64编码!
                      {{ copyFormula.PNG }}
                    </textarea>
                    </div>
                    <div class="spacer"></div>
                  </div>
                </el-card>
              </div>
            </transition>
          </div>
        </el-main>
      </el-container>
    </div>
  </div>
</template>
<script lang="ts">
// @ts-ignore
import {PageEnum} from "@/enums/pageEnum";

export default {
  name: PageEnum.BASE_HISTORY_NAME,
};
</script>
<script lang="ts" setup>
// @ts-nocheck
import {nextTick, onMounted, ref, reactive, watch} from 'vue'
import {Back, Search, CopyDocument, Check, InfoFilled, Delete} from "@element-plus/icons-vue";
import {ClipboardSetText} from "@wailsjs/runtime";
import {ElNotification} from "element-plus";
import {GetTex, ClipboardSetImage, GetLatestDate, GetDataByTimeOrLatex,DeleteByUuid} from "@wailsjs/go/utils/Utils";
import {convertTex2Mml, convertTex2SvgThenPng, renderLatex} from "@/utils/util/tex.ts";

const isCopy = ref(false);
const formulaType = ref('Latex');
const selectData = ref('');
const showScrollbar = ref(false);
const latexInput = ref("")
const copyFormula = reactive({
  Latex: "",
  Tex: "",
  MathML: "",
  PNG: "",
});
const textarea = ref(``);
const cardsDate = ref([]);
const detailOriginImg = ref('')
const baseImgUrl = ref('http://localhost:8080?image=');

// 初始化数据
onMounted(async () => {
  if (cardsDate.value.length == 0) {
    cardsDate.value = await GetLatestDate();
  }
});

// 监控路由变化以便重新获取数据
// const initLoadDate = async () => {
//   if (route.name == "history") {
//     cardsDate.value = await GetLatestDate();
//   }
//   console.log(route.name)
// }
// watch(() => route.params, initLoadDate)

watch([selectData, latexInput], async ([newFormulaType, newLatexInput], [oldFormulaType, oldLatexInput]) => {
  cardsDate.value = await GetDataByTimeOrLatex(newFormulaType,newLatexInput);
});

const confirmEvent = async (uuid: string) => {
  const deleteStatus = await DeleteByUuid(uuid)
  if (deleteStatus == null){
    ElNotification({
      title: '图片删除操作！',
      message: `图片删除成功！`,
      type: 'success',
      duration: 1500,
    });
    cardsDate.value = await GetLatestDate();
  } else {
    ElNotification({
      title: '图片删除操作！',
      message: `图片删除失败！`,
      type: 'error',
      duration: 1500,
    })
  }
}

const cancelEvent = () => {
  console.log('cancel!')
}

const onSerch = async () => {
  if ((selectData.value.length>0) || (latexInput.value.length >0)) {
    cardsDate.value = await GetDataByTimeOrLatex(selectData.value,latexInput.value);
  } else {
    ElNotification({
      title: '数据查询！',
      message: `时间与内容不能全为空！`,
      type: 'error',
      duration: 1500,
    })
  }
}

const disabledDate = (time: Date) => {
  return time.getTime() > Date.now()
}

const onCopy = async () => {
  isCopy.value = true; // 设置为 true
  switch (formulaType.value) {
    case "Latex":
      await ClipboardSetText(copyFormula.Latex);
      break
    case "Tex":
      await ClipboardSetText(copyFormula.Tex);
      break
    case "MathML":
      await ClipboardSetText(copyFormula.MathML);
      break;
    case "PNG":
      const pngBase64 = await convertTex2SvgThenPng(copyFormula.Tex);
      await ClipboardSetImage(pngBase64)
      break;
  }
  ElNotification({
    title: '复制通知!',
    message: `${formulaType.value} 复制成功!`,
    type: 'success',
    zIndex: 10000,
    duration: 1500,
  })
  // 在 3 秒后将 isCopy.value 设置为 false
  setTimeout(() => {
    isCopy.value = false;
  }, 3000);
};

// 复制监听
watch(copyFormula, async (newValue, oldValue) => {
  await renderLatex('history-preview', newValue.Latex);
  const texStr = await GetTex(newValue.Latex)
  copyFormula.Tex = texStr;
  copyFormula.MathML = await convertTex2Mml(texStr)
  // copyFormula.PNG = await convertTex2SvgThenPng(texStr)
}, {immediate: true})

// 观察器
onMounted(() => {
  // MathJax 版本问题导致右键菜单禁止失败
  // const showMathMenu = false
  // window.MathJax.options.enableMenu = showMathMenu;
  // 创建一个观察器实例，并定义观察的回调函数
  const observer = new MutationObserver(function (mutationsList, observer) {
    for (let mutation of mutationsList) {
      if (mutation.type === 'childList') {
        // 遍历所有新增子节点
        mutation.addedNodes.forEach(node => {
          // 如果新增子节点是你想处理的 .CtxtMenu_MenuFrame
          if (node.className === 'CtxtMenu_MenuFrame') {
            node.style.zIndex = "100000";
          }
        });
      }
    }
  });

  // 配置 observer 需要观察的变化类型，以及观察的目标节点
  observer.observe(document.body, {
    childList: true,
    subtree: true
  });

});

//--------------------仿小红书展示--------------------
const showOverlay = ref(false);
const overlayX = ref(0); // 覆盖层的水平位置
const overlayY = ref(0); // 覆盖层的垂直位置
let style = null;

const onDetail = async (event: MouseEvent,cardDate: Map<string,string>) => {
  overlayX.value = event.pageX;
  overlayY.value = event.pageY;
  showOverlay.value = true;
  nextTick(async () => {
    detailOriginImg.value = baseImgUrl.value + cardDate.Image
    const texStr = await GetTex(cardDate.Latex)
    await renderLatex('history-preview', cardDate.Latex);
    copyFormula.Latex = cardDate.Latex;
    copyFormula.Tex = texStr;
    copyFormula.MathML = await convertTex2Mml(texStr)
    // copyFormula.PNG = await convertTex2SvgThenPng(texStr)
  })
}

//----------仿小红书动画------------
const onBeforeEnter = () => {
  style = document.createElement('style')
  style.innerHTML =
      `@keyframes scale-up-center {
    0% {
      transform: scale(0.5);
      transform-origin: ${overlayX.value}px ${overlayY.value}px;
    }
    10% {
      transform: scale(0.5);
    }
    100% {
      transform: scale(1);
    }
  }`
  document.head.appendChild(style);
}

const onAfterEnter = (el) => {
  el.style = 'background-color: #fff'
  const button = el.querySelector('.backPage')
  button.style.display = ''
}

const onBeforeLeave = (el) => {
  const button = el.querySelector('.backPage')
  button.style.display = 'none'
  el.style = 'background-color: transparent'
}

const onAfterLeave = () => {
  if (style) {
    document.head.removeChild(style);
    style = null;
  }
}
</script>

<style lang="less" scoped>
.container {
  display: flex;
  flex-direction: column;
  height: 100%;
  overflow: hidden;
  --el-fill-color-blank: var(--background-primary);
  --el-text-color-regular: var(--text-primary);
  --el-border-color-light: var(--border-primary);
  --el-text-color-primary: var(--text-primary);
  --el-bg-color-overlay: var(--background-primary);

  .header-container {
    margin-top: 6px;
    height: auto;
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 4vw; /* 设置子元素之间的间距 */

    .latex-input {
      width: 220px;

      .el-button {
        background-color: var(--button-hover);
      }
    }
  }

  .main-container {
    margin: auto;

    .card-container {
      justify-items: center;
      max-height: 86vh;
      overflow-y: auto;
      margin: 0 auto;
      display: grid;
      grid-gap: 16px; // adjust to your preference
      grid-template-columns: repeat(1, 1fr);
      transition: grid-template-columns 0.3s cubic-bezier(0.2, 1, 0.8, 1);

      @media (min-width: 768px) {
        grid-template-columns: repeat(2, 1fr);
      }

      @media (min-width: 992px) {
        grid-template-columns: repeat(3, 1fr);
      }

      @media (min-width: 1200px) {
        grid-template-columns: repeat(4, 1fr);
      }

      @media (min-width: 1500px) {
        grid-template-columns: repeat(5, 1fr);
      }

      .card {
        border-radius: 10px; //如果你需要一个具有圆角的卡片样式，可以添加此行
        box-sizing: border-box;
        aspect-ratio: 16 / 14;
        position: relative;  /* 让 .card 成为 .bottom 的定位参考点 */
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-items: flex-start;
        .card-content {
          display: flex;
          flex-direction: column;
          justify-content: space-between;
          flex: 1;
        }
        .card-image {
          width: auto; /* 让图片保持原始比例 */
          max-width: 100%; /* 限制图片的最大宽度 */
          //height: auto;
          cursor: pointer; //悬浮时变手指
          // 居中显示
          object-fit: cover;
          display: flex;
          align-items: center;
          justify-content: center;
          height: 10vh;
        }

        .card-api-type {
          margin: auto;
          //margin-bottom: 0;
        }

        .card-bottom {
          position: absolute;  /* 让 .bottom 能够定位到 .card 底部 */
          bottom: 13px;     /* 把 .bottom 置于 .card 的底部 */
          line-height: 25px;
          display: flex;
          justify-content: space-between;
          align-items: center;
          .card-time {
            font-size: 12px;
            color: #999;
            word-wrap: break-word;  /* 允许在字符间断开 */
            white-space: normal;  /* 保持单词完整并在必要时换行 */
            flex: 1 1 0;  /* 元素可以自由增长、缩小，并且它的基础大小是0 */
            max-width: 50%;  /* 限制元素的最大宽度为整个父容器宽度的50% */
            text-align: left; //将时间靠左对齐
          }
          .card-btn {
            text-align: right; //将按钮靠右对齐
            flex: 1 1 0;  /* 元素可以自由增长、缩小，并且它的基础大小是0 */
            max-width: 50%;  /* 限制元素的最大宽度为整个父容器宽度的50% */
          }
        }

      }

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
        border-radius: 4px;
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

    .overlay-box {
      position: fixed;
      display: flex;
      justify-content: center;
      align-items: center;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      background-color: var(--mask-paper) !important;
      z-index: 9999;

      .backPage {
        position: fixed;
        top: 5%;
        left: 3%;
        justify-content: center;
        align-items: center;
        width: 40px;
        height: 40px;
        border-radius: 40px;
        border: 1px solid var(--color-border);
        cursor: pointer;
        transition: all .3s;
      }

      .overlay-card {
        overflow: hidden;
        transform-origin: left top;
        //
        box-sizing: border-box;
        width: 70%;
        height: 80%;
        border-radius: 20px; /* 这里定义圆角的像素值 */
        display: flex;
        flex-direction: column; // 确保子元素垂直
        .overlay-card-upper-container {
          flex: 1; // 使用 flex 1 来进行相等分配
          display: flex;
          align-items: center;
          justify-content: space-between;

          .center-image {
            display: flex;
            justify-content: center;
            align-items: center;
            width: 30%;
            border-radius: 20px; /* 这里定义圆角的像素值 */
            border: 1px dashed var(--border-secondary);
            height: 40vh;

            .origin-image {
              border-radius: 20px; /* 这里定义圆角的像素值 */
              max-width: 100%; /* 限制图片的最大宽度 */
              max-height: 100%;

            }
          }

          .el-scrollbar {
            box-sizing: border-box;
            border: 1px dashed var(--border-secondary);
            border-radius: 20px; /* 这里定义圆角的像素值 */
            width: 70%;

            .latex-box {
              width: 100%;
              height: 100%;
              //禁止右键菜单
              pointer-events: none;
              //overflow: hidden; // 这样超出的部分就会显示滚动条
              svg {
                max-width: 100%;
                max-height: 100%;
              }
            }
          }

        }

        .overlay-card-lower-container {
          flex: 1;
          display: flex;
          flex-direction: column;
          align-items: center;
          justify-content: space-between;
          gap: 5px; /* Control the vertical spacing between elements */
          padding: 14px;

          .spacer {
            flex-grow: 1;
            display: flex;
          }

          .el-radio-group,
          .textarea-box {
            padding: 10px 0;
          }

          .textarea-box {
            //margin-bottom: auto;
            position: relative;
            width: 100%;
            height: 28vh;

            .el-button {
              position: absolute;
              top: 5px;
              right: 5px;
              //margin: 5px; /* 可以调整按钮与容器边缘的距离 */
              //color: transparent;/
              z-index: 9999;
              border: none;
              outline: none; /* 去除选中时的边框 */
            }

            .text-textarea {
              border-radius: 5px; /* 这里定义圆角的像素值 */
              border: 1px dashed var(--border-secondary);
              background-color: transparent;
              width: 100%;
              height: 100%;
              //border: none;
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

            .bottom {
              margin-top: 13px;
              line-height: 12px;
              display: flex;
              justify-content: space-between;
              align-items: center;

              .time {
                font-size: 12px;
                color: #999;
              }
            }
          }
        }
      }

    }

    .fade-enter-active {
      animation: scale-up-center 0.5s cubic-bezier(0.250, 0.460, 0.450, 0.940) both;
    }

    .fade-leave-active {
      animation: scale-up-center 0.3s ease-out both reverse;
    }
  }
}
</style>