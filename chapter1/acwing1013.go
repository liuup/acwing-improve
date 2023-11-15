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

	nums := make([][]int, n+1)
	for i := range nums {
		nums[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			nums[i][j] = ri()
		}
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// dp
	for i := 1; i <= n; i++ { // n组
		for j := 1; j <= m; j++ {
			for k := 0; k <= j; k++ {
				dp[i][j] = max_i(dp[i][j], dp[i-1][j-k]+nums[i][k])
			}
		}
	}

	// dfs 最短路
	path := []int{}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i == 0 {
			return
		}

		// 寻找当前状态dp[i][j]是从哪一个状态转移过来的
		for a := 0; a <= j; a++ {
			if dp[i-1][j-a]+nums[i][a] == dp[i][j] {
				path = append(path, a)
				dfs(i-1, j-a)
				return
			}
		}
	}

	dfs(n, m)

	fmt.Fprintln(out, dp[n][m])

	// 逆序输出path
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Fprintln(out, len(path)-i, path[i])
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
