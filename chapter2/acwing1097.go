package main

import (
	"bufio"
	"fmt"
	"os"
)

func _debug() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	_ = []interface{}{in}

	// acwing 1097
	var n, m int
	fmt.Fscan(in, &n, &m)

	var s string
	grid := make([][]byte, n)
	for i := range grid {
		fmt.Fscan(in, &s)
		grid[i] = []byte(s)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if !isok(grid, i, j) {
			return
		}

		if grid[i][j] != 'W' {
			return
		}

		grid[i][j] = '+' // mark

		for _, d := range directions {
			di := i + d[0]
			dj := j + d[1]
			dfs(di, dj)
		}
	}

	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 'W' {
				dfs(i, j)
				cnt++
			}
		}
	}
	fmt.Fprintln(out, cnt)
}

var directions = [][]int{
	{1, 0}, // 上下左右
	{-1, 0},
	{0, 1},
	{0, -1},

	{-1, -1}, // 斜四方
	{-1, 1},
	{1, -1},
	{1, 1},
}

func isok(grid [][]byte, i, j int) bool { // 判断是否在二维数组越界
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

func main() {
	_debug()
}
