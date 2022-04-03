/*
 * @lc app=leetcode.cn id=palindromic-substrings lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [647] 回文子串
 *
 * https://leetcode-cn.com/problems/palindromic-substrings/description/
 *
 * algorithms
 * Medium (66.09%)
 * Total Accepted:    160.1K
 * Total Submissions: 242.1K
 * Testcase Example:  '"abc"'
 *
 * 给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
 *
 * 回文字符串 是正着读和倒过来读一样的字符串。
 *
 * 子字符串 是字符串中的由连续字符组成的一个序列。
 *
 * 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "abc"
 * 输出：3
 * 解释：三个回文子串: "a", "b", "c"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "aaa"
 * 输出：6
 * 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 由小写英文字母组成
 *
 *
 */

const MAXN int = 1e3

var dp [MAXN + 10][MAXN + 10]bool

func countSubstrings(s string) int {
	for i := 0; i < MAXN+10; i++ {
		for j := 0; j < MAXN+10; j++ {
			dp[i][j] = false
		}
	}
	arr := []byte(s)
	n := len(arr)
	ans := 0
	for i := 0; i < n; i++ {
		dp[i][i] = true
		if i+1 < n && arr[i] == arr[i+1] {
			dp[i][i+1] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if arr[i] == arr[j] {
				dp[i][j] = dp[i][j] || dp[i+1][j-1]
			}
		}
	}
	ans = 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if dp[i][j] {
				ans++
			}
		}
	}
	return ans
}
