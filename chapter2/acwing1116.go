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

	// acwing 1116

	var t int
	fmt.Fscan(in, &t)

	var n, m, x, y int
	for ; t > 0; t-- {
		fmt.Fscan(in, &n, &m, &x, &y)
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}

		ans := 0

		var dfs func(int, int, int)
		dfs = func(i, j int, step int) {
			if step == n*m {
				ans++
				return
			}

			for _, d := range directions {
				di := i + d[0]
				dj := j + d[1]
				if isok(di, dj, n, m) && !vis[di][dj] {
					vis[di][dj] = true
					dfs(di, dj, step+1)
					vis[di][dj] = false
				}
			}
		}

		vis[x][y] = true
		dfs(x, y, 1)

		fmt.Fprintln(out, ans)
	}
}

// 马走的规则 日字形状
var directions = [][]int{
	{-2, -1}, // 左上两个
	{-1, -2},
	{2, -1}, // 左下两个
	{1, -2},
	{-1, 2}, // 右上两个
	{-2, 1},
	{2, 1}, // 右下两个
	{1, 2},
}

func isok(x, y int, n, m int) bool {
	return x >= 0 && y >= 0 && x < n && y < m
}

func main() {
	_debug()
}
