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

	m, n := ri(), ri()

	nums := make([][]int, m+1)
	for i := range nums {
		nums[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			nums[i][j] = ri()
		}
	}

	dp := make([][][]int, n+m+10)
	for i := range dp {
		dp[i] = make([][]int, m+10)
		for j := range dp[i] {
			dp[i][j] = make([]int, m+10)
		}
	}

	for k := 2; k <= (n + m); k++ {
		for i := 1; i <= m; i++ { // i1
			for j := 1; j <= m; j++ { // i2
				// 越界
				if k-i <= 0 || k-j <= 0 || k-i > n || k-j > n {
					continue
				}

				if i != j {
					dp[k][i][j] = get(dp, k, i, j) + nums[i][k-i] + nums[j][k-j]
				} else {
					dp[k][i][j] = get(dp, k, i, j) + nums[i][k-i]
				}
			}
		}
	}
	fmt.Fprintln(out, dp[n+m][m][m])
}

func get(dp [][][]int, k, i, j int) int {
	return max_i(max_i(dp[k-1][i-1][j-1], dp[k-1][i][j-1]), max_i(dp[k-1][i-1][j], dp[k-1][i][j]))
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
