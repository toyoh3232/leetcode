package p44

import "strings"

/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 */

// Test ...
func Test() string {
	//	s := "abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb"
	p := "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"
	return rub(p)
}

// @lc code=start
func isMatch(s string, p string) bool {
	p = rub(p)
	if s == "" && p == "" {
		return true
	}
	if len(p) > 0 {
		switch p[0] {
		case '?':
			if len(s) > 0 {
				return isMatch(s[1:], p[1:])
			}
			break
		case '*':
			for i := 0; i <= len(s); i++ {
				if isMatch(s[i:], p[1:]) {
					return true
				}
			}
			break
		default:
			if len(s) > 0 && s[0] == p[0] {
				return isMatch(s[1:], p[1:])
			}
			break
		}
	}
	return false
}

func rub(s string) string {
	var sb strings.Builder

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			if i+1 < len(s) && s[i+1] == '*' {
				continue
			}
		}
		sb.WriteByte(s[i])
	}
	return sb.String()
}

// @lc code=end
