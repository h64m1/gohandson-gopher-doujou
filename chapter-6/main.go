package main

import "fmt"

func main() {
	// インターフェースの例
	// インターフェースはメソッドの集まり
	// - メソッドのリストがインターフェースを規定しているものと一致する型は、インターフェースを実装していることになる

	var s1 Stringer = Hex(100)
	fmt.Println(s1.String())

	// empty interface
	var v interface{}
	v = 100
	v = "hello world"
	fmt.Println(v)

	// 関数にメソッドを持たせる
	var s2 fmt.Stringer = Func(func() string {
		return "hi"
	})
	fmt.Println(s2)

	// 型アサーション
	// - インターフェース.(型)
	//   * インターフェース型の値を任意の型にキャスト
	//   * 第2戻り値にキャストできるかどうかが返る
	v = 100
	n, ok := v.(int)
	fmt.Println(n, ok)

	s, ok := v.(string)
	fmt.Println(s, ok)

	// 型スイッチ
	switch i := v.(type) {
	case int:
		fmt.Println(i * 2)
	case string:
		fmt.Println(i + "world")
	default:
		fmt.Println("default")
	}
}

// Stringer interface
type Stringer interface {
	String() string
}

// Hex int type
type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

// Func 関数型
type Func func() string

func (f Func) String() string { return f() }

// インターフェースの実装チェック
var _ fmt.Stringer = Func(nil)
