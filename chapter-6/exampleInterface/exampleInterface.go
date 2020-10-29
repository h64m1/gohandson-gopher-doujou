package exampleInterface

import "fmt"

// ShowExampleInterface show interface example
func ShowExampleInterface() {
	// インターフェースの例
	// インターフェースはメソッドの集まり
	// - メソッドのリストがインターフェースを規定しているものと一致する型は、インターフェースを実装していることになる
	exampleInterface()

	// 型アサーション
	// - インターフェース.(型)
	//   * インターフェース型の値を任意の型にキャスト
	//   * 第2戻り値にキャストできるかどうかが返る
	typeAssertion()

	// 型スイッチ
	typeSwitch()

	// Stringerインターフェース
	// - String() stringメソッドを持つインターフェースを作る
	// - そして3つ以上Stringerインターフェースを実装する型を作る
	// インターフェースを受け取る関数
	// - Stringerインターフェースを引数で受け取る関数を作る
	// - 受け取った値を上記の3つの具象型によって分岐し、具象型の型名と値を表示する
	stringerInterface()
}

func exampleInterface() {
	fmt.Println("# インターフェースの例")

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
}

func typeAssertion() {
	var v interface{}
	v = 100
	n, ok := v.(int)
	fmt.Println(n, ok)

	s, ok := v.(string)
	fmt.Println(s, ok)
}

func typeSwitch() {
	var v interface{}
	v = 100
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

func stringerInterface() {
	fmt.Println("Stringerインターフェース")
	var i = Int(1000)
	fmt.Println(i.String())
	showString(i)

	var f = Float(123.456)
	fmt.Println(f.String())
	showString(f)

	var b = Bool(true)
	fmt.Println(b.String())
	showString(b)
}

// Stringer interface
type Stringer interface {
	String() string
}

// Int int
type Int int

// Float float32
type Float float32

// Bool bool
type Bool bool

func (i Int) String() string {
	return fmt.Sprintf("%d", i)
}

func (f Float) String() string {
	return fmt.Sprintf("%f", f)
}

func (b Bool) String() string {
	return fmt.Sprintf("%t", b)
}

func showString(s Stringer) {
	// 型スイッチ
	switch v := s.(type) {
	case Int:
		fmt.Println(v, "Int")
	case Float:
		fmt.Println(v, "Float")
	case Bool:
		fmt.Println(v, "Bool")
	default:
		fmt.Println("default", v, s)
	}
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
