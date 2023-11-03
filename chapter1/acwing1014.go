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

	dp1 := make([]int, n)
	dp2 := make([]int, n)
	for i := range dp1 {
		dp1[i] = 1
		dp2[i] = 1
	}

	// 求上升
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp1[i] = max_i(dp1[i], dp1[j]+1)
			}
		}
	}

	// 求下降
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if nums[j] < nums[i] {
				dp2[i] = max_i(dp2[i], dp2[j]+1)
			}
		}
	}

	ans := -1
	for i := range dp1 {
		ans = max_i(ans, dp1[i]+dp2[i]-1)
	}
	fmt.Fprintln(out, ans)
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
