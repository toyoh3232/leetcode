package p40

import "sort"

/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// Test ...
func Test() [][]int {
	candidates, target := []int{1}, 1
	sort.Ints(candidates)
	return combinationSum2(candidates, target)
}

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
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
		if len(pers) == 0 || fold(pers[0], 0, add) > target {
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

func amountOf(array []int, value int) int {
	out := 0
	for _, v := range array {
		if v == value {
			out++
		}
	}
	return out
}

func rub(array []int) []int {
	dup := []int{}
	for _, v := range array {
		if amountOf(dup, v) == 0 {
			dup = append(dup, v)
		}
	}
	return dup
}
func permutation(array []int, times int) [][]int {
	out := [][]int{}
	if len(array) == 0 {
		return out
	}
	if times == 1 {
		for _, v := range rub(array) {
			out = append(out, []int{v})
		}
	} else {
		pre := permutation(array, times-1)
		for _, preList := range pre {
			for _, addValue := range rub(array) {

				if preList[times-2] <= addValue && amountOf(array, addValue) != amountOf(preList, addValue) {
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
