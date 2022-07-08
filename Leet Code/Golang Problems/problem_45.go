package main

import "math"

func jump(nums []int) int {
	target := len(nums) - 1
	dp := make([]int, len(nums))
	for index := 0; index < len(nums); index++ {
		current := dp[index] + 1
		for next := index + 1; next < len(nums) && next <= index+nums[index]; next++ {
			if dp[next] != 0 {
				dp[next] = int(math.Min(float64(dp[next]), float64(current)))
			} else {
				dp[next] = current
			}
			if next == target {
				break
			}
		}
	}
	return dp[target]
}
