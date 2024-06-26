/*
 * @lc app=leetcode.cn id=cut-off-trees-for-golf-event lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [675] 为高尔夫比赛砍树
 *
 * https://leetcode-cn.com/problems/cut-off-trees-for-golf-event/description/
 *
 * algorithms
 * Hard (41.53%)
 * Total Accepted:    6K
 * Total Submissions: 12.5K
 * Testcase Example:  '[[1,2,3],[0,0,4],[7,6,5]]'
 *
 * 你被请来给一个要举办高尔夫比赛的树林砍树。树林由一个 m x n 的矩阵表示， 在这个矩阵中：
 *
 *
 * 0 表示障碍，无法触碰
 * 1 表示地面，可以行走
 * 比 1 大的数 表示有树的单元格，可以行走，数值表示树的高度
 *
 *
 * 每一步，你都可以向上、下、左、右四个方向之一移动一个单位，如果你站的地方有一棵树，那么你可以决定是否要砍倒它。
 *
 * 你需要按照树的高度从低向高砍掉所有的树，每砍过一颗树，该单元格的值变为 1（即变为地面）。
 *
 * 你将从 (0, 0) 点开始工作，返回你砍完所有树需要走的最小步数。 如果你无法砍完所有的树，返回 -1 。
 *
 * 可以保证的是，没有两棵树的高度是相同的，并且你至少需要砍倒一棵树。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：forest = [[1,2,3],[0,0,4],[7,6,5]]
 * 输出：6
 * 解释：沿着上面的路径，你可以用 6 步，按从最矮到最高的顺序砍掉这些树。
 *
 * 示例 2：
 *
 *
 * 输入：forest = [[1,2,3],[0,0,0],[7,6,5]]
 * 输出：-1
 * 解释：由于中间一行被障碍阻塞，无法访问最下面一行中的树。
 *
 *
 * 示例 3：
 *
 *
 * 输入：forest = [[2,3,4],[0,0,5],[8,7,6]]
 * 输出：6
 * 解释：可以按与示例 1 相同的路径来砍掉所有的树。
 * (0,0) 位置的树，可以直接砍去，不用算步数。
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == forest.length
 * n == forest[i].length
 * 1 <= m, n <= 50
 * 0 <= forest[i][j] <= 109
 *
 *
 */
import "sort"

type pos struct {
	h, x, y int
}

var n, m int
var trees []pos
var directions = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func bfs(G [][]int, pre, cur pos) int {
	visit := make([][]bool, n)
	for i := range visit {
		visit[i] = make([]bool, m)
	}
	var Q []pos
	Q = append(Q, pre)
	visit[pre.x][pre.y] = true
	d := 0
	for len(Q) > 0 {
		for size := len(Q); size > 0; size-- {
			q := Q[0]
			Q = Q[1:]
			if q.x == cur.x && q.y == cur.y {
				G[cur.x][cur.y] = 1
				return d
			}
			for _, dir := range directions {
				nx, ny := q.x+dir[0], q.y+dir[1]
				if nx < 0 || nx >= n || ny < 0 || ny >= m || visit[nx][ny] || G[nx][ny] == 0 {
					continue
				}
				Q = append(Q, pos{G[nx][ny], nx, ny})
				visit[nx][ny] = true
			}
		}
		d++
	}
	return -1
}

func cutOffTree(forest [][]int) int {
	n, m = len(forest), len(forest[0])
	trees = nil
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if forest[i][j] > 1 {
				trees = append(trees, pos{forest[i][j], i, j})
			}
		}
	}
	sort.Slice(trees, func(p, q int) bool { return trees[p].h < trees[q].h })
	ans := 0
	prePos := pos{forest[0][0], 0, 0}
	for _, p := range trees {
		dist := bfs(forest, prePos, p)
		if dist < 0 {
			return -1
		}
		ans += dist
		prePos = p
	}
	return ans
}
