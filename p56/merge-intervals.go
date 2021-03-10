package p56

import "sort"

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// Test ...
func Test() [][]int {
	a := [][]int{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}}
	return merge(a)
}

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})
	for i, j := 0, 1; i < j && j < len(intervals); i, j = i+1, j+1 {
		fir, sec := intervals[i], intervals[j]
		if ok, v := overlap(fir, sec); ok {
			sec[0], sec[1] = v[0], v[1]
			return merge(append(intervals[0:i], intervals[j:]...))
		}
	}
	return intervals
}

func equal(interval1, interval2 []int) bool {

	return interval1[0] == interval2[0] && interval1[1] == interval2[1]
}

func overlap(interval1, interval2 []int) (bool, []int) {
	a, b, c, d := interval1[0], interval1[1], interval2[0], interval2[1]
	if b < c {
		return false, []int{}
	}
	if b <= d {
		return true, []int{a, d}
	}
	if b > d {
		return true, []int{a, b}
	}

	return false, []int{}
}

// @lc code=end
