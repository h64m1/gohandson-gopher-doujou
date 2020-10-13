package main

func main() {
	// 組み込みの関数
	printBuiltinFunction()

	// 多値の受け取り
	printMultipleReturnValues()
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
