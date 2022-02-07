package p5

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	m := make(chan string)
	ab := getAlphabet(s)
	for _, r := range ab {
		go func(r rune) {
			m <- getMaxPalindrome(s, r)
		}(r)
	}
	maxes := make([]string, len(ab))

	for i := 0; i < len(ab); i++ {
		maxes[i] = <-m
	}
	sort.Sort(byLength(maxes))
	return maxes[0]

}

func getAlphabet(s string) string {
	var m = make(map[rune]bool)
	var sb strings.Builder
	for _, c := range s {
		m[c] = true
	}
	for k := range m {
		sb.WriteString(string(k))
	}
	return sb.String()
}

func getMaxPalindrome(target string, c rune) string {

	indexes := make([]int, 0)
	max := 0
	var maxStr string
	for i, r := range target {
		if r == c {
			indexes = append(indexes, i)
		}
	}
	for _, i := range indexes {
		for _, j := range indexes {
			if j >= i && isPalindrome(target[i:j+1]) && max < j-i+1 {
				max = j - i + 1
				maxStr = target[i : j+1]
			}
		}
	}
	return maxStr

}

func isPalindrome(s string) bool {
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

// @lc code=end
