# Given an integer x, return true if x is palindrome integer.
# An integer is a palindrome when it reads the same backward as forward.
class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False
        else:
            x = str(x)
            for i in range(len(x)):
                if x[i] != x[len(x) - 1 - i]:
                    return False
            return True