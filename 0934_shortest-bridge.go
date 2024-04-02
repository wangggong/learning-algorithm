/*
 * @lc app=leetcode.cn id=shortest-bridge lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [934] 最短的桥
 *
 * https://leetcode.cn/problems/shortest-bridge/description/
 *
 * algorithms
 * Medium (48.02%)
 * Total Accepted:    53.9K
 * Total Submissions: 104.4K
 * Testcase Example:  '[[0,1],[1,0]]'
 *
 * 给你一个大小为 n x n 的二元矩阵 grid ，其中 1 表示陆地，0 表示水域。
 *
 * 岛 是由四面相连的 1 形成的一个最大组，即不会与非组内的任何其他 1 相连。grid 中 恰好存在两座岛 。
 *
 *
 *
 * 你可以将任意数量的 0 变为 1 ，以使两座岛连接起来，变成 一座岛 。
 *
 * 返回必须翻转的 0 的最小数目。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [[0,1],[1,0]]
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[0,1,0],[0,0,0],[0,0,1]]
 * 输出：2
 *
 *
 * 示例 3：
 *
 *
 * 输入：grid = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == grid.length == grid[i].length
 * 2 <= n <= 100
 * grid[i][j] 为 0 或 1
 * grid 中恰有两个岛
 *
 *
 */

var n int

func pos(x, y int) int { return x*n + y }

var id []int
var vis []bool

func query(k int) int {
	if k != id[k] {
		id[k] = query(id[k])
	}
	return id[k]
}

func merge(p, q int) { id[query(p)] = query(q) }

var directions = []int{0, 1, 0, -1, 0}

func shortestBridge(G [][]int) int {
	n = len(G)
	id = make([]int, n*n+1)
	vis = make([]bool, n*n+1)
	for i := range id {
		id[i] = i
	}
	start := -1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if G[i][j] != 0 {
				start = pos(i, j)
				break
			}
		}
	}
	var Q []int
	Q = append(Q, start)
	vis[start] = true
	for len(Q) > 0 {
		q := Q[0]
		Q = Q[1:]
		x, y := q/n, q%n
		for i := 0; i < 4; i++ {
			nx, ny := x+directions[i], y+directions[i+1]
			if k := pos(nx, ny); 0 <= nx && nx < n && 0 <= ny && ny < n && G[nx][ny] != 0 && !vis[k] {
				Q = append(Q, k)
				merge(q, k)
				vis[k] = true
			}
		}
	}
	vis = make([]bool, n*n+1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if k := pos(i, j); query(k) == query(start) {
				Q = append(Q, k)
				vis[k] = true
			}
		}
	}
	for level := 0; ; level++ {
		for size := len(Q); size > 0; size-- {
			q := Q[0]
			Q = Q[1:]
			x, y := q/n, q%n
			if G[x][y] != 0 && query(q) != query(start) {
				return level - 1
			}
			for i := 0; i < 4; i++ {
				nx, ny := x+directions[i], y+directions[i+1]
				if k := pos(nx, ny); 0 <= nx && nx < n && 0 <= ny && ny < n && query(k) != query(start) && !vis[k] {
					Q = append(Q, k)
					vis[k] = true
				}
			}
		}
	}
	return -1
}
