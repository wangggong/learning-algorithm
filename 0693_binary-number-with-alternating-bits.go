/*
 * @lc app=leetcode.cn id=binary-number-with-alternating-bits lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [693] 交替位二进制数
 *
 * https://leetcode-cn.com/problems/binary-number-with-alternating-bits/description/
 *
 * algorithms
 * Easy (60.87%)
 * Total Accepted:    29.2K
 * Total Submissions: 48K
 * Testcase Example:  '5'
 *
 * 给定一个正整数，检查它的二进制表示是否总是 0、1 交替出现：换句话说，就是二进制表示中相邻两位的数字永不相同。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 5
 * 输出：true
 * 解释：5 的二进制表示是：101
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 7
 * 输出：false
 * 解释：7 的二进制表示是：111.
 * 
 * 示例 3：
 * 
 * 
 * 输入：n = 11
 * 输出：false
 * 解释：11 的二进制表示是：1011.
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 2^31 - 1
 * 
 * 
 */
func hasAlternatingBits(n int) bool {
	return bits(1+(n ^ (n >> 1))) == 1
}

func bits(n int) int {
	ans := 0
	for n != 0 {
		ans++
		n &= n-1
	}
	return ans
}
