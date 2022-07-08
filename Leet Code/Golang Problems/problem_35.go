package main

//Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
func searchInsert(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		} else if v > target {
			return i
		}
	}
	return len(nums)
}
