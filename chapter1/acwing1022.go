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

	n, m, k := ri(), ri(), ri()

	ws := make([]int, k)
	vs := make([]int, k)
	for i := range vs {
		ws[i] = ri()
		vs[i] = ri()
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = -1 << 31
		}
	}

	dp[0][0] = 0
	for i := 1; i <= k; i++ {
		for j := n; j >= ws[i-1]; j-- {
			for k := m; k >= vs[i-1]; k-- {
				dp[j][k] = max_i(dp[j][k], dp[j-ws[i-1]][k-vs[i-1]]+1)
			}
		}
	}

	ans := -1
	t := 0
	for j := 0; j <= n; j++ {
		for k := 0; k < m; k++ {
			if dp[j][k] > ans || (ans == dp[j][k] && k < t) {
				ans = dp[j][k]
				t = k
			}
		}
	}

	fmt.Fprintln(out, ans, m-t)
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
