// local runtime version go1.20.2
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func _solve(_r io.Reader, _w io.Writer) {
	_in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const eof = 0
	_i, _n, buf := 0, 0, make([]byte, 1<<12) // 4KB

	rc := func() byte { // 读一个字符
		if _i == _n {
			_n, _ = _in.Read(buf)
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

	// dp(i, a) 第i行状态为a的种植方案数
	/*
		讨论每一行的状态 1<<n 种状态
		同一行:
		1. 只能种在肥沃的土地上
		2. 相邻不能种

		上下:
		1. 不能相邻
	*/

	grid := make([]int, m+2)

	// 把每一行转成二进制 方便后序处理
	for i := 1; i <= m; i++ {
		for j := 0; j < n; j++ {
			k := ri()
			if k == 1 {
				grid[i] |= (1 << (n - j - 1))
			}
		}
	}

	// 预处理每一行的所有可行状态
	s := make([]int, 0, 1<<n)
	for a := 0; a < (1 << n); a++ {
		if a&(a>>1) != 0 { // 有相邻 跳过
			continue
		}
		s = append(s, a)
	}

	// dp(i, a) 第i行状态为a时所得到的方案数
	dp := make([][]int64, m+5)
	for i := range dp {
		dp[i] = make([]int64, 1<<n)
	}

	dp[0][0] = 1                // 什么都不种也是一种方案
	for i := 1; i <= m+1; i++ { // 枚举行
		for a := 0; a < len(s); a++ { // 枚举第i行所有合法状态
			// 枚举第i-1行合法状态
			for b := 0; b < len(s); b++ {
				// a b种在肥沃土地上, a b同列不同时为1
				if ((s[a] & grid[i]) == s[a]) && (s[b]&grid[i-1] == s[b]) && (s[a]&s[b] == 0) {
					dp[i][a] = (dp[i][a] + dp[i-1][b]) % 10e8
				}
			}
		}
	}
	fmt.Fprintln(out, dp[m+1][0])
}

func main() {
	_solve(os.Stdin, os.Stdout)
}
