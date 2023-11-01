package main

import (
	"bufio"
	"fmt"
	"os"
)

func _debug() {
	in := bufio.NewReader(os.Stdin)
	const eof = 0
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	_ = []interface{}{in}

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

	// acwing 1100

	n, k := ri(), ri()

	if n >= k { // 特殊情况
		fmt.Fprintln(out, n-k)
		return
	}

	q := []int{}
	mp := map[int]struct{}{}

	q = append(q, n)
	mp[n] = struct{}{}

	sp := 0
	for len(q) != 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]

			if cur == k {
				fmt.Fprintln(out, sp)
				return
			}

			// 三种移动方式，要确定上界和下界
			if _, ok := mp[cur-1]; !ok && cur-1 >= 0 {
				q = append(q, cur-1)
				mp[cur-1] = struct{}{}
			}

			if _, ok := mp[cur+1]; !ok && cur+1 <= 1e5 {
				q = append(q, cur+1)
				mp[cur+1] = struct{}{}
			}

			if _, ok := mp[2*cur]; !ok && 2*cur <= 1e5 {
				q = append(q, 2*cur)
				mp[2*cur] = struct{}{}
			}
		}
		sp++
	}
}

func main() {
	_debug()
}
