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
	_ = []interface{}{rc, ri}

	// acwing 1252 并查集+01背包
	n, m, w := ri(), ri(), ri()

	vs := make([]int, n)
	ws := make([]int, n)
	for i := 0; i < n; i++ {
		vs[i] = ri()
		ws[i] = ri()
	}

	us := InitUnion(n + 1)

	for i := 0; i < m; i++ {
		a, b := ri(), ri()
		us.Union(a-1, b-1)
	}

	vs_mp := map[int]int{}
	ws_mp := map[int]int{}
	for i := 0; i < n; i++ {
		vs_mp[us.Find(i)] += vs[i]
		ws_mp[us.Find(i)] += ws[i]
	}

	// 重置为0
	vs = []int{}
	ws = []int{}
	for k, v := range vs_mp {
		vs = append(vs, v)
		ws = append(ws, ws_mp[k])
	}

	// 然后接下来就是01背包的模板了
	t := len(ws)

	dp := make([]int, w+1)
	for i := 1; i <= t; i++ {
		for j := w; j >= vs[i-1]; j-- {
			dp[j] = max_i(dp[j], dp[j-vs[i-1]]+ws[i-1])
		}
	}
	fmt.Println(dp[w])
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

func max_i(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	_debug()
}
