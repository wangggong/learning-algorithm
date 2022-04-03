/*
 * @lc app=leetcode.cn id=partition-labels lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [763] 划分字母区间
 *
 * https://leetcode-cn.com/problems/partition-labels/description/
 *
 * algorithms
 * Medium (76.43%)
 * Total Accepted:    103.2K
 * Total Submissions: 135.1K
 * Testcase Example:  '"ababcbacadefegdehijhklij"'
 *
 * 字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。
 *
 *
 *
 * 示例：
 *
 *
 * 输入：S = "ababcbacadefegdehijhklij"
 * 输出：[9,7,8]
 * 解释：
 * 划分结果为 "ababcbaca", "defegde", "hijhklij"。
 * 每个字母最多出现在一个片段中。
 * 像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。
 *
 *
 *
 *
 * 提示：
 *
 *
 * S的长度在[1, 500]之间。
 * S只包含小写字母 'a' 到 'z' 。
 *
 *
 */
func partitionLabels(s string) []int {
	arr := []byte(s)
	n := len(arr)
	right := make([]int, 26)
	for i := range right {
		right[i] = -1
	}
	for i, a := range arr {
		right[int(a-'a')] = i
	}
	var ans []int
	p, q, sum := 0, 0, 0
	for q < n {
		for ; p < n && p <= q; p++ {
			if q < right[int(arr[p]-'a')] {
				q = right[int(arr[p]-'a')]
			}
		}
		ans = append(ans, p-sum)
		q = p
		sum += ans[len(ans)-1]
	}
	return ans
}

