/*
 * @lc app=leetcode.cn id=ones-and-zeroes lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [474] 一和零
 *
 * https://leetcode-cn.com/problems/ones-and-zeroes/description/
 *
 * algorithms
 * Medium (62.31%)
 * Total Accepted:    93.9K
 * Total Submissions: 150.6K
 * Testcase Example:  '["10","0001","111001","1","0"]\n5\n3'
 *
 * 给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
 *
 *
 * 请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。
 *
 * 如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
 * 输出：4
 * 解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
 * 其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1
 * ，大于 n 的值 3 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：strs = ["10", "0", "1"], m = 1, n = 1
 * 输出：2
 * 解释：最大的子集是 {"0", "1"} ，所以答案是 2 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= strs.length <= 600
 * 1 <= strs[i].length <= 100
 * strs[i] 仅由 '0' 和 '1' 组成
 * 1 <= m, n <= 100
 *
 *
 */

// 这题求的是最大长度, 傻逼!
func findMaxForm(strs []string, m int, n int) int {
	lim := (m+1) * (n+1)
	dp := make([]int, lim)
	for _, s := range strs {
		c0, c1 := getCount([]byte(s))
		k := c0 * (n+1) + c1
		for j := lim-1; j >= 0; j-- {
			p, q := j / (n+1), j % (n+1)
			if c0 <= p && c1 <= q {
				dp[j] = max(dp[j], dp[j-k] + 1)
			}
		}
	}
	return dp[lim-1]
}

func getCount(arr []byte) (int, int) {
	p, q := 0, 0
	for _, a := range arr {
		if a == '0' {
			p++
		} else {
			q++
		}
	}
	return p, q
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

