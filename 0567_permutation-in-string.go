/*
 * @lc app=leetcode.cn id=permutation-in-string lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [567] 字符串的排列
 *
 * https://leetcode-cn.com/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (43.63%)
 * Total Accepted:    158.8K
 * Total Submissions: 364.1K
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
 *
 * 换句话说，s1 的排列之一是 s2 的 子串 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s1 = "ab" s2 = "eidbaooo"
 * 输出：true
 * 解释：s2 包含 s1 的排列之一 ("ba").
 *
 *
 * 示例 2：
 *
 *
 * 输入：s1= "ab" s2 = "eidboaoo"
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s1.length, s2.length <= 10^4
 * s1 和 s2 仅包含小写字母
 *
 *
 */

var count [1 << 8]int
var visit [1 << 8]bool

func checkInclusion(s1 string, s2 string) bool {
	bs1, bs2 := []byte(s1), []byte(s2)
	n := len(bs2)
	for i := range count {
		count[i] = 0
		visit[i] = false
	}
	tot := 0
	for _, b := range bs1 {
		count[int(b)]++
		visit[int(b)] = true
		tot++
	}
	for p, q := 0, 0; p < n; p++ {
		for ; q < n && tot > 0; q++ {
			index := int(bs2[q])
			if visit[index] {
				tot--
			}
			count[index]--
		}
		allZero := true
		for _, c := range count {
			if c != 0 {
				allZero = false
				break
			}
		}
		if allZero {
			return true
		}
		index := int(bs2[p])
		if visit[index] {
			tot++
		}
		count[index]++
	}
	return false
}
