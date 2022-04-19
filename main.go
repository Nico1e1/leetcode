package main

import (
	"fmt"
	"strconv"
)

type Reader interface {
	Read()
}
type Reader2 interface {
	Read()
}
type ReadWriter interface {
	Reader
	Write(str string)
}
type S struct {
	data string
}

func (s S) Read() {
	fmt.Println(s.data)
}

func (s S) setVal(str string) {
	s.data = str
	fmt.Println(s.data)
}

func (s *S) Write(str string) {
	s.data = str
}

type Ae struct {
	x int
}

func (a Ae) MethodA() {

}

type B struct {
	*Ae
}

type C struct {
	B
}

type P = *int

func main() {
	s1 := S{"a"}
	s2 := s1
	s2.data = "b"
	s1.Read()
	s2.Read()
}

func digitSum(s string, k int) string {
	var add = func(str string) string {
		num := 0
		for i := 0; i < len(str); i++ {
			num += int(str[i] - '0')
		}
		return strconv.Itoa(num)
	}
	for len(s) > k {
		tmp := ""
		for i := 0; i < len(s); {
			j := i + k
			if j > len(s) {
				j = len(s)
			}
			tmp += add(s[i:j])
			i = j
		}
		s = tmp
	}
	return s
}

func minimumRounds(tasks []int) int {
	count := make(map[int]int)
	for _, task := range tasks {
		count[task]++
	}
	res := 0
	for _, val := range count {
		if val == 1 {
			return -1
		}
		if val%3 == 1 {
			res += val/3 + 1
		} else {
			res += val / 3
			val %= 3
			res += val / 2
		}
	}
	return res
}

func maxTrailingZeros(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	res := 0
	count := func(num int) int {
		zero := 0
		str := strconv.Itoa(num)
		for i := len(str) - 1; i >= 0; i-- {
			if str[i] != 0 {
				break
			}
			zero++
		}
		return zero
	}
	var dfs func(i, j int, mul int, horizontal bool)
	dfs = func(i, j, mul int, horizontal bool) {
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] {
			return
		}
		mul *= grid[i][j]
		zero := count(mul)
		if zero > res {
			res = zero
		}
		visited[i][j] = true
		if horizontal {
			dfs(i, j+1, mul, true)
			dfs(i, j-1, mul, true)
		}
		dfs(i+1, j, mul, false)
		dfs(i-1, j, mul, false)
		visited[i][j] = false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, 1, true)
		}
	}
	return res
}
