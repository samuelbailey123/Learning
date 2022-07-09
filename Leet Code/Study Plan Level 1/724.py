class Solution:
    def __init__(self): 
        pass
    
    def pivotIndex(self, nums: list[int]) -> int:
        left_sum = 0
        right_sum = sum(nums)
        for i in range(len(nums)):
            right_sum -= nums[i]
            if left_sum == right_sum:
                return i
            left_sum += nums[i]
        return -1

print(Solution.pivotIndex([1,7,3,6,5,6])) # 3
