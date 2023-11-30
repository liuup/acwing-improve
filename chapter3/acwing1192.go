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

	n, m := ri(), ri()

	g := make([][]int, n+1)
	in := make([]int, n+1) // 统计入度

	salary := make([]int, n+1) // 工资

	for i := 1; i <= m; i++ {
		a, b := ri(), ri()
		g[b] = append(g[b], a)
		in[a]++
	}

	q := []int{}
	for i := 1; i <= n; i++ {
		if in[i] == 0 {
			q = append(q, i)
			salary[i] = 100
		}
	}

	count := 0

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		count++

		for _, e := range g[cur] {
			in[e]--
			if in[e] == 0 {
				salary[e] = salary[cur] + 1
				q = append(q, e)
			}
		}
	}

	if count != n {
		fmt.Fprintln(out, "Poor Xed")
	} else {
		sum := 0
		for i := 1; i <= n; i++ {
			sum += salary[i]
		}
		fmt.Fprintln(out, sum)
	}
}

func main() {
	_solve(os.Stdin, os.Stdout)
}
