/*
 * @lc app=leetcode.cn id=reverse-string-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [541] 反转字符串 II
 *
 * https://leetcode-cn.com/problems/reverse-string-ii/description/
 *
 * algorithms
 * Easy (60.07%)
 * Total Accepted:    97.9K
 * Total Submissions: 163.8K
 * Testcase Example:  '"abcdefg"\n2'
 *
 * 给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。
 *
 *
 * 如果剩余字符少于 k 个，则将剩余字符全部反转。
 * 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "abcdefg", k = 2
 * 输出："bacdfeg"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "abcd", k = 2
 * 输出："bacd"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^4
 * s 仅由小写英文组成
 * 1 <= k <= 10^4
 *
 *
 */
func reverseStr(s string, k int) string {
	ans := []byte(s)
	n := len(ans)
	p, q := 0, k-1
	for q < n {
		for i, j := p, q; i < j; i, j = i+1, j-1 {
			ans[i], ans[j] = ans[j], ans[i]
		}
		p += 2 * k
		q += 2 * k
	}
	q = n - 1
	if n-p > k {
		q = p + k - 1
	}
	for i, j := p, q; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}
