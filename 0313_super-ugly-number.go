/*
 * @lc app=leetcode.cn id=super-ugly-number lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [313] 超级丑数
 *
 * https://leetcode-cn.com/problems/super-ugly-number/description/
 *
 * algorithms
 * Medium (60.63%)
 * Total Accepted:    45.8K
 * Total Submissions: 76.2K
 * Testcase Example:  '12\n[2,7,13,19]'
 *
 * 超级丑数 是一个正整数，并满足其所有质因数都出现在质数数组 primes 中。
 *
 * 给你一个整数 n 和一个整数数组 primes ，返回第 n 个 超级丑数 。
 *
 * 题目数据保证第 n 个 超级丑数 在 32-bit 带符号整数范围内。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 12, primes = [2,7,13,19]
 * 输出：32
 * 解释：给定长度为 4 的质数数组 primes = [2,7,13,19]，前 12
 * 个超级丑数序列为：[1,2,4,7,8,13,14,16,19,26,28,32] 。
 *
 * 示例 2：
 *
 *
 * 输入：n = 1, primes = [2,3,5]
 * 输出：1
 * 解释：1 不含质因数，因此它的所有质因数都在质数数组 primes = [2,3,5] 中。
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 10^6
 * 1 <= primes.length <= 100
 * 2 <= primes[i] <= 1000
 * 题目数据 保证 primes[i] 是一个质数
 * primes 中的所有值都 互不相同 ，且按 递增顺序 排列
 *
 *
 *
 *
 *
 */

import "math"

const MAXN int = 1e6 + 10

var dp [MAXN]int

// 优先队列超时 / 超空间了, revenge!
func nthSuperUglyNumber(k int, primes []int) int {
	n := len(primes)
	indexes := make([]int, n)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 1
	for i := 0; i < k; i++ {
		for j := 0; j < n; j++ {
			if c := primes[j] * dp[indexes[j]]; c < dp[i] {
				dp[i] = c
			}
		}
		for j := 0; j < n; j++ {
			if c := primes[j] * dp[indexes[j]]; c == dp[i] {
				indexes[j]++
			}
		}
	}
	return dp[k-1]
}
