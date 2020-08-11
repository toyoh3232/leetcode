package p34

/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// Test ...
func Test() []int {
	a := []int{5, 7, 7, 8, 8, 10}
	target := 10
	return searchRange(a, target)
}

// @lc code=start
func searchRange(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			first := find(nums, mid, target, func(i int) int {
				return i - 1
			})
			last := find(nums, mid, target, func(i int) int {
				return i + 1
			})
			return []int{first, last}
		}
		if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}

	}
	return []int{-1, -1}
}

func find(nums []int, i int, target int, next func(int) int) int {
	pre := -1
	for i >= 0 && i < len(nums) && nums[i] == target {
		pre = i
		i = next(i)
	}
	return pre
}

// @lc code=end
