package p12

// import "fmt"

/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */
// func main() {
// 	fmt.Println(intToRoman(11))
// }

// @lc code=start
func intToRoman(num int) string {
	var res []byte
	for i := num / 1000; i > 0; i-- {
		res = append(res, 'M')
	}
	num = num - (num / 1000 * 1000)
	for i := num / 500; i > 0; i-- {
		res = append(res, 'D')
	}
	num = num - (num / 500 * 500)
	if num/100 == 4 {
		if len(res) > 0 && res[len(res)-1] == 'D' {
			res[len(res)-1] = 'C'
			res = append(res, 'M')
		} else {
			res = append(res, "CD"...)
		}
	} else {
		for i := num / 100; i > 0; i-- {
			res = append(res, 'C')
		}
	}
	num = num - (num / 100 * 100)
	for i := num / 50; i > 0; i-- {
		res = append(res, 'L')
	}
	num = num - (num / 50 * 50)
	if num/10 == 4 {
		if len(res) > 0 && res[len(res)-1] == 'L' {
			res[len(res)-1] = 'X'
			res = append(res, 'C')
		} else {
			res = append(res, "XL"...)
		}
	} else {
		for i := num / 10; i > 0; i-- {
			res = append(res, 'X')
		}
	}

	num = num - (num / 10 * 10)
	for i := num / 5; i > 0; i-- {
		res = append(res, 'V')
	}
	num = num - (num / 5 * 5)
	if num == 4 {
		if len(res) > 0 && res[len(res)-1] == 'V' {
			res[len(res)-1] = 'I'
			res = append(res, 'X')
		} else {
			res = append(res, "IV"...)
		}
	} else {
		for i := num; i > 0; i-- {
			res = append(res, 'I')
		}
	}
	return string(res)
}

// @lc code=end
