package p11

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	max := 0
	for i, v := range height {
		if getMax(v, i, height) > max {
			max = getMax(v, i, height)
		}
	}
	return max
}

func getMax(val, ind int, height []int) int {
	max, max2 := 0, 0
	for i, v := range height {
		if v >= val && i-ind > 0 {
			max2 = (i - ind) * val
		} else if v >= val && i-ind < 0 {
			max2 = (ind - i) * val
		}
		if max2 > max {
			max = max2
		}
	}
	return max
}

// @lc code=end
