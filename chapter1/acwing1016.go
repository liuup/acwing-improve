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

	var k int
	fmt.Fscan(in, &k)

	nums := make([]int, k)
	dp := make([]int, k)

	for i := range nums {
		fmt.Fscan(in, &nums[i])
		dp[i] = nums[i] // 初始化自己，自己就是一个上升序列
	}

	ans := -1
	for i := 0; i < k; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max_i(dp[i], dp[j]+nums[i])
			}
		}
		ans = max_i(ans, dp[i])
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
