class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        charMap = {}
        if len(s) != len(t):
            return False
        for i, ch in enumerate(s):
            if ch in charMap:
                if charMap[ch] != t[i]:
                    return False
            else:
                if t[i] in charMap.values():
                    return False
                charMap[ch] = t[i]
        return True
        