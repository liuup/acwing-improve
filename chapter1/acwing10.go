// local runtime version go1.20.2
package main

import (
	"bufio"
	"fmt"
	"os"
)

func _solve() {
	const eof = 0
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	_i, _n, buf := 0, 0, make([]byte, 1<<12) // 4KB

	rc := func() byte { // 读一个字符
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
	_ = []interface{}{ri}

	n, m := ri(), ri()

	vs := make([]int, n+1)
	ws := make([]int, n+1)

	root := 0
	g := make([][]int, n+1)

	for i := 1; i <= n; i++ {
		vs[i], ws[i] = ri(), ri()
		fa := ri()

		if fa == -1 {
			root = i
		} else {
			g[fa] = append(g[fa], i)
		}
	}

	// 表示的是以x为子树的物品,在容量不超过v时所获得的最大价值
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	var dfs func(x int)
	dfs = func(x int) {
		// 点x必须选 所以初始化dp[x][v[x]~m] = w[x]
		for i := vs[x]; i <= m; i++ {
			dp[x][i] = ws[x]
		}

		for i := 0; i < len(g[x]); i++ { // 遍历所有儿子
			y := g[x][i] // 儿子
			dfs(y)

			// j的范围为v[x]~m,小于v[x]无法选择以x为子树
			for j := m; j >= vs[x]; j-- {
				// 分给子树y的空间不能大于j-vs[x] 不然都无法选择根物品x
				for k := 0; k <= j-vs[x]; k++ {
					dp[x][j] = max_i(dp[x][j], dp[x][j-k]+dp[y][k])
				}
			}
		}
	}

	dfs(root)

	fmt.Fprintln(out, dp[root][m])
}

func max_i(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	_solve()
}
