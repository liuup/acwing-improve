package main

import (
	"bufio"
	"fmt"
	"os"
)

func _debug() {
	in := bufio.NewReader(os.Stdin) // bufferio 适用于绝大多数题目
	_ = []interface{}{in}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	nums := make([]int, n)
	for i := range nums {
		fmt.Fscan(in, &nums[i])
	}

	f := make([]int, n)
	g := make([]int, n)
	for i := range f {
		f[i] = 1
		g[i] = 1
	}

	// 求上升
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				f[i] = max_i(f[i], f[j]+1)
			}
		}
	}

	// 求下降
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if nums[i] > nums[j] {
				g[i] = max_i(g[i], g[j]+1)
			}
		}
	}

	ans := 1<<31 - 1
	for i := range f {
		ans = min_i(ans, n-(f[i]+g[i]-1))
	}
	fmt.Fprintln(out, ans)
}

func max_i(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min_i(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	_debug()
}
