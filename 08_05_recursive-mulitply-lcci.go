/*
 * Medium
 *
 * 递归乘法。 写一个递归函数，不使用 * 运算符， 实现两个正整数的相乘。可以使用加号、减号、位移，但要吝啬一些。
 *
 * 示例1:
 *
 *  输入：A = 1, B = 10
 *  输出：10
 * 示例2:
 *
 *  输入：A = 3, B = 4
 *  输出：12
 * 提示:
 *
 * 保证乘法范围不会溢出
 * 通过次数30,088
 * 提交次数44,810
 */

func multiply(A int, B int) int {
	ans := 0
	cur := A
	if B == 0 {
		return 0
	}
	for ; B > 0; B >>= 1 {
		if B&1 != 0 {
			ans += cur
		}
		cur += cur
	}
	return ans
}

/*
 * // 你要递归版本?
 * func multiply(A int, B int) int {
 * 	if B == 0 {
 * 		return 0
 * 	}
 * 	ans := 0
 * 	if B&1 != 0 {
 * 		ans += A
 * 	}
 * 	cur := multiply(A, B>>1)
 * 	ans += cur << 1
 * 	return ans
 * }
 */
