package main

import "fmt"

func main() {
	// if
	fmt.Println("--- ifで表示")
	for i := 0; i < 100; i++ {
		var number = i + 1
		fmt.Printf("%d-%s\n", number, evenOddStringIf(number))
	}

	// switch
	fmt.Println("--- switchで表示")
	for i := 0; i < 100; i++ {
		var number = i + 1
		fmt.Printf("%d-%s\n", number, evenOddStringSwitch(number))
	}
}
