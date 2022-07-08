class Solution:
    def getRow(self, rowIndex: int) -> List[int]:
        result = [[1], [1, 1]]
        def helper(row):
            if row > rowIndex:
                return result
            else:
                last_row = result[-1]
                middle = [last_row[i] + last_row[i+1] for i in range(0, len(last_row)-1)]
                result.append([1] + middle + [1]) 
                helper(row + 1)
        helper(0)
        return result[rowIndex]
        