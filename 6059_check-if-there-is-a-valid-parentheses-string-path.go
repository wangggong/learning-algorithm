/*
 * Hard
 *
 * 一个括号字符串是一个 非空 且只包含 '(' 和 ')' 的字符串。如果下面 任意 条件为 真 ，那么这个括号字符串就是 合法的 。
 *
 * 字符串是 () 。
 * 字符串可以表示为 AB（A 连接 B），A 和 B 都是合法括号序列。
 * 字符串可以表示为 (A) ，其中 A 是合法括号序列。
 * 给你一个 m x n 的括号网格图矩阵 grid 。网格图中一个 合法括号路径 是满足以下所有条件的一条路径：
 *
 * 路径开始于左上角格子 (0, 0) 。
 * 路径结束于右下角格子 (m - 1, n - 1) 。
 * 路径每次只会向 下 或者向 右 移动。
 * 路径经过的格子组成的括号字符串是 合法 的。
 * 如果网格图中存在一条 合法括号路径 ，请返回 true ，否则返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：grid = [["(","(","("],[")","(",")"],["(","(",")"],["(","(",")"]]
 * 输出：true
 * 解释：上图展示了两条路径，它们都是合法括号字符串路径。
 * 第一条路径得到的合法字符串是 "()(())" 。
 * 第二条路径得到的合法字符串是 "((()))" 。
 * 注意可能有其他的合法括号字符串路径。
 * 示例 2：
 *
 *
 *
 * 输入：grid = [[")",")"],["(","("]]
 * 输出：false
 * 解释：两条可行路径分别得到 "))(" 和 ")((" 。由于它们都不是合法括号字符串，我们返回 false 。
 *
 *
 * 提示：
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 100
 * grid[i][j] 要么是 '(' ，要么是 ')' 。
 */

const maxn int = 100

var directions = [][2]int{
	{0, 1},
	{1, 0},
}

var M [maxn + 10][maxn + 10]map[int]struct{}

// BFS. 注意小心爆内存, 只需要存一下已经出现的三元状态就可以了.
func hasValidPath(G [][]byte) bool {
	n, m := len(G), len(G[0])
	// assert n > 0 && m > 0
	if G[0][0] == ')' || G[n-1][m-1] == '(' {
		return false
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			M[i][j] = make(map[int]struct{})
		}
	}
	var Q [][3]int
	Q = append(Q, [3]int{0, 0, 1})
	M[0][0][1] = struct{}{}
	for len(Q) > 0 {
		q := Q[0]
		Q = Q[1:]
		y, x, c := q[0], q[1], q[2]
		if y == n-1 && x == m-1 && c == 0 {
			return true
		}
		for _, d := range directions {
			ny, nx, nc := y+d[0], x+d[1], c
			if ny >= n || nx >= m {
				continue
			}
			if G[ny][nx] == '(' {
				nc++
			} else {
				nc--
			}
			if nc < 0 {
				continue
			}
			if _, ok := M[ny][nx][nc]; ok {
				continue
			}
			Q = append(Q, [3]int{ny, nx, nc})
			M[ny][nx][nc] = struct{}{}
		}
	}
	return false
}
