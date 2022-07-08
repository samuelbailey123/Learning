package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isMatch(s string, p string) bool {
	var newPattern = []string{}
	for _, c := range p {
		if string(c) == "*" || string(c) == "?" {
			if string(c) == "*" {
				newPattern = append(newPattern, "[[:lower:]]")
			}
			if string(c) == "?" {
				newPattern = append(newPattern, "[[:lower:]]{1}")
			}
			newPattern = append(newPattern, string(c))
			continue
		}
		newPattern = append(newPattern, string(c))
	}
	p = strings.Join(newPattern, "")
	re := regexp.MustCompile(fmt.Sprintf("^%s$", p))
	return re.Match([]byte(s))
}
