package p8

/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start

const int32Max = int32(^uint32(0) >> 1)
const int32Min = -int32Max - 1

func myAtoi(s string) int {
	return strToInt32(toIntString(s))
}

func toIntString(s string) string {
	for len(s) > 0 && s[0] == ' ' {
		s = s[1:]
	}
	signed := true

	if len(s) > 0 {
		if s[0] == '-' {
			signed = false
			s = s[1:]
		} else if s[0] == '+' {
			s = s[1:]
		}
	}
	i := 0
	for i < len(s) && s[i] >= 48 && s[i] < 58 {
		i++
	}
	if i > 0 {
		if signed {
			return s[0:i]
		}
		return "-" + s[0:i]
	}
	return ""
}

func strToInt32(str string) int {
	if len(str) == 0 {
		return 0
	}
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
	if len(str) >= 11 {
		if signed {
			return int(int32Max)
		} else {
			return int(int32Min)
		}
	}
	var i int
	var fac int = 1
	for j := len(str) - 1; j >= 0; j-- {
		v := int(byte(str[j]) - 48)
		c := v * fac
		if signed {
			i += c
			if i > int(int32Max) {
				return int(int32Max)
			}
		} else {
			i -= c
			if i < int(int32Min) {
				return int(int32Min)
			}
		}
		fac = fac * 10
	}
	return i
}

// @lc code=end
