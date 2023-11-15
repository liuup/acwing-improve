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

	type item struct {
		value  int // 物品的价格
		weight int // 价格 * 权重
	}

	master := make([]item, m+1)
	components := make([][]item, m+1)
	for i := range components {
		components[i] = make([]item, 0, 2)
	}

	for i := 1; i <= m; i++ {
		v, p, q := ri(), ri(), ri()
		if q == 0 {
			master[i] = item{v, v * p}
		} else {
			components[q] = append(components[q], item{v, v * p})
		}
	}

	dp := make([]int, n+1)

	for i := 1; i <= m; i++ {
		if master[i].weight == 0 { // 忽略价值为0的
			continue
		}

		for j := n; j >= 0; j-- {
			solutions := 1 << len(components[i]) // 方案的个数

			for k := 0; k < solutions; k++ { // 枚举所有方案
				tmpv, tmpw := master[i].value, master[i].weight
				for u := 0; u < len(components[i]); u++ { // 有多少个附件
					if (k>>u)&1 != 0 { // 有没有选择买某一个附件
						tmpv += components[i][u].value
						tmpw += components[i][u].weight
					}
				}
				if j >= tmpv {
					dp[j] = max_i(dp[j], dp[j-tmpv]+tmpw)
				}
			}
		}
	}

	fmt.Fprintln(out, dp[n])
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
