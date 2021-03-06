package main

func canJump(nums []int) bool {
	prevGood := len(nums) - 1
	for i := prevGood - 1; i >= 0; i-- {
		if nums[i] >= prevGood-i {
			prevGood = i
		}
	}
	return prevGood == 0
}
