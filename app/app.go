package app

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.design/x/clipboard"
	"log"
	"strings"
)

// App struct
type App struct {
	ctx  context.Context
	file []byte
}

// NewApp creates a new App application struct
// NewApp 创建一个新的 App 应用程序
func NewApplication() *App {
	return &App{}
}

// startup is called at application startup
// startup 在应用程序启动时调用
func (a *App) Startup(ctx context.Context) {

	// Perform your setup here
	// 在这里执行初始化设置
	a.ctx = ctx

	// 窗口居中
	runtime.WindowCenter(a.ctx)

	// 获取 前端 发来的命令解析
	runtime.EventsOn(ctx, "parse_img", func(optionalData ...interface{}) {
		a.parse()
	})

	// 接收 前端 传来的base64图片
	runtime.EventsOn(ctx, "set_base64_img", func(optionalData ...interface{}) {
		base64Img := optionalData[0].(string)
		base64Img = strings.Split(base64Img, ",")[1]
		file, err := base64.StdEncoding.DecodeString(base64Img)
		if err != nil {
			fmt.Println(err)
		}
		a.file = file
	})

	a.initClipboard()

}

// domReady is called after the front-end dom has been loaded
// domReady 在前端Dom加载完毕后调用
func (a *App) DomReady(ctx context.Context) {
	// Add your action here
	// 在这里添加你的操作
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
// beforeClose在单击窗口关闭按钮或调用runtime.Quit即将退出应用程序时被调用.
// 返回 true 将导致应用程序继续，false 将继续正常关闭。
func (a *App) BeforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
// 在应用程序终止时被调用
func (a *App) Shutdown(ctx context.Context) {
	// Perform your teardown here
	// 在此处做一些资源释放的操作
}

func (a *App) initClipboard() {
	clipboard.Init()
	ch := clipboard.Watch(context.TODO(), clipboard.FmtImage)
	for img := range ch {
		// print out clipboard data whenever it is changed
		//println(string(data))
		//img := clipboard.Read(clipboard.FmtImage)
		a.file = img
		base64Img := fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(img))
		//fmt.Println(base64Img)
		//fmt.Println("go")
		//EventsEmit发消息
		//EventsOn接收消息
		runtime.EventsEmit(a.ctx, "clipboard_image", base64Img)
		runtime.WindowShow(a.ctx)
	}
}

func (a *App) parse() {
	//fmt.Println("parse")
	client := req.C()
	res, err := client.R().
		SetHeaders(map[string]string{
			"cookie":             "__atuvc=1^%^7C8; __atuvs=63f992ec1481c1e1000",
			"origin":             "https://p2t.behye.com",
			"referer":            "https://p2t.behye.com/",
			"sec-ch-ua-mobile":   "?0",
			"sec-fetch-dest":     "empty",
			"sec-ch-ua-platform": "Windows",
			"sec-fetch-mode":     "cors",
			"sec-fetch-site":     "same-origin",
			"user-agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
		}).
		SetFileBytes("image", "image", a.file). // Set form param name and filename
		SetFormData(map[string]string{ // Set form data while uploading
			"session_id": "session-AiYJrY-bxFfzzCwOA9Kb4cyWfpdJnt6q",
		}).
		Post("https://p2t.behye.com/api/pix2text")
	if err != nil {
		log.Println(err)
	}
	result, _ := res.ToString()
	fmt.Println(result)
	runtime.EventsEmit(a.ctx, "parse_result", result)
	//fmt.Println("发送完成")
}
