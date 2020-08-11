package p31

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */

// Test ...
func Test() {
	a := []int{3, 2, 1}
	nextPermutation(a)
	fmt.Println(a)
	nextPermutation(a)
	fmt.Println(a)
	nextPermutation(a)
	fmt.Println(a)
	nextPermutation(a)
	fmt.Println(a)
	nextPermutation(a)
	fmt.Println(a)
	nextPermutation(a)
	fmt.Println(a)
	nextPermutation(a)
	fmt.Println(a)
	nextPermutation(a)
	fmt.Println(a)
}

// @lc code=start
func nextPermutation(nums []int) {

	i := len(nums) - 1
	for i > 0 {
		if nums[i-1] < nums[i] {
			break
		}
		i--
	}
	if i == 0 {
		reverse(nums)
		return
	}
	sort.Ints(nums[i:])
	idx := findNextGreater(nums[i:], nums[i-1]) + i
	nums[i-1], nums[idx] = nums[idx], nums[i-1]
}

func findNextGreater(nums []int, x int) int {
	for i, v := range nums {
		if v > x {
			return i
		}
	}
	return -1
}

func reverse(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i, j = i+1, j-1
	}

}

// @lc code=end
