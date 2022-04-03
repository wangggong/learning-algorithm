/*
 * @lc app=leetcode.cn id=k-th-smallest-in-lexicographical-order lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [440] 字典序的第K小数字
 *
 * https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order/description/
 *
 * algorithms
 * Hard (38.73%)
 * Total Accepted:    24K
 * Total Submissions: 59.5K
 * Testcase Example:  '13\n2'
 *
 * 给定整数 n 和 k，返回  [1, n] 中字典序第 k 小的数字。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: n = 13, k = 2
 * 输出: 10
 * 解释: 字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10。
 *
 *
 * 示例 2:
 *
 *
 * 输入: n = 1, k = 1
 * 输出: 1
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= k <= n <= 10^9
 *
 *
 */

// 参考: https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order/solution/by-ac_oier-m3zl/ & https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order/solution/ben-ti-shi-shang-zui-wan-zheng-ju-ti-de-shou-mo-sh/
//
// 计数模拟
func findKthNumber(n int, k int) int {
	ans := 1
	// 第 `k` 个 => 字典序数组中 `k-1` 号 index.
	k--
	for k > 0 {
		// 获取当前 ans 及其前缀树下所有节点数 `cnt`.
		// 如果 `cnt <= k`, 换棵树, 往后找 ---- 前缀++.
		// 如果 `cnt > k`, 在这棵树上, 往下找 ---- 后面直接加个 `0`.
		if cnt := getCnt(n, ans); cnt <= k {
			k -= cnt
			ans++
		} else {
			k--
			ans *= 10
		}
	}
	return ans
}

// `1..n` 的整数序列中以 `k` 为前缀的树的规模.
func getCnt(n, k int) int {
	ans := 0
	for curr, next := k, k+1; curr <= n; curr, next = curr*10, next*10 {
		ans += min(n+1, next) - curr
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
