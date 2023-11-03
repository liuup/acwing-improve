package main

import (
	"bufio"
	"fmt"
	"os"
)

func _solve(in *bufio.Reader, out *bufio.Writer) {
	nums := make([]int, 1010)
	n := 0 // 长度

	for {
		_, err := fmt.Fscan(in, &nums[n])
		if err != nil {
			break
		}
		n++
	}

	f := make([]int, n)
	g := make([]int, n)

	ans1 := -1
	ans2 := -1
	for i := 0; i < n; i++ {
		f[i], g[i] = 1, 1
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				f[i] = max_i(f[i], f[j]+1)
			} else {
				g[i] = max_i(g[i], g[j]+1)
			}
		}
		ans1 = max_i(ans1, f[i])
		ans2 = max_i(ans2, g[i])
	}

	fmt.Fprintln(out, ans1)
	fmt.Fprintln(out, ans2)
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
