package main

/*
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
func migratoryBirds(arr []int32) int32 {
	var birdMap = make(map[int32]int32)
	for _, v := range arr {
		if _, ok := birdMap[v]; ok {
			birdMap[v]++
		} else {
			birdMap[v] = 1
		}
	}
	var maxBird int32
	var maxCount int32
	for k, v := range birdMap {
		if v > maxCount {
			maxBird = k
			maxCount = v
		}
	}
	return maxBird
}
