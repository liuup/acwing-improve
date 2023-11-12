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

	// acwing 7 混合背包

	n, v := ri(), ri()

	vs := make([]int, 100010)
	ws := make([]int, 100010)

	cnt := 1
	for i := 1; i <= n; i++ {
		a, b, s := ri(), ri(), ri()

		k := 1
		if s < 0 {
			s = 1
		} else if s == 0 {
			s = v / a // 完全背包 最大容量为背包最大容量/物品体积 向下取整
		}

		for k <= s {
			vs[cnt] = a * k
			ws[cnt] = b * k
			s -= k
			k *= 2
			cnt++
		}

		if s > 0 {
			vs[cnt] = s * a
			ws[cnt] = s * b
			cnt++
		}
	}

	dp := make([]int, v+1)

	for i := 1; i <= cnt; i++ {
		for j := v; j >= vs[i]; j-- {
			dp[j] = max_i(dp[j], dp[j-vs[i]]+ws[i])
		}
	}

	fmt.Fprintln(out, dp[v])
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
