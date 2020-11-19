package p41

/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 */

// Test ...
func Test() int {
	nums := []int{0, 2, 3}
	return firstMissingPositive(nums)
}

// @lc code=start
func firstMissingPositive(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return len(nums) + 1
}

// @lc code=end
