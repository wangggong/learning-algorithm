/*
 * @lc app=leetcode.cn id=find-palindrome-with-fixed-length lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [5253] 找到指定长度的回文数
 *
 * https://leetcode-cn.com/problems/find-palindrome-with-fixed-length/description/
 *
 * algorithms
 * Medium (28.55%)
 * Total Accepted:    4.2K
 * Total Submissions: 14.7K
 * Testcase Example:  '[1,2,3,4,5,90]\n3'
 *
 * 给你一个整数数组 queries 和一个 正 整数 intLength ，请你返回一个数组 answer ，其中 answer[i] 是长度为
 * intLength 的 正回文数 中第 queries[i] 小的数字，如果不存在这样的回文数，则为 -1 。
 *
 * 回文数 指的是从前往后和从后往前读一模一样的数字。回文数不能有前导 0 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：queries = [1,2,3,4,5,90], intLength = 3
 * 输出：[101,111,121,131,141,999]
 * 解释：
 * 长度为 3 的最小回文数依次是：
 * 101, 111, 121, 131, 141, 151, 161, 171, 181, 191, 201, ...
 * 第 90 个长度为 3 的回文数是 999 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：queries = [2,4,6], intLength = 4
 * 输出：[1111,1331,1551]
 * 解释：
 * 长度为 4 的前 6 个回文数是：
 * 1001, 1111, 1221, 1331, 1441 和 1551 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= queries.length <= 5 * 10^4
 * 1 <= queries[i] <= 10^9
 * 1 <= intLength <= 15
 *
 *
 */

func kthPalindrome(queries []int, intLength int) []int64 {
	var ans []int64
	for _, q := range queries {
		ans = append(ans, getPalindrome(intLength, q-1))
	}
	return ans
}

// 可以参考: https://leetcode-cn.com/problems/find-palindrome-with-fixed-length/solution/fan-zhuan-hui-wen-shu-zuo-ban-bu-fen-by-4pvs0/
//
// 这题我做出来了, 列个答案.
func getPalindrome(n int, k int) int64 {
	m := (n + 1) >> 1
	v := 1
	for i := 0; i+1 < m; i++ {
		v *= 10
	}
	if v+k >= v*10 {
		return -1
	}
	v += k
	ans := int64(v)
	if n%2 != 0 {
		v /= 10
	}
	for v > 0 {
		ans = ans*10 + int64(v%10)
		v /= 10
	}
	return ans
}
