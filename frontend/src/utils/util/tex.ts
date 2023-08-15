// @ts-nocheck
import {convertSvgElement2Png} from "@/utils/util/svg2png.ts";


export const renderLatex = (elementId: string,latexStr:string): Promise<void> => {
    // const previewElement = document.getElementById('history-preview')
    // if (!previewElement) return;
    // previewElement.innerHTML = newValue;
    // // 重置
    // // window.MathJax.texReset();
    // // window.MathJax.typesetClear();
    // // window.MathJax.typesetPromise([previewElement]).then(() => {});
    return new Promise(async (resolve, reject) => {
        const previewElement = document.getElementById(elementId);
        if (!previewElement) {
            reject(new Error(`Element with id "${elementId}" not found.`));
            return;
        }

        try {
            previewElement.innerHTML = latexStr;
            window.MathJax.texReset();
            window.MathJax.typesetClear();
            await window.MathJax.typesetPromise([previewElement]);
            resolve();
        } catch(error) {
            reject(error);
        }
    });
};

export const convertTex2Mml = (texStr: string): Promise<string> => {
    return new Promise((resolve, reject) => {
        window.MathJax.texReset();
        window.MathJax.typesetClear();

        window.MathJax.tex2mmlPromise(texStr)
            .then((mml) => {
                resolve(mml);
            })
            .catch((err) => {
                reject(err);
            });
    });
};

export const convertTex2SvgThenPng = (texStr: string): Promise<string> => {
    return new Promise(async (resolve, reject) => {
        try {
            window.MathJax.texReset();
            window.MathJax.typesetClear();
            const htmlElement = await window.MathJax.tex2svgPromise(texStr);
            console.log(htmlElement)
            const svgElement = htmlElement.querySelector("svg"); // 查找SVG元素
            if (!svgElement) throw new Error("Unable to find SVG element.");

            // 确保convertSvgElement2Png已经被导入和定义
            const pngBase64 = await convertSvgElement2Png(svgElement.outerHTML);

            resolve(pngBase64);
        } catch(error) {
            reject(error);
        }
    });
};