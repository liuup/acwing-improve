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

	// acwing 1019 庆功会

	n, m := ri(), ri()

	vs := make([]int, n+1)
	ws := make([]int, n+1)
	ss := make([]int, n+1)

	for i := 1; i <= n; i++ {
		vs[i], ws[i], ss[i] = ri(), ri(), ri()
	}

	dp := make([]int, m+1)

	for i := 1; i <= n; i++ {
		for j := m; j >= vs[i]; j-- {
			for k := 0; k <= ss[i] && k*vs[i] <= j; k++ { // 限制个数
				dp[j] = max_i(dp[j], dp[j-k*vs[i]]+k*ws[i])
			}
		}
	}
	fmt.Fprintln(out, dp[m])
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

func abs_i(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 创建指定维度的二维数组
// n: rows; m: cols
func make2dimen(n, m int) (ans [][]int) {
	ans = make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}
	return
}

func main() {
	_solve()

}
