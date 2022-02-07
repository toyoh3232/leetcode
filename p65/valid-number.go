package p65

import (
	"regexp"
)

/*
 * @lc app=leetcode id=65 lang=golang
 *
 * [65] Valid Number
 */

// @lc code=start
func isNumber(s string) bool {
	d1 := `^[\-\+]{0,1}[0-9]+\.($|[eE][\+\-]{0,1}[0-9]+$)`
	d2 := `^[\-\+]{0,1}[0-9]+\.[0-9]+($|[eE][\+\-]{0,1}[0-9]+$)`
	d3 := `^[\-\+]{0,1}\.[0-9]+($|[eE][\+\-]{0,1}[0-9]+$)`
	i1 := `^[\-\+]{0,1}[0-9]+($|[eE][\+\-]{0,1}[0-9]+$)`
	r1 := regexp.MustCompile(d1)
	r2 := regexp.MustCompile(d2)
	r3 := regexp.MustCompile(d3)
	r4 := regexp.MustCompile(i1)
	return r1.MatchString(s) || r2.MatchString(s) || r3.MatchString(s) || r4.MatchString(s)
}

// @lc code=end
