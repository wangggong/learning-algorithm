/*
 * @lc app=leetcode.cn id=consecutive-numbers-sum lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [829] 连续整数求和
 *
 * https://leetcode.cn/problems/consecutive-numbers-sum/description/
 *
 * algorithms
 * Hard (37.02%)
 * Total Accepted:    14.9K
 * Total Submissions: 36.4K
 * Testcase Example:  '5'
 *
 * 给定一个正整数 n，返回 连续正整数满足所有数字之和为 n 的组数 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: n = 5
 * 输出: 2
 * 解释: 5 = 2 + 3，共有两组连续整数([5],[2,3])求和后为 5。
 *
 * 示例 2:
 *
 *
 * 输入: n = 9
 * 输出: 3
 * 解释: 9 = 4 + 5 = 2 + 3 + 4
 *
 * 示例 3:
 *
 *
 * 输入: n = 15
 * 输出: 4
 * 解释: 15 = 8 + 7 = 4 + 5 + 6 = 1 + 2 + 3 + 4 + 5
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= n <= 10^9
 *
 *
 */

func consecutiveNumbersSum(n int) int {
	ans := 1
	for i := 2; i+1 <= 2*n/i; i++ {
		if (2*n)%i != 0 {
			continue
		}
		s := 2*n/i - (i - 1)
		if s%2 != 0 || s/2 <= 0 {
			continue
		}
		ans++
	}
	return ans
}

// 想到了求和公式, 没想到通过枚举项数控制复杂度... 我是傻逼...
// S = \Sigma_{i=p}^{p+k-1}{i} = (2*p+k-1) * k / 2 >= (k+1) * k / 2
// k+1 <= 2*S / k
