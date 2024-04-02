/*
 * @lc app=leetcode.cn id=kth-smallest-number-in-multiplication-table lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [668] 乘法表中第k小的数
 *
 * https://leetcode-cn.com/problems/kth-smallest-number-in-multiplication-table/description/
 *
 * algorithms
 * Hard (52.13%)
 * Total Accepted:    11.6K
 * Total Submissions: 21.2K
 * Testcase Example:  '3\n3\n5'
 *
 * 几乎每一个人都用 乘法表。但是你能在乘法表中快速找到第k小的数字吗？
 *
 * 给定高度m 、宽度n 的一张 m * n的乘法表，以及正整数k，你需要返回表中第k 小的数字。
 *
 * 例 1：
 *
 *
 * 输入: m = 3, n = 3, k = 5
 * 输出: 3
 * 解释:
 * 乘法表:
 * 1    2    3
 * 2    4    6
 * 3    6    9
 *
 * 第5小的数字是 3 (1, 2, 2, 3, 3).
 *
 *
 * 例 2：
 *
 *
 * 输入: m = 2, n = 3, k = 6
 * 输出: 6
 * 解释:
 * 乘法表:
 * 1    2    3
 * 2    4    6
 *
 * 第6小的数字是 6 (1, 2, 2, 3, 4, 6).
 *
 *
 * 注意：
 *
 *
 * m 和 n 的范围在 [1, 30000] 之间。
 * k 的范围在 [1, m * n] 之间。
 *
 *
 */

// 参考: https://leetcode.cn/problems/kth-smallest-number-in-multiplication-table/solution/dong-tu-yan-shi-by-xiaohu9527-3k7s/
//
// 对结果二分, 统计 `矩阵中小于等于当前值的数目` 采用的是矩阵 https://leetcode.cn/problems/kth-smallest-element-in-a-sorted-matrix/ 中的技巧.
func findKthNumber(m int, n int, k int) int {
	p, q := 1, m*n
	for p < q {
		mid := (p + q) >> 1
		// 注意这里找的是 `矩阵中小于等于当前值的数目` 这个指标 **大于等于** `k` 的第一个值.
		if getCnt(m, n, mid) >= k {
			q = mid
		} else {
			p = mid + 1
		}
	}
	return p
}

func getCnt(m, n, k int) int {
	ans := 0
	for p, q := 1, n; p <= m && q > 0; {
		if p*q <= k {
			p++
			ans += q
		} else {
			q--
		}
	}
	return ans
}
