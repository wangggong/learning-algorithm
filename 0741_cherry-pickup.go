/*
 * @lc app=leetcode.cn id=cherry-pickup lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [741] 摘樱桃
 *
 * https://leetcode-cn.com/problems/cherry-pickup/description/
 *
 * algorithms
 * Hard (37.76%)
 * Total Accepted:    4.8K
 * Total Submissions: 12.6K
 * Testcase Example:  '[[0,1,-1],[1,0,-1],[1,1,1]]'
 *
 * 一个N x N的网格(grid) 代表了一块樱桃地，每个格子由以下三种数字的一种来表示：
 *
 *
 * 0 表示这个格子是空的，所以你可以穿过它。
 * 1 表示这个格子里装着一个樱桃，你可以摘到樱桃然后穿过它。
 * -1 表示这个格子里有荆棘，挡着你的路。
 *
 *
 * 你的任务是在遵守下列规则的情况下，尽可能的摘到最多樱桃：
 *
 *
 * 从位置 (0, 0) 出发，最后到达 (N-1, N-1) ，只能向下或向右走，并且只能穿越有效的格子（即只可以穿过值为0或者1的格子）；
 * 当到达 (N-1, N-1) 后，你要继续走，直到返回到 (0, 0) ，只能向上或向左走，并且只能穿越有效的格子；
 * 当你经过一个格子且这个格子包含一个樱桃时，你将摘到樱桃并且这个格子会变成空的（值变为0）；
 * 如果在 (0, 0) 和 (N-1, N-1) 之间不存在一条可经过的路径，则没有任何一个樱桃能被摘到。
 *
 *
 * 示例 1:
 *
 *
 * 输入: grid =
 * [[0, 1, -1],
 * ⁠[1, 0, -1],
 * ⁠[1, 1,  1]]
 * 输出: 5
 * 解释：
 * 玩家从（0,0）点出发，经过了向下走，向下走，向右走，向右走，到达了点(2, 2)。
 * 在这趟单程中，总共摘到了4颗樱桃，矩阵变成了[[0,1,-1],[0,0,-1],[0,0,0]]。
 * 接着，这名玩家向左走，向上走，向上走，向左走，返回了起始点，又摘到了1颗樱桃。
 * 在旅程中，总共摘到了5颗樱桃，这是可以摘到的最大值了。
 *
 *
 * 说明:
 *
 *
 * grid 是一个 N * N 的二维数组，N的取值范围是1 <= N <= 50。
 * 每一个 grid[i][j] 都是集合 {-1, 0, 1}其中的一个数。
 * 可以保证起点 grid[0][0] 和终点 grid[N-1][N-1] 的值都不会是 -1。
 *
 *
 */

import "math"

const maxn int = 50

var dp [maxn + 10][maxn + 10][maxn + 10]int
var visit [maxn + 10][maxn + 10][maxn + 10]bool
var n, m int

/*
 * func cherryPickup(G [][]int) int {
 * 	n, m = len(G), len(G[0])
 * 	for i := 0; i < maxn+10; i++ {
 * 		for j := 0; j < maxn+10; j++ {
 * 			for k := 0; k < maxn+10; k++ {
 * 				dp[i][j][k] = math.MinInt32
 * 				visit[i][j][k] = false
 * 			}
 * 		}
 * 	}
 * 	ans := dfs(G, 0, 0, 0)
 * 	if ans < 0 {
 * 		return 0
 * 	}
 * 	return ans
 * }
 */

func cherryPickup(G [][]int) int {
	n := len(G)
	dp := make([][][]int, 2*n-1)
	for k := range dp {
		dp[k] = make([][]int, n+1)
		for x1 := range dp[k] {
			dp[k][x1] = make([]int, n+1)
		}
	}
	for k := range dp {
		for x1 := range dp[k] {
			for x2 := range dp[k][x1] {
				dp[k][x1][x2] = math.MinInt32
			}
		}
	}
	dp[0][1][1] = G[0][0]
	for k := 1; k <= 2*(n-1); k++ {
		for x1 := 0; x1 <= k; x1++ {
			if x1 >= n || k-x1 >= n || G[x1][k-x1] == -1 {
				continue
			}
			for x2 := 0; x2 <= k; x2++ {
				if x2 >= n || k-x2 >= n || G[x2][k-x2] == -1 {
					continue
				}
				c := G[x1][k-x1]
				if x1 != x2 {
					c += G[x2][k-x2]
				}
				dp[k][x1+1][x2+1] = max(
					max(dp[k-1][x1][x2], dp[k-1][x1+1][x2+1]),
					max(dp[k-1][x1+1][x2], dp[k-1][x1][x2+1]),
				) + c
			}
		}
	}
	return max(dp[2*(n-1)][n][n], 0)
}

func dfs(G [][]int, x1, y, x2 int) int {
	y1 := y
	y2 := x1 + y - x2
	if x1 < 0 || x1 >= n || x2 < 0 || x2 >= n || y1 < 0 || y1 >= m || y2 < 0 || y2 >= m {
		return math.MinInt32
	}
	if G[x1][y1] < 0 || G[x2][y2] < 0 {
		return math.MinInt32
	}
	if x1 == n-1 && y1 == m-1 {
		return G[x1][y1]
	}
	if visit[x1][y][x2] {
		return dp[x1][y][x2]
	}
	ans := math.MinInt32
	defer func() {
		dp[x1][y][x2] = ans
		visit[x1][y][x2] = true
	}()
	ans = max(
		max(dfs(G, x1+1, y, x2), dfs(G, x1, y+1, x2+1)),
		max(dfs(G, x1+1, y, x2+1), dfs(G, x1, y+1, x2)),
	)
	if ans < 0 {
		return ans
	}
	ans += G[x1][y1]
	if x1 != x2 {
		ans += G[x2][y2]
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
