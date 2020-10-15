package main

import "fmt"

func main() {
	// メソッド
	method()

	// レシーバ
	receiver()

	increment()

	// メソッド値とメソッド式
	var hex Hex = 100
	f1 := hex.String
	fmt.Println(f1())

	f2 := Hex.String
	fmt.Printf("%T\n%s\n", f2, f2(hex))
}

// メソッド: レシーバと紐付けられた関数
// - データとそれに対する操作を紐付けるために用いる
// - ドットでメソッドでアクセスする

// Hex type
type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func method() {
	var hex Hex = 100
	fmt.Println(hex.String())
}

// メソッドに関連付けられた変数
// - メソッド呼び出し時には通常の引数と同じような扱いになる
//   * コピーが発生
// - ポインタを使うことでレシーバへの変更を呼び出し元に伝えることができる
//   * レシーバがポインタの場合もドットでアクセスする

// T type
type T int

func (t *T) f() { println("hi") }

func receiver() {
	var v T
	(&v).f()
	v.f()
}

// MyInt インクリメント用のタイプ
type MyInt int

func (n *MyInt) Inc() {
	(*n)++
}

func increment() {
	var n MyInt
	println(n)
	n.Inc()
	println(n)
}
