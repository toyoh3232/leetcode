package p29

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers

 */

// Test ...
func Test() {
	fmt.Println(divide(3, 4))
}

// @lc code=start
func divide(dividend int, divisor int) int {
	x, y := int64(dividend), int64(divisor)
	if y < 0 {
		x, y = -x, -y
	}
	var c int64
	c = 0
	if x > 0 && y > 0 {
		for x = x - y; x >= 0; {
			c++
			x = x - y
		}
	} else if x < 0 && y > 0 {
		for x = x + y; x <= 0; {
			c--
			x = x + y
		}
	}
	if c > math.MaxInt32 {
		c = math.MaxInt32
	}
	return int(c)
}

// @lc code=end
