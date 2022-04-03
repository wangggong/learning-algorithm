/*
 * @lc app=leetcode.cn id=sqrtx lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [69] x 的平方根
 *
 * https://leetcode-cn.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (38.99%)
 * Total Accepted:    462.1K
 * Total Submissions: 1.2M
 * Testcase Example:  '4'
 *
 * 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
 *
 * 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
 *
 * 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：x = 4
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：x = 8
 * 输出：2
 * 解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= x <= 2^31 - 1
 *
 *
 */

/*
 * // 二分法
 * func mySqrt(x int) int {
 * 	// assert x >= 0;
 * 	if x <= 1 {
 * 		return x
 * 	}
 * 	// 1 2 3 4 5 6 7 8 9
 * 	// 1 1 1 0 0 0 0 0 0
 * 	p, q := 1, x
 * 	for p < q {
 * 		mid := (p + q + 1) >> 1
 * 		if mid*mid > x {
 * 			q = mid - 1
 * 		} else {
 * 			p = mid
 * 		}
 * 	}
 * 	return p
 * }
 */

// Newton 迭代
func mySqrt(x int) int {
	// assert x >= 0;
	if x <= 1 {
		return x
	}
	// p = sqrt(x)  =>  p^2 - x = 0  =>  f(p) = 0
	// p<n> = p<n-1> - f(p<n-1>) / f'(p<n-1>)
	c, p := x, 0
	for c != p {
		p, c = c, c-(c*c-x)/(2*c)
	}
	if c*c == x {
		return c
	}
	return c - 1
}
