# Given a string s, find the length of the longest substring without repeating characters.
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        max_len = 0
        start = 0
        d = {}
        for i in range(len(s)):
            if s[i] in d and start <= d[s[i]]:
                start = d[s[i]] + 1
            else:
                max_len = max(max_len, i - start + 1)
            d[s[i]] = i
        return max_len
