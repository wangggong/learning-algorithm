/*
 * @lc app=leetcode.cn id=dungeon-game lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [174] 地下城游戏
 *
 * https://leetcode-cn.com/problems/dungeon-game/description/
 *
 * algorithms
 * Hard (48.16%)
 * Total Accepted:    44.6K
 * Total Submissions: 92.5K
 * Testcase Example:  '[[-2,-3,3],[-5,-10,1],[10,30,-5]]'
 *
 *
 * table.dungeon, .dungeon th, .dungeon td {
 * ⁠ border:3px solid black;
 * }
 *
 * ⁠.dungeon th, .dungeon td {
 * ⁠   text-align: center;
 * ⁠   height: 70px;
 * ⁠   width: 70px;
 * }
 *
 *
 * 一些恶魔抓住了公主（P）并将她关在了地下城的右下角。地下城是由 M x N
 * 个房间组成的二维网格。我们英勇的骑士（K）最初被安置在左上角的房间里，他必须穿过地下城并通过对抗恶魔来拯救公主。
 *
 * 骑士的初始健康点数为一个正整数。如果他的健康点数在某一时刻降至 0 或以下，他会立即死亡。
 *
 * 有些房间由恶魔守卫，因此骑士在进入这些房间时会失去健康点数（若房间里的值为负整数，则表示骑士将损失健康点数）；其他房间要么是空的（房间里的值为
 * 0），要么包含增加骑士健康点数的魔法球（若房间里的值为正整数，则表示骑士将增加健康点数）。
 *
 * 为了尽快到达公主，骑士决定每次只向右或向下移动一步。
 *
 *
 *
 * 编写一个函数来计算确保骑士能够拯救到公主所需的最低初始健康点数。
 *
 * 例如，考虑到如下布局的地下城，如果骑士遵循最佳路径 右 -> 右 -> 下 -> 下，则骑士的初始健康点数至少为 7。
 *
 *
 * ⁠
 * -2 (K)
 * -3
 * 3
 * ⁠
 * ⁠
 * -5
 * -10
 * 1
 * ⁠
 * ⁠
 * 10
 * 30
 * -5 (P)
 * ⁠
 *
 *
 *
 *
 *
 * 说明:
 *
 *
 *
 * 骑士的健康点数没有上限。
 *
 * 任何房间都可能对骑士的健康点数造成威胁，也可能增加骑士的健康点数，包括骑士进入的左上角房间以及公主被监禁的右下角房间。
 *
 */

func calculateMinimumHP(G [][]int) int {
	n, m := len(G), len(G[0])
	dp := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		dp = append(dp, make([]int, m))
	}
	dp[n-1][m-1] = max(-G[n-1][m-1], 0) + 1
	for i := n - 2; i >= 0; i-- {
		dp[i][m-1] = max(dp[i+1][m-1]-G[i][m-1], 1)
	}
	for i := m - 2; i >= 0; i-- {
		dp[n-1][i] = max(dp[n-1][i+1]-G[n-1][i], 1)
	}
	for i := n - 2; i >= 0; i-- {
		for j := m - 2; j >= 0; j-- {
			dp[i][j] = min(
				max(dp[i+1][j]-G[i][j], 1),
				max(dp[i][j+1]-G[i][j], 1),
			)
		}
	}
	// fmt.Println(dp)
	return dp[0][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

