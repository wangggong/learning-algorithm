/*
 * @lc app=leetcode.cn id=prime-number-of-set-bits-in-binary-representation lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [762] 二进制表示中质数个计算置位
 *
 * https://leetcode-cn.com/problems/prime-number-of-set-bits-in-binary-representation/description/
 *
 * algorithms
 * Easy (70.48%)
 * Total Accepted:    19.5K
 * Total Submissions: 27.6K
 * Testcase Example:  '6\n10'
 *
 * 给定两个整数 L 和 R ，找到闭区间 [L, R] 范围内，计算置位位数为质数的整数个数。
 *
 * （注意，计算置位代表二进制表示中1的个数。例如 21 的二进制表示 10101 有 3 个计算置位。还有，1 不是质数。）
 *
 * 示例 1:
 *
 *
 * 输入: L = 6, R = 10
 * 输出: 4
 * 解释:
 * 6 -> 110 (2 个计算置位，2 是质数)
 * 7 -> 111 (3 个计算置位，3 是质数)
 * 9 -> 1001 (2 个计算置位，2 是质数)
 * 10-> 1010 (2 个计算置位，2 是质数)
 *
 *
 * 示例 2:
 *
 *
 * 输入: L = 10, R = 15
 * 输出: 5
 * 解释:
 * 10 -> 1010 (2 个计算置位, 2 是质数)
 * 11 -> 1011 (3 个计算置位, 3 是质数)
 * 12 -> 1100 (2 个计算置位, 2 是质数)
 * 13 -> 1101 (3 个计算置位, 3 是质数)
 * 14 -> 1110 (3 个计算置位, 3 是质数)
 * 15 -> 1111 (4 个计算置位, 4 不是质数)
 *
 *
 * 注意:
 *
 *
 * L, R 是 L <= R 且在 [1, 10^6] 中的整数。
 * R - L 的最大值为 10000。
 *
 *
 */
var primes = map[int]struct{}{2: struct{}{}, 3: struct{}{}, 5: struct{}{}, 7: struct{}{}, 11: struct{}{}, 13: struct{}{}, 17: struct{}{}, 19: struct{}{}, 23: struct{}{}, 29: struct{}{}, 31: struct{}{}}

func countPrimeSetBits(left int, right int) int {
	ans := 0
	for k := left; k <= right; k++ {
		if _, ok := primes[bits(k)]; ok {
			ans++
		}
	}
	return ans
}

func bits(n int) int {
	ans := 0
	for n != 0 {
		ans++
		n &= n - 1
	}
	return ans
}
