package p17

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	letters := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	return letterCombinationsR(digits, letters)
}

func letterCombinationsR(digits string, letters []string) []string {
	if len(digits) == 0 {
		return []string{""}
	}
	idx := int(digits[0] - 48)
	others := letterCombinationsR(digits[1:], letters)
	res := make([]string, len(letters[idx])*len(others))
	for i, v2 := range letters[idx] {
		for j, v3 := range others {
			res[i*len(others)+j] = string(v2) + v3
		}
	}
	return res
}

// @lc code=end
