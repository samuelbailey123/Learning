package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	list := strings.Split(s, " ")
	return len(list[len(list)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}
