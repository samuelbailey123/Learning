package main

//You are given a string s and an array of strings words of the same length. Return all starting indices of substring(s) in s that is a concatenation of each word in words exactly once, in any order, and without any intervening characters.
func findSubstring(s string, words []string) []int {
	out := []int{}
	totalWords := len(words)
	if totalWords == 0 {
		return out
	}
	eachWordLength := len(words[0])
	stringLength := len(s)
	wordsCountMap := map[string]int{}
	for _, word := range words {
		wordsCountMap[word]++
	}
	for i := 0; i < eachWordLength; i++ {
		start := i
		for start+(eachWordLength*totalWords) <= stringLength {
			sub := s[start : start+(eachWordLength*totalWords)]
			visited := map[string]int{}
			j := totalWords
			for j > 0 {
				subString := sub[eachWordLength*(j-1) : eachWordLength*j]
				visited[subString]++
				if visited[subString] > wordsCountMap[subString] {
					break
				}
				j--
			}
			if j == 0 {
				out = append(out, start)
			}
			max := 1
			if j > max {
				max = j
			}
			start += eachWordLength * max
		}
	}
	return out
}
