package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func _solve() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	ri := func() int { // 读一个整数
		in.Scan()
		x, _ := strconv.Atoi(string(in.Bytes()))
		return x
	}
	_ = []interface{}{ri}

	n := ri()

	nums := make([]int, n)
	for i := range nums {
		nums[i] = ri()
	}

	last2 := make([]int, 2)
	last1 := make([]int, 2)
	dp := make([]int, 2)

	for i := 0; i < n; i++ {

		if i-1 == -1 {
			last2[0] = 0
			last2[1] = -nums[i]
			continue
		}
		if i-2 == -1 {
			last1[0] = max_i(last2[0], last2[1]+nums[i])
			last1[1] = max_i(last2[1], -nums[i])
			continue
		}

		dp[0] = max_i(last1[0], last1[1]+nums[i])
		dp[1] = max_i(last1[1], last2[0]-nums[i])
		copy(last2, last1)
		copy(last1, dp)
	}
	fmt.Fprintln(out, dp[0])
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
