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

	top := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		poker := nums[i].end

		left, right := 0, ans
		for left < right {
			mid := (left + right) / 2
			if top[mid] >= poker {
				right = mid
			} else {
				left = mid + 1
			}
		}

		if left == ans {
			ans++
		}
		top[left] = poker
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
