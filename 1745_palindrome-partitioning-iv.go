/*
 * @lc app=leetcode.cn id=palindrome-partitioning-iv lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1745] 回文串分割 IV
 *
 * https://leetcode-cn.com/problems/palindrome-partitioning-iv/description/
 *
 * algorithms
 * Hard (48.67%)
 * Total Accepted:    3.8K
 * Total Submissions: 7.9K
 * Testcase Example:  '"abcbdd"'
 *
 * 给你一个字符串 s ，如果可以将它分割成三个 非空 回文子字符串，那么返回 true ，否则返回 false 。
 *
 * 当一个字符串正着读和反着读是一模一样的，就称其为 回文字符串 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "abcbdd"
 * 输出：true
 * 解释："abcbdd" = "a" + "bcb" + "dd"，三个子字符串都是回文的。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "bcbddxy"
 * 输出：false
 * 解释：s 没办法被分割成 3 个回文子字符串。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= s.length <= 2000
 * s 只包含小写英文字母。
 *
 *
 */

const maxn int = 2000

var dp [maxn + 10][maxn + 10]bool

func checkPartitioning(s string) bool {
	arr := []byte(s)
	n := len(arr)
	for i := 0; i < maxn+10; i++ {
		for j := 0; j < maxn+10; j++ {
			dp[i][j] = false
		}
	}
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
	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if dp[0][i-1] && dp[i][j-1] && dp[j][n-1] {
				return true
			}
		}
	}
	return false
}
