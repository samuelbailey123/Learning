package main

import "fmt"

// func buildArray(nums []int) []int {
// 	res := make([]int, 0)
// 	for i := 1; i <= len(nums); i++ {
// 		res = append(res, i)
// 	}
// 	return res
// }

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range nums {
		ans[i] = nums[nums[i]]
	}
	return ans
}

func main() {
	fmt.Println(buildArray([]int{0, 2, 1, 5, 3, 4}))
}
