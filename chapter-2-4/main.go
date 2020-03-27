package main

import "fmt"

func evenOddStringIf(n int) string {
	if n%2 == 0 {
		return "偶数"
	} else {
		return "奇数"
	}
}

func evenOddStringSwitch(n int) string {
	switch {
	case n%2 == 0:
		return "偶数"
	default:
		return "奇数"
	}
}

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
