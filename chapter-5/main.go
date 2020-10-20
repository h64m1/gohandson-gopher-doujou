package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var msg = flag.String("msg", "デフォルト値", "説明")
var n int

func init() {
	// ポインタを指定して設定を予約
	flag.IntVar(&n, "n", 1, "回数")
}

func main() {
	// os.Args: プログラム引数が入った文字列型のスライス
	// 要素のひとつめはプログラム名
	fmt.Println(os.Args)

	// フラグパッケージ
	flag.Parse()
	fmt.Println(strings.Repeat(*msg, n))

	// flag.Args関数はフラグの分は除外される
	fmt.Println(flag.Args())
}
