// local runtime version go1.20.2
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func _solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const eof = 0
	_i, _n, buf := 0, 0, make([]byte, 1<<12) // 4KB

	rc := func() byte { // 读一个字符
		if _i == _n {
			_n, _ = in.Read(buf)
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

	t := ri()

	type node struct {
		s, e, l int
	}

	for cc := 1; cc <= t; cc++ {
		n := ri()

		t := 0
		pms := make([]node, n)
		for i := 0; i < n; i++ {
			// s 吃完花费的时间
			// e 最初包含的能量
			// l 每秒失去的能量

			s, e, l := ri(), ri(), ri()
			t += s
			pms[i] = node{s, e, l}
		}

		// 先走一步贪心
		sort.Slice(pms, func(i, j int) bool {
			return pms[i].s*pms[j].l < pms[i].l*pms[j].s
		})

		dp := make([]int, t+1)
		for i := 1; i <= n; i++ {
			dp[i] = -1 << 31
		}
		dp[0] = 0

		for i := 0; i < n; i++ {
			for j := t; j >= pms[i].s; j-- {
				dp[j] = max_i(dp[j], dp[j-pms[i].s]+max_i(0, pms[i].e-(j-pms[i].s)*pms[i].l))
			}
		}

		ans := 0
		for i := 1; i <= t; i++ {
			ans = max_i(ans, dp[i])
		}

		fmt.Fprintf(out, "Case #%d: %d\n", cc, ans)
	}
}

func max_i(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	_solve(os.Stdin, os.Stdout)

}
