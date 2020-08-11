package p28

/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {

	index := -1
	if strStartWith(haystack, needle) {
		return 0
	}
	for i := range haystack {
		if strStartWith(haystack[i:], needle) {
			return i
		}
	}
	return index
}

func strStartWith(haystack string, needle string) bool {
	if len(needle) == 0 {
		return true
	}
	if len(needle) <= len(haystack) {
		for i := range needle {
			if needle[i] != haystack[i] {
				return false
			}
		}
		return true
	}
	return false
}

// @lc code=end
