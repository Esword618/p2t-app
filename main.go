package main

import (
	"bytes"
	"embed"
	myconfig "github.com/Esword618/p2t/config"
	"github.com/Esword618/p2t/global"
	"github.com/Esword618/p2t/model"
	"github.com/Esword618/p2t/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"golang.design/x/clipboard"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

// GetCurrPath 获取程序运行路径
func GetCurrPath() string {
	f, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(f)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {

	// Create an instance of the app structure
	app := NewApp()
	utilsBind := utils.NewUtils()

	// Create application with options
	err := wails.Run(&options.App{
		Title:       "p2t",
		Width:       1024,
		Height:      768,
		Frameless:   true,
		StartHidden: false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		OnStartup:        app.startup,
		OnBeforeClose:    app.beforeClose,
		Bind: []interface{}{
			app,
			utilsBind,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			ZoomFactor:                        1.0,
		},
		// Mac platform specific options
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "p2t",
				Message: "",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	// 剪贴板初始化
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	//	数据库初始化、
	initDatabase()
	// 配置初始化
	initConfig()
	// 启动图片服务
	go initServeImageFile()
}

func initConfig() {
	basePath, _ := utils.GetCurrentExecDirectory()
	// 配置文件
	configPath := filepath.Join(basePath, myconfig.ConfigName)
	viper.SetConfigFile(configPath)
	viper.WatchConfig()
	b, _ := utils.PathExist(configPath)
	if !b {
		//// 获取当前用户信息
		//currentUser, err := user.Current()
		//if err != nil {
		//	fmt.Println("无法获取当前用户信息：", err)
		//	return
		//}
		//
		//// 构建下载路径
		//savePath := filepath.Join(currentUser.HomeDir, "Documents")
		savePath := filepath.Join(basePath, "FileStorage", "Image")
		viper.Set("savedir", savePath)
		viper.ReadConfig(bytes.NewBuffer([]byte(myconfig.DefaultYaml)))
		viper.WriteConfig()
	}
	viper.ReadInConfig()
}

func initDatabase() {
	basePath, _ := utils.GetCurrentExecDirectory()
	DbPath := filepath.Join(basePath, myconfig.DbName)
	db, err := gorm.Open(sqlite.Open(DbPath))
	log.Println("DbPath：", DbPath)
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.LatexModel{})
	global.GlobalDb = db
}

func initServeImageFile() {
	// http://localhost:8080?image=2023-08\12\78dc5761ab8c4b55b940967ae24ebe3a.jpg
	r := gin.Default()
	// 设置CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	r.GET("/", func(c *gin.Context) {
		// 获取图片路径
		relaticePath := c.Query("image")
		imagePath := path.Join(viper.GetString("savedir"), relaticePath)
		if imagePath == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Image path is missing"})
			return
		}

		// 读取并返回图片
		c.File(imagePath)
	})
	r.Run(":8080")
}
