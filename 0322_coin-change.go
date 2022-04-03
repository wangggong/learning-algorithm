/*
 * @lc app=leetcode.cn id=coin-change lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [322] 零钱兑换
 *
 * https://leetcode-cn.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (45.23%)
 * Total Accepted:    394.7K
 * Total Submissions: 872.4K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
 *
 * 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
 *
 * 你可以认为每种硬币的数量是无限的。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：coins = [1, 2, 5], amount = 11
 * 输出：3
 * 解释：11 = 5 + 5 + 1
 *
 * 示例 2：
 *
 *
 * 输入：coins = [2], amount = 3
 * 输出：-1
 *
 * 示例 3：
 *
 *
 * 输入：coins = [1], amount = 0
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= coins.length <= 12
 * 1 <= coins[i] <= 2^31 - 1
 * 0 <= amount <= 10^4
 *
 *
 */
import "math"

func coinChange(C []int, target int) int {
	dp := make([]int, target+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for _, c := range C {
		for j := 0; j <= target; j++ {
			if j >= c {
				dp[j] = min(dp[j], dp[j-c]+1)
			}
		}
	}
	if dp[target] == math.MaxInt32 {
		return -1
	}
	return dp[target]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
