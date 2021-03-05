package p48

/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */

//Test ...
func Test() [][]int {
	m := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}
	rotate(m)
	return m
}

// @lc code=start
func rotate(matrix [][]int) {
	for d := len(matrix); d > 0; d = d - 2 {
		rotateEach(matrix, d)
	}
}

func rotateEach(matrix [][]int, d int) {
	n := len(matrix) - 1
	od := n + 1
	if d == 1 {
		return
	}
	i := 0 + (od-d)/2
	for j := i; j < i+d-1; j++ {
		var temp int
		temp = matrix[j][n-i]
		matrix[j][n-i] = matrix[i][j]
		matrix[i][j] = temp

		temp = matrix[n-i][n-j]
		matrix[n-i][n-j] = matrix[i][j]
		matrix[i][j] = temp

		temp = matrix[n-j][i]
		matrix[n-j][i] = matrix[i][j]
		matrix[i][j] = temp
	}
}

// @lc code=end
