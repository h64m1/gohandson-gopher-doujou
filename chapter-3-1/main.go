package main

import "fmt"

func existingType() {
	// 組み込み型（数値）
	var sum int
	sum = 5 + 6 + 3
	avg := float32(sum) / 3
	if avg > 4.5 {
		println("good", avg)
	}

	// 組み込み型（真偽値）
	var a, b, c bool
	for i := 0; i < 2; i++ {
		a = true
		if i == 1 {
			a = false
		}
		for j := 0; j < 2; j++ {
			b = true
			if j == 1 {
				b = false
			}
			for k := 0; k < 2; k++ {
				c = true
				if k == 1 {
					c = false
				}
				if a && b || !c {
					println(a, b, c, ": true")
				} else {
					println(a, b, c, ": false")
				}
			}
		}
	}
}

func compositType() {
	// 配列
	println("配列")
	showArray()
}

func showArray() {
	// 配列
	// ゼロ値で初期化
	var ns1 [5]int
	printArray("ns1", ns1[0:5])

	// 配列リテラルで初期化
	var ns2 = [5]int{10, 20, 30, 40, 50}
	printArray("ns2", ns2[0:5])

	// 要素数の値から型推論
	ns3 := [...]int{60, 70, 80, 90, 100}
	printArray("ns3", ns3[0:5])

	// 6番目が50, 11番目が100, 他が0
	ns4 := [...]int{5: 50, 10: 100}
	printArray("ns4", ns4[0:11])

	println(ns1[3], ns2[4], ns3[2], ns4[5])
}

func printArray(title string, array []int) {
	print(title, ": ")
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d ", array[i])
	}
	print("\n")
}

func main() {
	// 組み込み型
	println("組み込み型")
	existingType()

	// コンポジット型
	println("コンポジット型")
	compositType()
}
