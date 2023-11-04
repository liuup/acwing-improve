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

	// 读一个整数
	ri := func() int {
		in.Scan()
		x, _ := strconv.Atoi(string(in.Bytes()))
		return x
	}

	n, maxk := ri(), ri()

	nums := make([]int, n)
	for i := range nums {
		nums[i] = ri()
	}

	dp := make([][2]int, maxk+1)

	for j := 1; j <= maxk; j++ {
		dp[j][1] = -nums[0]
	}

	for i := 1; i < n; i++ {
		for k := maxk; k >= 1; k-- {
			dp[k][0] = max_i(dp[k][0], dp[k][1]+nums[i])
			dp[k][1] = max_i(dp[k][1], dp[k-1][0]-nums[i])
		}
	}
	fmt.Fprintln(out, dp[maxk][0])
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
