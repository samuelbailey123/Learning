package main

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	// special case of empty matrix
	if len(matrix) == 0 {
		return result
	}

	// get number of layered rectangles in matrix
	numR := (len(matrix)) / 2
	if numR2 := (len(matrix[0])) / 2; numR2 < numR {
		numR = numR2
	}

	// for each rectangle add to results the outside edge in spiral order
	i := 0
	for ; i < numR; i++ {
		for x := i; x < len(matrix[0])-i; x++ {
			result = append(result, matrix[i][x])
		}
		for x := i + 1; x < len(matrix)-i; x++ {
			result = append(result, matrix[x][len(matrix[0])-1-i])
		}
		for x := i + 1; x < len(matrix[0])-i; x++ {
			result = append(result, matrix[len(matrix)-1-i][len(matrix[0])-1-x])
		}
		for x := i + 1; x < len(matrix)-i-1; x++ {
			result = append(result, matrix[len(matrix)-1-x][i])
		}
	}

	// if odd rows or cols, then add remaining line
	oddRows := len(matrix)%2 == 1
	oddCols := len(matrix[0])%2 == 1
	if oddRows && (!oddCols || len(matrix) <= len(matrix[0])) {
		for x := i; x < len(matrix[0])-i; x++ {
			result = append(result, matrix[i][x])
		}
	} else if oddCols {
		for x := i; x < len(matrix)-i; x++ {
			result = append(result, matrix[x][i])
		}
	}

	return result
}
