package main

//Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:
// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if !checkRow(board, i, j) || !checkColumn(board, i, j) || !checkBox(board, i, j) {
					return false
				}
			}
		}
	}
	return true
}

func checkRow(board [][]byte, row, col int) bool {
	var set [9]bool
	for j := 0; j < 9; j++ {
		if board[row][j] != '.' {
			if set[board[row][j]-'1'] {
				return false
			}
			set[board[row][j]-'1'] = true
		}
	}
	return true
}

func checkColumn(board [][]byte, row, col int) bool {
	var set [9]bool
	for i := 0; i < 9; i++ {
		if board[i][col] != '.' {
			if set[board[i][col]-'1'] {
				return false
			}
			set[board[i][col]-'1'] = true
		}
	}
	return true
}

func checkBox(board [][]byte, row, col int) bool {
	var set [9]bool
	for i := row / 3 * 3; i < row/3*3+3; i++ {
		for j := col / 3 * 3; j < col/3*3+3; j++ {
			if board[i][j] != '.' {
				if set[board[i][j]-'1'] {
					return false
				}
				set[board[i][j]-'1'] = true
			}
		}
	}
	return true
}
