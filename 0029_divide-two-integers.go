/*
 * @lc app=leetcode.cn id=divide-two-integers lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [29] 两数相除
 *
 * https://leetcode-cn.com/problems/divide-two-integers/description/
 *
 * algorithms
 * Medium (22.13%)
 * Total Accepted:    149.2K
 * Total Submissions: 674.1K
 * Testcase Example:  '10\n3'
 *
 * 给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
 *
 * 返回被除数 dividend 除以除数 divisor 得到的商。
 *
 * 整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) =
 * -2
 *
 *
 *
 * 示例 1:
 *
 * 输入: dividend = 10, divisor = 3
 * 输出: 3
 * 解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
 *
 * 示例 2:
 *
 * 输入: dividend = 7, divisor = -3
 * 输出: -2
 * 解释: 7/-3 = truncate(-2.33333..) = -2
 *
 *
 *
 * 提示：
 *
 *
 * 被除数和除数均为 32 位有符号整数。
 * 除数不为 0。
 * 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。
 *
 *
 */

import "math"

/*
 * func divide(dividend int, divisor int) int {
 * 	// assert divisor != 0;
 * 	if dividend == 0 {
 * 		return 0
 * 	}
 * 	if divisor == -(1 << 31) {
 * 		if dividend == divisor {
 * 			return 1
 * 		}
 * 		return 0
 * 	}
 * 	ans := 0
 * 	flag := ((dividend > 0) && (divisor < 0)) || ((dividend < 0) && (divisor > 0))
 * 	dividend, divisor = abs(dividend), abs(divisor)
 * 	for c := dividend; c >= divisor; {
 * 		k := 0
 * 		for ; c > divisor << k; k++ {}
 * 		if c == divisor << k {
 * 			ans += 1 << k
 * 			c -= divisor << k
 * 		} else {
 * 			ans += 1 << (k-1)
 * 			c -= divisor << (k-1)
 * 		}
 * 	}
 * 	if flag {
 * 		return -ans
 * 	}
 * 	if ans > (1<<31)-1 {
 * 		ans = (1 << 31) - 1
 * 	}
 * 	return ans
 * }
 */

var powers []int

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	// assert divisor != 0;
	neg := (dividend > 0) != (divisor > 0)
	dividend, divisor = abs(dividend), abs(divisor)
	powers = nil
	for k := divisor; k <= dividend; {
		powers = append(powers, k)
		k += k
	}
	// fmt.Println(powers)
	n := len(powers) - 1
	ans := 0
	for ; n >= 0; n-- {
		if dividend >= powers[n] {
			ans += 1 << n
			dividend -= powers[n]
		}
	}
	if neg {
		ans = -ans
	}
	if ans > math.MaxInt32 {
		ans = math.MaxInt32
	}
	if ans < math.MinInt32 {
		ans = math.MinInt32
	}
	return ans
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
