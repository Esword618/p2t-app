package main

import (
	"context"
	"fmt"
	"github.com/Esword618/p2t/config"
	"github.com/Esword618/p2t/consts"
	"github.com/Esword618/p2t/global"
	"github.com/Esword618/p2t/utils"
	"github.com/energye/systray"
	"github.com/spf13/viper"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
)

// App struct
type App struct {
	ctx      context.Context
	hide     bool // 是否不退出程序仅隐藏窗口
	isHidden bool // 是否已隐藏窗口
	isOnTop  bool // 是否已置顶窗口
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

/* *****************生命周期***************** */
// startup 启动时生命周期
func (a *App) startup(ctx context.Context) {
	// 初始化运行目录
	err := os.Chdir(GetCurrPath())
	if err != nil {
		panic(err)
	}
	a.ctx = ctx
	a.hide = true
	global.GlobalCtx = ctx
	runtime.WindowSetDarkTheme(a.ctx)
	// 托盘
	go func() {
		systray.Run(a.mySystray, func() {
			fmt.Println("托盘运行")
			//a.hideWindow()
		})
	}()
	// 剪贴板
	go func() {
		utils.WatchClipboardImg()
	}()
}

// domReady 在前端Dom加载完毕后调用
func (a *App) domReady(ctx context.Context) {
	// Add your action here
	// 在这里添加你的操作
}

// beforeClose 退出前生命周期
func (a *App) beforeClose(ctx context.Context) bool {
	//if a.hide {
	//	a.hideWindow()
	//} else {
	//	//
	//	// 配置保存
	//	viper.WriteConfig()
	//	fmt.Println("退出前的操作", a)
	//}
	//return a.hide
	// 返回 true 将阻止程序关闭
	if !a.hide {
		return false
	}
	hide := viper.GetBool("min2tray")
	if hide {
		runtime.WindowHide(a.ctx)
	} else {
		// TODO: 退出前的操作
		// 配置保存
		viper.WriteConfig()
		fmt.Println("退出前的操作", a)
	}
	return hide
}

// shutdown is called at application termination
// 在应用程序终止时被调用
func (a *App) shutdown(ctx context.Context) {
	// 在此处做一些资源释放的操作
}

/* *****************托盘***************** */

// mySystray 托盘
func (a *App) mySystray() {
	systray.SetTemplateIcon(consts.Appicon, consts.Appicon)
	systray.SetTitle("p2t-app")
	systray.SetTooltip("p2t-app")
	systray.SetOnClick(func(menu systray.IMenu) {
		a.showWindow()
	})
	systray.SetOnDClick(func(menu systray.IMenu) {
		a.showWindow()
	})
	systray.SetOnRClick(func(menu systray.IMenu) {
		menu.ShowMenu()
	})
	// 主页
	mHome := systray.AddMenuItem("主页", "主页")
	mHome.Click(func() {
		runtime.EventsEmit(a.ctx, "try_msg", "/home")
		a.showWindow()

	})
	// 历史记录
	mDownload := systray.AddMenuItem("历史记录", "历史记录")
	mDownload.Click(func() {
		runtime.EventsEmit(a.ctx, "try_msg", "/history")
		a.showWindow()
	})
	// 偏好设置
	mSetting := systray.AddMenuItem("设置", "设置")
	mSetting.Click(func() {
		runtime.EventsEmit(a.ctx, "try_msg", "/setting")
		a.showWindow()
	})
	// 分隔符
	systray.AddSeparator()
	//版本
	version := fmt.Sprintf("版本(%s)", config.Version)
	systray.AddMenuItem(version, version)
	// 项目github地址
	mUrl := systray.AddMenuItem("项目地址", "Github")
	mUrl.Click(func() {
		runtime.BrowserOpenURL(a.ctx, consts.GithubUrl)
	})
	// 关于
	mAbout := systray.AddMenuItem("关于p2t-app", "关于p2t-app")
	mAbout.Click(func() {
		runtime.EventsEmit(a.ctx, "try_msg", "about")
		a.showWindow()
	})
	// 分隔符
	systray.AddSeparator()
	// 退出
	mQuit := systray.AddMenuItem("退出", "退出")
	mQuit.Click(func() {
		a.hide = false
		runtime.Quit(a.ctx)
		return
	})
}

/* *****************私有方法***************** */
// showWindow 显示窗口
func (a *App) showWindow() {
	runtime.WindowShow(a.ctx)
	a.isHidden = false
}

// hideWindow 隐藏窗口
func (a *App) hideWindow() {
	runtime.WindowHide(a.ctx)
	a.isHidden = true
}

// toggleWindow 切换窗口显示状态
func (a *App) toggleWindow() {
	if a.isHidden {
		a.showWindow()
		return
	}
	a.hideWindow()
}

/* *****************窗口控制方法(暴露给前端)***************** */

// WindowMinimise 窗口最小化
func (a *App) WindowMinimise() {
	runtime.WindowMinimise(a.ctx)
}

// WindowMaximise 窗口最大化(切换状态)
func (a *App) WindowMaximise() {
	runtime.WindowToggleMaximise(a.ctx)
}

// WindowClose 窗口关闭(控制是否不退出仅隐藏)
func (a *App) WindowClose(isHidden bool) {
	//a.hide = isHidden
	a.hide = true
	runtime.Quit(a.ctx)
}

// WindowOnTop 窗口置顶(切换状态)
func (a *App) WindowOnTop() {
	a.isOnTop = !a.isOnTop // 切换置顶flag状态
	runtime.WindowSetAlwaysOnTop(a.ctx, a.isOnTop)
}

// WindowFullScreen 窗口全屏(切换状态)
func (a *App) WindowFullScreen() {
	if !a.WindowIsFullScreen() {
		runtime.WindowFullscreen(a.ctx)
	} else {
		runtime.WindowUnfullscreen(a.ctx)
	}
}

// WindowToggle 切换窗口显示/隐藏
func (a *App) WindowShowOrHide() {
	a.toggleWindow()
}

/* *****************窗口状态(暴露给前端)***************** */

// WindowIsMaximised 窗口是否最大化
func (a *App) WindowIsMaximised() bool {
	return runtime.WindowIsMaximised(a.ctx)
}

// WindowIsMinimised 窗口是否最小化
func (a *App) WindowIsMinimised() bool {
	return runtime.WindowIsMinimised(a.ctx)
}

// WindowIsOnToped 窗口是否置顶
func (a *App) WindowIsOnToped() bool {
	return a.isOnTop
}

// WindowIsHidden 窗口是否隐藏
func (a *App) WindowIsHidden() bool {
	return a.isHidden
}

// WindowIsFullScreen 窗口是否全屏
func (a *App) WindowIsFullScreen() bool {
	return runtime.WindowIsFullscreen(a.ctx)
}

/* *****************窗口位置(暴露给前端)***************** */

// WindowSetPosition 设置窗口位置
func (a *App) WindowSetPosition(x, y int) {
	runtime.WindowSetPosition(a.ctx, x, y)
}

// WindowGetPosition 获取窗口位置
func (a *App) WindowGetPosition() (x, y int) {
	return runtime.WindowGetPosition(a.ctx)
}

// WindowCenter 设置窗口居中
func (a *App) WindowCenter() {
	runtime.WindowCenter(a.ctx)
}

/* *****************窗口大小(暴露给前端)***************** */

// WindowSetSize 设置窗口大小
func (a *App) WindowSetSize(width, height int) {
	runtime.WindowSetSize(a.ctx, width, height)
}

// WindowGetSize 获取窗口大小
func (a *App) WindowGetSize() (width, height int) {
	return runtime.WindowGetSize(a.ctx)
}
