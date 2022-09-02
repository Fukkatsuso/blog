package main

import (
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fogleman/gg"
)

const (
	fontPath = "/usr/share/fonts/Kinto Sans/KintoSans-Bold.ttf"
)

type Config struct {
	TemplatePath string   `json:"template"`
	OutputPath   string   `json:"path"`
	Title        []string `json:"title"`
}

func CreateImage(config Config) (image.Image, error) {
	// ベース
	width := 1200 * 3 / 4 // 1200px
	height := 630 * 3 / 4 // 630px
	dc := gg.NewContext(width, height)

	// テンプレート画像をベースにのせる
	backgroundImage, err := gg.LoadImage(config.TemplatePath)
	if err != nil {
		return nil, err
	}
	dc.DrawImage(backgroundImage, 0, 0)

	// フォントサイズ
	if err := dc.LoadFontFace(fontPath, 50); err != nil {
		return nil, err
	}
	// フォントカラー
	dc.SetHexColor("232F3E")

	// 記事タイトルをベースにのせる
	title := strings.Join(config.Title, "\n")
	maxWidth := float64(dc.Width()) * 0.85
	dc.DrawStringWrapped(title, float64(width)/2, float64(height)/2, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)

	return dc.Image(), nil
}

func readConfig(path string) ([]Config, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var configs []Config
	if err := json.Unmarshal(raw, &configs); err != nil {
		return nil, err
	}

	return configs, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please pass the config file path as an argument")
		os.Exit(1)
	}

	configFilePath := os.Args[1]
	configs, err := readConfig(configFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, config := range configs {
		img, err := CreateImage(config)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if err := gg.SavePNG(config.OutputPath, img); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
