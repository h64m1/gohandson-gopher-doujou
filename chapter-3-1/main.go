package main

func main() {
	// 組み込み型（数値）
	var sum int
	sum = 5 + 6 + 3
	avg := float32(sum) / 3
	if avg > 4.5 {
		println("good", avg)
	}

	// 組み込み型（真偽値）
	var a, b, c bool
	for i := 0; i < 2; i++ {
		a = true
		if i == 1 {
			a = false
		}
		for j := 0; j < 2; j++ {
			b = true
			if j == 1 {
				b = false
			}
			for k := 0; k < 2; k++ {
				c = true
				if k == 1 {
					c = false
				}
				if a && b || !c {
					println(a, b, c, ": true")
				} else {
					println(a, b, c, ": false")
				}
			}
		}
	}

}
