/*
 * @lc app=leetcode.cn id=longest-repeating-character-replacement lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [424] 替换后的最长重复字符
 *
 * https://leetcode-cn.com/problems/longest-repeating-character-replacement/description/
 *
 * algorithms
 * Medium (53.45%)
 * Total Accepted:    58K
 * Total Submissions: 108.4K
 * Testcase Example:  '"ABAB"\n2'
 *
 * 给你一个字符串 s 和一个整数 k 。你可以选择字符串中的任一字符，并将其更改为任何其他大写英文字符。该操作最多可执行 k 次。
 *
 * 在执行上述操作后，返回包含相同字母的最长子字符串的长度。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "ABAB", k = 2
 * 输出：4
 * 解释：用两个'A'替换为两个'B',反之亦然。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "AABABBA", k = 1
 * 输出：4
 * 解释：
 * 将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
 * 子串 "BBBB" 有最长重复字母, 答案为 4。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^5
 * s 仅由大写英文字母组成
 * 0 <= k <= s.length
 *
 *
 */
func characterReplacement(s string, k int) int {
	bytes := []byte(s)
	n := len(bytes)
	p, q := 0, 0
	count := make([]int, 1<<8)
	ans := 0
	index := byte('A')
	for ; q < n; q++ {
		count[int(bytes[q])]++
		// 最多的字符数 **仅** 取决于 `q` 指针.
		if count[int(bytes[q])] > count[int(index)] {
			index = bytes[q]
		}
		// 在当前窗口中与最多数目字符不同的字符数目 = 窗口长度 - 最多数目字符的数目.
		if (q-p+1)-count[int(index)] <= k {
			if ans < q-p+1 {
				ans = q - p + 1
			}
		} else {
			count[int(bytes[p])]--
			p++
		}
		// fmt.Printf("%v %v %v %v\n", p, q, diff, ans)
	}
	return ans
}
