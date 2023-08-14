<template>
  <div class="container">
    <div class="form-box">
      <el-form size="large" label-position="right" :model="formSetting" label-width="auto">
<!--        <el-text size="large">-->
<!--          <el-icon color="#1579e3"><WindPower /></el-icon>-->
<!--          模式选择-->
<!--        </el-text>-->
        <el-form-item label="模式选择">
          <el-radio-group v-model="formSetting.formulatype">
            <el-radio label="pix2text" />
            <el-radio label="simpletex" />
          </el-radio-group>
        </el-form-item>
        <el-form-item label="自动检测图片打开软件">
          <el-switch style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
                     active-text="on"
                     inactive-text="off"
                     v-model="formSetting.autoshow" />
        </el-form-item>
        <el-form-item label="Simpletex AppId">
          <el-input
              v-model="formSetting.appid"
              placeholder="Please input Simpletex AppId"
              clearable show-password
          />
        </el-form-item>
        <el-form-item label="Simpletex AppSecret">
          <el-input
              v-model="formSetting.appsecret"
              placeholder="Please input Simpletex AppSecret"
              clearable show-password
          />
        </el-form-item>
        <el-form-item label="自动更新(开发中)">
          <el-switch style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
                     active-text="on"
                     inactive-text="off"
                     v-model="formSetting.selfupdate" />
        </el-form-item>
        <el-form-item label="最小化至托盘">
          <el-switch style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
                     active-text="on"
                     inactive-text="off"
                     v-model="formSetting.min2tray" />
        </el-form-item>
        <!--        <el-form-item label="主题设置">-->
        <!--          <el-radio-group v-model="formSetting.theme">-->
        <!--            <el-radio label="light" />-->
        <!--            <el-radio label="dark" />-->
        <!--          </el-radio-group>-->
        <!--        </el-form-item>-->
      </el-form>
    </div>
  </div>
</template>

<script lang="ts" setup>
//ElementPlus
import {onMounted, ref, watch} from 'vue'
import {GetConfig,UpdateConfig} from "@wailsjs/go/utils/Utils";
import {useRoute} from "vue-router";

interface FormSettingType {
  savedir: string;
  theme: string;
  formulatype: string;
  autoshow: boolean;
  selfupdate: boolean;
  min2tray: boolean;
  appid: string;
  appsecret: string;
}

const route = useRoute();

const formSetting = ref(<FormSettingType>{
  savedir:"",
  theme:'light',
  formulatype:'pix2text',
  autoshow: true,
  selfupdate:false,
  min2tray:false,
  appid: "",
  appsecret: "",
})

// ts-ignore
watch(formSetting,async (newFormSetting) => {
  await UpdateConfig(newFormSetting)
},{deep:true})

// 监控路由变化以便重新获取数据
const initLoadSetting = async () => {
  if (route.name == "Setting") {
    formSetting.value = await GetConfig() as FormSettingType
  }
}
watch(() => route.params, initLoadSetting)

onMounted(async () => {
  formSetting.value = await GetConfig() as FormSettingType
})

</script>
<script lang="ts">
import { PageEnum } from "@/enums/pageEnum";
export default {
  name: PageEnum.BASE_SETTING_NAME,
};
</script>

<style lang="less" scoped>
.container {
  display: flex;
  flex-direction: column;
  height: 100%;
  width: 100%;
  align-items: center;
  --el-fill-color-blank: var(--background-primary);
  --el-text-color-primary: var(--text-primary);
  --el-bg-color: var(--background-primary);
  --el-dialog-bg-color: var(--background-primary);
  --el-text-color-regular: var(--text-primary);
  --el-fill-color-light:var(--fill-color-light);
  --el-border-color-light:var(--background-toc-active);
  padding: 50px;
  .form-box {
    margin-bottom: 20px;
    box-sizing: border-box;
    margin-top: 20px;   /* 修改此处的值以调整 form-box 距离顶部的距离 */
    width: 80%;
    padding: 20px;
    border-radius: 10px;
  }
}
</style>
