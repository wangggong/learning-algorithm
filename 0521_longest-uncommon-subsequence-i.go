/*
 * @lc app=leetcode.cn id=longest-uncommon-subsequence-i lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [521] 最长特殊序列 Ⅰ
 *
 * https://leetcode-cn.com/problems/longest-uncommon-subsequence-i/description/
 *
 * algorithms
 * Easy (71.35%)
 * Total Accepted:    28K
 * Total Submissions: 39.2K
 * Testcase Example:  '"aba"\n"cdc"'
 *
 * 给你两个字符串 a 和 b，请返回 这两个字符串中 最长的特殊序列  。如果不存在，则返回 -1 。
 * 
 * 「最长特殊序列」 定义如下：该序列为 某字符串独有的最长子序列（即不能是其他字符串的子序列） 。
 * 
 * 字符串 s 的子序列是在从 s 中删除任意数量的字符后可以获得的字符串。
 * 
 * 
 * 例如，“abc” 是 “aebdc” 的子序列，因为您可以删除 “aebdc” 中的下划线字符来得到 “abc” 。 “aebdc” 的子序列还包括
 * “aebdc” 、 “aeb” 和 “” (空字符串)。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入: a = "aba", b = "cdc"
 * 输出: 3
 * 解释: 最长特殊序列可为 "aba" (或 "cdc")，两者均为自身的子序列且不是对方的子序列。
 * 
 * 示例 2：
 * 
 * 
 * 输入：a = "aaa", b = "bbb"
 * 输出：3
 * 解释: 最长特殊序列是“aaa”和“bbb”。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：a = "aaa", b = "aaa"
 * 输出：-1
 * 解释: 字符串a的每个子序列也是字符串b的每个子序列。同样，字符串b的每个子序列也是字符串a的子序列。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= a.length, b.length <= 100
 * a 和 b 由小写英文字母组成
 * 
 * 
 */
func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	n, m := len(a), len(b)
	if n < m {
		return m
	}
	return n
}
