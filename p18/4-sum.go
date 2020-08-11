package p18

import "sort"

/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if i+2 < length && j+1 < length {
				m, n := twoSum(nums, target, i, j, j+1, length-1)
				for m != n {
					r := []int{nums[i], nums[j], nums[m], nums[n]}
					if !has(res, r) {
						res = append(res, r)
					}
					m, n = twoSum(nums, target, i, j, m+1, n-1)
				}
			}
		}
	}
	return res
}

func twoSum(nums []int, target, i, j, m, n int) (int, int) {

	for m < n && n < len(nums) {

		if nums[i]+nums[j]+nums[m]+nums[n] == target {
			return m, n
		}
		if nums[i]+nums[j]+nums[m]+nums[n] > target {
			n--
			continue
		}
		if nums[i]+nums[j]+nums[m]+nums[n] < target {
			m++
			continue
		}
	}
	return m, m
}

func has(nums [][]int, num []int) bool {
	for _, v := range nums {
		if equal(v, num) {
			return true
		}
	}
	return false
}

func equal(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i, v := range nums1 {
		if v != nums2[i] {
			return false
		}
	}
	return true
}

// @lc code=end
