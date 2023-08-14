package utils

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/Esword618/p2t/config"
	"github.com/Esword618/p2t/global"
	"github.com/Esword618/p2t/model"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gofrs/uuid"
	"github.com/imroc/req/v3"
	"github.com/spf13/viper"
	"github.com/sunshineplan/imgconv"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.design/x/clipboard"
	"golang.org/x/net/html"
	"image/jpeg"
	"io"
	"log"
	"math/rand"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func SaveFormattedSVGToFile(filename, svgContent string) {
	// 解析SVG内容
	node, err := html.Parse(bytes.NewReader([]byte(svgContent)))
	if err != nil {
		fmt.Println("Error parsing SVG:", err)
		return
	}

	// 格式化解析后的节点
	var formattedSVG bytes.Buffer
	err = html.Render(&formattedSVG, node)
	if err != nil {
		fmt.Println("Error rendering formatted SVG:", err)
		return
	}

	// 创建文件
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// 将格式化后的SVG内容写入文件
	_, err = formattedSVG.WriteTo(file)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Formatted SVG saved to", filename)
}

func Base64ToBytes(base64String string) ([]byte, error) {
	imgBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}
	return imgBytes, nil
}

func BytesToBase64(bytes []byte) string {
	base64String := base64.StdEncoding.EncodeToString(bytes)
	return base64String
}

// GetTimestamp13 返回当前时间的 13 位 Unix 时间戳
func GetTimestamp13() string {
	return strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
}

// Uuid uuid 生成
func Uuid() string {
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
	}
	return u2.String()
}

// GetFormattedYearMonthAndDay 返回格式化的年月和日期。
func GetFormattedYearMonthAndDay() (formattedYearMonth, formattedDay string) {
	now := time.Now()

	year, month, day := now.Year(), int(now.Month()), now.Day()

	// 将月份和日期格式化为至少两位的字符串
	formattedYearMonth = fmt.Sprintf("%d-%02d", year, month)
	formattedDay = fmt.Sprintf("%02d", day)

	return formattedYearMonth, formattedDay
}

// EnsureDir 判断文件的目录路径是否存在，不存在就创建
func EnsureDir(filePath string) error {
	dirName := filepath.Dir(filePath) // 获取 image 的目录路径
	fmt.Println(dirName)
	// 如果目录不存在，创建它
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dirName, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// SaveImage 将[]byte格式的图片保存到指定路径。
func SaveImage(imageData []byte, filePath string) error {
	err := EnsureDir(filePath)
	if err != nil {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("unable to create file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(imageData)
	if err != nil {
		return fmt.Errorf("unable to write data to file: %v", err)
	}
	return nil
}

func GenerateRandomStr(length int) string {
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GetTimestamp10() string {
	return strconv.FormatInt(time.Now().UTC().Unix(), 10)
}

func CalculateMD5(input string) string {
	hash := md5.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

// PathExist 判断一个文件或文件夹是否存在
// 输入文件路径，根据返回的bool值来判断文件或文件夹是否存在
func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetCurrentExecDirectory() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	} else {
		fmt.Println("Current Directory: ", dir)
		return dir, nil
	}
}

// 剪贴板监听
func WatchClipboardImg() {
	ch := clipboard.Watch(context.TODO(), clipboard.FmtImage)
	for img := range ch {
		imgBs64 := base64.StdEncoding.EncodeToString(img)
		imgBytes, err := Base64ToBytes(imgBs64)
		mtype := mimetype.Detect(imgBytes)
		if err != nil {
			log.Println(err, "Base64ToBytes")
		}
		base64Img := fmt.Sprintf("data:%s;base64,%s", mtype.String(), imgBs64)
		//EventsEmit发消息
		//EventsOn接收消息
		runtime.EventsEmit(global.GlobalCtx, "watch_clipboard_img", base64Img)
		// 监听自动展示软件
		autoShow := viper.GetBool("autoshow")
		if autoShow {
			runtime.WindowShow(global.GlobalCtx)
		}
	}

}

// 图片格式转换
func Png2Jpeg(pngBs64 string) string {
	imgBytes, err := Base64ToBytes(strings.ReplaceAll(pngBs64, "data:image/png;base64,", ""))
	if err != nil {
		fmt.Println(err)
	}

	img, err := imgconv.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		fmt.Println(err)
	}

	// Write the resulting image as TIFF.
	err = imgconv.Write(io.Discard, img, &imgconv.FormatOption{Format: imgconv.JPEG})
	if err != nil {
		log.Fatalf("failed to write image: %v", err)
	}

	// Create a new in-memory buffer
	buf := new(bytes.Buffer)

	// Write the image to the buffer as JPEG
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get the []byte representation
	imgToByte := buf.Bytes()

	return BytesToBase64(imgToByte)
}

func Png2Bmp(pngBs64 string) string {
	imgBytes, err := Base64ToBytes(strings.ReplaceAll(pngBs64, "data:image/png;base64,", ""))
	if err != nil {
		fmt.Println(err)
	}

	img, err := imgconv.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		fmt.Println(err)
	}

	// Write the resulting image as TIFF.
	err = imgconv.Write(io.Discard, img, &imgconv.FormatOption{Format: imgconv.BMP})
	if err != nil {
		log.Fatalf("failed to write image: %v", err)
	}

	// Create a new in-memory buffer
	buf := new(bytes.Buffer)

	// Write the image to the buffer as JPEG
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get the []byte representation
	imgToByte := buf.Bytes()

	return BytesToBase64(imgToByte)
}

func Png2Tiff(pngBs64 string) string {
	imgBytes, err := Base64ToBytes(strings.ReplaceAll(pngBs64, "data:image/png;base64,", ""))
	if err != nil {
		fmt.Println(err)
	}

	img, err := imgconv.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		fmt.Println(err)
	}

	// Write the resulting image as TIFF.
	err = imgconv.Write(io.Discard, img, &imgconv.FormatOption{Format: imgconv.TIFF})
	if err != nil {
		log.Fatalf("failed to write image: %v", err)
	}

	// Create a new in-memory buffer
	buf := new(bytes.Buffer)

	// Write the image to the buffer as JPEG
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get the []byte representation
	imgToByte := buf.Bytes()

	return BytesToBase64(imgToByte)
}

// TODO 修复报错为pdf时图片为黑色问题

func Png2Pdf(pngBs64 string) {
	imgBytes, err := Base64ToBytes(strings.ReplaceAll(pngBs64, "data:image/png;base64,", ""))
	if err != nil {
		fmt.Println(err)
	}

	img, err := imgconv.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		fmt.Println(err)
	}

	// 获取当前用户信息
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("无法获取当前用户信息：", err)
		return
	}

	// 构建下载路径
	downloadPath := filepath.Join(currentUser.HomeDir, "Downloads")

	savePath, err := runtime.SaveFileDialog(global.GlobalCtx, runtime.SaveDialogOptions{
		DefaultDirectory: downloadPath,
		DefaultFilename:  GetTimestamp13() + ".pdf",
		Title:            "保存pdf文件",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "*.pdf",
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(savePath)

	imgconv.Save(savePath, img, &imgconv.FormatOption{Format: imgconv.PDF})
}

func GetLatestReleases() []any {
	var latestReleases model.LatestReleasesModel
	_, err := req.SetTLSFingerprintEdge().
		R().
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36").
		SetSuccessResult(&latestReleases).
		Get(config.GetReleasesUrl)
	if err != nil {
		fmt.Println(err)
	}
	if latestReleases.TagName != config.Version {
		return []any{true, latestReleases.HtmlUrl}
	} else {
		return []any{false}
	}
}
