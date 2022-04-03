/*
 * @lc app=leetcode.cn id=best-time-to-buy-and-sell-stock-iv lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [188] 买卖股票的最佳时机 IV
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/description/
 *
 * algorithms
 * Hard (40.79%)
 * Total Accepted:    113.9K
 * Total Submissions: 279.2K
 * Testcase Example:  '2\n[2,4,1]'
 *
 * 给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
 *
 * 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
 *
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：k = 2, prices = [2,4,1]
 * 输出：2
 * 解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
 *
 * 示例 2：
 *
 *
 * 输入：k = 2, prices = [3,2,6,5,0,3]
 * 输出：7
 * 解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
 * ⁠    随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 =
 * 3 。
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= k <= 100
 * 0 <= prices.length <= 1000
 * 0 <= prices[i] <= 1000
 *
 *
 */
func maxProfit(k int, P []int) int {
	n := len(P)
	if n == 0 {
		return 0
	}
	if k == 0 {
		return 0
	}
	dp := make([]int, k*2)
	for i := 0; i < k; i++ {
		dp[i*2] = -P[0]
	}
	for i := 1; i < n; i++ {
		tmp := make([]int, k*2)
		copy(tmp, dp[:])
		for j := 0; j < k; j++ {
			if j == 0 {
				dp[j*2] = max(tmp[j*2], -P[i])
			} else {
				dp[j*2] = max(tmp[j*2], dp[j*2-1]-P[i])
			}
			dp[j*2+1] = max(tmp[j*2+1], P[i]+tmp[j*2])
		}
	}
	return dp[2*k-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
