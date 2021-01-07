package p44

import (
	"fmt"
	"leetcode/automata"
)

/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 */

// Test ...
func Test() bool {
	s := "abba"
	p := "abba"
	return isMatch(s, p)
}

// @lc code=start

//Krauss's way
func isMatch(s string, p string) bool {
	nfa := automata.NewNFA("0", "2")
	nfa.AddTrasition("0", 'a', "0")
	nfa.AddTrasition("0", 'b', "1")
	nfa.AddTrasition("1", 'b', "2")
	fmt.Println(nfa.Discription())
	return false
}

func nub(p string) string {
	var ns []byte
	for i := 0; i < len(p); i++ {
		if i+1 < len(p) && p[i] == '*' && p[i+1] == '*' {
			continue
		}
		ns = append(ns, p[i])
	}
	return string(ns)
}

func lenP(p string) int {
	i := 0
	for _, v := range p {
		if v == '*' {
			continue
		}
		i++
	}
	return i
}

// @lc code=end
