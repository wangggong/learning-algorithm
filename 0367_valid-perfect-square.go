/*
 * @lc app=leetcode.cn id=valid-perfect-square lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [367] 有效的完全平方数
 *
 * https://leetcode-cn.com/problems/valid-perfect-square/description/
 *
 * algorithms
 * Easy (44.85%)
 * Total Accepted:    131.7K
 * Total Submissions: 293.7K
 * Testcase Example:  '16'
 *
 * 给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
 *
 * 进阶：不要 使用任何内置的库函数，如  sqrt 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：num = 16
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：num = 14
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= num <= 2^31 - 1
 *
 *
 */

/*
 * func isPerfectSquare(num int) bool {
 * 	// assert num >= 0;
 * 	if num < 0 {
 * 		return false
 * 	}
 * 	if num <= 1 {
 * 		return true
 * 	}
 * 	p, q := 1, num
 * 	for p < q {
 * 		mid := (p + q + 1) >> 1
 * 		if mid*mid > num {
 * 			q = mid - 1
 * 		} else {
 * 			p = mid
 * 		}
 * 	}
 * 	return p*p == num
 * }
 */

// n^2 = 1 + 3 + 5 + ... + (2*n-1)
func isPerfectSquare(num int) bool {
	i := 1
	for num > 0 {
		num -= i
		i += 2
	}
	return num == 0
}
