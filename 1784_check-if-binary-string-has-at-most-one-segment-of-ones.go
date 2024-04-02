/*
 * @lc app=leetcode.cn id=check-if-binary-string-has-at-most-one-segment-of-ones lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1784] 检查二进制字符串字段
 *
 * https://leetcode.cn/problems/check-if-binary-string-has-at-most-one-segment-of-ones/description/
 *
 * algorithms
 * Easy (42.94%)
 * Total Accepted:    29.9K
 * Total Submissions: 52.7K
 * Testcase Example:  '"1001"'
 *
 * 给你一个二进制字符串 s ，该字符串 不含前导零 。
 *
 * 如果 s 包含 零个或一个由连续的 '1' 组成的字段 ，返回 true​​​ 。否则，返回 false 。
 *
 * 如果 s 中 由连续若干个 '1' 组成的字段 数量不超过 1，返回 true​​​ 。否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "1001"
 * 输出：false
 * 解释：由连续若干个 '1' 组成的字段数量为 2，返回 false
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "110"
 * 输出：true
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 100
 * s[i]​​​​ 为 '0' 或 '1'
 * s[0] 为 '1'
 *
 *
 */
func checkOnesSegment(s string) bool {
	cnt := 0
	for p, q, n := 0, 0, len(s); p < n; p = q {
		for ; p < n && s[p] != '1'; p++ {
		}
		if p == n {
			break
		}
		for q = p; q < n && s[q] == '1'; q++ {
		}
		cnt++
	}
	return cnt <= 1
}
