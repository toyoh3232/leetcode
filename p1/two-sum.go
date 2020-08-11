package p1

import "fmt"

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// Test ...
func Test() {
	a := []int{2, 7, 11, 15}
	fmt.Println(twoSum(a, 9))
}

// @lc code=start
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// @lc code=end
