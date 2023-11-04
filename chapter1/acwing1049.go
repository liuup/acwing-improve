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

	t := ri()
	for ; t > 0; t-- {
		n := ri()
		nums := make([]int, n)
		for i := range nums {
			nums[i] = ri()
		}

		a, b, c := 0, nums[0], 0

		for i := 2; i <= n; i++ {
			c = max_i(b, a+nums[i-1])
			a = b
			b = c
		}

		fmt.Fprintln(out, c)
	}
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
