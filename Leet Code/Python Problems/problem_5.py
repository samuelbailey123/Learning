# Given a string s, return the longest palindromic substring in s.
class Solution:
    def longestPalindrome(self, s: str) -> str:
        if not s:
            return ''
        if len(s) == 1:
            return s
        max_length = 1
        longest_str = s[0]
        for i in range(len(s)):
            if i - max_length >= 1 and s[i - max_length - 1:i + 1] == s[i - max_length - 1:i + 1][::-1]:
                max_length += 2
                longest_str = s[i - max_length + 1:i + 1]
            elif i - max_length >= 0 and s[i - max_length:i + 1] == s[i - max_length:i + 1][::-1]:
                max_length += 1
                longest_str = s[i - max_length:i + 1]
        return longest_str
        
#incorrect solution 3 attempts