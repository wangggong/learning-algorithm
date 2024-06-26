/*
 * @lc app=leetcode.cn id=optimal-partition-of-string lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2405] 子字符串的最优划分
 *
 * https://leetcode.cn/problems/optimal-partition-of-string/description/
 *
 * algorithms
 * Medium (72.76%)
 * Total Accepted:    5.9K
 * Total Submissions: 8.1K
 * Testcase Example:  '"abacaba"'
 *
 * 给你一个字符串 s ，请你将该字符串划分成一个或多个 子字符串 ，并满足每个子字符串中的字符都是 唯一
 * 的。也就是说，在单个子字符串中，字母的出现次数都不超过 一次 。
 *
 * 满足题目要求的情况下，返回 最少 需要划分多少个子字符串。
 *
 * 注意，划分后，原字符串中的每个字符都应该恰好属于一个子字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "abacaba"
 * 输出：4
 * 解释：
 * 两种可行的划分方法分别是 ("a","ba","cab","a") 和 ("ab","a","ca","ba") 。
 * 可以证明最少需要划分 4 个子字符串。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "ssssss"
 * 输出：6
 * 解释：
 * 只存在一种可行的划分方法 ("s","s","s","s","s","s") 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^5
 * s 仅由小写英文字母组成
 *
 *
 */
func partitionString(s string) int {
	ans := 0
	vis := make([]bool, 'z'-'a'+1)
	for _, b := range s {
		i := int(b - 'a')
		if vis[i] {
			ans++
			for j := range vis {
				vis[j] = false
			}
		}
		vis[i] = true
	}
	return ans + 1
}
