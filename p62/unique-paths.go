package p62

/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */

// Test ...
func Test() int {
	m, n := 20, 10
	return uniquePaths(m, n)
}

// @lc code=start
func uniquePaths(m int, n int) int {
	m, n = m-1, n-1
	if n > m {
		m, n = n, m

	}
	a := make([]int, m+n)
	for i := range a {
		a[i] = i + 1
	}
	for i := range a {
		if a[i] <= m {
			a[i] = 1
		}
	}
	b := make([]int, n)
	for i := range b {
		b[i] = i + 1
	}
	for i := range b {
		if b[i] > 1 {
			j, v := get(a, b[i])
			if j > -1 {
				a[j] = v / b[i]
				b[i] = 1
			}
		}
	}
	return 0
}

func get(a []int, b int) (int, int) {
	for i := range a {
		if a[i]%b == 0 {
			return i, a[i]
		}
	}
	return -1, 0
}

// @lc code=end
