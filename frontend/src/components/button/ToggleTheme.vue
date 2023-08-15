<template>
    <button class="theme-toggle" :class="{ 'dark': isDark }" @click="toggleTheme">
        <svg class="sun-and-moon" aria-hidden="true" style="height: fit-content;width: fit-content;" viewBox="0 0 24 24"
            ref="svg">
            <mask class="moon" id="moon-mask">
                <rect x="0" y="0" width="100%" height="100%" fill="white"></rect>
                <circle cx="24" cy="10" r="6" fill="black"></circle>
            </mask>
            <circle class="sun" cx="12" cy="12" r="6" mask="url(#moon-mask)" fill="currentColor"></circle>
            <g class="sun-beams" stroke="currentColor">
                <line x1="12" y1="1" x2="12" y2="3"></line>
                <line x1="12" y1="21" x2="12" y2="23"></line>
                <line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line>
                <line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line>
                <line x1="1" y1="12" x2="3" y2="12"></line>
                <line x1="21" y1="12" x2="23" y2="12"></line>
                <line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line>
                <line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line>
            </g>
        </svg>
    </button>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
const props = defineProps({
    size: {
        type: Number,
    },
    duration: {
        type: Number,
    }
})

const isDark = ref(false)
const toggleTheme = (event: MouseEvent) => {
    const x = event.clientX;
    const y = event.clientY;
    const endRadius = Math.hypot(
        Math.max(x, innerWidth - x),
        Math.max(y, innerHeight - y)
    );

    let isDark: boolean;
    const switchTheme = (() => {
        const root = document.documentElement;
        isDark = root.classList.contains("dark");
        root.classList.remove(isDark ? "dark" : "light");
        root.classList.add(isDark ? "light" : "dark");
    })
    // @ts-ignore 适配不支持startViewTransition的浏览器
    if (!document?.startViewTransition) {
        switchTheme();
        return;
    }
    // @ts-ignore
    const transition = document.startViewTransition(switchTheme);

    transition.ready.then(() => {
        const clipPath = [
            `circle(0px at ${x}px ${y}px)`,
            `circle(${endRadius}px at ${x}px ${y}px)`,
        ];
        document.documentElement.animate(
            {
                clipPath: isDark ? [...clipPath].reverse() : clipPath,
            },
            {
                duration: props.duration || 500,
                easing: "ease-in",
                pseudoElement: isDark
                    ? "::view-transition-old(root)"
                    : "::view-transition-new(root)",
            }
        );
    });

}

const svg = ref<SVGElement>();
onMounted(() => {
    if (svg.value) {
        const width = svg.value.clientWidth;
        const height = svg.value.clientHeight;
        svg.value.style.height = (props.size || width) + "px";
        svg.value.style.width = (props.size || height) + "px";
    }
})
</script>
<style>
::view-transition-old(root),
::view-transition-new(root) {
    animation: none;
    mix-blend-mode: normal;
}

/* 进入dark模式和退出dark模式时，两个图像的位置顺序正好相反 */
.dark::view-transition-old(root) {
    z-index: 1;
}

.dark::view-transition-new(root) {
    z-index: 999;
}

::view-transition-old(root) {
    z-index: 999;
}

::view-transition-new(root) {
    z-index: 1;
}

