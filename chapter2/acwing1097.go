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
	_ = []interface{}{ri}

	n, m := ri(), ri()

	grid := make([][]byte, n)
	for i := range grid {
		grid[i] = make([]byte, m)

		b := rc()
		for ; b == '\n'; b = rc() {
		}

		for j := range grid[i] {
			grid[i][j] = b
			b = rc()
		}
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if !isok(grid, i, j) {
			return
		}
		if grid[i][j] != 'W' {
			return
		}
		grid[i][j] = '=' // mark visited

		for _, d := range directions {
			di := i + d[0]
			dj := j + d[1]

			dfs(di, dj)
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 'W' {
				dfs(i, j)
				count++
			}
		}
	}
	fmt.Fprintln(out, count)
}

var directions = [][]int{
	{1, 0}, // 上下左右
	{-1, 0},
	{0, 1},
	{0, -1},

	{-1, -1}, // 斜四方，按需注释
	{-1, 1},
	{1, -1},
	{1, 1},
}

func isok(grid [][]byte, i, j int) bool { // 判断是否在二维数组越界
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

func main() {
	_solve(os.Stdin, os.Stdout)
}
