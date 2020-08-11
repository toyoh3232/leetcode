package p35

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// Test ...
func Test() int {
	a := []int{1, 3, 5, 6}
	target := 7
	return searchInsert(a, target)
}

// @lc code=start
func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2

		if nums[mid] == target {
			return mid
		}

		if i == mid-1 {
			if target < nums[i] {
				return i
			}
			if target > nums[mid] {
				return mid + 1
			}
			return mid
		}

		if j == mid+1 {
			if target < nums[mid] {
				return mid
			}
			if target > nums[j] {
				return j + 1
			}
			return mid
		}

		if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}

	}
	return -1
}

// @lc code=end
