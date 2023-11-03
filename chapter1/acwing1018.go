package main

import (
	"bufio"
	"fmt"
	"os"
)

func _debug() {
	in := bufio.NewScanner(os.Stdin)
	_ = []interface{}{in}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const eof = 0
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

	n := ri()
	nums := make([][]int, n)
	for i := range nums {
		nums[i] = make([]int, n)
		for j := range nums[i] {
			nums[i][j] = ri()
		}
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = 1<<31 - 1
		}
	}

	dp[0][0] = nums[0][0]

	// 处理第一行
	for j := 1; j < n; j++ {
		dp[0][j] = min_i(dp[0][j], dp[0][j-1]+nums[0][j])
	}
	// 第一列
	for i := 1; i < n; i++ {
		dp[i][0] = min_i(dp[i][0], dp[i-1][0]+nums[i][0])
	}

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min_i(dp[i-1][j], dp[i][j-1]) + nums[i][j]
		}
	}
	fmt.Fprintln(out, dp[n-1][n-1])
}

func min_i(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	_debug()
}
