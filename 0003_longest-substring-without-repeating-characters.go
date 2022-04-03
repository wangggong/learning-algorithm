/*
 * @lc app=leetcode.cn id=longest-substring-without-repeating-characters lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (38.48%)
 * Total Accepted:    1.5M
 * Total Submissions: 4M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 5 * 10^4
 * s 由英文字母、数字、符号和空格组成
 *
 *
 */

var visit [1 << 8]bool

/*
 * func lengthOfLongestSubstring(s string) int {
 * 	bs := []byte(s)
 * 	n := len(bs)
 * 	for i := range visit {
 * 		visit[i] = false
 * 	}
 * 	p, q := 0, 0
 * 	ans := 0
 * 	for ; p < n; p++ {
 * 		for ; q < n; q++ {
 * 			idx := int(bs[q])
 * 			if visit[idx] {
 * 				break
 * 			}
 * 			visit[idx] = true
 * 		}
 * 		if q-p > ans {
 * 			ans = q - p
 * 		}
 * 		// NOTE: 诶你特么复原 visit 的时候能不能小心点?
 * 		idx := int(bs[p])
 * 		visit[idx] = false
 * 	}
 * 	return ans
 * }
 */

func lengthOfLongestSubstring(s string) int {
	bytes := []byte(s)
	n := len(bytes)
	var count [1 << 8]int
	p, q := 0, 0
	ans := 0
	for ; q < n; q++ {
		ind := int(bytes[q])
		count[ind]++
		for ; count[ind] > 1; p++ {
			count[int(bytes[p])]--
		}
		if q-p+1 > ans {
			ans = q - p + 1
		}
	}
	return ans
}
