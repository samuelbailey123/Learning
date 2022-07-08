package main

//Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

func permute(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	if len(nums) == 1 {
		res = append(res, nums)
		return res
	}
	for i := 0; i < len(nums); i++ {
		tmp := make([]int, len(nums)-1)
		copy(tmp, nums[:i])
		copy(tmp[i:], nums[i+1:])
		for _, v := range permute(tmp) {
			res = append(res, append([]int{nums[i]}, v...))
		}
	}
	return res
}
