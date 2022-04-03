/*
 * @lc app=leetcode.cn id=climbing-stairs lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [70] 爬楼梯
 *
 * https://leetcode-cn.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (53.54%)
 * Total Accepted:    752.8K
 * Total Submissions: 1.4M
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 * 
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 2
 * 输出：2
 * 解释：有两种方法可以爬到楼顶。
 * 1. 1 阶 + 1 阶
 * 2. 2 阶
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 3
 * 输出：3
 * 解释：有三种方法可以爬到楼顶。
 * 1. 1 阶 + 1 阶 + 1 阶
 * 2. 1 阶 + 2 阶
 * 3. 2 阶 + 1 阶
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 45
 * 
 * 
 */

/*
 * func climbStairs(n int) int {
 *     if n == 1 {
 *         return 1
 *     }
 *     dp := make([]int, n+1)
 *     dp[1] = 1
 *     dp[2] = 2
 *     for i := 3; i <= n; i++ {
 *         dp[i] = dp[i-1] + dp[i-2]
 *     }
 *     return dp[n]
 * }
 */


func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	a, b := 1, 1
	for n > 0 {
		a, b = b, a+b
		n--
	}
	return a

}
