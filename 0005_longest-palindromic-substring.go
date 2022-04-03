/*
 * @lc app=leetcode.cn id=longest-palindromic-substring lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (36.18%)
 * Total Accepted:    938.4K
 * Total Submissions: 2.6M
 * Testcase Example:  '"babad"'
 *
 * 给你一个字符串 s，找到 s 中最长的回文子串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "babad"
 * 输出："bab"
 * 解释："aba" 同样是符合题意的答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbbd"
 * 输出："bb"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 仅由数字和英文字母组成
 *
 *
 */

const MAXN int = 1e3

var dp [MAXN + 10][MAXN + 10]bool

/*
 * func longestPalindrome(s string) string {
 * 	for i := 0; i < MAXN+10; i++ {
 * 		for j := 0; j < MAXN+10; j++ {
 * 			dp[i][j] = false
 * 		}
 * 	}
 * 	arr := []byte(s)
 * 	n := len(arr)
 * 	var ans []byte
 * 	for i := 0; i < n; i++ {
 * 		dp[i][i] = true
 * 		if len(ans) < 1 {
 * 			ans = arr[i : i+1]
 * 		}
 * 		if i+1 < n && arr[i] == arr[i+1] {
 * 			dp[i][i+1] = true
 * 			if len(ans) < 2 {
 * 				ans = arr[i : i+2]
 * 			}
 * 		}
 * 	}
 * 	for i := n - 1; i >= 0; i-- {
 * 		for j := i + 1; j < n; j++ {
 * 			if arr[i] == arr[j] {
 * 				dp[i][j] = dp[i][j] || dp[i+1][j-1]
 * 			}
 * 			if dp[i][j] && len(ans) < j-i+1 {
 * 				ans = arr[i : j+1]
 * 			}
 * 		}
 * 	}
 * 	return string(ans)
 * }
 */

func longestPalindrome(s string) string {
	bytes := []byte(s)
	n := len(bytes)
	var ans []byte
	for i := 0; i < n; i++ {
		for j := 0; j <= i && i+j < n && bytes[i+j] == bytes[i-j]; j++ {
			if len(ans) < len(bytes[i-j:i+j+1]) {
				ans = bytes[i-j:i+j+1]
			}
		}
		if i+1 < n && bytes[i] == bytes[i+1] {
			for j := 0; j <= i && i+1+j < n && bytes[i+1+j] == bytes[i-j]; j++ {
				if len(ans) < len(bytes[i-j:i+j+2]) {
					ans = bytes[i-j:i+j+2]
				}
			}
		}
	}
	return string(ans)
}
