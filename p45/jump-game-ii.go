package p45

/*
 *@lc app=leetcode id=45 lang=golang
 *
 * 45] Jump Game II
 */

// Test ...
func Test() int {
	a := []int{2, 3, 0, 1, 4}
	return jump(a)
}

// @lc code=start

func jump(nums []int) int {
	steps := make([]int, len(nums))
	steps[0] = 0
	for i := 1; i < len(steps); i++ {
		substeps := []int{}
		for j := i - 1; j >= 0; j-- {
			if i-j <= nums[j] {
				substeps = append(substeps, steps[j])
			}
		}
		steps[i] = 1 + min(substeps)
	}
	return steps[len(nums)-1]
}
func min(nums []int) int {
	v := nums[0]
	for i := range nums {
		if nums[i] < v {
			v = nums[i]
		}
	}
	return v
}

// @lc code=end
