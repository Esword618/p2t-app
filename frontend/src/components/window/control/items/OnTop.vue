<template>
    <!-- 置顶——状态切换 -->
    <div :title="isOntoped ? '取消置顶' : '置顶'" class="sticky" @click="onTopWindow">
        <svg viewBox="0 0 1024 1024" width="15" height="15">
            <path v-for="(d, _index) in onTopPath" :d="d"></path>
        </svg>
    </div>
</template>

<script lang="ts" setup>
import { PropType, ref } from 'vue';
import { onTopPath } from "../assets";

const props = defineProps({
    WindowOnTop: {
        type: Function as PropType<() => Promise<void> | void>,
        required: true
    },
    WindowIsOnToped: {
        type: Function as PropType<() => Promise<boolean> | boolean>,
        required: true
    },
    style: {
        type: Object as PropType<{
            fill?: string;
            activeColor?: string;
            hoverColor?: string;
        }>,
        required: false,
    }
})

let isOntoped = ref(await props.WindowIsOnToped());
const style = ref({
    stickyRotate: isOntoped.value ? "-45deg" : "0deg",
    stickyColor: isOntoped.value ? (props.style?.activeColor || "limegreen") : props.style?.fill,
})

async function onTopWindow() {
    // OnTopWindow
    await props.WindowOnTop()
    isOntoped.value = await props.WindowIsOnToped()
    if (isOntoped.value) {
        style.value.stickyRotate = "-45deg"
        style.value.stickyColor = (props.style?.activeColor || "limegreen")
    } else {
        style.value.stickyRotate = "0deg"
        style.value.stickyColor = props.style?.fill
    }
}
</script>



<style lang="less" scoped>
span {
    cursor: pointer;
    display: inline-block;
    height: 100%;
    //width: 47px;
    overflow: hidden;
    padding: 2px;
    position: relative;
    transition: background-color 300ms;

    &::before {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
    }
}

.sticky {
    position: relative;
    display: block;
    width: 43px;
    fill: v-bind("style.stickyColor");
    cursor: pointer;

    svg {
        position: absolute;
        display: inline-flex;
        top: 50%;
        left: 50%;
        transition: transform 300ms;
        transform: translateX(-50%) translateY(-50%) rotate(v-bind("style.stickyRotate"));
    }

    &:hover {
        background-color: v-bind("props.style?.hoverColor");
    }
}
</style>