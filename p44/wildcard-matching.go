package p44

import "strings"

/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 */

// Test ...
func Test() bool {
	s := "mississippi"
	p := "m*issip*"
	return isMatch(s, p)
}

// @lc code=start

//Krauss's way
func isMatch(s string, p string) bool {
	s, p = rub(s), rub(p)
	wd, pr := false, ""
	wdn, sn := false, ""
	for len(s) >= 0 {
		if p == "*" {
			return true
		}
		if len(s) == 0 && len(p) == 0 {
			return true
		}
		if (len(s) > 0 && p == "" && !wd) || (len(s) == 0 && len(p) >= 1) {
			return false
		}

		if wd && !wdn && s[0] == pr[0] {
			sn = s[0:]
			wdn = true
		}

		if p == "" && wd {
			p = pr
			continue
		} else if s[0] != p[0] {
			if p[0] == '*' {
				wd, pr = true, p[1:]
				p = pr
				continue
			}
			if p[0] == '?' {
				s, p = s[1:], p[1:]
				continue
			}
			if !wd {
				return false
			}
			if p != pr {
				p = pr
			} else if wdn {
				s, p = sn, pr
				wdn = false
			}
			continue
		} else {
			s, p = s[1:], p[1:]
		}

	}

	return true
}

//recursive way
func isMatch2(s string, p string) bool {
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
