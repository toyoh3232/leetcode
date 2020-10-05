package p50

/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// Test ...
func Test() float64 {
	return myPow(2.0, -2)
}

// @lc code=start
func myPow(x float64, n int) float64 {
	switch {
	case n < 0:
		return myPow(1/x, -n)
	case n == 0:
		return 1
	case n == 1:
		return x
	default:
		bn := n / 2
		bf := myPow(x, bn)
		if n%2 == 0 {
			return bf * bf
		}
		return bf * bf * x
	}
}

// @lc code=end
