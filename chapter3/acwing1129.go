package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func _debug() {
	in := bufio.NewReader(os.Stdin)
	_ = []interface{}{in}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const eof = 0
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

	t, c, ts, te := ri(), ri(), ri(), ri()

	g := make([][]edge, t+1)
	for i := 0; i < c; i++ {
		from, to, val := ri(), ri(), ri()
		g[from] = append(g[from], edge{to, val})
		g[to] = append(g[to], edge{from, val})
	}

	dist := make([]int, t+1)
	for i := range dist {
		dist[i] = 1<<31 - 1
	}

	pq := InitHeap(Minheap)

	pq.push_(pair{ts, 0})
	dist[ts] = 0

	for pq.Len() != 0 {
		cur := pq.pop_()

		if dist[cur.id] < cur.distance {
			continue
		}

		for _, e := range g[cur.id] {
			d := dist[cur.id] + e.val
			if d < dist[e.to] {
				dist[e.to] = d
				pq.push_(pair{e.to, d})
			}
		}
	}

	fmt.Fprintln(out, dist[te])
}

type edge struct {
	to, val int
}

// 优先队列/堆
// h := InitHeap(Maxheap)
// h.push_(pair{1})
// p := h.pop_()

const Maxheap = true
const Minheap = false

type pair struct {
	id       int
	distance int
}

// type hp []pair
type hp struct {
	vals []pair
	mode bool // true maxheap; false minheap
}

func InitHeap(mode bool) hp { // init heap
	h := hp{mode: mode}
	heap.Init(&h)
	return h
}

func (h hp) Less(i, j int) bool {
	if h.mode { // true maxheap; 修改对应的if分支
		// return h.vals[i].val > h.vals[j].val
		return h.vals[i].distance > h.vals[j].distance
	}

	return h.vals[i].distance < h.vals[j].distance
}

func (h hp) Len() int            { return len(h.vals) }
func (h hp) Swap(i, j int)       { h.vals[i], h.vals[j] = h.vals[j], h.vals[i] }
func (h *hp) Push(v interface{}) { h.vals = append(h.vals, v.(pair)) }
func (h *hp) Pop() interface{} {
	a := *h
	v := a.vals[len(a.vals)-1]
	*h = hp{a.vals[:len(a.vals)-1], a.mode}
	return v
}
func (h *hp) push_(v pair) { heap.Push(h, v) } // 两个自定义push pop
func (h *hp) pop_() pair   { return heap.Pop(h).(pair) }
func (h *hp) Peek() pair   { return (h.vals)[0] } // 有越界风险

func main() {
	_debug()
}
