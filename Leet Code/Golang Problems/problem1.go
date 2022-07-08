package main

func twoSum(nums []int, target int) []int {
	var record = make(map[int]int)
	var result []int
	for key, value := range nums {
		if _, ok := record[target-value]; ok {
			result = append(result, key)
			result = append(result, record[target-value])
			break
		}
		record[value] = key
	}
	return result
}
