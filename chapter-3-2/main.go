package main

func main() {
	// 組み込みの関数
	printBuiltinFunction()
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
