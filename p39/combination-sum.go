package p39

import (
	"sort"
)

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// Test ...
func Test() [][]int {
	candidates, target := []int{2}, 1

	return combinationSum(candidates, target)
}

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	newCandidates := filter(candidates, func(v int) bool {
		return v <= target
	})

	sort.Ints(newCandidates)
	var out [][]int
	if len(newCandidates) == 0 {
		return out
	}
	for i := 1; ; i++ {
		pers := permutation(newCandidates, i)
		if fold(pers[0], 0, add) > target {
			break
		}
		for _, p := range pers {
			if fold(p, 0, add) == target {
				out = append(out, p)
			}
		}
	}
	return out
}

// Filter ...
func filter(array []int, f func(int) bool) []int {
	out := []int{}
	for _, v := range array {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

func fold(array []int, initial int, f func(int, int) int) int {
	out := initial
	for _, v := range array {
		out = f(out, v)
	}
	return out
}

func add(x, y int) int {
	return x + y
}

func permutation(array []int, times int) [][]int {
	out := [][]int{}
	if len(array) == 0 {
		return out
	}
	if times == 1 {
		for _, v := range array {
			out = append(out, []int{v})
		}
	} else {
		pre := permutation(array, times-1)
		for _, preList := range pre {
			for _, addValue := range array {
				if preList[times-2] <= addValue {
					addList := make([]int, 0, times)
					addList = append(addList, preList...)
					addList = append(addList, addValue)
					out = append(out, addList)
				}
			}
		}
	}

	return out
}

// @lc code=end
