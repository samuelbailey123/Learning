package main

func divide(dividend int, divisor int) int {
	c := 0
	dividendS, r := SplitSignAndNum(dividend)
	divisorS, d := SplitSignAndNum(divisor)
	for r >= d {
		r = r - d
		c++
	}

	result := c

	if dividendS < 0 || divisorS < 0 {
		if !(dividendS < 0 && divisorS < 0) {
			result = -c
		}
	}

	if result < -2147483648 {
		return -2147483648
	} else if result > 2147483647 {
		return 2147483647
	}

	return result
}

func SplitSignAndNum(i int) (int, int) {
	if i < 0 {
		return -1, -i
	}
	return 1, i
}
