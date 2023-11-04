package main

import (
	"bufio"
	"fmt"
	"os"
)

func _solve() {
	in := bufio.NewReader(os.Stdin) // bufferio 适用于绝大多数题目
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	ri := func() (a int) { fmt.Fscan(in, &a); return }
	rs := func() (s string) { fmt.Fscan(in, &s); return }
	_ = []interface{}{ri, rs}

	n, v := ri(), ri()
	vs := make([]int, n)
	ws := make([]int, n)

	for i := range vs {
		vs[i], ws[i] = ri(), ri()
	}

	dp := make([]int, v+1)
	cnt := make([]int, v+1) // 背包容积为i时总价值最佳的方案数
	for i := range cnt {    // 什么都不装也是一种方案数
		cnt[i] = 1
	}

	const mod = 1e9 + 7
	for i := 1; i <= n; i++ { // n个物品
		for j := v; j >= vs[i-1]; j-- {
			value := dp[j-vs[i-1]] + ws[i-1]
			if value > dp[j] {
				dp[j] = value
				cnt[j] = cnt[j-vs[i-1]]
			} else if value == dp[j] {
				cnt[j] = (cnt[j] + cnt[j-vs[i-1]]) % mod
			}
		}
	}

	fmt.Fprintln(out, cnt[v])
}

func main() {
	_solve()
}
