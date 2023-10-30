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

	// acwing 1018
	n := ri()

	nums := make([][]int, n+1)
	for i := range nums {
		nums[i] = make([]int, n+1)
	}

	for {
		a, b, c := ri(), ri(), ri()
		if a == 0 && b == 0 && c == 0 {
			break
		}
		nums[a][b] += c
	}

	dp := make([][][]int, 2*n+10)
	for i := range dp {
		dp[i] = make([][]int, n+10)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+10)
		}
	}

	for k := 2; k <= 2*n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if k-i <= 0 || k-i > n || k-j <= 0 || k-j > n {
					continue
				}

				dp[k][i][j] = get(dp, i, j, k) + nums[i][k-i]
				if i != j {
					dp[k][i][j] += nums[j][k-j]
				}
			}
		}
	}

	fmt.Fprintln(out, dp[2*n][n][n])
}

func get(dp [][][]int, i, j, k int) int {
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
