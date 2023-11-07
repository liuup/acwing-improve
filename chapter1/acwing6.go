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

	n, v := ri(), ri()

	vs := make([]int, n+1) // 体积
	ws := make([]int, n+1) // 价值
	ss := make([]int, n+1) // 数量

	for i := 1; i <= n; i++ {
		vs[i] = ri()
		ws[i] = ri()
		ss[i] = ri()
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, v+1)
	}

	for i := 1; i <= n; i++ {
		for r := 0; r < vs[i]; r++ { // 体积
			q := []int{}
			for j := r; j <= v; j += vs[i] {
				for len(q) != 0 && j-q[0] > ss[i]*vs[i] {
					q = q[1:]
				}
				for len(q) != 0 && dp[i-1][q[len(q)-1]]+(j-q[len(q)-1])/vs[i]*ws[i] <= dp[i-1][j] {
					q = q[:len(q)-1]
				}
				q = append(q, j)
				dp[i][j] = dp[i-1][q[0]] + (j-q[0])/vs[i]*ws[i]
			}
		}
	}
	fmt.Fprintln(out, dp[n][v])
}

func main() {
	_solve()

}
