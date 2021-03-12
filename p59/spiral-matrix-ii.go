package p59

/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 */

// Test ...
func Test() [][]int {
	return generateMatrix(20)
}

// @lc code=start
func generateMatrix(n int) [][]int {
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}
	i := 1
	for _, pos := range next([][]int{{0, 0}}, 4, n, n) {
		m[pos[0]][pos[1]] = i
		i++
	}
	return m
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
