package main

import (
	"bufio"
	"fmt"
	"os"
)

func _debug() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	_ = []interface{}{in}

	// acwing 1098

	var n, m int
	fmt.Fscan(in, &n, &m)
	grid := make([][]int, 2*n+1)
	for i := range grid {
		grid[i] = make([]int, 2*m+1)
	}

	var a int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &a)
			set(grid, i, j, a)
		}
	}

	n = 2*n + 1
	m = 2*m + 1

	// 把边界处补为墙
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 || i == 2*n || j == 2*m {
				grid[i][j] = 1
			}
		}
	}

	node := func(i, j int) int { return i*m + j }

	us := InitUnion(n*m + 1)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				for _, d := range directions {
					di := i + d[0]
					dj := j + d[1]

					if isok(grid, di, dj) && grid[di][dj] == 0 {
						us.Union(node(i, j), node(di, dj))
					}
				}
			}
		}
	}

	mp := map[int]int{} // 统计每个连通分块的大小
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				mp[us.Find(node(i, j))]++
			}
		}
	}

	for i := 0; i < n; i += 2 { // 排除横向的墙壁
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				mp[us.Find(node(i, j))]--
				grid[i][j] = 2 //  mark 要不然会重复计算墙壁
			}
		}
	}
	for j := 0; j < m; j += 2 { // 排除纵向的墙壁
		for i := 0; i < n; i++ {
			if grid[i][j] == 0 {
				mp[us.Find(node(i, j))]--
				grid[i][j] = 2 // mark
			}
		}
	}

	area := 0
	for _, v := range mp {
		if v > area {
			area = v
		}
	}
	fmt.Fprintln(out, len(mp))
	fmt.Fprintln(out, area)
}

// union set 并查集
type unionset struct {
	father []int
	count  int // 连通分量的个数
}

func InitUnion(n int) unionset {
	fa := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fa[i] = i

	}
	return unionset{fa, n}
}

func (u *unionset) Find(i int) int {
	if u.father[i] == i {
		return i
	}
	u.father[i] = u.Find(u.father[i]) // 路径压缩
	return u.father[i]
}

func (u *unionset) Union(i, j int) {
	i_fa, j_fa := u.Find(i), u.Find(j)
	if i_fa == j_fa {
		return
	}
	u.father[i_fa] = j_fa
	u.count--
}

func (u *unionset) IsConnected(x, y int) bool {
	return u.Find(x) == u.Find(y)
}

func set(grid [][]int, i, j, a int) {
	if a == 1 {
		set1(grid, i, j, a)
	} else if a == 2 {
		set2(grid, i, j, a)
	} else if a == 4 {
		set4(grid, i, j, a)
	} else if a == 8 {
		set8(grid, i, j, a)
	} else if a == 3 { // 1 2
		set1(grid, i, j, a)
		set2(grid, i, j, a)
	} else if a == 5 { // 1 4
		set1(grid, i, j, a)
		set4(grid, i, j, a)
	} else if a == 9 { // 1 8
		set1(grid, i, j, a)
		set8(grid, i, j, a)
	} else if a == 6 { // 2 4
		set2(grid, i, j, a)
		set4(grid, i, j, a)
	} else if a == 10 { // 2 8
		set2(grid, i, j, a)
		set8(grid, i, j, a)
	} else if a == 12 { // 4 8
		set4(grid, i, j, a)
		set8(grid, i, j, a)
	} else if a == 7 { // 1 2 4
		set1(grid, i, j, a)
		set2(grid, i, j, a)
		set4(grid, i, j, a)
	} else if a == 11 { // 1 2 8
		set1(grid, i, j, a)
		set2(grid, i, j, a)
		set8(grid, i, j, a)
	} else if a == 13 { // 1 4 8
		set1(grid, i, j, a)
		set4(grid, i, j, a)
		set8(grid, i, j, a)
	} else if a == 14 { // 2 4 8
		set2(grid, i, j, a)
		set4(grid, i, j, a)
		set8(grid, i, j, a)
	} else if a == 15 { // 1 2 4 8
		set1(grid, i, j, a)
		set2(grid, i, j, a)
		set4(grid, i, j, a)
		set8(grid, i, j, a)
	}
}

func set1(grid [][]int, i, j, a int) {
	grid[i*2][j*2] = 1
	grid[i*2+1][j*2] = 1
	grid[i*2+2][j*2] = 1
}

func set2(grid [][]int, i, j, a int) {
	grid[i*2][j*2] = 1
	grid[i*2][j*2+1] = 1
	grid[i*2][j*2+2] = 1
}

func set4(grid [][]int, i, j, a int) {
	grid[i*2][j*2+2] = 1
	grid[i*2+1][j*2+2] = 1
	grid[i*2+2][j*2+2] = 1
}
func set8(grid [][]int, i, j, a int) {
	grid[i*2+2][j*2] = 1
	grid[i*2+2][j*2+1] = 1
	grid[i*2+2][j*2+2] = 1
}

var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func isok(grid [][]int, i, j int) bool { // 判断是否在二维数组越界
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

func main() {
	_debug()
}
