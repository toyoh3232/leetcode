package p33

/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

//Test ...
func Test() int {
	nums, target := []int{4, 5, 6, 7, 0, 1, 2}, 3
	return search(nums, target)
}

// @lc code=start
func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[i] < nums[mid] {
			switch {
			case i <= mid-1 && nums[i] <= target && target <= nums[mid-1]:
				j = mid - 1
				continue
			default:
				i = mid + 1
				continue
			}
		} else {
			switch {
			case mid+1 <= j && nums[mid+1] <= target && target <= nums[j]:
				i = mid + 1
				continue
			default:
				j = mid - 1
				continue
			}
		}
	}
	return -1
}

// @lc code=end
