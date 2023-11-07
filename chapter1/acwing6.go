package main

import (
	"bufio"
	"fmt"
	"os"
)

func _solve() {
	const eof = 0
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	_i, _n, buf := 0, 0, make([]byte, 1<<12) // 4KB

	rc := func() byte { // 读一个字符
		if _i == _n {
			_n, _ = os.Stdin.Read(buf)
			if _n == 0 { // EOF
				return eof
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}

	ri := func() (x int) { // 读一个整数，支持负数
		neg := false
		b := rc()
		for ; '0' > b || b > '9'; b = rc() {
			// 某些多组数据的题目，不告诉有多少组数据，那么需要额外判断是否读到了 EOF
			if b == eof {
				return
			}
			if b == '-' {
				neg = true
			}
		}
		for ; '0' <= b && b <= '9'; b = rc() {
			x = x*10 + int(b&15)
		}
		if neg {
			return -x
		}
		return
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

	f := make([]int, v+1)
	g := make([]int, v+1)

	// https://www.bilibili.com/video/BV1354y1C7SF

	for i := 1; i <= n; i++ {
		copy(g, f)
		for r := 0; r < vs[i]; r++ { // 分类

			// 单调队列的窗口是在g数组上滑动的 里面存放的是下标
			// 其下标所对应的值从大到小进行排列
			q := []int{}
			for j := r; j <= v; j += vs[i] {
				// q[0] 不在窗口[k-s*v,k-v]中，下标越界了
				for len(q) != 0 && q[0] < j-ss[i]*vs[i] { // q[0] < k-s*v
					q = q[1:]
				}

				// (j-q[len(q)-1])/vs[i] 是还能放入的物品的个数 排除尾元素
				for len(q) != 0 && g[j] >= g[q[len(q)-1]]+(j-q[len(q)-1])/vs[i]*ws[i] {
					q = q[:len(q)-1]
				}

				q = append(q, j)

				// 用队头最大值更新
				f[j] = g[q[0]] + (j-q[0])/vs[i]*ws[i]
			}
		}
	}
	fmt.Fprintln(out, f[v])
}

func main() {
	_solve()

}
