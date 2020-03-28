package main

func evenOddStringIf(n int) string {
	if n%2 == 0 {
		return "偶数"
	}

	return "奇数"
}

func evenOddStringSwitch(n int) string {
	switch {
	case n%2 == 0:
		return "偶数"
	default:
		return "奇数"
	}
}
