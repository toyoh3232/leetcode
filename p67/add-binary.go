package p67

/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a string, b string) string {
	if len(a) > len(b) {
		a, b = b, a
	}
	c := make([]rune, len(b))
	i, lhs, rhs, car := 1, 0, 0, 0
	for {
		if len(b)-i < 0 {
			break
		}
		if len(a)-i >= 0 && a[len(a)-i] == '1' {
			lhs = 1
		}
		if b[len(b)-i] == '1' {
			rhs = 1
		}
		switch lhs + rhs + car {
		case 0:
			c[len(c)-i] = '0'
		case 1:
			c[len(c)-i] = '1'
			lhs, rhs, car = 0, 0, 0
		case 2:
			c[len(c)-i] = '0'
			lhs, rhs, car = 0, 0, 1
		case 3:
			c[len(c)-i] = '1'
			lhs, rhs, car = 0, 0, 1
		}
		i += 1
	}
	if car == 1 {
		c = append([]rune{'1'}, c...)
	}
	return string(c)
}

// @lc code=end
