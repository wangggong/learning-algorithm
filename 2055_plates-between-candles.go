/*
 * @lc app=leetcode.cn id=plates-between-candles lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2055] 蜡烛之间的盘子
 *
 * https://leetcode-cn.com/problems/plates-between-candles/description/
 *
 * algorithms
 * Medium (38.99%)
 * Total Accepted:    7.1K
 * Total Submissions: 16.8K
 * Testcase Example:  '"**|**|***|"\n[[2,5],[5,9]]'
 *
 * 给你一个长桌子，桌子上盘子和蜡烛排成一列。给你一个下标从 0 开始的字符串 s ，它只包含字符 '*' 和 '|' ，其中 '*' 表示一个 盘子
 * ，'|' 表示一支 蜡烛 。
 *
 * 同时给你一个下标从 0 开始的二维整数数组 queries ，其中 queries[i] = [lefti, righti] 表示 子字符串
 * s[lefti...righti] （包含左右端点的字符）。对于每个查询，你需要找到 子字符串中 在 两支蜡烛之间 的盘子的 数目 。如果一个盘子在
 * 子字符串中 左边和右边 都 至少有一支蜡烛，那么这个盘子满足在 两支蜡烛之间 。
 *
 *
 * 比方说，s = "||**||**|*" ，查询 [3, 8] ，表示的是子字符串 "*||**|" 。子字符串中在两支蜡烛之间的盘子数目为 2
 * ，子字符串中右边两个盘子在它们左边和右边 都 至少有一支蜡烛。
 *
 *
 * 请你返回一个整数数组 answer ，其中 answer[i] 是第 i 个查询的答案。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 * 输入：s = "**|**|***|", queries = [[2,5],[5,9]]
 * 输出：[2,3]
 * 解释：
 * - queries[0] 有两个盘子在蜡烛之间。
 * - queries[1] 有三个盘子在蜡烛之间。
 *
 *
 * 示例 2:
 *
 *
 *
 * 输入：s = "***|**|*****|**||**|*", queries =
 * [[1,17],[4,5],[14,17],[5,11],[15,16]]
 * 输出：[9,0,0,0,0]
 * 解释：
 * - queries[0] 有 9 个盘子在蜡烛之间。
 * - 另一个查询没有盘子在蜡烛之间。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= s.length <= 10^5
 * s 只包含字符 '*' 和 '|' 。
 * 1 <= queries.length <= 10^5
 * queries[i].length == 2
 * 0 <= lefti <= righti < s.length
 *
 *
 */
func platesBetweenCandles(s string, queries [][]int) []int {
	bytes := []byte(s)
	n := len(bytes)
	left, right := make([]int, n), make([]int, n)
	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i]
		if bytes[i] == '*' {
			prefixSum[i+1]++
		}
	}
	if bytes[0] == '*' {
		left[0] = -1
	}
	for i := 1; i < n; i++ {
		left[i] = left[i-1]
		if bytes[i] == '|' {
			left[i] = i
		}
	}
	if bytes[n-1] == '*' {
		right[n-1] = n
	} else {
		right[n-1] = n - 1
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1]
		if bytes[i] == '|' {
			right[i] = i
		}
	}
	// fmt.Printf("%v\n%v\n%v\n", prefixSum, left, right)
	var ans []int
	for _, q := range queries {
		l, r := right[q[0]], left[q[1]]
		if l < 0 || r >= n || r < q[0] || l > q[1] {
			ans = append(ans, 0)
			continue
		}
		ans = append(ans, prefixSum[r]-prefixSum[l])
	}
	return ans
}
