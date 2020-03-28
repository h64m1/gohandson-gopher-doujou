package main

import (
	"fmt"
	"math/rand"
	"time"
)

func omikuji(n int) string {
	switch n {
	case 1:
		return "凶"
	case 2, 3:
		return "吉"
	case 4, 5:
		return "中吉"
	case 6:
		return "大吉"
	default:
		return ""
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		var dice = rand.Intn(6) + 1
		var unsei = omikuji(dice)
		fmt.Println("サイコロの目:", dice, ", 運勢:", unsei)
	}
}
