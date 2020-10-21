package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/h64m1/gohandson-gopher-doujou/chapter-5/jpg2png"
)

// var msg = flag.String("msg", "デフォルト値", "説明")
// var n int

func init() {
	// ポインタを指定して設定を予約
	// flag.IntVar(&n, "n", 1, "回数")
}

func printStd() {
	// os.Args: プログラム引数が入った文字列型のスライス
	// 要素のひとつめはプログラム名
	fmt.Println(os.Args)

	// フラグパッケージ
	flag.Parse()
	// fmt.Println(strings.Repeat(*msg, n))

	// flag.Args関数はフラグの分は除外される
	fmt.Println(flag.Args())

	// 入出力
	fmt.Fprintf(os.Stdout, "標準出力\n")
	fmt.Fprintf(os.Stderr, "標準エラー出力\n")
	// プログラムの終了
	// os.Exit(1)

	// log.Fatal: 標準エラー出力にエラーメッセージを表示
	// os.Exit(1)で以上終了させる
	// sf, err := os.Open("file")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // 関数終了時に閉じる
	// defer sf.Close()

	// defer: 関数の遅延実行
	// - 関数終了時に呼び出し
	// - 引数の評価はdefer呼び出し時
	// - スタック形式で実行、最後に呼び出されたものが最初に実行
	msg := "!!!"
	defer fmt.Println(msg)
	msg = "world"
	defer fmt.Println(msg)
	fmt.Println("hello")

	// forの中でdeferは避ける
	// - 予約した関数呼び出しはreturn時に実行される
	// - forの中を関数に分ければ良い
}

func main() {
	// mycat.Cat()
	jpg2png.Convert()
}
