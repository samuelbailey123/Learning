package main

// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var res []string
	var m = map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
	var dfs func(string, int)
	dfs = func(s string, i int) {
		if i == len(digits) {
			res = append(res, s)
			return
		}
		for _, c := range m[digits[i]] {
			dfs(s+c, i+1)
		}
	}
	dfs("", 0)
	return res
}
