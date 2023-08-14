import Lenis from '@studio-freight/lenis'

export const initSmoothScrolling = (target: HTMLElement) => {
    const lenis = new Lenis({
        lerp: 0.1,
        smoothWheel: true,
        content: target,
        orientation: 'vertical',
    });

    lenis.on('scroll', (_e: any) => {
        // console.log(e);
        // // https://github.com/terwanerik/ScrollTrigger
        // ScrollTrigger.update()
    });

    const scrollFn = (time: number) => {
        lenis.raf(time);
        requestAnimationFrame(scrollFn);
    };
    requestAnimationFrame(scrollFn);
}

export function smoothScroll(target: Window | HTMLElement, targetPosition: number, duration: number) {
    console.log(target, targetPosition, duration);
    let startPosition: number;
    if (target instanceof Window) {
        startPosition = target.scrollY;
    } else {
        startPosition = target.scrollTop;
    }

    const distance = targetPosition - startPosition;
    let startTime: number | null = null;

    function animation(currentTime: number) {
        if (startTime === null) startTime = currentTime;
        const timeElapsed = currentTime - startTime;
        const run = ease(timeElapsed, startPosition, distance, duration);
        target.scrollTo(0, run);
        if (timeElapsed < duration) requestAnimationFrame(animation);
    }

    function ease(t: number, b: number, c: number, d: number) {
        t /= d / 2;
        if (t < 1) return c / 2 * t * t + b;
        t--;
        return -c / 2 * (t * (t - 2) - 1) + b;
    }
    requestAnimationFrame(animation);
}

export class SmoothScroller {
    private target: Window | HTMLElement;
    private duration: number;
    private animationId: number | undefined;
    // @ts-ignore
    private flag: boolean = true;

    constructor(target: Window | HTMLElement, duration: number) {
        this.target = target;
        this.duration = duration;
    }

    public scrollTo(targetPosition: number) {
        // if (!this.flag) return;
        // this.flag = false;
        // this.target.scrollTo({
        //     top: targetPosition,
        //     behavior: 'smooth'
        // });
        // setTimeout(() => {
        //     this.flag = true;
        // }, this.duration);

        let startPosition: number;
        if (this.target instanceof Window) {
            startPosition = this.target.scrollY;
        } else {
            startPosition = this.target.scrollTop;
        }

        const distance = targetPosition - startPosition;
        let startTime: number | null = null;
        const animation = (currentTime: number) => {
            if (startTime === null) startTime = currentTime;
            const timeElapsed = currentTime - startTime;
            const run = ease(timeElapsed, startPosition, distance, this.duration);
            this.target.scrollTo(0, run);
            if (timeElapsed < this.duration) {
                this.animationId = requestAnimationFrame(animation);
            };
        }
        this.animationId = requestAnimationFrame(animation.bind(this));
    }

    public stop() {
        // console.log(this.flag);
        if (this.animationId !== undefined) {
            // console.log(this.animationId);
            cancelAnimationFrame(this.animationId);
        }
    }
}

// 缓动函数 timeElapsed, startPosition, distance, Allduration
function ease(t: number, b: number, c: number, d: number) {
    t /= d / 2;
    if (t < 1) return c / 2 * t * t + b;
    t--;
    return -c / 2 * (t * (t - 2) - 1) + b;
}
// function easeInOutQuad(t: number, b: number, c: number, d: number) {
//     t /= d / 2;
//     if (t < 1) return c / 2 * t * t + b;
//     t--;
//     return -c / 2 * (t * (t - 2) - 1) + b;
// }
// function easeInOutCubic(t: number, b: number, c: number, d: number) {
//     t /= d / 2;
//     if (t < 1) return c / 2 * t * t * t + b;
//     t -= 2;
//     return c / 2 * (t * t * t + 2) + b;
// }
// function easeInOutQuart(t: number, b: number, c: number, d: number) {
//     t /= d / 2;
//     if (t < 1) return c / 2 * t * t * t * t + b;
//     t -= 2;
//     return -c / 2 * (t * t * t * t - 2) + b;
// }
// function easeInOutQuint(t: number, b: number, c: number, d: number) {
//     t /= d / 2;
//     if (t < 1) return c / 2 * t * t * t * t * t + b;
//     t -= 2;
//     return c / 2 * (t * t * t * t * t + 2) + b;
// }
// // @ts-ignore
// function easeInOutSine(t: number, b: number, c: number, d: number) {
//     return -c / 2 * (Math.cos(Math.PI * t / d) - 1) + b;
// }
// function easeInOutExpo(t: number, b: number, c: number, d: number) {
//     t /= d / 2;
//     if (t < 1) return c / 2 * Math.pow(2, 10 * (t - 1)) + b;
//     t--;
//     return c / 2 * (-Math.pow(2, -10 * t) + 2) + b;
// }
// function easeInOutCirc(t: number, b: number, c: number, d: number) {
//     t /= d / 2;
//     if (t < 1) return -c / 2 * (Math.sqrt(1 - t * t) - 1) + b;
//     t -= 2;
//     return c / 2 * (Math.sqrt(1 - t * t) + 1) + b;
// }