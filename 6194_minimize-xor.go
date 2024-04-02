/*
 * @lc app=leetcode.cn id=minimize-xor lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6194] 最小 XOR
 *
 * https://leetcode.cn/problems/minimize-xor/description/
 *
 * algorithms
 * Medium (38.07%)
 * Total Accepted:    4.2K
 * Total Submissions: 10.9K
 * Testcase Example:  '3\n5'
 *
 * 给你两个正整数 num1 和 num2 ，找出满足下述条件的整数 x ：
 *
 *
 * x 的置位数和 num2 相同，且
 * x XOR num1 的值 最小
 *
 *
 * 注意 XOR 是按位异或运算。
 *
 * 返回整数 x 。题目保证，对于生成的测试用例， x 是 唯一确定 的。
 *
 * 整数的 置位数 是其二进制表示中 1 的数目。
 *
 *
 *
 * 示例 1：
 *
 * 输入：num1 = 3, num2 = 5
 * 输出：3
 * 解释：
 * num1 和 num2 的二进制表示分别是 0011 和 0101 。
 * 整数 3 的置位数与 num2 相同，且 3 XOR 3 = 0 是最小的。
 *
 *
 * 示例 2：
 *
 * 输入：num1 = 1, num2 = 12
 * 输出：3
 * 解释：
 * num1 和 num2 的二进制表示分别是 0001 和 1100 。
 * 整数 3 的置位数与 num2 相同，且 3 XOR 1 = 2 是最小的。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= num1, num2 <= 10^9
 *
 *
 */
func minimizeXor(n1, n2 int) (ans int) {
	c := cnt1(n2)
	for d := 31; d >= 0; d-- {
		if n1&(1<<d) != 0 {
			if c > 0 {
				ans |= 1 << d
				c--
			}
		}
	}
	for d := 0; d <= 31; d++ {
		if n1&(1<<d) == 0 {
			if c > 0 {
				ans |= 1 << d
				c--
			}
		}
	}
	return
}

func cnt1(n int) (ans int) {
	for ; n > 0; n &= n - 1 {
		ans++
	}
	return
}
