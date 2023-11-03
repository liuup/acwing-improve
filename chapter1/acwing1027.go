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

	nums := make([][]int, n+1)
	for i := range nums {
		nums[i] = make([]int, n+1)
	}

	for {
		a, b, c := ri(), ri(), ri()
		if a == 0 && b == 0 && c == 0 {
			break
		}
		nums[a][b] = c
	}

	dp := make([][][]int, 2*n+10)
	for i := range dp {
		dp[i] = make([][]int, n+10)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+10)
		}
	}

	// k == i+j
	for k := 2; k <= 2*n; k++ {
		for i := 1; i <= n; i++ { // i1
			for j := 1; j <= n; j++ { // i2
				if k-i <= 0 || k-i > n || k-j <= 0 || k-j > n {
					continue
				}

				if i != j { // 两条路线没有重合
					dp[k][i][j] = get(dp, k, i, j) + nums[j][k-j] + nums[i][k-i]
				} else {
					dp[k][i][j] = get(dp, k, i, j) + nums[j][k-j]
				}
			}
		}
	}

	fmt.Fprintln(out, dp[2*n][n][n])

	/*

		f[i][j]表示所有从1，1走到i,j的最大值
		f(i,j) = max(f[i-1, j] + w[i, j], f[i,j-1]+w[i,j])

		走两次：
		f(i1,j1,i2,j2)表示所有从(1,,1)(1,1)分别走到(i1,j1)(i2,j2)的路径最大值

		如何处理同一个格子不能被重复选择

		只有当i1+j1 == j2+j2 时，两条路径的格子才可能重合

		f[k, i1, i2]表示所有从(1, 1)(1, 1)分别走到(i1, k-i1)(i2, k-i2)的路径的最大值

		第一条 下；第二条 下	max(f(k-1, i1-1, i2-1))
		第一条 下；第二条 右	max(f(k-1, i1-1, i2))
		第一条 右；第二条 右	max(f(k-1, i1, i2))
		第一条 右；第二条 下	max(f(k-1, i1, i2-1))

	*/

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
	_debug()
}
