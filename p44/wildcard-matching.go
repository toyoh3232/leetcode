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
	nfa := automata.NewNFA("1", "2")
	nfa.Add("1", 'a', "2")
	fmt.Println(nfa.String())
	dfa := nfa.Compile()
	return dfa.Execute(("a"))
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
