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

	m, n := ri(), ri()

	k := ri()

	a := make([]int, k+1)
	b := make([]int, k+1)
	c := make([]int, k+1)

	for i := 1; i <= k; i++ {
		a[i], b[i], c[i] = ri(), ri(), ri() // 氧气 氮气 重量
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = 1<<31 - 1
		}
	}

	dp[0][0] = 0
	for i := 1; i <= k; i++ { // k个物品
		for j := m; j >= 0; j-- {
			for p := n; p >= 0; p-- {
				dp[j][p] = min_i(dp[j][p], dp[max_i(0, j-a[i])][max_i(0, p-b[i])]+c[i])
			}
		}
	}

	// 当氧气和氮气容量分别为m n时，所达到的重量最小值
	fmt.Println(dp[m][n]) // 最后输出的是这个
}

func max_i(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min_i(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	_solve()

}
