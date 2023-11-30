// local runtime version go1.20.2
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func _solve(_r io.Reader, _w io.Writer) {
	_in := bufio.NewReader(_r) // io速度基本够用了
	out := bufio.NewWriter(_w)
	defer out.Flush()

	ri := func() (ans int) { fmt.Fscan(_in, &ans); return ans }
	rs := func() (ans string) { fmt.Fscan(_in, &ans); return ans }
	_ = []interface{}{ri, rs}

	n := ri()

	g := make([][]int, n+1)

	in := make([]int, n+1) // 入度统计

	for i := 1; i <= n; i++ {
		for {
			ch := ri()
			if ch == 0 {
				break
			}
			g[i] = append(g[i], ch)
			in[ch]++
		}
	}

	q := []int{}
	for i := 1; i <= n; i++ {
		if in[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		fmt.Fprint(out, cur, " ")

		for _, e := range g[cur] { // 遍历所有孩子
			in[e]--
			if in[e] == 0 {
				q = append(q, e)
			}
		}
	}
}

func main() {
	_solve(os.Stdin, os.Stdout)
}
