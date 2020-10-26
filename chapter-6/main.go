package main

import "fmt"

func main() {
	// インターフェースの例
	// インターフェースはメソッドの集まり
	// - メソッドのリストがインターフェースを規定しているものと一致する型は、インターフェースを実装していることになる

	var s Stringer = Hex(100)
	fmt.Println(s.String())

	// empty interface
	var v interface{}
	v = 100
	v = "hello world"
	fmt.Println(v)
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
