/*
 * @lc app=leetcode.cn id=k-similar-strings lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [854] 相似度为 K 的字符串
 *
 * https://leetcode.cn/problems/k-similar-strings/description/
 *
 * algorithms
 * Hard (37.15%)
 * Total Accepted:    17.4K
 * Total Submissions: 39.1K
 * Testcase Example:  '"ab"\n"ba"'
 *
 * 对于某些非负整数 k ，如果交换 s1 中两个字母的位置恰好 k 次，能够使结果字符串等于 s2 ，则认为字符串 s1 和 s2 的 相似度为 k 。
 *
 * 给你两个字母异位词 s1 和 s2 ，返回 s1 和 s2 的相似度 k 的最小值。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s1 = "ab", s2 = "ba"
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 * 输入：s1 = "abc", s2 = "bca"
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s1.length <= 20
 * s2.length == s1.length
 * s1 和 s2  只包含集合 {'a', 'b', 'c', 'd', 'e', 'f'} 中的小写字母
 * s2 是 s1 的一个字母异位词
 *
 *
 */
// 贪逼版 BFS, 疯狂剪枝
func kSimilarity(s1 string, s2 string) (ans int) {
	n := len(s1)
	type info struct {
		s string
		k int
	}
	var Q []info
	m := make(map[string]struct{})
	Q = append(Q, info{s1, 0})
	m[s1] = struct{}{}
	for ; len(Q) > 0; ans++ {
		for size := len(Q); size > 0; size-- {
			q := Q[0]
			Q = Q[1:]
			if q.s == s2 {
				return
			}
			i := q.k
			for ; i < n && q.s[i] == s2[i]; i++ {
			}
			for j := i + 1; j < n; j++ {
				if q.s[j] != s2[i] {
					continue
				}
				if q.s[j] == s2[j] {
					continue
				}
				bs := []byte(q.s)
				bs[i], bs[j] = bs[j], bs[i]
				ns := string(bs)
				if _, ok := m[ns]; ok {
					continue
				}
				Q = append(Q, info{ns, i + 1})
				m[ns] = struct{}{}
			}
		}
	}
	// unreachable
	return -1
}
