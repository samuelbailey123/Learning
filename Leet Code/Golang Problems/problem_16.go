package main

import "sort"

// Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.
// Return the sum of the three integers. You may assume that each input would have exactly one solution.
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	n := len(nums)
	for i := 0; i < n; i++ {
		j, k := i+1, n-1
		for j < k {
			v := nums[i] + nums[j] + nums[k]
			if abs(target-v) < abs(target-ans) {
				ans = v
			}
			if v <= target {
				j++
			} else {
				k--
			}
		}
	}
	return ans
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
