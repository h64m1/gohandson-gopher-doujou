package main

import "fmt"

// 構造体に匿名フィールドを埋め込む
type Hoge interface {
	M()
	N()
}

// fuga型にHoge型を埋め込む
type fuga struct {
	Hoge // 名前のないフィールドになる
}

func (f fuga) M() {
	// Mの振る舞いを変える
	fmt.Println("Hi")
	f.Hoge.M() // 元のメソッドを呼ぶ
}

func HiHoge(h Hoge) Hoge {
	return fuga{h} // 構造体を作る
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

// インターフェースの合成
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}

// ReaderとWriterを埋め込む
type ReadWriter interface {
	Reader
	Writer
}

func main() {
	h := Hex2{10}
	fmt.Println(h.String())
}
