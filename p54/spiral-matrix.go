package p54

/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// Test ...
func Test() []int {
	a := [][]int{{1, 2}}
	return spiralOrder(a)
}

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	position := next([][]int{{0, 0}}, 4, len(matrix), len(matrix[0]))
	r := []int{}
	for _, v := range position {
		r = append(r, matrix[v[0]][v[1]])
	}
	return r
}

func next(visited [][]int, direction, rows, cols int) [][]int {
	//base case
	if len(visited) == rows*cols {
		return visited
	}

	current := visited[len(visited)-1]
	i, j := current[0], current[1]
	ni, nj := i, j
	// south:1, west:2,north:3, east:4
	switch direction {
	case 4:
		nj = nj + 1
		break
	case 1:
		ni = ni + 1
		break
	case 2:
		nj = nj - 1
		break
	case 3:
		ni = ni - 1
		break
	}

	if has(visited, ni, nj) || ni == rows || nj == cols || ni < 0 || nj < 0 {
		// turn
		// south:1, west:2,north:3, east:4
		switch direction {
		case 4:
			direction = 1
		case 1:
			direction = 2
		case 2:
			direction = 3
		case 3:
			direction = 4
		}
	} else {
		visited = append(visited, []int{ni, nj})
	}
	return next(visited, direction, rows, cols)
}

func has(visited [][]int, i, j int) bool {
	for _, arr := range visited {
		if i == arr[0] && j == arr[1] {
			return true
		}
	}
	return false
}

// @lc code=end
