/*
 * Medium
 *
 * 给定两个字符串 s 和 t，判断他们的编辑距离是否为 1。
 *
 * 注意：
 *
 * 满足编辑距离等于 1 有三种可能的情形：
 *
 * 往 s 中插入一个字符得到 t
 * 从 s 中删除一个字符得到 t
 * 在 s 中替换一个字符得到 t
 * 示例 1：
 *
 * 输入: s = "ab", t = "acb"
 * 输出: true
 * 解释: 可以将 'c' 插入字符串 s 来得到 t。
 * 示例 2:
 *
 * 输入: s = "cab", t = "ad"
 * 输出: false
 * 解释: 无法通过 1 步操作使 s 变为 t。
 * 示例 3:
 *
 * 输入: s = "1203", t = "1213"
 * 输出: true
 * 解释: 可以将字符串 s 中的 '0' 替换为 '1' 来得到 t。
 * 通过次数9,313
 * 提交次数27,593
 */

func isOneEditDistance(s string, t string) bool {
	if s == t {
		return false
	}
	n, m := len(s), len(t)
	if n-m > 1 || m-n > 1 {
		return false
	}
	cntDiff := n - m
	oneEdit := false
	for p, q := 0, 0; p < n && q < m; p, q = p+1, q+1 {
		if s[p] == t[q] {
			continue
		}
		if oneEdit {
			return false
		}
		if cntDiff == 0 {
			oneEdit = true
			continue
		}
		if cntDiff == -1 {
			if q+1 < m && s[p] == t[q+1] {
				oneEdit = true
				q++
				continue
			}
			return false
		}
		if cntDiff == 1 {
			if p+1 < n && s[p+1] == t[q] {
				oneEdit = true
				p++
				continue
			}
			return false
		}
	}
	return true
}
