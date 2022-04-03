/*
 * @lc app=leetcode.cn id=number-of-1-bits lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [191] 位1的个数
 *
 * https://leetcode-cn.com/problems/number-of-1-bits/description/
 *
 * algorithms
 * Easy (75.32%)
 * Total Accepted:    219.7K
 * Total Submissions: 291.7K
 * Testcase Example:  '00000000000000000000000000001011'
 *
 * 编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。
 *
 *
 *
 * 提示：
 *
 *
 * 请注意，在某些语言（如
 * Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
 * 在 Java 中，编译器使用二进制补码记法来表示有符号整数。因此，在上面的 示例 3 中，输入表示有符号整数 -3。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：00000000000000000000000000001011
 * 输出：3
 * 解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。
 *
 *
 * 示例 2：
 *
 *
 * 输入：00000000000000000000000010000000
 * 输出：1
 * 解释：输入的二进制串 00000000000000000000000010000000 中，共有一位为 '1'。
 *
 *
 * 示例 3：
 *
 *
 * 输入：11111111111111111111111111111101
 * 输出：31
 * 解释：输入的二进制串 11111111111111111111111111111101 中，共有 31 位为 '1'。
 *
 *
 *
 * 提示：
 *
 *
 * 输入必须是长度为 32 的 二进制串 。
 *
 *
 *
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 如果多次调用这个函数，你将如何优化你的算法？
 *
 *
 */

/*
 * func hammingWeight(num uint32) int {
 * 	ans := 0
 * 	for num != 0 {
 * 		ans++
 * 		num = num & (num - 1)
 * 	}
 * 	return ans
 * }
 */

func hammingWeight(num uint32) int {
	//     0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 1 1
	//
	//       X   X   X   X   X   X   X   X   X   X   X   X   X   X   X   X
	// ->  | X | X | X | X | X | X | X | X | X | X | X | X | X | X | X | X
	//
	//         X X     X X     X X     X X     X X     X X     X X     X X
	// ->  | | X X | | X X | | X X | | X X | | X X | | X X | | X X | | X X
	//
	//             X X X X         X X X X         X X X X         X X X X
	// ->  | | | | X X X X | | | | X X X X | | | | X X X X | | | | X X X X
	//
	//                     X X X X X X X X                 X X X X X X X X
	// ->  X X X X X X X X | | | | | | | | X X X X X X X X | | | | | | | |
	//
	//                                     X X X X X X X X X X X X X X X X
	// ->  X X X X X X X X X X X X X X X X | | | | | | | | | | | | | | | |
	num = num&0x55555555 + ((num >> 1) & 0x55555555)
	num = num&0x33333333 + ((num >> 2) & 0x33333333)
	num = num&0x0f0f0f0f + ((num >> 4) & 0x0f0f0f0f)
	num = num&0x00ff00ff + ((num >> 8) & 0x00ff00ff)
	num = num&0x0000ffff + ((num >> 16) & 0x0000ffff)
	return int(num & 0x3f)
}
