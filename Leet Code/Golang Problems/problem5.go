package main

type InputString struct {
	S      string
	Length int16
}

func (is *InputString) checkRevers(left, right int16) int16 {
	for left >= 0 && right < is.Length && is.S[left] == is.S[right] {
		left--
		right++
	}
	return right - left - 1
}

func longestPalindrome(s string) string {

	is := InputString{s, int16(len(s))}
	start := int16(0)
	end := int16(0)
	len := int16(0)

	if s == "" || is.Length < 1 {
		return ""
	}

	for i := int16(0); i < is.Length; i++ {
		len1 := is.checkRevers(i, i)
		len2 := is.checkRevers(i, i+1)
		if len1 > len2 {
			len = len1
		} else {
			len = len2
		}
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}

	return s[start : end+1]
}
