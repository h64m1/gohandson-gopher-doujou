package main

import "fmt"

// 構造体に匿名フィールドを埋め込む
type Hoge struct {
	N int
}

// Fuga型にHoge型を埋め込む
type Fuga struct {
	Hoge // 名前のないフィールドになる
}

// 型リテラルでなければ埋め込められる
// - typeで定義したものや組み込み型
// - インターフェースも埋め込められる
// Stringerを実装
type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

// Hex2もStringerを実装
type Hex2 struct{ Hex }

func main() {
	f := Fuga{Hoge{N: 100}}

	// Hoge型のフィールドにアクセス出来る
	fmt.Println(f.N)

	// 型名を指定してアクセス出来る
	fmt.Println(f.Hoge.N)

	h := Hex2{10}
	fmt.Println(h.String())
}
