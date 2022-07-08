package main

import "sort"

//Given an array of strings strs, group the anagrams together. You can return the answer in any order.
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		key := sortString(str)
		m[key] = append(m[key], str)
	}
	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func sortString(str string) string {
	s := []byte(str)
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return string(s)
}
