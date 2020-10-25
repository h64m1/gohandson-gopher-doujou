// Package jpg2png jpg or gif画像をpngに変換
package jpg2png

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"  // gifのDecode実行用
	_ "image/jpeg" // jpegのDecode実行用
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var from = flag.String("from", "jpg", "変換前の画像形式（デフォルトはjpg: jpg or gif）")

type imageFile struct {
	Path      string // ファイル名
	Extension string // 拡張子
}

// Convert jpg or gif画像をpngに変換
//
// jpg2pngを実行したディレクトリ配下にある全ての画像ファイルを対象（デフォルトではjpgが対象）
//
// オプション
//	from [option] optionで画像の拡張子を指定（e.g. ./jpg2png -from gif）
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
	flag.Parse()

	// 画像の拡張子
	ext := "." + *from
	if !isExtensionOk(ext) {
		fmt.Println("画像の拡張子はjpg or gifを指定してください")
		os.Exit(1)
	}

	err := filepath.Walk("figure",
		func(path string, info os.FileInfo, err error) error {
			image := imageFile{Path: path, Extension: ext}
			convert(image)
			return nil
		})
	if err != nil {
		fmt.Println(err)
	}
}

// isExtensionOk 拡張子がjpg or gifの場合true
func isExtensionOk(ext string) bool {
	return ext == ".jpg" || ext == ".gif"
}

func convert(imgFile imageFile) {
	// 指摘拡張子以外は何もしない (defaultはjpg)
	if filepath.Ext(imgFile.Path) != imgFile.Extension {
		return
	}

	file, err := os.Open(imgFile.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	pngExt := ".png"
	pngFileName := strings.Replace(imgFile.Path, filepath.Ext(imgFile.Path), pngExt, 1)
	pngFile, err := os.Create(pngFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer pngFile.Close()

	// jpg -> png
	png.Encode(pngFile, img)

	fmt.Println("convert from", filepath.Base(imgFile.Path), "to", pngFileName)
}
