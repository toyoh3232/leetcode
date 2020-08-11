package p15

import "sort"

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start

func threeSum(nums []int) [][]int {
	count := len(nums)
	sort.Ints(nums)
	var res [][]int
	for i := range nums {
		if i-1 >= 0 && nums[i-1] == nums[i] {
			continue
		}
		if nums[i] > 0 || i+2 == count {
			break
		}
		for m, n := i+1, count-1; m < n; {
			if m-1 > i && nums[m-1] == nums[m] {
				m++
				continue
			}
			if n+1 < count && nums[n] == nums[n+1] {
				n--
				continue
			}
			if nums[m]+nums[n] == -nums[i] {
				res = append(res, []int{nums[i], nums[m], nums[n]})
				m++
				n--
				continue
			}
			if nums[m]+nums[n] > -nums[i] {
				n--
				continue
			}
			if nums[m]+nums[n] < -nums[i] {
				m++
				continue
			}

		}
	}
	return res
}

// @lc code=end
