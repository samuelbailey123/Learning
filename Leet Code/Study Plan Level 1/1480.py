class Solution:
    def runningSum(self, nums: list[int]) -> list[int]: 
        running_sum = []
        x = 0 
        for i in nums: 
            x += i
            running_sum.append(x)
        return running_sum

nums = [1,2,3,4]
print(Solution.runningSum([1,2,3,4])) # 1 3 6 10