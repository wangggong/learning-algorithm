/*
 * @lc app=leetcode.cn id=powx-n lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [50] Pow(x, n)
 *
 * https://leetcode-cn.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (37.82%)
 * Total Accepted:    257.5K
 * Total Submissions: 680.7K
 * Testcase Example:  '2.00000\n10'
 *
 * 实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，x^n^ ）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：x = 2.00000, n = 10
 * 输出：1024.00000
 *
 *
 * 示例 2：
 *
 *
 * 输入：x = 2.10000, n = 3
 * 输出：9.26100
 *
 *
 * 示例 3：
 *
 *
 * 输入：x = 2.00000, n = -2
 * 输出：0.25000
 * 解释：2^-2 = 1/2^2 = 1/4 = 0.25
 *
 *
 *
 *
 * 提示：
 *
 *
 * -100.0 < x < 100.0
 * -2^31 <= n <= 2^31-1
 * -10^4 <= x^n <= 10^4
 *
 *
 */
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n == 1 {
		return x
	}
	// assert n > 1;
	if x == 0 {
		return 0
	}
	if x < 0 {
		if n&1 != 0 {
			return -myPow(-x, n)
		}
		return myPow(-x, n)
	}
	// assert x > 0;
	ans := float64(1)
	for n > 0 {
		if n&1 != 0 {
			ans *= x
		}
		// NOTE: 这里是底自己在平方
		x *= x
		n >>= 1
	}
	return ans
}
