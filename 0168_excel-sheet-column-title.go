/*
 * @lc app=leetcode.cn id=excel-sheet-column-title lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [168] Excel表列名称
 *
 * https://leetcode-cn.com/problems/excel-sheet-column-title/description/
 *
 * algorithms
 * Easy (43.44%)
 * Total Accepted:    98.9K
 * Total Submissions: 227.8K
 * Testcase Example:  '1'
 *
 * 给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。
 *
 * 例如：
 *
 *
 * A -> 1
 * B -> 2
 * C -> 3
 * ...
 * Z -> 26
 * AA -> 27
 * AB -> 28
 * ...
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：columnNumber = 1
 * 输出："A"
 *
 *
 * 示例 2：
 *
 *
 * 输入：columnNumber = 28
 * 输出："AB"
 *
 *
 * 示例 3：
 *
 *
 * 输入：columnNumber = 701
 * 输出："ZY"
 *
 *
 * 示例 4：
 *
 *
 * 输入：columnNumber = 2147483647
 * 输出："FXSHRXW"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= columnNumber <= 1 << 31 - 1
 *
 *
 */

/*
 * func convertToTitle(N int) string {
 * 	var ans []byte
 * 	for N > 26 {
 * 		b := byte(N%26-1) + 'A'
 * 		if N%26 == 0 {
 * 			b = 'Z'
 * 			N -= 26
 * 		}
 * 		ans = append(ans, b)
 * 		N /= 26
 * 	}
 * 	ans = append(ans, byte(N-1)+'A')
 * 	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
 * 		ans[i], ans[j] = ans[j], ans[i]
 * 	}
 * 	return string(ans)
 * }
 */

func convertToTitle(N int) string {
	var ans []byte
	for N > 0 {
		N--
		ans = append(ans, byte(N%26)+'A')
		/**
		 * 这里和下面的写法还是有区别的, N 的确需要减 1:
		 *
		 * ans = append(ans, byte((N-1) % 26) + 'A')
		 */
		N /= 26
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}

// 二刷反而不会了, 丢人啊!
