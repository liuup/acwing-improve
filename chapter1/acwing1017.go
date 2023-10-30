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
	_ = []interface{}{rc, ri}

	// acwing 1017

	k := ri()

	for ; k > 0; k-- {
		n := ri()
		nums := make([]int, n)
		for i := range nums {
			nums[i] = ri()
		}

		// 正着找一遍
		top := make([]int, n)
		piles1 := 0
		for i := 0; i < n; i++ {
			poker := nums[i]

			left, right := 0, piles1
			for left < right {
				mid := (left + right) / 2
				if top[mid] >= poker {
					right = mid
				} else if top[mid] < poker {
					left = mid + 1
				}
			}
			if left == piles1 {
				piles1++
			}
			top[left] = poker
		}

		// 逆着再找一遍
		top = make([]int, n)
		piles2 := 0
		for i := n - 1; i >= 0; i-- {
			poker := nums[i]

			left, right := 0, piles2
			for left < right {
				mid := (left + right) / 2
				if top[mid] >= poker {
					right = mid
				} else if top[mid] < poker {
					left = mid + 1
				}
			}
			if left == piles2 {
				piles2++
			}
			top[left] = poker
		}
		fmt.Fprintln(out, max_i(piles1, piles2))
	}
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
