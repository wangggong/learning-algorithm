/*
 * Hard
 * 给你一个大小为 m x n 的二进制网格 grid 。网格表示一个地图，其中，0 表示水，1 表示陆地。最初，grid 中的所有单元格都是水单元格（即，所有单元格都是 0）。
 *
 * 可以通过执行 addLand 操作，将某个位置的水转换成陆地。给你一个数组 positions ，其中 positions[i] = [ri, ci] 是要执行第 i 次操作的位置 (ri, ci) 。
 *
 * 返回一个整数数组 answer ，其中 answer[i] 是将单元格 (ri, ci) 转换为陆地后，地图中岛屿的数量。
 *
 * 岛屿 的定义是被「水」包围的「陆地」，通过水平方向或者垂直方向上相邻的陆地连接而成。你可以假设地图网格的四边均被无边无际的「水」所包围。
 *
 *
 * 示例 1：
 *
 *
 * 输入：m = 3, n = 3, positions = [[0,0],[0,1],[1,2],[2,1]]
 * 输出：[1,1,2,3]
 * 解释：
 * 起初，二维网格 grid 被全部注入「水」。（0 代表「水」，1 代表「陆地」）
 * - 操作 #1：addLand(0, 0) 将 grid[0][0] 的水变为陆地。此时存在 1 个岛屿。
 * - 操作 #2：addLand(0, 1) 将 grid[0][1] 的水变为陆地。此时存在 1 个岛屿。
 * - 操作 #3：addLand(1, 2) 将 grid[1][2] 的水变为陆地。此时存在 2 个岛屿。
 * - 操作 #4：addLand(2, 1) 将 grid[2][1] 的水变为陆地。此时存在 3 个岛屿。
 * 示例 2：
 *
 * 输入：m = 1, n = 1, positions = [[0,0]]
 * 输出：[1]
 *
 *
 * 提示：
 *
 * 1 <= m, n, positions.length <= 1e4
 * 1 <= m * n <= 1e4
 * positions[i].length == 2
 * 0 <= ri < m
 * 0 <= ci < n
 *
 *
 * 进阶：你可以设计一个时间复杂度 O(k log(mn)) 的算法解决此问题吗？（其中 k == positions.length）
 *
 * 通过次数5,408
 * 提交次数13,565
 */

const MAXN = 1e4 + 10

var visit [MAXN][MAXN]bool
var directions = [4][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

type UnionSet struct {
	Father map[int]int
	Size   map[int]int
	count  int
}

func (us UnionSet) Get(id int) int {
	// assert us.Father != nil;
	if us.Father[id] == id {
		return id
	}
	v := us.Get(us.Father[id])
	us.Father[id] = v
	return v
}

func (us *UnionSet) Add(id int) {
	// assert us.Father != nil;
	// assert us.Size != nil;
	us.Father[id] = id
	us.Size[id] = 1
	us.count++
}

func (us *UnionSet) Join(p, q int) {
	// assert us.Father != nil;
	// assert us.Size != nil;
	fp := us.Get(p)
	fq := us.Get(q)
	if fp == fq {
		return
	}
	if sizeFp, sizeFq := us.Size[fp], us.Size[fq]; sizeFp > sizeFq {
		us.Father[fq] = fp
		us.Size[fp] += sizeFq
	} else {
		us.Father[fp] = fq
		us.Size[fq] += sizeFp
	}
	us.count--
}

func (us UnionSet) Count() int { return us.count }

func getID(x, y int) int {
	return x * MAXN + y
}

func numIslands2(m int, n int, positions [][]int) []int {
	us := &UnionSet{Father: make(map[int]int), Size: make(map[int]int)}
	var ans []int
	for i := 0; i < MAXN; i++ {
		for j := 0; j < MAXN; j++ {
			visit[i][j] = false
		}
	}
	for _, p := range positions {
		x, y := p[0], p[1]
		if visit[y][x] {
			ans = append(ans, us.Count())
			continue
		}
		curr := getID(x, y)
		us.Add(curr)
		visit[y][x] = true
		for _, d := range directions {
			nx, ny := x+d[0], y+d[1]
			if !valid(m, n, nx, ny) {
				continue
			}
			if !visit[ny][nx] {
				continue
			}
			next := getID(nx, ny)
			if p, q := us.Get(curr), us.Get(next); p == q {
				continue
			}
			us.Join(curr, next)
		}
		ans = append(ans, us.Count())
	}
	return ans
}

func valid(m, n, x, y int) bool {
	return 0 <= x && x < m && 0 <= y && y < n
}

