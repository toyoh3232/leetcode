package p27

// import "fmt"

/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// func main() {
// 	a := []int{0, 1, 2, 2, 3, 0, 4, 2}
// 	fmt.Println(removeElement(a, 2))
// 	fmt.Println(a)
// }

// @lc code=start
func removeElement(nums []int, val int) int {
	lt := len(nums)
	for removeOne(nums, lt, val) {
		lt--
	}
	return lt
}

func removeOne(nums []int, lt, val int) bool {
	for i := 0; i < lt; i++ {
		if nums[i] == val {
			for j := i; j < lt-1; j++ {
				nums[j] = nums[j+1]
			}
			return true
		}
	}
	return false
}

// @lc code=end
