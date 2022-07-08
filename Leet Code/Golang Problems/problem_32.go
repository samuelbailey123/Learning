package main

//Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.
func longestValidParentheses(s string) int {
	var stack []int
	stack = append(stack, -1)
	max := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				max = maxInt(max, i-stack[len(stack)-1])
			}
		}
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
