package main

//Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.
func searchRange(nums []int, target int) []int {
	var left, right int
	for left < len(nums) && nums[left] < target {
		left++
	}
	for right < len(nums) && nums[right] <= target {
		right++
	}
	if left == right {
		return []int{-1, -1}
	}
	return []int{left, right - 1}
}
