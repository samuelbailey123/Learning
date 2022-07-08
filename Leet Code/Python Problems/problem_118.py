class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        if numRows == 0:
            return []
        if numRows == 1:
            return [[1]]
        if numRows == 2:
            return [[1], [1, 1]]
        result = [[1], [1, 1]]
        for i in range(2, numRows):
            result.append([1] + [result[i - 1][j] + result[i - 1][j + 1] for j in range(i - 1)] + [1])
        return result
        