package mycat

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var n = flag.Bool("n", false, "行番号を表示")
var lineNumber int = 1

// Cat my cat command
func Cat() {
	// catコマンドの仕様
	// - 引数でファイルパスの一覧をもらい、そのファイルに与えられた順に標準出力にそのまま出力されるコマンドを作ってください
	// - また、-nオプションを指定すると、行番号を各行につけて表示されるようにしてください
	// - なお行番号はすべてのファイルで通し番号にしてください

	flag.Parse()
	var fileNames = flag.Args()
	for _, fileName := range fileNames {
		err := readFile(fileName)
		if err != nil {
		}
	}
}

func readFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	var showLineNumber = *n

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if showLineNumber {
			fmt.Println(lineNumber, scanner.Text())
			lineNumber++
		} else {
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました:", err)
	}

	return err
}
