// local runtime version go1.20.2
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

	n, v := ri(), ri()

	dp := make([]int, v+1)
	cnt := make([]int, v+1) // 背包容积为i时总价值最佳的方案数
	for i := range cnt {
		cnt[i] = 1
	}

	const mod = 1e9 + 7
	for i := 1; i <= n; i++ {
		vs, ws := ri(), ri()

		for j := v; j >= vs; j-- {
			value := dp[j-vs] + ws
			if value > dp[j] { // 新方案价值更大
				dp[j] = value
				cnt[j] = cnt[j-vs]
			} else if value == dp[j] {
				cnt[j] = (cnt[j] + cnt[j-vs]) % mod
			}
		}
	}
	fmt.Fprintln(out, cnt[v])
}

func main() {
	_solve()
}
