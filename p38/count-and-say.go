package p38

import "strings"

/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

// Test ...
func Test() string {
	return countAndSay(5)
}

// @lc code=start
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	preStr := countAndSay(n - 1)
	i, j := 0, 1
	var newStr strings.Builder
	for i < len(preStr) {
		if i+1 < len(preStr) {
			if preStr[i+1] != preStr[i] {
				newStr.WriteByte(byte(j + 48))
				newStr.WriteByte(preStr[i])
				j = 1
				i++
				continue
			}
			if preStr[i+1] == preStr[i] {
				i++
				j++
				continue
			}
		}
		if i+1 == len(preStr) {
			newStr.WriteByte(byte(j + 48))
			newStr.WriteByte(preStr[i])
			break
		}
	}
	return newStr.String()
}

// @lc code=end
