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

	n, v := ri(), ri()

	vs := make([]int, n+1)
	ws := make([]int, n+1)
	for i := 1; i <= n; i++ {
		vs[i], ws[i] = ri(), ri()
	}

	// 考虑前i个物品 体积不超过j的方案最大值
	dp := make([][]int, n+10)
	for i := range dp {
		dp[i] = make([]int, v+10)
	}

	// 从后往前看
	for i := n; i >= 1; i-- {
		for j := 0; j <= v; j++ {
			if j >= vs[i] {
				dp[i][j] = max_i(dp[i+1][j], dp[i+1][j-vs[i]]+ws[i])
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}

	j := v
	for i := 1; i <= n; i++ {
		if j >= vs[i] && dp[i][j] == dp[i+1][j-vs[i]]+ws[i] { // 可以选就必选
			fmt.Fprint(out, i, " ")
			j -= vs[i]
		}
	}
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
