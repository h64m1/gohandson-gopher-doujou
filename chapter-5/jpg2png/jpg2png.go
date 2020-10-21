package jpg2png

import (
	"fmt"
	"os"
	"path/filepath"
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
			if filepath.Ext(path) == ".jpg" {
				fmt.Println(path)
			}
			return nil
		})
	if err != nil {
		fmt.Println(err)
	}
}
