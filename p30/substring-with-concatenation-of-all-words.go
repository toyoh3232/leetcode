package p30

import "fmt"

/*
 * @lc app=leetcode id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 */

// Test ...
func Test() {
	s := "abcdefcdabef"
	words := []string{}
	fmt.Println(findSubstring(s, words))
}

// @lc code=start
func findSubstring(s string, words []string) []int {
	var out []int
	if len(words) != 0 {
		for i := range s {
			if iterateString(s[i:], []string{}, words) {
				out = append(out, i)
			}
		}
	}
	return out
}

func iterateString(s string, leftWords, rightWords []string) bool {

	if len(leftWords) == 0 && len(rightWords) == 0 {
		return true
	}

	for i, lpre := range leftWords {
		lt := len(lpre)
		if startWith(s, lpre) {
			return iterateString(s[lt:], leftWords[0:i], append(rightWords, leftWords[i+1:]...))
		}
	}
	for i, rpre := range rightWords {
		lt := len(rpre)
		if startWith(s, rpre) {
			return iterateString(s[lt:], append(leftWords, rightWords[0:i]...), rightWords[i+1:])
		}
	}
	return false
}

func startWith(s string, pre string) bool {
	if len(pre) == 0 {
		return true
	}
	if len(pre) > len(s) {
		return false
	}
	for i := range pre {
		if pre[i] != s[i] {
			return false
		}
	}
	return true

}

// @lc code=end
