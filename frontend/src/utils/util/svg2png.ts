// export const convertSvg2Png = (svgBase64: string): Promise<string> => {
//     return new Promise((resolve, reject) => {
//         let image = new Image();
//
//         image.onload = function() {
//             let canvas = document.createElement('canvas');
//             canvas.width = image.width;
//             canvas.height = image.height;
//             let ctx = canvas.getContext("2d");
//
//             if(ctx === null) {
//                 reject(new Error("Unable to get canvas context"));
//                 return;
//             }
//
//             ctx.drawImage(image, 0, 0);
//             let imageData = ctx.getImageData(0, 0, image.width, image.height).data;
//
//             let minX = Infinity, maxX = 0, minY = Infinity, maxY = 0;
//
//             for (let y = 0; y < image.height; y++) {
//                 for (let x = 0; x < image.width; x++) {
//                     let idx = 4 * (x + image.width * y);
//
//                     if (imageData[idx] > 0 || imageData[idx + 1] > 0 || imageData[idx + 2] > 0 || imageData[idx + 3] > 0) {
//                         minX = Math.min(minX, x);
//                         maxX = Math.max(maxX, x);
//                         minY = Math.min(minY, y);
//                         maxY = Math.max(maxY, y);
//                     }
//                 }
//             }
//
//             let finalCanvas = document.createElement('canvas');
//             finalCanvas.width = maxX - minX;
//             finalCanvas.height = maxY - minY;
//             let context = finalCanvas.getContext('2d');
//
//             if(context === null) {
//                 reject(new Error("Unable to get final canvas context"));
//                 return;
//             }
//
//             context.drawImage(
//                 canvas, minX, minY, maxX - minX, maxY - minY, 0, 0, maxX - minX, maxY - minY
//             );
//
//             resolve(finalCanvas.toDataURL("image/png"));
//         };
//
//         image.onerror = function() {
//             reject(new Error("Error loading image"));
//         };
//
//         image.src = "data:image/svg+xml;base64," + svgBase64;
//     });
// };

export const convertSvg2Png = (svgBase64: string): Promise<string> => {
    return new Promise((resolve, reject) => {
        let image = new Image();
        image.onload = async function () {
            let canvas = document.createElement('canvas');
            canvas.width = image.width;
            canvas.height = image.height;
            let ctx = canvas.getContext("2d");

            if (ctx === null) {
                reject(new Error("Unable to get canvas context"));
                return;
            }

            ctx.drawImage(image, 0, 0);
            let imageData = ctx.getImageData(0, 0, image.width, image.height);

            let minX = Infinity, maxX = 0, minY = Infinity, maxY = 0;

            for (let i = 0; i < imageData.data.length; i += 4) {

                // calculate coordinates
                let x = (i / 4) % canvas.width;
                let y = Math.floor((i / 4) / canvas.width);

                if (imageData.data[i] > 0 || imageData.data[i + 1] > 0 || imageData.data[i + 2] > 0 || imageData.data[i + 3] > 0) {
                    minX = Math.min(minX, x);
                    maxX = Math.max(maxX, x);
                    minY = Math.min(minY, y);
                    maxY = Math.max(maxY, y);
                }
            }

            let finalCanvas = document.createElement('canvas');
            finalCanvas.width = maxX - minX;
            finalCanvas.height = maxY - minY;
            let context = finalCanvas.getContext('2d');

            if (context === null) {
                reject(new Error("Unable to get final canvas context"));
                return;
            }

            context.drawImage(
                canvas, minX, minY, maxX - minX, maxY - minY, 0, 0, maxX - minX, maxY - minY
            );

            resolve(finalCanvas.toDataURL("image/png"));
        };

        image.onerror = function () {
            reject(new Error("Error loading image"));
        };

        image.src = "data:image/svg+xml;base64," + svgBase64;
    });
};
export const convertSvgElement2Png = (svgString: string): Promise<string> => {
    const svgBase64 = window.btoa(unescape(encodeURIComponent(svgString)));
    return convertSvg2Png(svgBase64);
};