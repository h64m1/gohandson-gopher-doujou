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
	print("配列| ", title, ": ")
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d ", array[i])
	}
	print("\n")
}

func showSlice() {
	// ゼロ値はnil
	var ns1 []int
	printSlice("ns1", ns1)

	// 長さと容量を指定して初期化
	// 各要素はゼロ値で初期化
	ns1 = make([]int, 3, 10)
	printSlice("ns1", ns1)

	// スライスリテラルで初期化
	// 要素数は指定しなくてよい
	// 自動で配列は作られる
	var ns2 = []int{10, 20, 30, 40, 50}
	printSlice("ns2", ns2)

	// 5番目が50、10番目が100で他が0の要素数11のスライス
	ns3 := []int{5: 50, 10: 100}
	printSlice("ns3", ns3)

	// スライスの操作
	ns4 := []int{10, 20, 30, 40, 50}

	println("要素にアクセス: ", ns4[3])
	println("長さ: ", len(ns4))

	// 要素の追加
	ns4 = append(ns4, 60, 70)
	// 容量
	println("長さ: ", len(ns4), ", 容量: ", cap(ns4))

	// appendの挙動
	ns5 := []int{10, 20}
	fmt.Println("ns5: ", ns5, cap(ns5))

	ns6 := append(ns5, 30)
	ns5[0] = 100
	fmt.Println("ns5: ", ns5, cap(ns5))
	fmt.Println("ns6: ", ns6, cap(ns6))

	ns7 := append(ns6, 40)
	ns6[1] = 200
	fmt.Println("ns5: ", ns5, cap(ns5))
	fmt.Println("ns6: ", ns6, cap(ns6))
	fmt.Println("ns7: ", ns7, cap(ns7))

	// 配列・スライスへのスライス演算
	ns8 := []int{10, 20, 30, 40, 50}
	n, m := 2, 4
	// n番目以降のスライスを取得
	fmt.Println(ns8[n:])
	// 先頭からm-1番目までのスライスを取得する
	fmt.Println(ns8[:m])

	ms := ns8[:m:m]
	fmt.Println(ms, cap(ms))

	// slice tricks
	// Cut
	printSlice("ns2", ns2)
	ns2 = append(ns2[:1], ns2[3:]...)
	printSlice("cut ns2", ns2)

	// Delete
	i := 1
	printSlice("ns2", ns2)
	ns2 = append(ns2[:i], ns2[i+1:]...)
	printSlice("delete ns2", ns2)
	// or
	printSlice("ns2", ns2)
	ns2 = ns2[:i+copy(ns2[i:], ns2[i+1:])]
	printSlice("delete ns2", ns2)
}

func printSlice(title string, slice []int) {
	print("スライス | ", title, ": ")
	if slice == nil {
		println("nil")
		return
	}

	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d ", slice[i])
	}
	print("\n")
}

func compositType() {
	// 配列
	println("配列")
	showArray()

	// スライス
	println("スライス")
	showSlice()
}

func useSlice() {
	// 3つの変数しか使わないように修正してください
	// - プログラムの動作はそのままにすること
	n1 := 19
	n2 := 86
	n3 := 1
	n4 := 12

	sum1 := n1 + n2 + n3 + n4
	println("sum1", sum1)

	// 修正版
	n := []int{19, 86, 1, 12}
	sum2 := 0
	for i := 0; i < len(n); i++ {
		sum2 += n[i]
	}
	println("sum2", sum2)
}

func mapType() {
	// マップの初期化
	mapInit()

	// マップの操作
	mapManipulation()
}

func mapInit() {
	// マップの初期化
	// ゼロ値はnil
	var m map[string]int
	printMap("ゼロ値はnil", m)

	// makeで初期化
	m = make(map[string]int)
	printMap("makeで初期化", m)

	// 容量の指定
	m = make(map[string]int, 10)
	printMap("容量の指定", m)

	// リテラルで初期化
	m1 := map[string]int{"x": 10, "y": 20}
	printMap("リテラルで初期化", m1)

	// 空で初期化
	m2 := map[string]int{}
	printMap("空で初期化", m2)
}

func mapManipulation() {
	println("マップの操作")

	// マップの操作
	m := map[string]int{"x": 10, "y": 20}

	// キーを指定してアクセス
	println(m["x"])

	// キーを指定して入力
	m["z"] = 30

	// 存在を確認する
	n, ok := m["z"]
	println(n, ok)

	n, ok = m["a"]
	println(n, ok)

	// キーを指定して削除
	delete(m, "z")

	// 削除されていることを確認
	n, ok = m["z"]
	println(n, ok)
}

func printMap(title string, m map[string]int) {
	println(title, " | len=", len(m))
	for k, v := range m {
		println("key", k, ", value", v)
	}
}

func main() {
	// 組み込み型
	println("組み込み型")
	existingType()

	// コンポジット型
	println("コンポジット型")
	compositType()

	// スライスの利用
	println("スライスの利用")
	useSlice()

	// マップ
	println("マップ")
	mapType()
}
