package main

import "fmt"

func main() {

	// 変数定義と代入が一緒
	var helloWorld1 string = "Hello world !"
	fmt.Println("変数定義と代入が一緒", helloWorld1)

	// 変数定義と代入が別
	var helloWorld2 string
	helloWorld2 = "Hello world !"
	fmt.Println("変数定義と代入が別", helloWorld2)

	// 型を省略
	var helloWorld3 = "Hello world !"
	fmt.Println("型を省略", helloWorld3)

	// varを省略
	helloWorld4 := "Hello world !"
	fmt.Println("varを省略", helloWorld4)
}
