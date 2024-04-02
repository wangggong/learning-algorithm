/*
 * @lc app=leetcode.cn id=count-different-palindromic-subsequences lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [730] 统计不同回文子序列
 *
 * https://leetcode.cn/problems/count-different-palindromic-subsequences/description/
 *
 * algorithms
 * Hard (50.18%)
 * Total Accepted:    6.7K
 * Total Submissions: 11.8K
 * Testcase Example:  '"bccb"'
 *
 * 给定一个字符串 s，返回 s 中不同的非空「回文子序列」个数 。
 *
 * 通过从 s 中删除 0 个或多个字符来获得子序列。
 *
 * 如果一个字符序列与它反转后的字符序列一致，那么它是「回文字符序列」。
 *
 * 如果有某个 i , 满足 ai != bi ，则两个序列 a1, a2, ... 和 b1, b2, ... 不同。
 *
 * 注意：
 *
 *
 * 结果可能很大，你需要对 10^9 + 7 取模 。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = 'bccb'
 * 输出：6
 * 解释：6 个不同的非空回文子字符序列分别为：'b', 'c', 'bb', 'cc', 'bcb', 'bccb'。
 * 注意：'bcb' 虽然出现两次但仅计数一次。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = 'abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba'
 * 输出：104860361
 * 解释：共有 3104860382 个不同的非空回文子序列，104860361 对 10^9 + 7 取模后的值。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s[i] 仅包含 'a', 'b', 'c' 或 'd'
 *
 *
 */

const mod int = 1e9 + 7

/*
 * // 三维 dp
 * func countPalindromicSubsequences(s string) int {
 * 	n := len(s)
 * 	dp := make([][][4]int, n+1)
 * 	for i := range dp {
 * 		dp[i] = make([][4]int, n+1)
 * 	}
 * 	for i := 0; i < n; i++ {
 * 		dp[i][i+1][int(s[i]-'a')]++
 * 	}
 * 	for i := n; i >= 0; i-- {
 * 		for j := i + 2; j <= n; j++ {
 * 			for b := byte('a'); b <= 'd'; b++ {
 * 				k := int(b - 'a')
 * 				if s[i] == b && s[j-1] == b {
 * 					dp[i][j][k] = 2
 * 					for _, d := range dp[i+1][j-1] {
 * 						dp[i][j][k] = (dp[i][j][k] + d) % mod
 * 					}
 * 				} else if s[i] == b && s[j-1] != b {
 * 					dp[i][j][k] = dp[i][j-1][k]
 * 				} else if s[i] != b && s[j-1] == b {
 * 					dp[i][j][k] = dp[i+1][j][k]
 * 				} else { // s[i] != b && s[j-1] != b
 * 					dp[i][j][k] = dp[i+1][j-1][k]
 * 				}
 * 			}
 * 		}
 * 	}
 * 	ans := 0
 * 	for _, d := range dp[0][n] {
 * 		ans = (ans + d) % mod
 * 	}
 * 	return ans
 * }
 */

func countPalindromicSubsequences(s string) int {
	n := len(s)
	prev, next := make([]int, n), make([]int, n)
	{
		idx := []int{-1, -1, -1, -1}
		for i := 0; i < n; i++ {
			prev[i] = idx[int(s[i]-'a')]
			idx[int(s[i]-'a')] = i
		}
	}
	{
		idx := []int{n, n, n, n}
		for i := n - 1; i >= 0; i-- {
			next[i] = idx[int(s[i]-'a')]
			idx[int(s[i]-'a')] = i
		}
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < n; i++ {
		dp[i][i+1]++
	}
	for i := n; i >= 0; i-- {
		for j := i + 2; j <= n; j++ {
			if s[i] != s[j-1] {
				// 容斥原理. 注意为了避免溢出, 减去某值后再取模别忘了先加一个模数上去.
				dp[i][j] = (dp[i+1][j] + dp[i][j-1] - dp[i+1][j-1] + mod) % mod
				continue
			}
			// 这里去重的技巧:
			// `dp[i][j]` 代表 `s[i:j]` 范围内的答案, 已经去了重的结果数.
			// 如果 `s[i] == s[j]`, 那么在 `dp[i+1][j-1]` 的基础上每个结果前后补个 `s[i]` 也是一个结果. 再加上 `{s[i]}` & `{s[i], s[j]}` 两个结果就是总结果数.
			// 在这个基础上去重, 我们需要找到 `p, q`, 其中位置关系如下:
			//
			// ... ... x ... ... x ... ... x ... ... x ... ...
			//           \      /  \      /  \      /
			//          其他字母   其他字母  其他字母
			// ... ... i ... ... p ... ... q-1 q ... j-1 j... ...
			//
			// 这时每个在 `s[p:q]` 范围内的结果前后再加上 `s[i]` 就会重复, 所以要减掉该部分. 特别的, 如果 `s[p:q]` 长度为 `1`, 只减一个 `1` 即可.
			p, q := next[i], prev[j-1]+1
			switch c := q - p; {
			case c <= 0:
				dp[i][j] = (dp[i+1][j-1]*2 + 2) % mod
			case c == 1:
				dp[i][j] = (dp[i+1][j-1]*2 + 1) % mod
			default:
				dp[i][j] = (dp[i+1][j-1]*2 - dp[p+1][q-1] + mod) % mod
			}
		}
	}
	return dp[0][n]
}
