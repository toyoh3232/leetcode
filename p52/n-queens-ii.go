package p52

import "sync"

/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 */

// Test ...
func Test() int {
	return totalNQueens(14)
}

// @lc code=start
func totalNQueens(n int) int {
	s := make([][]bool, n)
	var number int
	for r := range s {
		s[r] = make([]bool, n)
	}
	done, ans := make(chan bool), make(chan [][]bool)
	go do(done, ans, s)
	for {
		select {
		case <-done:
			return number
		case <-ans:
			number++
		}
	}
}

func do(done chan bool, ans chan [][]bool, start [][]bool) {
	var wg sync.WaitGroup
	wg.Add(1)
	go iterate(start, 0, len(start), &wg, ans)
	wg.Wait()
	done <- true
}

func iterate(s [][]bool, r int, n int, wg *sync.WaitGroup, ans chan [][]bool) {
	defer wg.Done()
	if r == n {
		ans <- s
		return
	}
	for c := range s[r] {
		on(s[r], c)
		if check(s, r) {
			ns := cp(s)
			wg.Add(1)
			go iterate(ns, r+1, n, wg, ans)
		}
	}
}

func cp(s [][]bool) [][]bool {
	ns := make([][]bool, len(s))
	for r := range ns {
		ns[r] = make([]bool, len(s[r]))
		copy(ns[r], s[r])
	}
	return ns
}

func on(r []bool, c int) {
	for cc := range r {
		if cc == c {
			r[cc] = true
		} else {
			r[cc] = false
		}
	}
}

func check(s [][]bool, r int) bool {
	return horizon(s, r) && vertical(s, r) && diagonal(s, r)
}

func horizon(s [][]bool, r int) bool {
	c := find(s[r])
	for i := 0; i < len(s); i++ {
		if i != c && s[r][i] {
			return false
		}
	}
	return true
}

func vertical(s [][]bool, r int) bool {
	c := find(s[r])
	for i := 0; i < len(s[0]); i++ {
		if i != r && s[i][c] {
			return false
		}
	}
	return true
}

func diagonal(s [][]bool, r int) bool {
	c := find(s[r])
	n := len(s)
	for i, j := r-1, c-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if s[i][j] {
			return false
		}
	}
	for i, j := r+1, c+1; i < n && j < n; i, j = i+1, j+1 {
		if s[i][j] {
			return false
		}
	}
	for i, j := r+1, c-1; i < n && j >= 0; i, j = i+1, j-1 {
		if s[i][j] {
			return false
		}
	}
	for i, j := r-1, c+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if s[i][j] {
			return false
		}
	}
	return true
}

func find(row []bool) int {
	for i := range row {
		if row[i] {
			return i
		}
	}
	return -1
}

// @lc code=end
