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

	// acwing 1113

	var w, h int
	for {
		fmt.Fscan(in, &w, &h)
		if w == 0 && h == 0 {
			break
		}

		grid := make([][]byte, h)
		for i := 0; i < h; i++ {
			grid[i] = make([]byte, w)
		}

		x, y := 0, 0
		find := false
		var s string
		for i := 0; i < h; i++ {
			fmt.Fscan(in, &s)
			if !find {
				for j := 0; j < len(s); j++ {
					if s[j] == '@' {
						x, y = i, j
						find = true
						break
					}
				}
			}
			grid[i] = []byte(s)
		}

		cnt := 0
		var dfs func(int, int)
		dfs = func(i, j int) {
			if !isok(grid, i, j) {
				return
			}
			if grid[i][j] != '.' {
				return
			}

			cnt++
			grid[i][j] = '+' // mark

			for _, d := range directions {
				di := i + d[0]
				dj := j + d[1]

				dfs(di, dj)
			}
		}

		grid[x][y] = '.'
		dfs(x, y)

		fmt.Fprintln(out, cnt)
	}
}

var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} // 四个遍历方向
func isok(grid [][]byte, i, j int) bool { // 判断是否在二维数组越界
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

func main() {
	_debug()
}
