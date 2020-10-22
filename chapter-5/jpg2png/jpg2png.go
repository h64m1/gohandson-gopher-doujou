package jpg2png

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Convert convert jpeg to png
func Convert() {
	// 次の仕様を満たすコマンド
	// - ディレクトリを指定する
	// - 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
	// - ディレクトリ以下は再帰的に処理する
	// - 変換前の画像形式を指定できる（オプション）
	// 以下を満たすように開発
	// - mainパッケージと分離する
	// - 自作パッケージと標準パッケージと準標準パッケージのみを使う
	// - ユーザ定義型を作ってみる
	// - GoDocを生成してみる

	err := filepath.Walk("figure",
		func(path string, info os.FileInfo, err error) error {
			convert(path)
			return nil
		})
	if err != nil {
		fmt.Println(err)
	}
}

func convert(path string) {
	// jpg以外は何もしない
	if filepath.Ext(path) != ".jpg" {
		return
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	pngFileName := strings.Replace(path, "jpg", "png", 1)
	pngFile, err := os.Create(pngFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer pngFile.Close()

	// jpg -> png
	png.Encode(pngFile, img)

	fmt.Println(img.Bounds().String(), pngFileName)
}
