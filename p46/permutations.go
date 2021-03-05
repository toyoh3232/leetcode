package p46

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// Test ...
func Test() [][]int {
	a := []int{1, 2, 3}
	return permute(a)
}

// @lc code=start
func permute(nums []int) [][]int {
	length := len(nums)
	if length == 1 {
		return [][]int{nums}
	}
	r := [][]int{}
	for i := range nums {
		head := []int{nums[i]}
		tail := []int{}
		tail = append(tail, nums[0:i]...)
		tail = append(tail, nums[i+1:]...)
		for _, v := range permute(tail) {
			r = append(r, append(head, v...))
		}
	}
	return r
}

// @lc code=end
