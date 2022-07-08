package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	sub := []int{}
	sort.Ints(candidates)
	helper(candidates, target, &result, sub)
	return result
}

func helper(candidates []int, target int, result *[][]int, sub []int) {

	if len(candidates) == 0 {
		return
	}
	if candidates[0] == target {
		sub = append(sub, candidates[0])
		*result = append(*result, sub)
		return
	} else if candidates[0] < target {
		for i := 0; i < len(candidates)-1; i++ {
			if candidates[i+1] != candidates[i] {
				helper(candidates[i+1:], target, result, sub)
				sub2 := make([]int, len(sub))
				copy(sub2, sub)
				sub2 = append(sub2, candidates[0])
				helper(candidates[1:], target-candidates[0], result, sub2)
				return
			}
		}
		sub3 := make([]int, len(sub)) //the case when all the rest numbers are the same
		copy(sub3, sub)
		sub3 = append(sub3, candidates[0])
		helper(candidates[1:], target-candidates[0], result, sub3)
	} else {
		return
	}
}
