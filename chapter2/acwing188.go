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

	// acwing 188

	var m, n int
	fmt.Fscan(in, &m, &n)

	var s string
	x, y := 0, 0 // 牛的坐标
	find := false

	grid := make([][]byte, n)
	for i := range grid {
		fmt.Fscan(in, &s)
		grid[i] = []byte(s)
		for j := range grid[i] {
			if grid[i][j] == 'K' && !find {
				x, y = i, j
				find = true
				break
			}
		}
	}

	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}

	q := []node{}
	q = append(q, node{x, y})
	vis[x][y] = true

	step := 0
	for len(q) != 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]

			if grid[cur.x][cur.y] == 'H' {
				fmt.Fprintln(out, step)
				return
			}

			for _, d := range directions {
				dx := cur.x + d[0]
				dy := cur.y + d[1]

				if !isok(grid, dx, dy) || vis[dx][dy] || grid[dx][dy] == '*' {
					continue
				}

				q = append(q, node{dx, dy})
				vis[dx][dy] = true
			}
		}
		step++
	}
}

type node struct {
	x, y int
}

var directions = [][]int{
	{-1, -2},
	{-2, -1},
	{1, -2},
	{2, -1},
	{-2, 1},
	{-1, 2},
	{2, 1},
	{1, 2},
}

func isok(grid [][]byte, i, j int) bool { // 判断是否在二维数组越界
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

func main() {
	_debug()
}
