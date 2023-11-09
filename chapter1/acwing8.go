\package main

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

	n, v, m := ri(), ri(), ri()

	vv := make([]int, n+1) // 体积
	mm := make([]int, n+1) // 重量
	ww := make([]int, n+1) // 价值
	for i := 1; i <= n; i++ {
		vv[i], mm[i], ww[i] = ri(), ri(), ri()
	}

	dp := make([][]int, v+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := v; j >= vv[i]; j-- {
			for k := m; k >= mm[i]; k-- {
				dp[j][k] = max_i(dp[j][k], dp[j-vv[i]][k-mm[i]]+ww[i])
			}
		}
	}

	fmt.Fprintln(out, dp[v][m])
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
