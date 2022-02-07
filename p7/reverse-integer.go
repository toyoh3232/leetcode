package p7

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	if x < 0 {
		return int(strToInt32("-" + reverseString(int32ToStr(int32(x))[1:])))
	}
	return int(strToInt32(reverseString(int32ToStr(int32(x)))))

}

func reverseString(str string) string {
	a, lth := []rune(str), len(str)
	for i := 0; i < lth/2; i++ {
		t := a[i]
		a[i] = a[lth-i-1]
		a[lth-i-1] = t
	}
	return string(a)
}

func int32ToStr(i int32) string {
	if i == 0 {
		return "0"
	}
	signed, j := i > 0, i
	var digits int32

	for j != 0 {
		j /= 10
		digits++
	}
	str := make([]rune, digits)
	for k := digits - 1; k >= 0; k-- {
		if signed {
			str[k] = i%10 + 48
		} else {
			str[k] = -(i % 10) + 48
		}

		i /= 10
	}
	if signed {
		return string(str)
	}
	return "-" + string(str)

}

func strToInt32(str string) int32 {
	signed := str[0] != '-'
	if !signed {
		str = str[1:]
	}
	for str[0] == '0' && len(str) > 1 {
		str = str[1:]
	}
	if str == "0" {
		return 0
	}
	var i int32
	var fac int32 = 1
	for j := len(str) - 1; j >= 0; j-- {
		v := int32(byte(str[j]) - 48)
		c := v * fac
		if c/fac != v {
			return 0
		}
		if signed {
			i += c
			if i < 0 {
				return 0
			}
		} else {
			i -= c
			if i > 0 {
				return 0
			}
		}
		fac = fac * 10
	}
	return i
}

// @lc code=end
