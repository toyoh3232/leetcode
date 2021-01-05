package p42

/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// Test ...
func Test() int {
	height := []int{4}
	return trap(height)
}

// @lc code=start
func trap(height []int) int {
	if len(height) < 2 {
		return 0
	}
	v := 0
	y, x := max(height), len(height)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if check(i, j, height) {
				v++
			}
		}
	}
	return v
}

func max(s []int) int {
	v := s[0]
	for i := range s {
		if s[i] > v {
			v = s[i]
		}
	}
	return v
}

func check(x, y int, data []int) bool {
	if data[x] > y {
		return false
	}
	leftChk := func(xx int) bool {
		for i := xx - 1; i >= 0; i-- {
			if data[i] > y {
				return true
			}
		}
		return false
	}(x)
	rightChk := func(xx int) bool {
		for j := xx + 1; j < len(data); j++ {
			if data[j] > y {
				return true
			}
		}
		return false
	}(x)

	return leftChk && rightChk
}

// @lc code=end
