package p9

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	d, y := 0, x
	for y != 0 {
		y = y / 10
		d++
	}
	xs := make([]int, d)
	for i := range xs {
		xs[i] = x % 10
		x = x / 10
	}
	for len(xs) > 1 {
		if xs[0] != xs[len(xs)-1] {
			return false
		}
		xs = xs[1 : len(xs)-1]
	}
	return true
}

// @lc code=end
