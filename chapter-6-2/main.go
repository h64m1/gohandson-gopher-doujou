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

func main() {
	f := Fuga{}

	// Hoge型のフィールドにアクセス出来る
	fmt.Println(f.N)
}
