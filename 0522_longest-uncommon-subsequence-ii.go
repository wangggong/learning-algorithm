/*
 * @lc app=leetcode.cn id=longest-uncommon-subsequence-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [522] 最长特殊序列 II
 *
 * https://leetcode.cn/problems/longest-uncommon-subsequence-ii/description/
 *
 * algorithms
 * Medium (37.53%)
 * Total Accepted:    10.1K
 * Total Submissions: 26.6K
 * Testcase Example:  '["aba","cdc","eae"]'
 *
 * 给定字符串列表 strs ，返回其中 最长的特殊序列 。如果最长特殊序列不存在，返回 -1 。
 *
 * 特殊序列 定义如下：该序列为某字符串 独有的子序列（即不能是其他字符串的子序列）。
 *
 * s 的 子序列可以通过删去字符串 s 中的某些字符实现。
 *
 *
 * 例如，"abc" 是 "aebdc" 的子序列，因为您可以删除"aebdc"中的下划线字符来得到 "abc"
 * 。"aebdc"的子序列还包括"aebdc"、 "aeb" 和 "" (空字符串)。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: strs = ["aba","cdc","eae"]
 * 输出: 3
 *
 *
 * 示例 2:
 *
 *
 * 输入: strs = ["aaa","aaa","aa"]
 * 输出: -1
 *
 *
 *
 *
 * 提示:
 *
 *
 * 2 <= strs.length <= 50
 * 1 <= strs[i].length <= 10
 * strs[i] 只包含小写英文字母
 *
 *
 */
func findLUSlength(strs []string) int {
	var index [51][]int
	for i, s := range strs {
		index[len(s)] = append(index[len(s)], i)
	}
	var cand []int
	for l := 50; l >= 0; l-- {
		cand = append(cand, index[l]...)
		for _, i := range index[l] {
			valid := true
			for _, j := range cand {
				if i == j {
					continue
				}
				if isSubseq(strs[i], strs[j]) {
					valid = false
					break
				}
			}
			if valid {
				return l
			}
		}
	}
	return -1
}

func isSubseq(s, t string) bool {
	for i, j, n, m := 0, 0, len(s), len(t); j < m; j++ {
		if s[i] == t[j] {
			i++
		}
		if i == n {
			return true
		}
	}
	return false
}
