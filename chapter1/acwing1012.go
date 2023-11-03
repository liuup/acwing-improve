package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func _solve(in *bufio.Reader, out *bufio.Writer) {

	var n int
	fmt.Fscan(in, &n)

	var s, e int
	nums := make([]node, n)
	for i := range nums {
		fmt.Fscan(in, &s, &e)
		nums[i] = node{s, e}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i].start < nums[j].start
	})

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	ans := -1
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i].start > nums[j].start && nums[i].end > nums[j].end {
				dp[i] = max_i(dp[i], dp[j]+1)
			}
		}
		ans = max_i(ans, dp[i])
	}

	fmt.Fprintln(out, ans)
}

type node struct {
	start, end int
}

func max_i(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {

	in := bufio.NewReader(os.Stdin) // bufferio 适用于绝大多数题目
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	_solve(in, out)

}
