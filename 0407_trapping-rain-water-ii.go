/*
 * @lc app=leetcode.cn id=trapping-rain-water-ii lang=golang
 *
 * [407] 接雨水 II
 *
 * https://leetcode-cn.com/problems/trapping-rain-water-ii/description/
 *
 * algorithms
 * Hard (58.02%)
 * Total Accepted:    25.6K
 * Total Submissions: 44.2K
 * Testcase Example:  '[[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]'
 *
 * 给你一个 m x n 的矩阵，其中的值均为非负整数，代表二维高度图每个单元的高度，请计算图中形状最多能接多少体积的雨水。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入: heightMap = [[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]
 * 输出: 4
 * 解释: 下雨后，雨水将会被上图蓝色的方块中。总的接雨水量为1+2+1=4。
 *
 *
 * 示例 2:
 *
 *
 *
 *
 * 输入: heightMap =
 * [[3,3,3,3,3],[3,2,2,2,3],[3,2,1,2,3],[3,2,2,2,3],[3,3,3,3,3]]
 * 输出: 10
 *
 *
 *
 *
 * 提示:
 *
 *
 * m == heightMap.length
 * n == heightMap[i].length
 * 1 <= m, n <= 200
 * 0 <= heightMap[i][j] <= 2 * 10^4
 *
 *
 *
 *
 */
import "container/heap"

const MAXN = 200 + 10
const DIRS = 4

type cell struct{ h, x, y int }

type hp struct {
	C []cell
}

func (h hp) Len() int            { return len(h.C) }
func (h hp) Less(p, q int) bool  { return h.C[p].h < h.C[q].h }
func (h hp) Swap(p, q int)       { h.C[p], h.C[q] = h.C[q], h.C[p] }
func (h *hp) Push(c interface{}) { h.C = append(h.C, c.(cell)) }
func (h *hp) Pop() interface{}   { v := h.C[len(h.C)-1]; h.C = h.C[:len(h.C)-1]; return v }

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func trapRainWater(H [][]int) int {
	n, m := len(H), len(H[0])
	if n <= 2 || m <= 2 {
		return 0
	}

	ans := 0

	var vis [MAXN][MAXN]bool
	h := &hp{}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || i == n-1 || j == 0 || j == m-1 {
				heap.Push(h, cell{H[i][j], i, j})
				vis[i][j] = true
			}
		}
	}

	dirs := [DIRS][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	for h.Len() > 0 {
		c := heap.Pop(h).(cell)
		for i := 0; i < DIRS; i++ {
			nx, ny := c.x+dirs[i][0], c.y+dirs[i][1]
			if 0 <= nx && nx < n && 0 <= ny && ny < m && !vis[nx][ny] {
				if H[nx][ny] < c.h {
					ans += c.h - H[nx][ny]
				}
				vis[nx][ny] = true
				heap.Push(h, cell{max(H[nx][ny], c.h), nx, ny})
			}
		}
	}

	return ans
}
