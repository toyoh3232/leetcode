package p44

/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 */

// Test ...
func Test() bool {
	s := "abba"
	p := "abba"
	return isMatch(s, p)
}

// @lc code=start

//Krauss's way
func isMatch(s string, p string) bool {
	lenY, lenX := len(s)+1, len(p)+1
	dp := make([][]bool, lenY)
	for y := 0; y < lenY; y++ {
		dp[y] = make([]bool, lenX)
	}
	dp[0][0] = true
	for x := 1; x < lenX; x++ {
		if p[x-1] == '*' {
			dp[0][x] = dp[0][x-1]
		}
	}
	for y := 1; y < lenY; y++ {
		for x := 1; x < lenX; x++ {
			switch p[x-1] {
			case '*':
				dp[y][x] = dp[y][x-1] || dp[y-1][x]
			case '?', s[y-1]:
				dp[y][x] = dp[y-1][x-1]
			}
		}
	}
	return dp[len(s)][len(p)]
}

// @lc code=end
