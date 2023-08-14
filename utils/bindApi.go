package utils

import (
	"fmt"
	"github.com/Esword618/p2t/global"
	"github.com/Esword618/p2t/model"
	"github.com/gabriel-vasile/mimetype"
	"github.com/imroc/req/v3"
	"github.com/spf13/viper"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	"net/http"
	"path"
	"regexp"
)

func (u *Utils) Pix2TextApi(bs64Img string) string {
	re := regexp.MustCompile(`data:image/.+;base64,`)
	bs64Img = re.ReplaceAllString(bs64Img, "")
	imgBytes, err := Base64ToBytes(bs64Img)
	if err != nil {
		log.Println(err, "Base64ToBytes")
	}
	var pix2textModel model.Pix2TextModel
	mtype := mimetype.Detect(imgBytes)
	timestamp := GetTimestamp13()
	fileType := mtype.Extension()
	fileName := timestamp + fileType
	res, err := req.SetTLSFingerprintEdge().
		R().
		SetHeaders(map[string]string{
			"origin":             "https://p2t.behye.com",
			"referer":            "https://p2t.behye.com/",
			"sec-ch-ua-mobile":   "?0",
			"sec-fetch-dest":     "empty",
			"sec-ch-ua-platform": "Windows",
			"sec-fetch-mode":     "cors",
			"sec-fetch-site":     "same-origin",
			"user-agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
		}).
		SetFileBytes("image", fileName, imgBytes). // Set form param name and filename
		SetFormData(map[string]string{             // Set form data while uploading
			"session_id": "",
		}).
		SetSuccessResult(&pix2textModel).
		Post("https://p2t.breezedeus.com/api/pix2text")
	if err != nil {
		return err.Error()
	} else {
		if res.StatusCode == http.StatusOK {
			formattedYearMonth, formattedDay := GetFormattedYearMonthAndDay()
			relaticePath := path.Join(formattedYearMonth, formattedDay, fileName)
			// 保存到sqlite数据库
			global.GlobalDb.Create(&model.LatexModel{
				Uuid:    Uuid(),
				ApiType: model.Pix2text,
				Latex:   pix2textModel.Results,
				Image:   relaticePath,
			})
			savePath := path.Join(viper.GetString("savedir"), relaticePath)
			//保存到文件夹
			SaveImage(imgBytes, savePath)
			return pix2textModel.Results
		} else {
			return res.Err.Error()
		}
	}
}

func (u *Utils) SimpletexApi(bs64Img string) string {
	AppId := viper.GetString("appid")
	AppSecret := viper.GetString("appsecret")

	if len(AppId) == 0 || len(AppSecret) == 0 {
		runtime.EventsEmit(global.GlobalCtx, "watch_notification", map[string]any{
			"title":   "Sipmletex 模式识别！",
			"message": "AppId与AppSecret不可为空",
			"type":    "error",
		})
		return "AppId与AppSecret不可为空"
	}
	re := regexp.MustCompile(`data:image/.+;base64,`)
	bs64Img = re.ReplaceAllString(bs64Img, "")
	imgBytes, err := Base64ToBytes(bs64Img)
	if err != nil {
		log.Println(err, "Base64ToBytes")
	}
	var simpletexModel model.SimpletexModel
	mtype := mimetype.Detect(imgBytes)
	timestamp13 := GetTimestamp13()
	fileType := mtype.Extension()
	fileName := timestamp13 + fileType

	randomStr := GenerateRandomStr(16)
	timestamp10 := GetTimestamp10()
	str := fmt.Sprintf("app-id=%s&random-str=%s&timestamp=%s&secret=%s", AppId, randomStr, timestamp10, AppSecret)
	sign := CalculateMD5(str)
	_, err = req.R().SetHeaders(map[string]string{
		"app-id":     AppId,
		"random-str": randomStr,
		"timestamp":  timestamp10,
		"sign":       sign,
	}).SetFileBytes("file", "file.jpg", imgBytes).
		SetSuccessResult(&simpletexModel).
		Post("https://server.simpletex.cn/api/latex_ocr")

	if err != nil {
		return err.Error()
	} else {
		if simpletexModel.Status {
			formattedYearMonth, formattedDay := GetFormattedYearMonthAndDay()
			relaticePath := path.Join(formattedYearMonth, formattedDay, fileName)
			// 保存到sqlite数据库
			latexStr := "$$\n" + simpletexModel.Res.Latex + "\n$$"
			global.GlobalDb.Create(&model.LatexModel{
				Uuid:    Uuid(),
				ApiType: model.Simpletex,
				Latex:   latexStr,
				Image:   relaticePath,
			})
			savePath := path.Join(viper.GetString("savedir"), relaticePath)
			//保存到文件夹
			SaveImage(imgBytes, savePath)
			return latexStr
		} else {
			runtime.EventsEmit(global.GlobalCtx, "watch_notification", map[string]any{
				"title":   "Sipmletex 模式识别！",
				"message": "AppId与AppSecret填写有误，请检查！",
				"type":    "error",
			})
			return "请检查你的AppId 与 AppSecret"
		}
	}
}

func (u *Utils) ParseImg(bs64Img string) string {
	formulaType := viper.GetString("formulatype")
	switch formulaType {
	case "pix2text":
		fmt.Println("pix2text")
		return u.Pix2TextApi(bs64Img)
	case "simpletex":
		fmt.Println("simpletex")
		return u.SimpletexApi(bs64Img)
	default:
		fmt.Println("default")
		return u.SimpletexApi(bs64Img)
	}
}
