// local runtime version go1.20.2
package main

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
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

	/*
		集合: dp(i, j, s) 所有摆在前i行,已经摆了j个国王,并且第j行摆放的状态
		是s的所有方案的集合 s是一个二进制的数字

		属性: count
		枚举上一层状态是什么

		状态计算: dp(i, j, a) = 累加dp[i-1, j-c[a], b]	// c[a]是第i行放的国王个数

		最终求得: 累加state dp[n][k][state]
	*/
	// 遍历所有状态

	n, k := ri(), ri() // 棋盘行数/列数 国王总数

	// 每一行的合法状态都是相同的,所有一个一维数组就够了
	s := []int{}             // 同一行的合法状态集
	num := make([]int, 1<<n) // 每个合法状态包含的国王数

	// 前i行放了j个国王,第j行第a个状态时的方案数
	dp := make([][][]int64, n+2)
	for i := range dp {
		dp[i] = make([][]int64, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int64, 1<<n)
		}
	}

	// 预处理
	for a := 0; a < (1 << n); a++ { // 枚举一行的所有状态
		if a&(a>>1) > 0 { // 存在相邻的1 跳过
			continue
		}
		s = append(s, a)
		num[a] = bits.OnesCount(uint(a))
	}

	cnt := len(s) // 同一行的合法状态总数

	dp[0][0][0] = 1             // 不放也是一种方案
	for i := 1; i <= n+1; i++ { // 枚举行
		for j := 0; j <= k; j++ { // 枚举国王数
			for a := 0; a < cnt; a++ { // 枚举所有状态
				for b := 0; b < cnt; b++ { // 枚举i-1行的合法状态
					c := num[s[a]] // 第i行 a状态的国王数

					// 假如前面已经放了c个国王,那么现在还能放j-c个国王
					if j >= c && (s[b]&s[a] == 0) && (s[b]&(s[a]>>1) == 0) && (s[b]&(s[a]<<1) == 0) {
						dp[i][j][a] += dp[i-1][j-c][b]
					}
				}
			}
		}
	}

	// 第n+1行,恰好放了k个国王,第n+1行不存在国王,也就是状态为0
	fmt.Fprintln(out, dp[n+1][k][0])
}

func main() {
	_solve(os.Stdin, os.Stdout)
}
