// 获取指定id设备的视频流
export function getInitStream(
    source: { id: string },
    audio?: boolean
): Promise<MediaStream | null> {
    console.log(source);

    return new Promise((resolve, _reject) => {
        // 获取指定窗口的媒体流
        // 此处遵循的是webRTC的接口类型  暂时TS类型没有支持  只能断言成any
        (navigator.mediaDevices as any)
            .getUserMedia({
                audio: audio
                    ? {
                        mandatory: {
                            chromeMediaSource: "desktop",
                        },
                    }
                    : false,
                video: {
                    mandatory: {
                        chromeMediaSource: "desktop",
                        chromeMediaSourceId: source.id,
                    },
                },
            })
            .then((stream: MediaStream) => {
                resolve(stream);
            })
            .catch((error: any) => {
                handleError(error);
                resolve(null);
            });
    });
}
function handleError(e: any) {
    console.log(e);
}
// 将视频流推送到video标签
export function handleStream(
    stream: MediaProvider | null,
    selectors: string,
    mode?: "query" | "create"
) {
    if (stream === null) return;
    if (!mode) mode = "query";
    let video: HTMLVideoElement | null = null;
    switch (mode) {
        case "query":
            if (selectors === "" || document.querySelector(selectors) === null)
                return null;
            video = document.querySelector(selectors);
            break;
        case "create":
            video = document.createElement("video");
            break;
        default:
            break;
    }
    if (video === null) return null;
    video.srcObject = stream;
    video.onloadedmetadata = (_e) => video!.play();
    return video;
}

export function getVideoDevices(): Promise<MediaDeviceInfo[]> {
    return new Promise((resolve, reject) => {
        navigator.mediaDevices
            .enumerateDevices()
            .then((devices) => {
                resolve(devices.filter((device) => device.kind == "videoinput"));
            })
            .catch((error) => {
                reject(error);
            });
    });
}
/**
 * @param event dom绑定事件时的默认事件传参(vue的dom上使用$event传入)
 * @param selectors video标签的css选择器描述
 * @param deviceid 视频输入设备Id
 * @returns 返回video的dom,根据设备id存在与否,给出错误提示
 */
export function getVideoStream(
    event: MouseEvent,
    selectors: string,
    deviceid?: string
) {
    const target = (
        event.target === event.currentTarget ? event.target : event.currentTarget
    ) as HTMLElement;
    deviceid = deviceid || target.dataset["deviceid"];

    const mediaStream = navigator.mediaDevices.getUserMedia({
        video: {
            deviceId: deviceid,
        },
    });
    const video = document.querySelector(selectors) as HTMLVideoElement;
    mediaStream.then((stream) => {
        video.srcObject = stream;
        video.play();
    });

    return {
        video: video,
        message: deviceid ? "" : "没有获取到设备ID",
        error: deviceid ? 0 : 1,
    };
}

export async function getScreenStream(
    selectors: string,
    id: string,
    audio?: boolean
) {
    const mediaStream = await getInitStream({ id: id }, audio);
    const video = document.querySelector(selectors) as HTMLVideoElement;
    video.srcObject = mediaStream;
    // video.play();
    console.log(mediaStream, video);

    return {
        video: video,
        stream: mediaStream,
        message: mediaStream ? "" : "没有获取到设备ID",
        error: mediaStream ? 0 : 1,
    };
}

// 灰度图
export function getGrayData(data: ImageData) {
    // sRGB颜色模式的彩色图像数据  数据长度是data.width*data.height的4倍
    // 也就是说彩色sRGB有4个颜色通道
    // 除r、g、b三色之外  还有一个叫做"白色点"的色值  它的色值为[255,255,255]，这是用来表示纯白色的，而不是真正的灰度色
    const imgData = new ImageData(data.width, data.height);

    for (let i = 0; i < data.data.length; i += 4) {
        const r = data.data[i];
        const g = data.data[i + 1];
        const b = data.data[i + 2];
        const gray = (r + g + b) / 3;
        imgData.data[i] = gray;
        imgData.data[i + 1] = gray;
        imgData.data[i + 2] = gray;
        imgData.data[i + 3] = data.data[i + 3];
    }
    return imgData;
}