/*
 * @lc app=leetcode.cn id=number-of-substrings-containing-all-three-characters lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1358] 包含所有三种字符的子字符串数目
 *
 * https://leetcode-cn.com/problems/number-of-substrings-containing-all-three-characters/description/
 *
 * algorithms
 * Medium (50.02%)
 * Total Accepted:    7K
 * Total Submissions: 13.8K
 * Testcase Example:  '"abcabc"'
 *
 * 给你一个字符串 s ，它只包含三种字符 a, b 和 c 。
 *
 * 请你返回 a，b 和 c 都 至少 出现过一次的子字符串数目。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "abcabc"
 * 输出：10
 * 解释：包含 a，b 和 c 各至少一次的子字符串为 "abc", "abca", "abcab", "abcabc", "bca", "bcab",
 * "bcabc", "cab", "cabc" 和 "abc" (相同字符串算多次)。
 *
 *
 * 示例 2：
 *
 * 输入：s = "aaacb"
 * 输出：3
 * 解释：包含 a，b 和 c 各至少一次的子字符串为 "aaacb", "aacb" 和 "acb" 。
 *
 *
 * 示例 3：
 *
 * 输入：s = "abc"
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= s.length <= 5 x 10^4
 * s 只包含字符 a，b 和 c 。
 *
 *
 */
func numberOfSubstrings(s string) int {
	arr := []byte(s)
	n := len(arr)
	p, q := 0, 0
	count := []int{0, 0, 0}
	ans := 0
	for ; q < n; q++ {
		count[int(arr[q]-'a')]++
		for count[0] > 0 && count[1] > 0 && count[2] > 0 {
			count[int(arr[p]-'a')]--
			p++
		}
		ans += p
	}
	return ans
}
