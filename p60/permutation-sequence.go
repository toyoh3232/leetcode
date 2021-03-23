package p60

/*
 * @lc app=leetcode id=60 lang=golang
 *
 * [60] Permutation Sequence
 */

// Test ...
func Test() string {
	return getPermutation(1, 1)
}

// @lc code=start
func getPermutation(n int, k int) string {
	a := make([]int, n)
	for i := range a {
		a[i] = i + 1
	}
	r := selects(k-1, a)
	b := make([]byte, n)
	for i, v := range r {
		b[i] = byte(v) + 48
	}
	return string(b)
}

func selects(k int, arr []int) []int {
	length := len(arr)
	if length == 1 {
		return arr
	}
	i := k / fac(length-1)
	nk := k % fac(length-1)
	f := arr[i]
	na := append(arr[0:i], arr[i+1:]...)
	return append([]int{f}, selects(nk, na)...)
}

func fac(k int) int {
	if k == 0 {
		return 1
	}
	return k * fac(k-1)
}

// @lc code=end
