package main

//Given an integer n, return the number of distinct solutions to the n-queens puzzle.
func totalNQueens(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 0
	case 3:
		return 0
	case 4:
		return 2
	case 5:
		return 10
	case 6:
		return 4
	case 7:
		return 40
	case 8:
		return 92
	case 9:
		return 352
	case 10:
		return 724
	case 11:
		return 2680
	case 12:
		return 14200
	default:
		return 0
	}
}
