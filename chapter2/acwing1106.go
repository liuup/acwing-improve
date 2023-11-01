package main

import (
	"bufio"
	"fmt"
	"os"
)

func _debug() {
	in := bufio.NewReader(os.Stdin)
	const eof = 0
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	_ = []interface{}{in}

	_i, _n, buf := 0, 0, make([]byte, 1<<12) // 4KB
	rc := func() byte {                      // 读一个字符
		if _i == _n {
			_n, _ = os.Stdin.Read(buf)
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
	_ = []interface{}{rc, ri}

	// acwing 1106

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

	peak := 0
	valley := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !vis[i][j] {
				has_highter := false
				has_lower := false

				// bfs
				q := []node{}
				q = append(q, node{i, j})

				for len(q) != 0 {
					cur := q[0]
					q = q[1:]

					for _, d := range directions {
						dx := cur.x + d[0]
						dy := cur.y + d[1]

						if !isok(grid, dx, dy) {
							continue
						}

						if grid[dx][dy] != grid[cur.x][cur.y] {
							if grid[dx][dy] > grid[cur.x][cur.y] {
								has_highter = true
							} else {
								has_lower = true
							}
						} else if !vis[dx][dy] {
							q = append(q, node{dx, dy})
							vis[dx][dy] = true
						}
					}
				}

				if !has_highter {
					peak++
				}
				if !has_lower {
					valley++
				}
			}
		}
	}
	fmt.Fprintln(out, peak, valley)
}

type node struct {
	x, y int
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

func isok(grid [][]int, i, j int) bool { // 判断是否在二维数组越界
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

func main() {
	_debug()
}
