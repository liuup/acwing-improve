// local runtime version go1.20.2
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func _solve(_r io.Reader, _w io.Writer) {
	_in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const eof = 0
	_i, _n, buf := 0, 0, make([]byte, 1<<12) // 4KB

	rc := func() byte { // 读一个字符
		if _i == _n {
			_n, _ = _in.Read(buf)
			if _n == 0 { // EOF
				return eof
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}

	ri := func() (x int) { // 读一个整数，支持负数
		neg := false
		b := rc()
		for ; '0' > b || b > '9'; b = rc() {
			// 某些多组数据的题目，不告诉有多少组数据，那么需要额外判断是否读到了 EOF
			if b == eof {
				return
			}
			if b == '-' {
				neg = true
			}
		}
		for ; '0' <= b && b <= '9'; b = rc() {
			x = x*10 + int(b&15)
		}
		if neg {
			return -x
		}
		return
	}

	n := ri()
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
		for j := range grid[i] {
			grid[i][j] = ri()
		}
	}
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	type node struct{ i, j int }
	pre := make([][]node, n)
	for i := range pre {
		pre[i] = make([]node, n)
	}

	q := []node{}
	q = append(q, node{0, 0})
	vis[0][0] = true

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		if cur.i == n-1 && cur.j == n-1 {
			break
		}

		for _, d := range directions {
			di := cur.i + d[0]
			dj := cur.j + d[1]

			if !isok(grid, di, dj) || vis[di][dj] || grid[di][dj] == 1 {
				continue
			}
			vis[di][dj] = true
			pre[di][dj] = node{cur.i, cur.j}
			q = append(q, node{di, dj})
		}
	}

	var showpath func(i, j int)
	showpath = func(i, j int) {
		if i == 0 && j == 0 { // 终止条件
			fmt.Fprintln(out, i, j)
			return
		}
		showpath(pre[i][j].i, pre[i][j].j)
		fmt.Fprintln(out, i, j)
	}

	showpath(n-1, n-1)
}

var directions = [][]int{
	{1, 0}, {-1, 0}, {0, 1}, {0, -1}, // 上下左右
	// {-1, -1}, {-1, 1}, {1, -1}, {1, 1}, // 斜四方
}

func isok(grid [][]int, i, j int) bool { // 判断是否在二维数组越界
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

func main() {
	_solve(os.Stdin, os.Stdout)
}
