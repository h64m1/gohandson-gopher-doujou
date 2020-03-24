package main

import "fmt"

func main() {
	const helloWorld1 = "Hello world !"
	fmt.Println(helloWorld1)

	// 右辺の省略
	// 全部3
	const (
		a = 1 + 2
		b
		c
	)
	fmt.Println(a, b, c)

	// iotaで1ずつincrement
	const (
		d = iota
		e
	)
	// 式の中でも使える
	const (
		f = 1 << iota
		g
		h
	)
	fmt.Println(d, e, f, g, h)
}
