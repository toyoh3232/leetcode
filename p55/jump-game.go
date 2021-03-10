package p55

/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// Test ...
func Test() bool {
	a := []int{1, 0, 1, 0}
	return canJump(a)
}

// @lc code=start
func canJump(nums []int) bool {
	items := make([]bool, len(nums))
	if len(nums) == 1 {
		return true
	}
	if nums[0] == 0 {
		return false
	}
	items[0] = true
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if i-j <= nums[j] && items[j] {
				items[i] = true
				break
			}
			items[i] = false
		}
	}

	return items[len(items)-1]
}

// @lc code=end
