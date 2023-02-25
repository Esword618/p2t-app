// 想要使用必须先引入 defineStore；
import { defineStore } from "pinia";
import { reactive, toRefs } from "vue";

// defineStore 方法有两个参数，第一个参数是模块化名字（也就相当于身份证一样，不能重复）

// 定义状态的类型 一般情况下Pinia会自动推断出state的类型，你就无需定义接口。如果state的数据类型比较复杂，推荐你定义一个接口来描述state的类型
interface ImgState {
  imgBase64Url: string;
  showUploadImg: boolean;
}

const userStore = defineStore("userStore", () => {
  const state = reactive<ImgState>({
    imgBase64Url: "",
    showUploadImg: false,
  });

  const getImgBase64Url = (): string => state.imgBase64Url;
  const getShowUploadImg = (): boolean => state.showUploadImg;

  return {
    ...toRefs(state),
    getImgBase64Url,
    getShowUploadImg,
  };
});

export default userStore;
