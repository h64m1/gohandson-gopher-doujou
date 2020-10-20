package mycat

import (
	"flag"
	"fmt"
)

var n = flag.Bool("n", false, "行番号を表示")

// Cat my cat command
func Cat() {
	// catコマンドの仕様
	// - 引数でファイルパスの一覧をもらい、そのファイルに与えられた順に標準出力にそのまま出力されるコマンドを作ってください
	// - また、-nオプションを指定すると、行版後を各行につけて表示されるようにしてください
	// - なお行番号はすべてのファイルで通し番号にしてください

	flag.Parse()
	var showLineNumber = *n

	for k, v := range flag.Args() {
		if showLineNumber {
			fmt.Println(k, v)
		} else {
			fmt.Println(v)
		}
	}
}
