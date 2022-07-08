package main

//Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.

func permuteUnique(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	} else {
		var result [][]int
		dict := make(map[int]struct{})
		for i, n := range nums {
			if _, ok := dict[n]; ok {
				continue
			}
			dict[n] = struct{}{}
			nums2 := make([]int, len(nums))
			copy(nums2, nums)
			temp := permuteUnique(append(nums2[:i], nums2[i+1:]...))
			for _, m := range temp {
				m = append(m, n)
				result = append(result, m)
			}
		}
		return result
	}
}
