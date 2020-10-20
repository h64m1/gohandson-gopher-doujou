package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args: プログラム引数が入った文字列型のスライス
	// 要素のひとつめはプログラム名
	fmt.Println(os.Args)
}
