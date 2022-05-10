/*
 * @lc app=leetcode.cn id=bricks-falling-when-hit lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [803] 打砖块
 *
 * https://leetcode-cn.com/problems/bricks-falling-when-hit/description/
 *
 * algorithms
 * Hard (47.68%)
 * Total Accepted:    13K
 * Total Submissions: 27.4K
 * Testcase Example:  '[[1,0,0,0],[1,1,1,0]]\n[[1,0]]'
 *
 * 有一个 m x n 的二元网格 grid ，其中 1 表示砖块，0 表示空白。砖块 稳定（不会掉落）的前提是：
 *
 *
 * 一块砖直接连接到网格的顶部，或者
 * 至少有一块相邻（4 个方向之一）砖块 稳定 不会掉落时
 *
 *
 * 给你一个数组 hits ，这是需要依次消除砖块的位置。每当消除 hits[i] = (rowi, coli)
 * 位置上的砖块时，对应位置的砖块（若存在）会消失，然后其他的砖块可能因为这一消除操作而 掉落 。一旦砖块掉落，它会 立即 从网格 grid
 * 中消失（即，它不会落在其他稳定的砖块上）。
 *
 * 返回一个数组 result ，其中 result[i] 表示第 i 次消除操作对应掉落的砖块数目。
 *
 * 注意，消除可能指向是没有砖块的空白位置，如果发生这种情况，则没有砖块掉落。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [[1,0,0,0],[1,1,1,0]], hits = [[1,0]]
 * 输出：[2]
 * 解释：网格开始为：
 * [[1,0,0,0]，
 * ⁠[1,1,1,0]]
 * 消除 (1,0) 处加粗的砖块，得到网格：
 * [[1,0,0,0]
 * ⁠[0,1,1,0]]
 * 两个加粗的砖不再稳定，因为它们不再与顶部相连，也不再与另一个稳定的砖相邻，因此它们将掉落。得到网格：
 * [[1,0,0,0],
 * ⁠[0,0,0,0]]
 * 因此，结果为 [2] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[1,0,0,0],[1,1,0,0]], hits = [[1,1],[1,0]]
 * 输出：[0,0]
 * 解释：网格开始为：
 * [[1,0,0,0],
 * ⁠[1,1,0,0]]
 * 消除 (1,1) 处加粗的砖块，得到网格：
 * [[1,0,0,0],
 * ⁠[1,0,0,0]]
 * 剩下的砖都很稳定，所以不会掉落。网格保持不变：
 * [[1,0,0,0],
 * ⁠[1,0,0,0]]
 * 接下来消除 (1,0) 处加粗的砖块，得到网格：
 * [[1,0,0,0],
 * ⁠[0,0,0,0]]
 * 剩下的砖块仍然是稳定的，所以不会有砖块掉落。
 * 因此，结果为 [0,0] 。
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 200
 * grid[i][j] 为 0 或 1
 * 1 <= hits.length <= 4 * 10^4
 * hits[i].length == 2
 * 0 <= xi <= m - 1
 * 0 <= yi <= n - 1
 * 所有 (xi, yi) 互不相同
 *
 *
 */
// 并查集+逆向思维. 照着答案刷不一定能刷对的玩意, 细节暴多.
const (
	blank = iota
	brick
)

var n, m, k int
var directions = [4 + 1]int{0, 1, 0, -1, 0}

type UF struct {
	fa   []int
	size []int
}

func NewUF(n int) *UF {
	uf := &UF{
		fa:   make([]int, n+1),
		size: make([]int, n+1),
	}
	for i := 0; i <= n; i++ {
		uf.fa[i], uf.size[i] = i, 1
	}
	return uf
}

func (uf *UF) Query(x int) int {
	if x != uf.fa[x] {
		uf.fa[x] = uf.Query(uf.fa[x])
	}
	return uf.fa[x]
}

func (uf *UF) Merge(p, q int) {
	fp, fq := uf.Query(p), uf.Query(q)
	// 如果已经在同一个联通域, 不做处理.
	if fp != fq {
		uf.fa[fp] = fq
		uf.size[fq] += uf.size[fp]
	}
	return
}

func (uf *UF) Size(x int) int {
	return uf.size[uf.Query(x)]
}

func index(x, y int) int {
	return x*m + y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func hitBricks(grid [][]int, hits [][]int) []int {
	n, m, k = len(grid), len(grid[0]), len(hits)
	src := n * m
	uf := NewUF(src)
	G := make([][]int, n)
	for i := range G {
		G[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			G[i][j] = grid[i][j]
		}
	}
	// 先把所有被打下来的砖块置空.
	for _, h := range hits {
		x, y := h[0], h[1]
		G[x][y] = blank
	}
	// 建立并查集.
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if G[i][j] == blank {
				continue
			}
			if i == 0 {
				uf.Merge(j, src)
				continue
			}
			if G[i-1][j] == brick {
				uf.Merge(index(i-1, j), index(i, j))
			}
			if j > 0 && G[i][j-1] == brick {
				uf.Merge(index(i, j-1), index(i, j))
			}
		}
	}
	// 反向填充结果.
	ans := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		x, y := hits[i][0], hits[i][1]
		// 如果该位置本来就没砖块, 跳过.
		if grid[x][y] == blank {
			continue
		}
		// 先记录一下屋顶 `src` 有多少砖块.
		prev := uf.Size(src)
		if x == 0 {
			uf.Merge(y, src)
		}
		for j := 0; j < 4; j++ {
			nx, ny := x+directions[j], y+directions[j+1]
			if nx < 0 || nx >= n || ny < 0 || ny >= m || G[nx][ny] == blank {
				continue
			}
			uf.Merge(index(nx, ny), index(x, y))
		}
		// 再记录一下填充砖块后屋顶 `src` 有多少砖块.
		curr := uf.Size(src)
		ans[i] = max(0, curr-prev-1)
		// 别忘了把砖块填上.
		G[x][y] = brick
	}
	return ans
}
