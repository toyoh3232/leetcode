package p3

import "strings"

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	si := 0
	ei := 0
	max := 0
	for ei < len(s) {
		prefix := s[si:ei]
		cur_ch := s[ei : ei+1]
		if !strings.Contains(prefix, cur_ch) {
			ei++
			if len(prefix)+1 > max {
				max = len(prefix) + 1
			}
		} else {
			si = si + strings.Index(prefix, cur_ch) + 1
			ei = si
		}
	}
	return max
}

// @lc code=end
