package main

import (
	"fmt"

	"github.com/tenntenn/greeting"
)

func main() {
	// パッケージ
	usePackage()

	// スコープ
	scope()
}

func usePackage() {
	fmt.Println(greeting.Do())
}

func scope() {
	// ブロックスコープ
	f()
	g()

	// パッケージ変数
	h()
	msg = "hi, gophers"
	h()
}

func f() {
	n := 100
	println(n)
}

func g() {
	// f()のnとは別物
	n := 200
	println(n)
}

// :=は関数内でのみ使用可能
var msg string = "hello"

// var msg := "hello" // これはエラー

func h() {
	println(msg)
}
