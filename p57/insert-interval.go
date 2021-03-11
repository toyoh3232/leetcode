package p57

import "sort"

/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */

// Test ...
func Test() [][]int {
	a := [][]int{{1, 3}, {6, 9}}
	b := []int{2, 5}
	return insert(a, b)
}

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	for i := 0; i < len(intervals); i++ {
		if ok, v := overlap(intervals[i], newInterval); ok {
			return insert(append(intervals[0:i], intervals[i+1:]...), v)
		}
	}
	newIntervals := append(intervals, newInterval)
	sorts(newIntervals)
	return newIntervals
}

func sorts(intervals [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})
}

func overlap(interval1, interval2 []int) (bool, []int) {
	intervals := [][]int{interval1, interval2}
	sorts(intervals)
	a, b, c, d := intervals[0][0], intervals[0][1], intervals[1][0], intervals[1][1]
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
