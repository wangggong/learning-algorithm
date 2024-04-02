/*
 * @lc app=leetcode.cn id=maximum-number-of-points-with-cost lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1937] 扣分后的最大得分
 *
 * https://leetcode-cn.com/problems/maximum-number-of-points-with-cost/description/
 *
 * algorithms
 * Medium (26.43%)
 * Total Accepted:    4K
 * Total Submissions: 15.1K
 * Testcase Example:  '[[1,2,3],[1,5,1],[3,1,1]]'
 *
 * 给你一个 m x n 的整数矩阵 points （下标从 0 开始）。一开始你的得分为 0 ，你想最大化从矩阵中得到的分数。
 *
 * 你的得分方式为：每一行 中选取一个格子，选中坐标为 (r, c) 的格子会给你的总得分 增加 points[r][c] 。
 *
 * 然而，相邻行之间被选中的格子如果隔得太远，你会失去一些得分。对于相邻行 r 和 r + 1 （其中 0 ），选中坐标为 (r, c1) 和 (r +
 * 1, c2) 的格子，你的总得分 减少 abs(c1 - c2) 。
 *
 * 请你返回你能得到的 最大 得分。
 *
 * abs(x) 定义为：
 *
 *
 * 如果 x >= 0 ，那么值为 x 。
 * 如果 x  ，那么值为 -x 。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：points = [[1,2,3],[1,5,1],[3,1,1]]
 * 输出：9
 * 解释：
 * 蓝色格子是最优方案选中的格子，坐标分别为 (0, 2)，(1, 1) 和 (2, 0) 。
 * 你的总得分增加 3 + 5 + 3 = 11 。
 * 但是你的总得分需要扣除 abs(2 - 1) + abs(1 - 0) = 2 。
 * 你的最终得分为 11 - 2 = 9 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：points = [[1,5],[2,3],[4,2]]
 * 输出：11
 * 解释：
 * 蓝色格子是最优方案选中的格子，坐标分别为 (0, 1)，(1, 1) 和 (2, 0) 。
 * 你的总得分增加 5 + 3 + 4 = 12 。
 * 但是你的总得分需要扣除 abs(1 - 1) + abs(1 - 0) = 1 。
 * 你的最终得分为 12 - 1 = 11 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == points.length
 * n == points[r].length
 * 1 <= m, n <= 1e5
 * 1 <= m * n <= 1e5
 * 0 <= points[r][c] <= 1e5
 *
 *
 */

// 参考: https://leetcode.cn/problems/maximum-number-of-points-with-cost/solution/dp-you-hua-ji-qiao-chai-xiang-qian-hou-z-5vvc/
//
// 针对之前 `O(nm^2)` 的 dp 进行了拆项优化. 奇怪的知识增加了.jpg
func maxPoints(P [][]int) int64 {
	n, m := len(P), len(P[0])
	dp := make([]int, m)
	copy(dp, P[0])
	for i := 1; i < n; i++ {
        /*
         * dp[j] = max_{k=0}^{m}{P[i][j] + dp[k] - abs(k-j)
         *       = max \left\{\begin{array}{rcl}
         *       P[i][j] + dp[k] + k - j, $    $ {k < j} \\
         *       P[i][j] + dp[k] - k + j, $    $ {k >= j} \\
         *       \end{array}\right.
         *       = max \left\{\begin{array}{rcl}
         *       (P[i][j] - j) + (dp[k] + k), $    $ {k < j} \\
         *       (P[i][j] + j) + (dp[k] - k), $    $ {k >= j} \\
         *       \end{array}\right.
         *       = max(P[i][j] - j + left[j], P[i][j] + j + right[j])
         * Where
         * left[j] = max_{k=0}^{j-1}{dp[k] + k},
         * right[j] = max_{k=j}^{m}{dp[k] - k}.
         */
		left, right := make([]int, m), make([]int, m)
		left[0] = dp[0]
		for j := 1; j < m; j++ {
			left[j] = max(dp[j]+j, left[j-1])
		}
		right[m-1] = dp[m-1] - (m - 1)
		for j := m - 2; j >= 0; j-- {
			right[j] = max(dp[j]-j, right[j+1])
		}
		for j := 0; j < m; j++ {
			dp[j] = dp[j] + P[i][j]
			dp[j] = max(dp[j], P[i][j]-j+left[j])
			dp[j] = max(dp[j], P[i][j]+j+right[j])
		}
	}
	ans := int64(0)
	for _, d := range dp {
		if ans < int64(d) {
			ans = int64(d)
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

/*
 * // 暴力 `O(nm^2)` 不出意料的超时了.
 * func maxPoints(P [][]int) int64 {
 * 	n, m := len(P), len(P[0])
 * 	dp := make([]int64, m)
 * 	for i, p := range P[0] {
 * 		dp[i] = int64(p)
 * 	}
 * 	for i := 1; i < n; i++ {
 * 		pre := make([]int64, m)
 * 		copy(pre, dp)
 * 		for j := 0; j < m; j++ {
 * 			for k := 0; k < m; k++ {
 * 				dp[j] = max(dp[j], pre[k]+int64(P[i][j])-int64(abs(k-j)))
 * 			}
 * 		}
 * 	}
 * 	ans := int64(0)
 * 	for _, d := range dp {
 * 		ans = max(ans, d)
 * 	}
 * 	return ans
 * }
 *
 * func abs(x int) int {
 * 	if x >= 0 {
 * 		return x
 * 	}
 * 	return -x
 * }
 */
