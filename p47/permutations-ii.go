package p47

/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// Test ...
func Test() [][]int {
	a := []int{1, 1, 2}
	return permuteUnique(a)
}

// @lc code=start
func permuteUnique(nums []int) [][]int {

	return rub(permute(nums))
}

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

func rub(nums [][]int) [][]int {
	r := [][]int{}
	for i := range nums {
		if !has(r, nums[i]) {
			r = append(r, nums[i])
		}
	}
	return r

}

func has(nums [][]int, num []int) bool {
	if len(nums) == 0 {
		return false
	}
	for _, v := range nums {
		i := 0
		for ; i < len(num); i++ {
			if num[i] != v[i] {
				break
			}
		}
		if i == len(num) {
			return true
		}
	}
	return false
}

// @lc code=end