:where(html) {
    --ease-1: cubic-bezier(.25, 0, .5, 1);
    --ease-2: cubic-bezier(.25, 0, .4, 1);
    --ease-3: cubic-bezier(.25, 0, .3, 1);
    --ease-4: cubic-bezier(.25, 0, .2, 1);
    --ease-5: cubic-bezier(.25, 0, .1, 1);
    --ease-in-1: cubic-bezier(.25, 0, 1, 1);
    --ease-in-2: cubic-bezier(.50, 0, 1, 1);
    --ease-in-3: cubic-bezier(.70, 0, 1, 1);
    --ease-in-4: cubic-bezier(.90, 0, 1, 1);
    --ease-in-5: cubic-bezier(1, 0, 1, 1);
    --ease-out-1: cubic-bezier(0, 0, .75, 1);
    --ease-out-2: cubic-bezier(0, 0, .50, 1);
    --ease-out-3: cubic-bezier(0, 0, .3, 1);
    --ease-out-4: cubic-bezier(0, 0, .1, 1);
    --ease-out-5: cubic-bezier(0, 0, 0, 1);
    --ease-in-out-1: cubic-bezier(.1, 0, .9, 1);
    --ease-in-out-2: cubic-bezier(.3, 0, .7, 1);
    --ease-in-out-3: cubic-bezier(.5, 0, .5, 1);
    --ease-in-out-4: cubic-bezier(.7, 0, .3, 1);
    --ease-in-out-5: cubic-bezier(.9, 0, .1, 1);
    --ease-elastic-1: cubic-bezier(.5, .75, .75, 1.25);
    --ease-elastic-2: cubic-bezier(.5, 1, .75, 1.25);
    --ease-elastic-3: cubic-bezier(.5, 1.25, .75, 1.25);
    --ease-elastic-4: cubic-bezier(.5, 1.5, .75, 1.25);
    --ease-elastic-5: cubic-bezier(.5, 1.75, .75, 1.25);
    --ease-squish-1: cubic-bezier(.5, -.1, .1, 1.5);
    --ease-squish-2: cubic-bezier(.5, -.3, .1, 1.5);
    --ease-squish-3: cubic-bezier(.5, -.5, .1, 1.5);
    --ease-squish-4: cubic-bezier(.5, -.7, .1, 1.5);
    --ease-squish-5: cubic-bezier(.5, -.9, .1, 1.5);
    --ease-step-1: steps(2);
    --ease-step-2: steps(3);
    --ease-step-3: steps(4);
    --ease-step-4: steps(7);
    --ease-step-5: steps(10)
}
</style>
<style lang="less" scoped>
.theme-toggle {
    --size: 1.25rem;
    --icon-fill: #404040;
    --icon-fill-hover: #d97706;
    background: none;
    border: none;
    padding: 0;
    inline-size: var(--size);
    block-size: var(--size);
    aspect-ratio: 1;
    border-radius: 50%;
    cursor: pointer;
    touch-action: manipulation;
    -webkit-tap-highlight-color: transparent;
    outline-offset: 5px;

    &:is(:hover, :focus-visible)>.sun-and-moon>:is(.moon, .sun) {
        fill: var(--icon-fill-hover)
    }

    &:is(:hover, :focus-visible) .sun-and-moon>.sun-beams {
        stroke: var(--icon-fill-hover)
    }

    .sun-and-moon {
        inline-size: 100%;
        block-size: 100%;
        stroke-linecap: round;

        .sun-beams {
            stroke: var(--icon-fill);
            stroke-width: 2px
        }

        &>:is(.moon, .sun, .sun-beams) {
            transform-origin: center center
        }

        &>:is(.moon, .sun) {
            fill: var(--icon-fill);
        }
    }
}

.dark {

    &.theme-toggle,
    .theme-toggle {
        --icon-fill: #e5e5e5;
        --icon-fill-hover: #3b82f6;

        .sun-and-moon {
            .sun {
                transform: scale(1.75);
            }

            .sun-beams {
                opacity: 0;
            }

            .moon {
                circle {
                    transform: translate(-7px);

                    @supports (cx: 1) {
                        transform: translate(0);
                        cx: 17;
                    }
                }
            }
        }
    }
}

@media (prefers-reduced-motion: no-preference) {
    .sun-and-moon {
        .sun {
            transition: transform .5s var(--ease-elastic-3);
        }

        .sun-beams {
            transition: transform .5s var(--ease-elastic-4), opacity .5s var(--ease-3);
        }

        .moon {
            circle {
                transition: transform .25s var(--ease-out-5);

                @supports (cx: 1) {
                    transition: cx .25s var(--ease-out-5);
                }
            }
        }
    }

    .dark {
        .sun-and-moon {
            .sun {
                transform: scale(1.75);
                transition-timing-function: var(--ease-3);
                transition-duration: .25s
            }

            .sun-beams {
                transform: rotate(-25deg);
                transition-duration: .15s;
            }

            .moon {
                circle {
                    transition-delay: .25s;
                    transition-duration: .5s;
                }
            }
        }
    }

}
</style>