package p36

/*
 * @lc app=leetcode id=36 lang=golang
 *
 * {36} Valid Sudoku
 */

// Test ...
func Test() bool {
	a := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	return isValidSudoku(a)
}

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	return isValidCage(board) && isValidCol(board) && isValidRow(board)
}

func isValidCol(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if !isValid([]byte{board[0][i], board[1][i], board[2][i],
			board[3][i], board[4][i], board[5][i],
			board[6][i], board[7][i], board[8][i]}) {
			return false
		}
	}
	return true
}

func isValidRow(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if !isValid(board[i]) {
			return false
		}
	}
	return true
}

func isValidCage(board [][]byte) bool {
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !isValid([]byte{board[i][j], board[i][j+1], board[i][j+2],
				board[i+1][j], board[i+1][j+1], board[i+1][j+2],
				board[i+2][j], board[i+2][j+1], board[i+2][j+2]}) {
				return false
			}
		}
	}
	return true
}

func isValid(array []byte) bool {
	m := make(map[byte]bool)
	for i := range array {
		if array[i] != '.' {
			if _, ok := m[array[i]]; ok {
				return false
			}
			m[array[i]] = true
		}
	}
	return true
}

// @lc code=end
