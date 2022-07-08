package main

import "fmt"

func addBinary(a string, b string) string {
	var result string
	var carry int
	var i, j int
	for i, j = len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var sum int
		if i >= 0 {
			sum += int(a[i] - '0')
		}
		if j >= 0 {
			sum += int(b[j] - '0')
		}
		sum += carry
		carry = sum / 2
		result = fmt.Sprintf("%d%s", sum%2, result)
	}
	if carry > 0 {
		result = fmt.Sprintf("%d%s", carry, result)
	}
	return result
}

func main() {
	fmt.Println(addBinary("11", "1"))
}
