package p22

/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	var out []string
	i, j := 1, n-1
	for j > 0 {
		out = append(out, permu(generateParenthesis(i), generateParenthesis(j))...)
		i++
		j--
	}
	out = append(out, maps(generateParenthesis(n-1), func(s string) string {
		return "(" + s + ")"
	})...)
	return nub(out)
}

func maps(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func selects(vs []string, f func(string) bool) []string {
	var vsm []string
	for _, v := range vs {
		if f(v) {
			vsm = append(vsm, v)
		}
	}
	return vsm
}

func nub(xs []string) []string {
	ys := make([]string, len(xs))
	copy(ys, xs)
	for _, x := range xs {
		ys = selects(ys, func(s string) bool {
			return x != s
		})
		ys = append(ys, x)
	}
	return ys
}

func permu(xs []string, ys []string) []string {
	var out []string
	for _, x := range xs {
		for _, y := range ys {
			out = append(out, x+y)
		}
	}
	return out
}

// @lc code=end
// func main() {
// 	fmt.Println(generateParenthesis(8))
// }
