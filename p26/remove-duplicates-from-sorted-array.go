package p26

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] != nums[i] {
			nums[count] = nums[i+1]
			count++
		}
	}
	return count
}

func findNextIndex(nums []int, value int) int {
	for i, v := range nums {
		if v > value {
			return i
		}
	}
	return -1
}

// @lc code=end
