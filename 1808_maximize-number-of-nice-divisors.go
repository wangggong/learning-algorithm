/*
 * @lc app=leetcode.cn id=maximize-number-of-nice-divisors lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1808] 好因子的最大数目
 *
 * https://leetcode-cn.com/problems/maximize-number-of-nice-divisors/description/
 *
 * algorithms
 * Hard (27.18%)
 * Total Accepted:    2.7K
 * Total Submissions: 9.8K
 * Testcase Example:  '5'
 *
 * 给你一个正整数 primeFactors 。你需要构造一个正整数 n ，它满足以下条件：
 *
 *
 * n 质因数（质因数需要考虑重复的情况）的数目 不超过 primeFactors 个。
 * n 好因子的数目最大化。如果 n 的一个因子可以被 n 的每一个质因数整除，我们称这个因子是 好因子 。比方说，如果 n = 12 ，那么它的质因数为
 * [2,2,3] ，那么 6 和 12 是好因子，但 3 和 4 不是。
 *
 *
 * 请你返回 n 的好因子的数目。由于答案可能会很大，请返回答案对 10^9 + 7 取余 的结果。
 *
 * 请注意，一个质数的定义是大于 1 ，且不能被分解为两个小于该数的自然数相乘。一个数 n 的质因子是将 n 分解为若干个质因子，且它们的乘积为 n
 * 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：primeFactors = 5
 * 输出：6
 * 解释：200 是一个可行的 n 。
 * 它有 5 个质因子：[2,2,2,5,5] ，且有 6 个好因子：[10,20,40,50,100,200] 。
 * 不存在别的 n 有至多 5 个质因子，且同时有更多的好因子。
 *
 *
 * 示例 2：
 *
 *
 * 输入：primeFactors = 8
 * 输出：18
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= primeFactors <= 1e9
 *
 *
 */
const mod int = 1e9 + 7

// 参考: https://leetcode-cn.com/problems/maximize-number-of-nice-divisors/solution/fan-yi-wan-zhi-hou-jiu-xiang-dang-yu-yua-113z/
//
// 找规律做出来了, 但确实需要正本清源.
func maxNiceDivisors(primeFactors int) int {
	ans := 1
	if primeFactors <= 4 {
		return primeFactors
	}
	// assert primeFactors > 4
	k := (primeFactors - 2) / 3
	ans = pow(3, k, mod)
	// fmt.Println(primeFactors, k, ans)
	primeFactors -= k * 3
	ans = (ans * primeFactors) % mod
	return ans
}

// 快速幂可还行...
func pow(n, k, mod int) int {
	ans := 1
	cur := n
	for ; k > 0; k >>= 1 {
		if k&1 != 0 {
			ans = (ans * cur) % mod
		}
		cur = (cur * cur) % mod
	}
	return ans
}

/*
1 -> 2
2 -> 2*2 = 4
3 -> 2*2*2 = 8
4 -> 2*2*2*2 = 16
5 -> 2*2*2*5*5 = 200
6 -> 2*2*2*5*5*5 = 1000
7 -> 2*2*2*2*5*5*5 = 2000
8 -> 2*2*2*3*3*3*5*5 = 1600
*/
