/*
 * @lc app=leetcode.cn id=count-unguarded-cells-in-the-grid lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2257] 统计网格图中没有被保卫的格子数
 *
 * https://leetcode-cn.com/problems/count-unguarded-cells-in-the-grid/description/
 *
 * algorithms
 * Medium (48.17%)
 * Total Accepted:    2.9K
 * Total Submissions: 6K
 * Testcase Example:  '4\n6\n[[0,0],[1,1],[2,3]]\n[[0,1],[2,2],[1,4]]'
 *
 * 给你两个整数 m 和 n 表示一个下标从 0 开始的 m x n 网格图。同时给你两个二维整数数组 guards 和 walls ，其中
 * guards[i] = [rowi, coli] 且 walls[j] = [rowj, colj] ，分别表示第 i 个警卫和第 j
 * 座墙所在的位置。
 *
 * 一个警卫能看到 4 个坐标轴方向（即东、南、西、北）的 所有 格子，除非他们被一座墙或者另外一个警卫 挡住 了视线。如果一个格子能被 至少
 * 一个警卫看到，那么我们说这个格子被 保卫 了。
 *
 * 请你返回空格子中，有多少个格子是 没被保卫 的。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：m = 4, n = 6, guards = [[0,0],[1,1],[2,3]], walls = [[0,1],[2,2],[1,4]]
 * 输出：7
 * 解释：上图中，被保卫和没有被保卫的格子分别用红色和绿色表示。
 * 总共有 7 个没有被保卫的格子，所以我们返回 7 。
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：m = 3, n = 3, guards = [[1,1]], walls = [[0,1],[1,0],[2,1],[1,2]]
 * 输出：4
 * 解释：上图中，没有被保卫的格子用绿色表示。
 * 总共有 4 个没有被保卫的格子，所以我们返回 4 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= m, n <= 10^5
 * 2 <= m * n <= 10^5
 * 1 <= guards.length, walls.length <= 5 * 10^4
 * 2 <= guards.length + walls.length <= m * n
 * guards[i].length == walls[j].length == 2
 * 0 <= rowi, rowj < m
 * 0 <= coli, colj < n
 * guards 和 walls 中所有位置 互不相同 。
 *
 *
 */
var G [][]int

const (
	_default = iota
	guarded
	wall
)

var directions = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func countUnguarded(m int, n int, Gs [][]int, Ws [][]int) int {
	G = make([][]int, m)
	for i := range G {
		G[i] = make([]int, n)
	}
	var Q [][]int
	for _, g := range Gs {
		y, x := g[0], g[1]
		for i := 0; i < 4; i++ {
			cur := []int{y, x, i}
			Q = append(Q, cur)
		}
		G[y][x] = wall
	}
	for _, w := range Ws {
		y, x := w[0], w[1]
		G[y][x] = wall
	}
	for len(Q) > 0 {
		q := Q[0]
		Q = Q[1:]
		d := directions[q[2]]
		ny, nx := q[0]+d[0], q[1]+d[1]
		if ny < 0 || ny >= m || nx < 0 || nx >= n || G[ny][nx] == wall {
			continue
		}
		Q = append(Q, []int{ny, nx, q[2]})
		G[ny][nx] = guarded
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if G[i][j] == _default {
				ans++
			}
		}
	}
	return ans
}
