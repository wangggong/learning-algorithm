/*
 * @lc app=leetcode.cn id=longest-arithmetic-subsequence-of-given-difference lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1218] 最长定差子序列
 *
 * https://leetcode-cn.com/problems/longest-arithmetic-subsequence-of-given-difference/description/
 *
 * algorithms
 * Medium (51.50%)
 * Total Accepted:    37.7K
 * Total Submissions: 73.1K
 * Testcase Example:  '[1,2,3,4]\n1'
 *
 * 给你一个整数数组 arr 和一个整数 difference，请你找出并返回 arr 中最长等差子序列的长度，该子序列中相邻元素之间的差等于
 * difference 。
 *
 * 子序列 是指在不改变其余元素顺序的情况下，通过删除一些元素或不删除任何元素而从 arr 派生出来的序列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：arr = [1,2,3,4], difference = 1
 * 输出：4
 * 解释：最长的等差子序列是 [1,2,3,4]。
 *
 * 示例 2：
 *
 *
 * 输入：arr = [1,3,5,7], difference = 1
 * 输出：1
 * 解释：最长的等差子序列是任意单个元素。
 *
 *
 * 示例 3：
 *
 *
 * 输入：arr = [1,5,7,8,5,3,4,2,1], difference = -2
 * 输出：4
 * 解释：最长的等差子序列是 [7,5,3,1]。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10^4
 *
 *
 */
/*
 * func longestSubsequence(arr []int, difference int) int {
 * 	n := len(arr)
 * 	m := make(map[int]int)
 * 	dp := make([]int, n)
 * 	ans := 0
 * 	for i, v := range arr {
 * 		idx, ok := m[v-difference]
 * 		if ok {
 * 			dp[i] = dp[idx] + 1
 * 		} else {
 * 			dp[i] = 1
 * 		}
 * 		m[v] = i
 * 		if ans < dp[i] {
 * 			ans = dp[i]
 * 		}
 * 	}
 * 	return ans
 * }
 */

func longestSubsequence(arr []int, difference int) int {
	dp := make(map[int]int)
	ans := 0
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		if dp[v] > ans {
			ans = dp[v]
		}
	}
	return ans
}
