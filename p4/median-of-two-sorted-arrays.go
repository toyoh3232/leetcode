package p4

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	merged := make([]int, length)
	var v int
	for i := 0; i <= length/2; i++ {
		v, nums1, nums2 = merge(nums1, nums2)
		merged[i] = v

	}
	if length%2 == 0 {
		return float64(merged[length/2]+merged[length/2-1]) / 2.0
	}
	return float64(merged[length/2])
}

func merge(nums1 []int, nums2 []int) (int, []int, []int) {
	if len(nums1) > 0 && len(nums2) == 0 {
		return nums1[0], nums1[1:], nil
	}
	if len(nums1) == 0 && len(nums2) > 0 {
		return nums2[0], nil, nums2[1:]
	}
	if nums1[0] < nums2[0] {
		return nums1[0], nums1[1:], nums2
	}
	return nums2[0], nums1, nums2[1:]

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2, m, n = nums2, nums1, n, m
	}
	imin, imax, half := 0, m, (m+n+1)/2
	var leftMax, rightMin int
	for imin <= imax {
		i := (imin + imax) / 2
		j := half - i
		if i < m && nums2[j-1] > nums1[i] {
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			imax = i - 1
		} else {
			if i == 0 {
				leftMax = nums2[j-1]
			} else if j == 0 {
				leftMax = nums1[i-1]
			} else {
				leftMax = max(nums1[i-1], nums2[j-1])
			}
			if (m+n)%2 == 1 {
				return float64(leftMax)
			}
			if i == m {
				rightMin = nums2[j]
			} else if j == n {
				rightMin = nums1[i]
			} else {
				rightMin = min(nums1[i], nums2[j])
			}
			break
		}
	}
	return float64(leftMax+rightMin) / 2.0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
