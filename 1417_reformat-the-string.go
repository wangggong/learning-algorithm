/*
 * @lc app=leetcode.cn id=reformat-the-string lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1417] 重新格式化字符串
 *
 * https://leetcode.cn/problems/reformat-the-string/description/
 *
 * algorithms
 * Easy (51.99%)
 * Total Accepted:    39K
 * Total Submissions: 70.3K
 * Testcase Example:  '"a0b1c2"'
 *
 * 给你一个混合了数字和字母的字符串 s，其中的字母均为小写英文字母。
 *
 * 请你将该字符串重新格式化，使得任意两个相邻字符的类型都不同。也就是说，字母后面应该跟着数字，而数字后面应该跟着字母。
 *
 * 请你返回 重新格式化后 的字符串；如果无法按要求重新格式化，则返回一个 空字符串 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "a0b1c2"
 * 输出："0a1b2c"
 * 解释："0a1b2c" 中任意两个相邻字符的类型都不同。 "a0b1c2", "0a1b2c", "0c2a1b" 也是满足题目要求的答案。
 *
 *
 * 示例 2：
 *
 * 输入：s = "leetcode"
 * 输出：""
 * 解释："leetcode" 中只有字母，所以无法满足重新格式化的条件。
 *
 *
 * 示例 3：
 *
 * 输入：s = "1229857369"
 * 输出：""
 * 解释："1229857369" 中只有数字，所以无法满足重新格式化的条件。
 *
 *
 * 示例 4：
 *
 * 输入：s = "covid2019"
 * 输出："c2o0v1i9d"
 *
 *
 * 示例 5：
 *
 * 输入：s = "ab123"
 * 输出："1a2b3"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 500
 * s 仅由小写英文字母和/或数字组成。
 *
 *
 */
func reformat(s string) string {
	var alpha, num []byte
	for _, b := range []byte(s) {
		if 'a' <= b && b <= 'z' {
			alpha = append(alpha, b)
		} else {
			num = append(num, b)
		}
	}
	n, m := len(alpha), len(num)
	if n > m+1 || m > n+1 {
		return ""
	}
	var ans []byte
	getAlpha := true
	if m >= n {
		getAlpha = false
	}
	for p, q := 0, 0; p < n || q < m; getAlpha = !getAlpha {
		if getAlpha {
			ans = append(ans, alpha[p])
			p++
		} else {
			ans = append(ans, num[q])
			q++
		}
	}
	return string(ans)
}
