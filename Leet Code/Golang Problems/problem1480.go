package main

func runningSum(nums []int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			res = append(res, nums[i])
		} else {
			res = append(res, res[i-1]+nums[i])
		}
	}
	return res
}
