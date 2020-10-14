package main

import "fmt"

func main() {
	// 組み込みの関数
	printBuiltinFunction()

	// 多値の受け取り
	printMultipleReturnValues()

	// 無名関数
	printClosure()

	// 値のコピー、ポインタ
	printCopyValue()
}

func printBuiltinFunction() {
	println("print/println: 表示を行う")
	println("make: コンポジット型の初期化")
	println("new: 指定した型のメモリを確保")
	println("len/cap: スライスなどの長さ/容量を返す")
	println("copy: スライスのコピーを行う")
	println("delete: マップから指定したキーのエントリを削除")
	println("complex: 複素数型を作成")
	println("imag/real: 複素数の虚部/実数部を取得")
	println("panic/recover: パニックを起こす/回復する")
}

func printMultipleReturnValues() {
	x, y := swap(10, 20)
	println("swap (x,y)=(10,20):", x, y)
	// 省略したい場合は _ (ブランク変数)
	x, _ = swap(10, 20)
	_, y = swap(10, 20)
}

func swap(x, y int) (int, int) {
	return y, x
}

// 名前付き戻り値
func swapNamedReturnValues(x, y int) (x2, y2 int) {
	// 一時変数なしで値の入れ替え
	y2, x2 = x, y
	return
}

func printClosure() {
	msg := "hello world"
	func() {
		// 関数の外のmsgを参照できる
		println(msg)
	}()

	// 関数型
	fs1 := make([]func() string, 2)
	fs1[0] = func() string { return "hello" }
	fs1[1] = func() string { return "world" }

	for k, f := range fs1 {
		fmt.Println(k, f())
	}

	// クロージャのよくあるバグ
	// - 関数外の変数を参照している場合
	// - 実行のタイミングでは値が変わっている可能性がある
	fs2 := make([]func(), 3)
	for i := range fs2 {
		fs2[i] = func() { fmt.Println(i) }
	}
	for _, f := range fs2 {
		f()
	}
}

func printCopyValue() {
	p1 := struct {
		age  int
		name string
	}{age: 10, name: "Gopher"}

	// copy
	p2 := p1
	p2.age = 20
	println(p1.age, p1.name)
	println(p2.age, p2.name)

	// ポインタ
	var x int
	set100(&x)
	println(x)

	// 内部でポインタが用いられているデータ型
	// - スライス
	// - マップ
	// - チャネル
	ns1 := []int{10, 20, 30}
	ns2 := ns1
	ns1[1] = 200
	println(ns1[0], ns1[1], ns1[2])
	println(ns2[0], ns2[1], ns2[2])
}

func set100(xp *int) {
	// ポインタに値を代入
	*xp = 100
}
