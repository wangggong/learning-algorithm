/*
 * @lc app=leetcode.cn id=unique-substrings-in-wraparound-string lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [467] 环绕字符串中唯一的子字符串
 *
 * https://leetcode-cn.com/problems/unique-substrings-in-wraparound-string/description/
 *
 * algorithms
 * Medium (44.86%)
 * Total Accepted:    17.2K
 * Total Submissions: 35.3K
 * Testcase Example:  '"a"'
 *
 * 把字符串 s 看作是 “abcdefghijklmnopqrstuvwxyz” 的无限环绕字符串，所以 s 看起来是这样的：
 *
 *
 * "...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd...." .
 *
 *
 * 现在给定另一个字符串 p 。返回 s 中 唯一 的 p 的 非空子串 的数量 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: p = "a"
 * 输出: 1
 * 解释: 字符串 s 中只有一个"a"子字符。
 *
 *
 * 示例 2:
 *
 *
 * 输入: p = "cac"
 * 输出: 2
 * 解释: 字符串 s 中的字符串“cac”只有两个子串“a”、“c”。.
 *
 *
 * 示例 3:
 *
 *
 * 输入: p = "zab"
 * 输出: 6
 * 解释: 在字符串 s 中有六个子串“z”、“a”、“b”、“za”、“ab”、“zab”。
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= p.length <= 10^5
 * p 由小写英文字母构成
 *
 *
 */
func findSubstringInWraproundString(s string) int {
	n := len(s)
	ans := 0
	// 维护以当前字母为首字母的最长子串长度, 然后你会发现这就是当前字母为首字母的去重后子串数目. 话说这不是个 dp 吗?
	dp := make([]int, 'z'-'a'+1)
	for p, q := 0, 0; p < n; p = q + 1 {
		for q = p; q+1 < n && ((s[q+1] == 'a' && s[q] == 'z') || (s[q+1]-'a')-(s[q]-'a') == 1); q++ {
		}
		for i := p; i <= q; i++ {
			b := s[i] - 'a'
			dp[b] = max(dp[b], q-i+1)
		}
	}
	for b := 'a'; b <= 'z'; b++ {
		ans += dp[b-'a']
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
