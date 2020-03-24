package main

import "fmt"

func evenOddString(n int) string {
	if n%2 == 0 {
		return "偶数"
	} else {
		return "奇数"
	}
}

func main() {
	for i := 0; i < 100; i++ {
		var number = i + 1
		fmt.Printf("%d-%s\n", number, evenOddString(number))
	}
}
