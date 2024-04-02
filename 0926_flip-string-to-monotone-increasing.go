/*
 * @lc app=leetcode.cn id=flip-string-to-monotone-increasing lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [926] 将字符串翻转到单调递增
 *
 * https://leetcode.cn/problems/flip-string-to-monotone-increasing/description/
 *
 * algorithms
 * Medium (55.87%)
 * Total Accepted:    15.1K
 * Total Submissions: 25.3K
 * Testcase Example:  '"00110"'
 *
 * 如果一个二进制字符串，是以一些 0（可能没有 0）后面跟着一些 1（也可能没有 1）的形式组成的，那么该字符串是 单调递增 的。
 *
 * 给你一个二进制字符串 s，你可以将任何 0 翻转为 1 或者将 1 翻转为 0 。
 *
 * 返回使 s 单调递增的最小翻转次数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "00110"
 * 输出：1
 * 解释：翻转最后一位得到 00111.
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "010110"
 * 输出：2
 * 解释：翻转得到 011111，或者是 000111。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "00011000"
 * 输出：2
 * 解释：翻转得到 00000000。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^5
 * s[i] 为 '0' 或 '1'
 *
 *
 */
func minFlipsMonoIncr(s string) int {
	n := len(s)
	ps0, ps1, ps0n, ps1n := 0, 0, 0, 0
	for _, b := range s {
		if b == '0' {
			ps0n++
		} else {
			ps1n++
		}
	}
	ans := min(ps0n, ps1n)
	for i, b := range s {
		if b == '0' {
			ps0++
		} else {
			ps1++
		}
		ans = min(ans, (i-ps0)+((n-i)-(ps1n-ps1)))
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
