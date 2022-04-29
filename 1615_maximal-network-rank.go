/*
 * @lc app=leetcode.cn id=maximal-network-rank lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1615] 最大网络秩
 *
 * https://leetcode-cn.com/problems/maximal-network-rank/description/
 *
 * algorithms
 * Medium (53.59%)
 * Total Accepted:    7K
 * Total Submissions: 13.1K
 * Testcase Example:  '4\n[[0,1],[0,3],[1,2],[1,3]]'
 *
 * n 座城市和一些连接这些城市的道路 roads 共同组成一个基础设施网络。每个 roads[i] = [ai, bi] 都表示在城市 ai 和 bi
 * 之间有一条双向道路。
 *
 * 两座不同城市构成的 城市对 的 网络秩 定义为：与这两座城市 直接 相连的道路总数。如果存在一条道路直接连接这两座城市，则这条道路只计算 一次 。
 *
 * 整个基础设施网络的 最大网络秩 是所有不同城市对中的 最大网络秩 。
 *
 * 给你整数 n 和数组 roads，返回整个基础设施网络的 最大网络秩 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：n = 4, roads = [[0,1],[0,3],[1,2],[1,3]]
 * 输出：4
 * 解释：城市 0 和 1 的网络秩是 4，因为共有 4 条道路与城市 0 或 1 相连。位于 0 和 1 之间的道路只计算一次。
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：n = 5, roads = [[0,1],[0,3],[1,2],[1,3],[2,3],[2,4]]
 * 输出：5
 * 解释：共有 5 条道路与城市 1 或 2 相连。
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 8, roads = [[0,1],[1,2],[2,3],[2,4],[5,6],[5,7]]
 * 输出：5
 * 解释：2 和 5 的网络秩为 5，注意并非所有的城市都需要连接起来。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= n <= 100
 * 0 <= roads.length <= n * (n - 1) / 2
 * roads[i].length == 2
 * 0 <= a_i, b_i <= n-1
 * a_i != b_i
 * 每对城市之间 最多只有一条 道路相连
 *
 *
 */
const maxn int = 100

var G [maxn + 5][maxn + 5]int
var count [maxn + 5]int

func maximalNetworkRank(n int, roads [][]int) int {
	for i := 0; i < maxn+5; i++ {
		count[i] = 0
		for j := 0; j < maxn+5; j++ {
			G[i][j] = 0
		}
	}
	for _, r := range roads {
		if r[0] > r[1] {
			r[0], r[1] = r[1], r[0]
		}
		count[r[0]]++
		count[r[1]]++
		G[r[0]][r[1]]++
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ans = max(ans, count[i]+count[j]-G[i][j])
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
