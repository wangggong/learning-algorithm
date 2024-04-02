/*
 * @lc app=leetcode.cn id=minimum-obstacle-removal-to-reach-corner lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6081] 到达角落需要移除障碍物的最小数目
 *
 * https://leetcode.cn/problems/minimum-obstacle-removal-to-reach-corner/description/
 *
 * algorithms
 * Hard (43.41%)
 * Total Accepted:    1.9K
 * Total Submissions: 4.5K
 * Testcase Example:  '[[0,1,1],[1,1,0],[1,1,0]]'
 *
 * 给你一个下标从 0 开始的二维整数数组 grid ，数组大小为 m x n 。每个单元格都是两个值之一：
 *
 *
 * 0 表示一个 空 单元格，
 * 1 表示一个可以移除的 障碍物 。
 *
 *
 * 你可以向上、下、左、右移动，从一个空单元格移动到另一个空单元格。
 *
 * 现在你需要从左上角 (0, 0) 移动到右下角 (m - 1, n - 1) ，返回需要移除的障碍物的 最小 数目。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：grid = [[0,1,1],[1,1,0],[1,1,0]]
 * 输出：2
 * 解释：可以移除位于 (0, 1) 和 (0, 2) 的障碍物来创建从 (0, 0) 到 (2, 2) 的路径。
 * 可以证明我们至少需要移除两个障碍物，所以返回 2 。
 * 注意，可能存在其他方式来移除 2 个障碍物，创建出可行的路径。
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：grid = [[0,1,0,0,0],[0,1,0,1,0],[0,0,0,1,0]]
 * 输出：0
 * 解释：不移除任何障碍物就能从 (0, 0) 到 (2, 4) ，所以返回 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 10^5
 * 2 <= m * n <= 10^5
 * grid[i][j] 为 0 或 1
 * grid[0][0] == grid[m - 1][n - 1] == 0
 *
 *
 */

import "container/heap"

type Heap [][3]int

func (h Heap) Len() int              { return len(h) }
func (h Heap) Less(p, q int) bool    { return h[p][2] < h[q][2] }
func (h Heap) Swap(p, q int)         { h[p], h[q] = h[q], h[p]; return }
func (h *Heap) Push(v interface{})   { *h = append(*h, v.([3]int)) }
func (h *Heap) Pop() (v interface{}) { l := len(*h); v = (*h)[l-1]; *h = (*h)[:l-1]; return }

var directions = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func minimumObstacles(G [][]int) int {
	n, m := len(G), len(G[0])
	visit := make([][]bool, n)
	for i := range visit {
		visit[i] = make([]bool, m)
	}
	Q := &Heap{}
	heap.Push(Q, [3]int{0, 0, 0})
	for Q.Len() > 0 {
		q := heap.Pop(Q).([3]int)
		x, y, d := q[0], q[1], q[2]
		if x == n-1 && y == m-1 {
			return d
		}
		if visit[x][y] {
			continue
		}
		visit[x][y] = true
		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || nx >= n || ny < 0 || ny >= m {
				continue
			}
			nd := d
			if G[nx][ny] == 1 {
				nd++
			}
			heap.Push(Q, [3]int{nx, ny, nd})
		}
	}
	// unreachable
	return n*m + 1
}
