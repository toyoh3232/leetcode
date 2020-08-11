package p37

import (
	"fmt"

	. "github.com/crillab/gophersat/bf"
)

/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 */

// Test ...
func Test() {
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
	solveSudoku(a)
	printBoard(a)
}

func printBoard(board [][]byte) {
	for i := 0; i < 9; i++ {
		for _, v := range board[i] {
			fmt.Print(v-48, " ")
		}
		fmt.Println()
	}
}

// @lc code=start
func solveSudoku(board [][]byte) {
	const varFmt = "line-%d-col-%d:%d"
	f := True
	// In each spot, exactly one number is written
	for line := 1; line <= 9; line++ {
		for col := 1; col <= 9; col++ {
			vars := make([]string, 9)
			for val := 1; val <= 9; val++ {
				vars[val-1] = fmt.Sprintf(varFmt, line, col, val)
			}
			f = And(f, Unique(vars...))
		}
	}
	// In each line, each number appears at least once.
	// Since there are 9 spots and 9 numbers, that means each number appears exactly once.
	for line := 1; line <= 9; line++ {
		for val := 1; val <= 9; val++ {
			var vars []Formula
			for col := 1; col <= 9; col++ {
				vars = append(vars, Var(fmt.Sprintf(varFmt, line, col, val)))
			}
			f = And(f, Or(vars...))
		}
	}
	// In each column, each number appears at least once.
	for col := 1; col <= 9; col++ {
		for val := 1; val <= 9; val++ {
			var vars []Formula
			for line := 1; line <= 9; line++ {
				vars = append(vars, Var(fmt.Sprintf(varFmt, line, col, val)))
			}
			f = And(f, Or(vars...))
		}
	}
	// In each 3x3 box, each number appears at least once.
	for lineB := 0; lineB < 3; lineB++ {
		for colB := 0; colB < 3; colB++ {
			for val := 1; val <= 9; val++ {
				var vars []Formula
				for lineOff := 1; lineOff <= 3; lineOff++ {
					line := lineB*3 + lineOff
					for colOff := 1; colOff <= 3; colOff++ {
						col := colB*3 + colOff
						vars = append(vars, Var(fmt.Sprintf(varFmt, line, col, val)))
					}
				}
				f = And(f, Or(vars...))
			}
		}
	}
	// Some spots already have a fixed value
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				f = And(f, Var(fmt.Sprintf(varFmt, i+1, j+1, board[i][j]-48)))
			}
		}
	}
	model := Solve(f)
	if model == nil {
		return
	}
	for line := 1; line <= 9; line++ {
		for col := 1; col <= 9; col++ {
			for val := 1; val <= 9; val++ {
				if model[fmt.Sprintf(varFmt, line, col, val)] {
					board[line-1][col-1] = byte(val + 48)
				}
			}
		}
	}
}

// @lc code=end
