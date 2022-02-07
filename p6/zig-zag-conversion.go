package p6

import (
	"strings"
)

/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	runes := make([][]rune, numRows)
	for i := range runes {
		runes[i] = make([]rune, len(s))
	}
	m, n := 0, numRows-1
	for i, c := range s {
		s := i % (numRows + numRows - 2)
		if s < numRows-1 {
			runes[s][m] = c
		} else if s == numRows-1 {
			runes[s][m] = c
			n, m = numRows-2, m+1
		} else {
			runes[n][m] = c
			m += 1
			n -= 1
		}
	}
	var sb strings.Builder
	for _, r := range runes {
		for _, c := range r {
			if c != 0 {
				sb.WriteString(string(c))
			}
		}
	}
	return sb.String()
}

// @lc code=end
