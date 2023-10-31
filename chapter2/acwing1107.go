package main

import (
	"bufio"
	"fmt"
	"os"
)

func _debug() {
	const eof = 0
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	_i, _n, buf := 0, 0, make([]byte, 1<<12) // 4KB
	rc := func() byte {                      // 读一个字符
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
	_ = []interface{}{rc, ri}

	ed := make([]byte, 8)
	for i := range ed {
		ed[i] = byte('0' + ri())
	}

	start := "12345678"

	q := []string{}
	dist := map[string]int{}
	pre := map[string]node{}

	q = append(q, start)

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		if cur == string(ed) {
			break
		}

		child := [3]string{}
		child[0] = moveA(cur)
		child[1] = moveB(cur)
		child[2] = moveC(cur)

		for i, c := range child {
			if _, ok := pre[c]; !ok {
				q = append(q, c)
				dist[c] = dist[cur] + 1
				pre[c] = node{cur, byte('A' + i)}
			}
		}
	}

	end := string(ed)

	fmt.Println(dist[end])
	if dist[end] != 0 {
		res := []byte{}
		for end != start {
			res = append(res, pre[end].op)
			end = pre[end].last
		}
		// 反转一下路径
		for i := 0; i < len(res)/2; i++ {
			res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
		}

		fmt.Println(string(res))
	}
}

type node struct {
	last string
	op   byte
}

func moveA(s string) string {
	sb := []byte(s)
	for i := 0; i < 4; i++ {
		sb[i], sb[7-i] = sb[7-i], sb[i]
	}
	return string(sb)
}

func moveB(s string) string {
	sb := []byte(s)
	for i := 0; i < 3; i++ {
		sb[3], sb[i] = sb[i], sb[3]
	}
	for i := 4; i < 7; i++ {
		sb[i], sb[i+1] = sb[i+1], sb[i]
	}
	return string(sb)
}

func moveC(s string) string {
	sb := []byte(s)
	sb[1], sb[2] = sb[2], sb[1]
	sb[5], sb[6] = sb[6], sb[5]
	sb[1], sb[5] = sb[5], sb[1]
	return string(sb)
}

func main() {
	_debug()
}
