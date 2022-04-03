/*
 * @lc app=leetcode.cn id=best-time-to-buy-and-sell-stock-with-transaction-fee lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [714] 买卖股票的最佳时机含手续费
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/
 *
 * algorithms
 * Medium (73.45%)
 * Total Accepted:    130.9K
 * Total Submissions: 178.1K
 * Testcase Example:  '[1,3,2,8,4,9]\n2'
 *
 * 给定一个整数数组 prices，其中 prices[i]表示第 i 天的股票价格 ；整数 fee 代表了交易股票的手续费用。
 *
 * 你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
 *
 * 返回获得利润的最大值。
 *
 * 注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：prices = [1, 3, 2, 8, 4, 9], fee = 2
 * 输出：8
 * 解释：能够达到的最大利润:
 * 在此处买入 prices[0] = 1
 * 在此处卖出 prices[3] = 8
 * 在此处买入 prices[4] = 4
 * 在此处卖出 prices[5] = 9
 * 总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8
 *
 * 示例 2：
 *
 *
 * 输入：prices = [1,3,7,5,10,3], fee = 3
 * 输出：6
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= prices.length <= 5 * 10^4
 * 1 <= prices[i] < 5 * 10^4
 * 0 <= fee < 5 * 10^4
 *
 *
 */

const maxn int = 5e4

var dp [maxn + 5][2]int

/*
 * // DP
 * func maxProfit(P []int, fee int) int {
 * 	n := len(P)
 * 	for i := 0; i < n; i++ {
 * 		dp[i][0] = 0
 * 		dp[i][1] = 0
 * 	}
 * 	dp[0][0] = -P[0]
 * 	for i := 1; i < n; i++ {
 * 		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-P[i])
 * 		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+P[i]-fee)
 * 	}
 * 	return dp[n-1][1]
 * }
 */

/*
 * func maxProfit(P []int, fee int) int {
 * 	n := len(P)
 * 	// assert n > 0;
 * 	buy, sell := -P[0], 0
 * 	for i := 1; i < n; i++ {
 * 		b, s := buy, sell
 * 		buy = max(b, s-P[i])
 * 		sell = max(s, b+P[i]-fee)
 * 	}
 * 	return sell
 * }
 */

// 贪心, 确实不如 DP 好懂.
func maxProfit(P []int, fee int) int {
	n := len(P)
	buy := P[0]
	ans := 0
	for i := 1; i < n; i++ {
		// 正式买入, 更新最低价.
		if buy > P[i] {
			buy = P[i]
		}
		// 如果卖出亏本就不如不卖, 但其实没这段逻辑不影响.
		// if P[i] > buy && P[i] <= buy + fee {
		// 	continue
		// }
		// 如果有利可图, 卖出去的同时更新最低价到 `P[i] - fee` (更低的价格要求更高的利润, 然后就能退化为基本的股神模型.)
		if P[i]-buy > fee {
			ans += P[i] - buy - fee
			buy = P[i] - fee
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
