package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var a, b, c int
	a, b, c = 1, 2, 0
	for i := 3; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

func main() {
	println(climbStairs(2))
	println(climbStairs(3))
}
