/*
 * @lc app=leetcode.cn id=minimum-window-substring lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [76] 最小覆盖子串
 *
 * https://leetcode-cn.com/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (43.42%)
 * Total Accepted:    237.4K
 * Total Submissions: 546K
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 ""
 * 。
 *
 *
 *
 * 注意：
 *
 *
 * 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
 * 如果 s 中存在这样的子串，我们保证它是唯一的答案。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "ADOBECODEBANC", t = "ABC"
 * 输出："BANC"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "a", t = "a"
 * 输出："a"
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "a", t = "aa"
 * 输出: ""
 * 解释: t 中两个字符 'a' 均应包含在 s 的子串中，
 * 因此没有符合条件的子字符串，返回空字符串。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length, t.length <= 1e5
 * s 和 t 由英文字母组成
 *
 *
 *
 * 进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？
 */

var count [1 << 8]int
var visit [1 << 8]bool

// 输入：s = "ADOBECODEBANC", t = "ABC"
func minWindow(s string, t string) string {
	bs, bt := []byte(s), []byte(t)
	n := len(bs)
	for i := range count {
		count[i] = 0
		visit[i] = false
	}
	tot := 0
	for _, b := range bt {
		count[int(b)]++
		visit[int(b)] = true
		tot++
	}
	ans := []byte("")
	p, q := 0, 0
	for ; p < n; p++ {
		for ; q < n && tot > 0; q++ {
			index := int(bs[q])
			if !visit[index] {
				continue
			}
			// NOTE: 这里先更新欠债表: 如果更新后的欠债表仍不小于 0, 更新 tot.
			count[index]--
			if count[index] >= 0 {
				tot--
			}
		}
		if tot == 0 {
			cur := bs[p:q]
			if la, lc := len(ans), len(cur); la == 0 || la > lc {
				ans = make([]byte, lc)
				copy(ans, cur)
			}
		}
		// for ; p < n && tot == 0; p++ {
		index := int(bs[p])
		if visit[index] {
			// NOTE: 这里先更新 tot: 如果欠债表项不小于 0, 更新 tot.
			if count[index] >= 0 {
				tot++
			}
			count[index]++
		}
	}
	return string(ans)
}
