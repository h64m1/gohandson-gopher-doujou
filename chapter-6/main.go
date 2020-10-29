package main

import (
	"fmt"

	"github.com/h64m1/gohandson-gopher-doujou/chapter-6/exampleInterface"
)

func main() {
	// インターフェースの例
	exampleInterface.ShowExampleInterface()

	// 構造体の埋め込み
	fmt.Println("# 構造体の埋め込み")
	f := Title{Name{name: "hello"}}
	// Name型のフィールドにアクセル
	fmt.Println(f.name)
	// 型名を指定してアクセス
	fmt.Println(f.Name.name)
}

// Name Name構造体
type Name struct {
	name string
}

// Title Title構造体
type Title struct {
	Name // 名前のないフィールド
}
