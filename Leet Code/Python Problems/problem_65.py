class Solution:
    def isNumber(self, s: str) -> bool:
        if s in ['inf', '-inf', '+inf', "Infinity", "-Infinity", "+Infinity"]:
            return False
        try:
            convert_s = float(s.strip())
            return True
        except ValueError:
            return False
