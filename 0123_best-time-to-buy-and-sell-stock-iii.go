/*
 * @lc app=leetcode.cn id=best-time-to-buy-and-sell-stock-iii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [123] 买卖股票的最佳时机 III
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/description/
 *
 * algorithms
 * Hard (54.38%)
 * Total Accepted:    150.2K
 * Total Submissions: 275.3K
 * Testcase Example:  '[3,3,5,0,0,3,1,4]'
 *
 * 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
 * 
 * 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
 * 
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入：prices = [3,3,5,0,0,3,1,4]
 * 输出：6
 * 解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
 * 随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
 * 
 * 示例 2：
 * 
 * 
 * 输入：prices = [1,2,3,4,5]
 * 输出：4
 * 解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4
 * 。   
 * 注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。   
 * 因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：prices = [7,6,4,3,1] 
 * 输出：0 
 * 解释：在这个情况下, 没有交易完成, 所以最大利润为 0。
 * 
 * 示例 4：
 * 
 * 
 * 输入：prices = [1]
 * 输出：0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= prices.length <= 1e5
 * 0 <= prices[i] <= 1e5
 * 
 * 
 */


// const MAXN int = 1e5

// var dp [4][MAXN]int
var dp [4]int

/*
 * func maxProfit(P []int) int {
 * 	n := len(P)
 * 	for i := 0; i < 4; i++ {
 * 		for j := 0; j < MAXN; j++ {
 * 			dp[i][j] = -MAXN
 * 		}
 * 	}
 * 	dp[0][0] = -P[0]
 * 	for i := 1; i < n; i++ {
 * 		dp[0][i] = max(dp[0][i-1], -P[i])
 * 		dp[1][i] = max(dp[1][i-1], P[i]+dp[0][i-1])
 * 		dp[2][i] = max(dp[2][i-1], dp[1][i-1] - P[i])
 * 		dp[3][i] = max(dp[3][i-1], P[i]+dp[2][i-1])
 * 	}
 * 	return max(max(dp[1][n-1], dp[3][n-1]), 0)
 * }
 */

func maxProfit(P []int) int {
	n := len(P)
	// assert n > 0;
	dp[0] = -P[0]
	dp[1] = 0
	// 傻逼, 第二次买入也需要初始化...
	dp[2] = -P[0]
	dp[3] = 0
	for i := 1; i < n; i++ {
		tmp := make([]int, 4)
		copy(tmp, dp[:])
		dp[0] = max(tmp[0], -P[i])
		dp[1] = max(tmp[1], P[i]+tmp[0])
		dp[2] = max(tmp[2], tmp[1] - P[i])
		dp[3] = max(tmp[3], P[i]+tmp[2])
	}
	return dp[3]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
