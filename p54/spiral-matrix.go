package p54

/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// Test ...
func Test() [][]int {
	return next([][]int{{0, 0}}, 4, 3, 3)
}

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	return []int{}
}

func next(visited [][]int, direction, rows, cols int) [][]int {
	//base case
	if len(visited) == rows*cols {
		return visited
	}

	current := visited[0]
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

	if has(visited, ni, nj) || ni == rows || nj == cols {
		// turn
		// south:1, west:2,north:3, east:4
		switch direction {
		case 4:
			visited = next(visited, 1, rows, cols)
			break
		case 1:
			visited = next(visited, 2, rows, cols)
			break
		case 2:
			visited = next(visited, 3, rows, cols)
			break
		case 3:
			visited = next(visited, 4, rows, cols)
			break
		}
	} else {
		visited = append(visited, []int{ni, nj})
	}

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
