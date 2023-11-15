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

	vs := make([]int, m+1)
	ws := make([]int, m+1)

	for i := 1; i <= m; i++ {
		vs[i], ws[i] = ri(), ri()
	}

	dp := make([]int64, n+1)

	for i := 1; i <= m; i++ {
		for j := n; j >= vs[i]; j-- {
			dp[j] = max_i64(dp[j], dp[j-vs[i]]+int64(vs[i]*ws[i]))
		}
	}
	fmt.Fprintln(out, dp[n])
}

func max_i64(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func main() {
	_solve()
}
