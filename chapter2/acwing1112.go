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

	// acwing 1112 迷宫
	var k int
	fmt.Fscan(in, &k)

	var s string
	var n int
	for ; k > 0; k-- {
		fmt.Fscan(in, &n)
		grid := make([][]byte, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &s)
			grid[i] = []byte(s)
		}

		var x1, y1, x2, y2 int
		fmt.Fscan(in, &x1, &y1, &x2, &y2)

		if grid[x1][y1] == '#' || grid[x2][y2] == '#' {
			fmt.Fprintln(out, "NO")
			continue
		}
		if x1 == x2 && y1 == y2 {
			fmt.Fprintln(out, "YES")
			continue
		}

		if dfs(grid, x1, y1, x2, y2) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func dfs(grid [][]byte, x1, y1, x2, y2 int) bool {
	if !isok(grid, x1, y1) {
		return false
	}

	if grid[x1][y1] != '.' {
		return false
	}

	if x1 == x2 && y1 == y2 {
		return true
	}

	grid[x1][y1] = '=' // 标记已访问

	for _, d := range directions {
		dx := x1 + d[0]
		dy := y1 + d[1]
		if dfs(grid, dx, dy, x2, y2) {
			return true
		}
	}
	return false
}

var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} // 四个遍历方向
func isok(grid [][]byte, i, j int) bool { // 判断是否在二维数组越界
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

func main() {
	_debug()
}
