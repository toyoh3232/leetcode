package p53

import "fmt"

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// Test ...
func Test() int {
	return maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
}

// @lc code=start
func maxSubArray(nums []int) int {
	m := max(nums)
	shadow := make([]int, len(nums))
	for j := 0; j < len(nums); j++ {
		for i := range shadow {
			if i+j < len(nums) {
				shadow[i] += nums[i+j]
			}
		}
		if m2 := max(shadow); m2 > m {
			m = m2
		}
	}
	fmt.Println(shadow)
	return m
}

func max(nums []int) int {
	m := nums[0]
	for _, v := range nums {
		if v > m {
			m = v
		}
	}
	return m
}

// @lc code=end
