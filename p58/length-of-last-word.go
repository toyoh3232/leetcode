package p58

import "regexp"

/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// Test ...
func Test() int {
	s := " "
	return lengthOfLastWord(s)
}

// @lc code=start
func lengthOfLastWord(s string) int {
	re := regexp.MustCompile(`[A-Za-z]+`)
	ss := re.FindAllString(s, -1)
	if len(ss) == 0 {
		return 0
	}
	return len(ss[len(ss)-1])

}

// @lc code=end
