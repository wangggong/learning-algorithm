/*
 * @lc app=leetcode.cn id=number-of-dice-rolls-with-target-sum lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1155] 掷骰子的N种方法
 *
 * https://leetcode-cn.com/problems/number-of-dice-rolls-with-target-sum/description/
 *
 * algorithms
 * Medium (48.32%)
 * Total Accepted:    11.5K
 * Total Submissions: 23.8K
 * Testcase Example:  '1\n6\n3'
 *
 * 这里有 n 个一样的骰子，每个骰子上都有 k 个面，分别标号为 1 到 k 。
 *
 * 给定三个整数 n ,  k 和 target ，返回可能的方式(从总共 k^n 种方式中)滚动骰子的数量，使正面朝上的数字之和等于 target 。
 *
 * 答案可能很大，你需要对 10^9 + 7 取模 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 1, k = 6, target = 3
 * 输出：1
 * 解释：你扔一个有6张脸的骰子。
 * 得到3的和只有一种方法。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 2, k = 6, target = 7
 * 输出：6
 * 解释：你扔两个骰子，每个骰子有6个面。
 * 得到7的和有6种方法1+6 2+5 3+4 4+3 5+2 6+1。
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 30, k = 30, target = 500
 * 输出：222616187
 * 解释：返回的结果必须是对 10^9 + 7 取模。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n, k <= 30
 * 1 <= target <= 1000
 *
 *
 */

const MAXN = 1000 + 10
const MAXCNT = 30 + 5
const MOD = 1e9 + 7

// var dp [MAXCNT][MAXN]int
var dp [MAXN]int

/*
 * func numRollsToTarget(n int, k int, target int) int {
 * 	for i := 0; i < MAXCNT; i++ {
 * 		for j := 0; j < MAXN; j++ {
 * 			dp[i][j] = 0
 * 		}
 * 	}
 * 	for i := 1; i <= k; i++ {
 * 		dp[1][i] = 1
 * 	}
 * 	for i := 2; i <= n; i++ {
 * 		for j := 1; j <= target; j++ {
 * 			for v := 1; v <= k; v++ {
 * 				if j-v < 0 {
 * 					continue
 * 				}
 * 				dp[i][j] = (dp[i][j] + dp[i-1][j-v]) % MOD
 * 			}
 * 		}
 * 	}
 * 	return dp[n][target]
 * }
 */

func numRollsToTarget(n int, k int, target int) int {
	for i := range dp {
		dp[i] = 0
	}
	for i := 1; i <= k; i++ {
		dp[i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := target; j > 0; j-- {
			// NOTE: 每次不要忘记清零状态, 之前状态不能累加.
			dp[j] = 0
			for v := 1; v <= k; v++ {
				if j-v < 0 {
					continue
				}
				dp[j] = (dp[j] + dp[j-v]) % MOD
			}
		}
	}
	return dp[target]
}
