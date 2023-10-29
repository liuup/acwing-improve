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
	rs := func() (s []byte) {
		b := rc()
		for ; '0' > b || b > '9'; b = rc() { // 'A' 'Z'
		}
		// for ; 'a' <= b && b <= 'z'; b = rc() { // 'A' 'Z'
		// 	s = append(s, b)
		// }
		for ; '0' <= b && b <= '9'; b = rc() { // 'A' 'Z'
			s = append(s, b)
		}
		return
	}
	_ = []interface{}{rc, ri, rs}

	// acwing 423
	// 01背包

	t, m := ri(), ri() // 时间 数目

	time := make([]int, m)  // 时间
	value := make([]int, m) // 价值

	for i := 0; i < m; i++ {
		time[i] = ri()
		value[i] = ri()
	}

	// 对于前m个物品，消耗t时间 所拿到的最大价值
	dp := make([]int, t+1)

	for i := 1; i <= m; i++ {
		for j := t; j >= time[i-1]; j-- {
			dp[j] = max_i(dp[j], dp[j-time[i-1]]+value[i-1])
		}
	}
	fmt.Fprintln(out, dp[t])
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
