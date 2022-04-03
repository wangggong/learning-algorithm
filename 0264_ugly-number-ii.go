import "math"

/*
 * @lc app=leetcode.cn id=ugly-number-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [264] 丑数 II
 *
 * https://leetcode-cn.com/problems/ugly-number-ii/description/
 *
 * algorithms
 * Medium (58.28%)
 * Total Accepted:    117.6K
 * Total Submissions: 201.7K
 * Testcase Example:  '10'
 *
 * 给你一个整数 n ，请你找出并返回第 n 个 丑数 。
 *
 * 丑数 就是只包含质因数 2、3 和/或 5 的正整数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 10
 * 输出：12
 * 解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：1
 * 解释：1 通常被视为丑数。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 1690
 *
 *
 */

// 优先队列显然.
/*
 * func nthUglyNumber(n int) int {
 * 	primes := [3]int{2, 3, 5}
 * 	candidates := [3]int{2, 3, 5}
 * 	indexes := [3]int{}
 * 	dp := make([]int, n)
 * 	dp[0] = 1
 * 	for i := 1; i < n; i++ {
 * 		dp[i] = math.MaxInt32
 * 		for _, c := range candidates {
 * 			if c < dp[i] {
 * 				dp[i] = c
 * 			}
 * 		}
 * 		// NOTE: 更新每个值. 只要与当前结果相同就往后推!
 * 		for j := 0; j < 3; j++ {
 * 			if candidates[j] == dp[i] {
 * 				indexes[j]++
 * 				ind := indexes[j]
 * 				candidates[j] = dp[ind] * primes[j]
 * 			}
 * 		}
 * 	}
 * 	return dp[n-1]
 * }
 */

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	primes := [3]int{2, 3, 5}
	indexes := [3]int{}
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt32
		for j := 0; j < 3; j++ {
			if c := primes[j] * dp[indexes[j]]; dp[i] > c {
				dp[i] = c
			}
		}
		for j := 0; j < 3; j++ {
			if c := primes[j] * dp[indexes[j]]; dp[i] == c {
				indexes[j]++
			}
		}
	}
	return dp[n-1]
}
