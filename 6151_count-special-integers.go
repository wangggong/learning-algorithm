/*
 * @lc app=leetcode.cn id=count-special-integers lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6151] 统计特殊整数
 *
 * https://leetcode.cn/problems/count-special-integers/description/
 *
 * algorithms
 * Hard (28.86%)
 * Total Accepted:    1.8K
 * Total Submissions: 5.8K
 * Testcase Example:  '20'
 *
 * 如果一个正整数每一个数位都是 互不相同 的，我们称它是 特殊整数 。
 *
 * 给你一个 正 整数 n ，请你返回区间 [1, n] 之间特殊整数的数目。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 20
 * 输出：19
 * 解释：1 到 20 之间所有整数除了 11 以外都是特殊整数。所以总共有 19 个特殊整数。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 5
 * 输出：5
 * 解释：1 到 5 所有整数都是特殊整数。
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 135
 * 输出：110
 * 解释：从 1 到 135 总共有 110 个整数是特殊整数。
 * 不特殊的部分数字为：22 ，114 和 131 。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 2 * 10^9
 *
 *
 */

import "strconv"

func countSpecialNumbers(n int) (ans int) {
	s := strconv.Itoa(n + 1)
	m := len(s)
	ans = getCntBefore(m)
	// fmt.Println(ans)
	for i := 0; i < m; i++ {
		cur := getCntAt(s, i)
		// fmt.Println(s, i, cur)
		ans += cur
	}
	return
}

func getCntBefore(k int) (ans int) {
	for i := 1; i < k; i++ {
		cur := 9
		for j := 1; j < i; j++ {
			cur *= 10 - j
		}
		ans += cur
	}
	return
}

func getCntAt(s string, k int) (ans int) {
	vis := make([]bool, 10)
	for _, b := range s[:k] {
		if vis[int(b-'0')] {
			return
		}
		vis[int(b-'0')] = true
	}
	cnt := 1
	for i, n := k+1, len(s); i < n; i++ {
		cnt *= 10 - i
	}
	for b := byte('0'); b < s[k]; b++ {
		if k == 0 && b == '0' {
			continue
		}
		if vis[int(b-'0')] {
			continue
		}
		ans += cnt
	}
	return
}

// f(135) = g(100) + g(130) + g(135) = 9 + 9*9 + 2*8 + 4
