/*
 * @lc app=leetcode.cn id=permutation-sequence lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [60] 排列序列
 *
 * https://leetcode.cn/problems/permutation-sequence/description/
 *
 * algorithms
 * Hard (53.05%)
 * Total Accepted:    106.7K
 * Total Submissions: 201.2K
 * Testcase Example:  '3\n3'
 *
 * 给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。
 *
 * 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
 *
 *
 * "123"
 * "132"
 * "213"
 * "231"
 * "312"
 * "321"
 *
 *
 * 给定 n 和 k，返回第 k 个排列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3, k = 3
 * 输出："213"
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 4, k = 9
 * 输出："2314"
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 3, k = 1
 * 输出："123"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 9
 * 1 <= k <= n!
 *
 *
 */
func getPermutation(n int, k int) string {
	var ans []byte
	all := make([]int, n)
	for i := range all {
		all[i] = i + 1
	}
	pow := 1
	for i := 1; i <= n; i++ {
		pow *= i
	}
	for k -= 1; k > 0; k %= pow {
		pow /= len(all)
		ind := k / pow
		ans = append(ans, byte(all[ind]+'0'))
		all = append(all[:ind], all[ind+1:]...)
	}
	for _, a := range all {
		ans = append(ans, byte(a+'0'))
	}
	return string(ans)
}
