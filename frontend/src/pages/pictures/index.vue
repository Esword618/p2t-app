<template>
  <div class="container">
    <el-container>
      <el-header class="header-container">
        <div class="block">
          <el-date-picker
              v-model="selectData"
              type="date"
              :disabled-date="disabledDate"
              placeholder="Pick a day"
          />
        </div>
        <el-input
            class="latex-input"
            v-model="latexInput"
            placeholder="Please input latex text"
        >
          <template #append>
            <el-button>
              <el-icon>
                <Search/>
              </el-icon>
            </el-button>
          </template>
        </el-input>
      </el-header>
      <el-main class="main-container">
        <div @mouseenter="showScrollbar = true"
             @mouseleave="showScrollbar = false"
             class="card-container"
        >
          <el-card
              v-for="(o) in cards"
              :key="o"
              class="card"
          >
            <el-image
                class="image"
                src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                :zoom-rate="1.2"
                :preview-src-list="['https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png']"
                :initial-index="4"
                fit="cover"
            />
            <div style="padding: 14px">
              <span>Yummy hamburger</span>
              <div class="bottom">
                <time class="time">{{ currentDate }}</time>
                <el-button class="detail-button"
                           type="primary"
                           round
                           size="small"
                           @click="onDetail('11111')"
                >
                  <el-icon>
                    <MagicStick/>
                  </el-icon>
                  详细
                </el-button>
              </div>
            </div>
          </el-card>
        </div>
      </el-main>
    </el-container>
  </div>
</template>
<script lang="ts">
import {PageEnum} from "@/enums/pageEnum";

export default {
  name: PageEnum.BASE_PICTURES_NAME,
};
</script>
<script lang="ts" setup>
import {ref} from 'vue'
import {Search, MagicStick} from "@element-plus/icons-vue";

const selectData = ref('')
const disabledDate = (time: Date) => {
  return time.getTime() > Date.now()
}

const currentDate = ref(new Date())
const showScrollbar = ref(false)
const cards = ref(Array.from({length: 20})) // Change the number as needed
const latexInput = ref("")


const onDetail = (uuid: String) => {
  console.log(uuid);
}

</script>

<style lang="less" scoped>
.container {
  display: flex;
  flex-direction: column;
  height: 100%;
  //width: 95vw;
  overflow: hidden;
  --el-fill-color-blank: var(--background-primary);
  --el-text-color-regular: var(--text-primary);
  --el-border-color-light: var(--border-primary);
  --el-text-color-primary: var(--text-primary);
  --el-bg-color-overlay:var(--background-primary);
  .header-container {
    margin-top: 6px;
    height: auto;
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 4vw; /* 设置子元素之间的间距 */

    .latex-input {
      //flex: none;
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
      //max-width: 100%;
      display: grid;
      grid-gap: 16px; // adjust to your preference
      grid-template-columns: repeat(1, 1fr);
      transition: grid-template-columns 0.3s cubic-bezier(0.2, 1, 0.8, 1);

      @media (min-width: 768px) {
        grid-template-columns: repeat(2, 1fr);
      }

      @media (min-width: 992px) {
        grid-template-columns: repeat(4, 1fr);
      }

      @media (min-width: 1200px) {
        grid-template-columns: repeat(5, 1fr);
      }

      @media (min-width: 1500px) {
        grid-template-columns: repeat(6, 1fr);
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

      .card {
        box-sizing: border-box;
        //margin: 8px;
        //transition: grid-column-end 0.3s cubic-bezier(0.2, 1, 0.8, 1);
        .time {
          font-size: 12px;
          color: #999;
        }

        .bottom {
          margin-top: 13px;
          line-height: 12px;
          display: flex;
          justify-content: space-between;
          align-items: center;
        }

        .image {
          width: auto; /* 让图片保持原始比例 */
          max-width: 100%; /* 限制图片的最大宽度 */
          height: auto;
          display: block;
        }
      }
    }
  }
}
</style>