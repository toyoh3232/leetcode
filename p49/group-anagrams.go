package p49

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

//Test ...
func Test() [][]string {
	s := []string{"a"}
	return groupAnagrams(s)
}

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _,str := range strs {
		key := hash(str)
		if _,ok := m[key];!ok {
			m[key] = []string{}
		}
		m[key] = append(m[key], str)
	}
	r := [][]string{};

	for _,v := range m {
		r = append(r, v)
	}
	return r
}

func hash(str string) string {
	nStr := make ([]byte, len(str))
	for i:=0;i<len(str);i++ {
		insert(nStr,str[i])
	}
	return string(nStr)
}

func insert(str []byte, b byte) {
	index := 0
	for i, v := range str {
		if v == 0 {
			index = i
			break
		}
		if v > b {
			index = i
			break
		}
	}
	if str[index] != 0 {
		for i := len(str) - 2; i >= index; i-- {
			str[i+1] = str[i]
		}
	}
	str[index] = b
}

// @lc code=end
