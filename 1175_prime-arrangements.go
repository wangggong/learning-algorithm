/*
 * @lc app=leetcode.cn id=prime-arrangements lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1175] 质数排列
 *
 * https://leetcode.cn/problems/prime-arrangements/description/
 *
 * algorithms
 * Easy (49.42%)
 * Total Accepted:    14.8K
 * Total Submissions: 27.3K
 * Testcase Example:  '5'
 *
 * 请你帮忙给从 1 到 n 的数设计排列方案，使得所有的「质数」都应该被放在「质数索引」（索引从 1 开始）上；你需要返回可能的方案总数。
 *
 * 让我们一起来回顾一下「质数」：质数一定是大于 1 的，并且不能用两个小于它的正整数的乘积来表示。
 *
 * 由于答案可能会很大，所以请你返回答案 模 mod 10^9 + 7 之后的结果即可。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 5
 * 输出：12
 * 解释：举个例子，[1,2,5,4,3] 是一个有效的排列，但 [5,2,3,4,1] 不是，因为在第二种情况里质数 5 被错误地放在索引为 1
 * 的位置上。
 *
 *
 * 示例 2：
 *
 * 输入：n = 100
 * 输出：682289015
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 100
 *
 *
 */

const mod int = 1e9 + 7

func numPrimeArrangements(n int) int {
	k := cntPrime(n)
	return ((A(k) % mod) * (A(n-k) % mod)) % mod
}

func cntPrime(n int) int {
	if n < 2 {
		return 0
	}
	ans := 1
	for i := 3; i <= n; i++ {
		if isPrime(i) {
			ans++
		}
	}
	return ans
}

func isPrime(n int) bool {
	if n%2 == 0 {
		return n == 2
	}
	for i := 3; i <= n/i; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func A(n int) int {
	ans := 1
	for i := 1; i <= n; i++ {
		ans = (ans * i) % mod
	}
	return ans
}
