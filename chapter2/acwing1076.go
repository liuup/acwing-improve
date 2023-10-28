package main

import (
	"bufio"
	"fmt"
	"os"
)

func _debug() {
	const eof = 0
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
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
	rs := func() (s []byte) {
		b := rc()
		for ; '0' > b || b > '9'; b = rc() { // 'A' 'Z'
		}
		// for ; 'a' <= b && b <= 'z'; b = rc() { // 'A' 'Z'
		// 	s = append(s, b)
		// }
		for ; '0' <= b && b <= '9'; b = rc() { // 'A' 'Z'
			s = append(s, b)
		}
		return
	}
	_ = []interface{}{rc, ri, rs}

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

		if cur.x == n-1 && cur.y == n-1 { // final
			break
		}

		for _, d := range directions {
			dx := cur.x + d[0]
			dy := cur.y + d[1]

			if !isok(grid, dx, dy) || grid[dx][dy] == 1 || vis[dx][dy] {
				continue
			}

			vis[dx][dy] = true
			pre[dx][dy] = node{cur.x, cur.y}
			q = append(q, node{dx, dy})
		}
	}

	showpath(pre, 0, 0, n-1, n-1, out)

}

func showpath(pre [][]node, x1, y1, x2, y2 int, out *bufio.Writer) {
	if x2 == x1 && y2 == y1 {
		fmt.Fprintln(out, x2, y2)
		return
	}

	showpath(pre, x1, y1, pre[x2][y2].x, pre[x2][y2].y, out)
	fmt.Fprintln(out, x2, y2)
}

var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} // 四个遍历方向
func isok(grid [][]int, i, j int) bool { // 判断是否在二维数组越界
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

type node struct {
	x, y int
}

func main() {
	_debug()
}
