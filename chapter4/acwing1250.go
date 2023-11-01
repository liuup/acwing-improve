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

	var n, m int
	fmt.Fscan(in, &n, &m)

	us := InitUnion(n*n + 1)

	var a, b int
	var c string

	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a, &b, &c)

		x := (a-1)*n + b
		y := 0

		if c == "D" {
			y = a*n + b // 向下
		} else {
			y = (a-1)*n + b + 1 // 向右
		}

		if us.IsConnected(x, y) {
			fmt.Fprintln(out, i+1)
			return
		}
		us.Union(x, y)
	}

	fmt.Fprintln(out, "draw")
}

// union set 并查集
type unionset struct {
	father []int
	count  int // 连通分量的个数
}

func InitUnion(n int) unionset {
	fa := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fa[i] = i

	}
	return unionset{fa, n}
}

func (u *unionset) Find(i int) int {
	if u.father[i] == i {
		return i
	}
	u.father[i] = u.Find(u.father[i]) // 路径压缩
	return u.father[i]
}

func (u *unionset) Union(i, j int) {
	i_fa, j_fa := u.Find(i), u.Find(j)
	if i_fa == j_fa {
		return
	}
	u.father[i_fa] = j_fa
	u.count--
}

func (u *unionset) IsConnected(x, y int) bool {
	return u.Find(x) == u.Find(y)
}

func main() {
	_debug()
}
