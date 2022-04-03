/*
 * @lc app=leetcode.cn id=longest-palindromic-subsequence lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [516] 最长回文子序列
 *
 * https://leetcode-cn.com/problems/longest-palindromic-subsequence/description/
 *
 * algorithms
 * Medium (66.05%)
 * Total Accepted:    109.1K
 * Total Submissions: 165K
 * Testcase Example:  '"bbbab"'
 *
 * 给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
 *
 * 子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "bbbab"
 * 输出：4
 * 解释：一个可能的最长回文子序列为 "bbbb" 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbbd"
 * 输出：2
 * 解释：一个可能的最长回文子序列为 "bb" 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 仅由小写英文字母组成
 *
 *
 */

const MAXN int = 1000

var dp [MAXN + 10][MAXN + 10]int

func longestPalindromeSubseq(s string) int {
	for i := 0; i < MAXN+10; i++ {
		for j := 0; j < MAXN+10; j++ {
			dp[i][j] = 0
		}
	}
	arr := []byte(s)
	n := len(arr)
	var rev []byte
	for i := n - 1; i >= 0; i-- {
		rev = append(rev, arr[i])
	}
	return lcs(arr, rev)
}

/*
 * // O(n^2) DP
 * func lcs(s, t []byte) int {
 * 	n, m := len(s), len(t)
 * 	for i := 1; i <= n; i++ {
 * 		for j := 1; j <= m; j++ {
 * 			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
 * 			if s[i-1] == t[j-1] {
 * 				dp[i][j] = max(dp[i][j], 1+dp[i-1][j-1])
 * 			}
 * 		}
 * 	}
 * 	return dp[n][m]
 * }
 */

// O(nlogn) DP 转为最长递增子序列
//
// 参考: https://www.cnblogs.com/forevergoodboy/p/11666830.html
func lcs(s, t []byte) int {
	n, m := len(s), len(t)
	var now []int
	var indexes [26][]int
	for i := m - 1; i >= 0; i-- {
		ind := int(t[i] - 'a')
		indexes[ind] = append(indexes[ind], i)
	}
	for i := 0; i < n; i++ {
		ind := int(s[i] - 'a')
		for _, j := range indexes[ind] {
			now = append(now, j)
		}
	}
	nn := len(now)
	if nn == 0 {
		return 0
	}
	var low []int
	for i := 0; i < nn; i++ {
		if len(low) == 0 || now[i] > low[len(low)-1] {
			low = append(low, now[i])
		} else {
			low[find(low, now[i])] = now[i]
		}
	}
	return len(low)
}

func find(arr []int, k int) int {
	p, q := 0, len(arr)
	for p < q {
		mid := (p + q) >> 1
		if arr[mid] >= k {
			q = mid
		} else {
			p = mid + 1
		}
	}
	return p
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
