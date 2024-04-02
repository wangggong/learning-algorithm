/*
 * @lc app=leetcode.cn id=making-a-large-island lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [827] 最大人工岛
 *
 * https://leetcode-cn.com/problems/making-a-large-island/description/
 *
 * algorithms
 * Hard (39.47%)
 * Total Accepted:    13.3K
 * Total Submissions: 33.6K
 * Testcase Example:  '[[1,0],[0,1]]'
 *
 * 给你一个大小为 n x n 二进制矩阵 grid 。最多 只能将一格 0 变成 1 。
 *
 * 返回执行此操作后，grid 中最大的岛屿面积是多少？
 *
 * 岛屿 由一组上、下、左、右四个方向相连的 1 形成。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: grid = [[1, 0], [0, 1]]
 * 输出: 3
 * 解释: 将一格0变成1，最终连通两个小岛得到面积为 3 的岛屿。
 *
 *
 * 示例 2:
 *
 *
 * 输入: grid = [[1, 1], [1, 0]]
 * 输出: 4
 * 解释: 将一格0变成1，岛屿的面积扩大为 4。
 *
 * 示例 3:
 *
 *
 * 输入: grid = [[1, 1], [1, 1]]
 * 输出: 4
 * 解释: 没有0可以让我们变成1，面积依然为 4。
 *
 *
 *
 * 提示：
 *
 *
 * n == grid.length
 * n == grid[i].length
 * 1 <= n <= 500
 * grid[i][j] 为 0 或 1
 *
 *
 */

// 并查集. 太恶心了...
var fa, size []int
var n, m int

var directions = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func query(x int) int {
	if x != fa[x] {
		fa[x] = query(fa[x])
	}
	return fa[x]
}

func merge(p, q int) {
	fp, fq := query(p), query(q)
	fa[fp] = fq
	size[fq] += size[fp]
	return
}

func largestIsland(G [][]int) int {
	n, m = len(G), len(G[0])
	ans := 0
	fa = make([]int, n*m)
	size = make([]int, n*m)
	visit := make([]bool, n*m)
	for i := 0; i < n*m; i++ {
		fa[i] = i
		size[i] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if G[i][j] == 0 {
				continue
			}
			if !visit[i*m+j] {
				dfs(G, visit, i, j)
			}
		}
	}
	for i := 0; i < n*m; i++ {
		ans = max(ans, size[query(i)])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if G[i][j] == 1 {
				continue
			}
			cur := 1
			fathers := make(map[int]struct{})
			for _, d := range directions {
				nx, ny := i+d[0], j+d[1]
				if nx < 0 || nx >= n || ny < 0 || ny >= m || G[nx][ny] == 0 {
					continue
				}
				fathers[query(nx*m+ny)] = struct{}{}
			}
			for f := range fathers {
				cur += size[f]
			}
			ans = max(ans, cur)
		}
	}
	return ans
}

func dfs(G [][]int, visit []bool, x, y int) {
	visit[x*m+y] = true
	for _, d := range directions {
		nx, ny := x+d[0], y+d[1]
		if nx < 0 || nx >= n || ny < 0 || ny >= m || G[nx][ny] == 0 || visit[nx*m+ny] {
			continue
		}
		merge(x*m+y, nx*m+ny)
		dfs(G, visit, nx, ny)
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}