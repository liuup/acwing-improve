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
	rsn := func(n int) []byte {
		b := rc()
		for ; 'A' > b || b > 'Z'; b = rc() { // 'A' 'Z'
		}
		s := make([]byte, 0, n)
		s = append(s, b)
		for i := 1; i < n; i++ {
			s = append(s, rc())
		}
		return s
	}
	_ = []interface{}{ri, rsn}

	n, m := ri(), ri()

	// P=1 H=1
	grid := make([]int, n+10) // 稍微开大点
	for i := 1; i <= n; i++ {
		sb := rsn(m)
		for j := 0; j < m; j++ {
			if sb[j] == 'P' {
				grid[i] |= (1 << (m - j - 1))
			}
		}
	}

	// 预处理每行的合法状态
	s := []int{}    // 合法状态
	nums := []int{} // 每个状态中1的个数
	for a := 0; a < (1 << m); a++ {
		if (a&(a>>1) != 0) || (a&(a>>2) != 0) {
			continue
		}
		s = append(s, a)
		nums = append(nums, bits.OnesCount(uint(a)))
	}

	// dp(i, a, b) 前i行已经摆好 第i行状态为a时 第i-1行第b个状态时 能摆放的最大数量
	dp := [2][][]int{}
	for i := range dp {
		dp[i] = make([][]int, 1<<m)
		for j := range dp[i] {
			dp[i][j] = make([]int, 1<<m)
		}
	}

	for i := 1; i <= n+2; i++ { // 遍历所有行
		for a := 0; a < len(s); a++ {
			for b := 0; b < len(s); b++ {
				for c := 0; c < len(s); c++ {
					if (s[a]&s[b] == 0) && (s[a]&s[c] == 0) && (s[b]&s[c] == 0) && (s[a]&grid[i] == s[a]) && (s[b]&grid[i-1] == s[b]) {
						dp[i&1][a][b] = max_i(dp[i&1][a][b], dp[(i-1)&1][b][c]+nums[a])
					}
				}
			}
		}
	}
	fmt.Fprintln(out, dp[(n+2)&1][0][0])
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
