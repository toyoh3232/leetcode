package p10

/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

// Run ...
func Run() {

}

// @lc code=start
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if len(p) == 1 || p[1] != '*' {
		if len(s) == 0 {
			return false
		}
		if s[0] != p[0] && p[0] != '.' {
			return false
		}
		return isMatch(s[1:], p[1:])
	}
	for i := 0; i <= len(s); i++ {
		str := make([]byte, i)
		for i := range str {
			str[i] = p[0]
		}
		if isMatch(s, string(str)+p[2:]) {
			return true
		}
	}
	return false
}

// @lc code=end
