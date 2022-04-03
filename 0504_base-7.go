/*
 * @lc app=leetcode.cn id=base-7 lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [504] 七进制数
 *
 * https://leetcode-cn.com/problems/base-7/description/
 *
 * algorithms
 * Easy (50.34%)
 * Total Accepted:    47.9K
 * Total Submissions: 92.6K
 * Testcase Example:  '100'
 *
 * 给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入: num = 100
 * 输出: "202"
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: num = -7
 * 输出: "-10"
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * -10^7 <= num <= 10^7
 * 
 * 
 */
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	neg := false
	if num < 0 {
		neg, num = true, -num
	}
	var ans []byte
	for num > 0 {
		ans = append(ans, byte(num % 7) + '0')
		num /= 7
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	if neg {
		return "-" + string(ans)
	}
	return string(ans)
}
