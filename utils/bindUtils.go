package utils

import (
	"fmt"
	"github.com/88250/lute"
	"github.com/Esword618/p2t/config"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/viper"
	"golang.design/x/clipboard"
	"log"
	"strings"
)

type Utils struct {
	Lute *lute.Lute
}

func NewUtils() *Utils {
	return &Utils{
		Lute: lute.New(),
	}
}

func (u *Utils) Md2Html(mdStr string) string {
	u.Lute.SetAutoSpace(true)
	u.Lute.SetInlineMathAllowDigitAfterOpenMarker(true)
	u.Lute.SetVditorMathBlockPreview(true)
	html := u.Lute.Md2HTML(mdStr)
	return html
}

func (u *Utils) GetTex(latexStr string) string {
	htmlContent := u.Md2Html(latexStr)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}

	dataMath := doc.Find(".language-math").Text()
	return dataMath
}

func (u *Utils) ClipboardSetImage(base64String string) {

	imgBytes, err := Base64ToBytes(strings.ReplaceAll(base64String, "data:image/png;base64,", ""))
	if err != nil {
		fmt.Println(err)
	}
	clipboard.Write(clipboard.FmtImage, imgBytes)
}

func (u *Utils) ClipboardGetImage() string {
	return "data:image/png;base64," + BytesToBase64(clipboard.Read(clipboard.FmtImage))
}

func (u *Utils) GetVersion() string {
	return config.Version
}

//base64HeaderArray: [
//    {value: 'data:image/*;base64,',    title: '*'},
//    {value: 'data:image/jpeg;base64,', title: 'jpg/jepg'},
//    {value: 'data:image/png;base64,',  title: 'png'},
//    {value: 'data:image/svg;base64,',  title: 'svg'},
//]

func (u *Utils) Png2Jpeg(pngBs64 string) string {
	return Png2Jpeg(pngBs64)
}

func (u *Utils) Png2Bmp(pngBs64 string) string {
	return Png2Bmp(pngBs64)
}

func (u *Utils) Png2Tiff(pngBs64 string) string {
	return Png2Tiff(pngBs64)
}

// TODO 修复报错为pdf时图片为黑色问题

func (u *Utils) Png2Pdf(pngBs64 string) {
	Png2Pdf(pngBs64)
}

// 配置
func (u *Utils) GetConfig() map[string]any {
	return viper.AllSettings()
}

func (u *Utils) UpdateConfig(updateCfg map[string]any) {
	for key, value := range updateCfg {
		viper.Set(key, value)
	}
	viper.WriteConfig()
}

func (u *Utils) GetLatestReleases() []any {
	return GetLatestReleases()
}
