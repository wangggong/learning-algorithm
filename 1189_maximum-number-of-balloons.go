/*
 * @lc app=leetcode.cn id=maximum-number-of-balloons lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1189] “气球” 的最大数量
 *
 * https://leetcode-cn.com/problems/maximum-number-of-balloons/description/
 *
 * algorithms
 * Easy (64.67%)
 * Total Accepted:    35.5K
 * Total Submissions: 52.2K
 * Testcase Example:  '"nlaebolko"'
 *
 * 给你一个字符串 text，你需要使用 text 中的字母来拼凑尽可能多的单词 "balloon"（气球）。
 *
 * 字符串 text 中的每个字母最多只能被使用一次。请你返回最多可以拼凑出多少个单词 "balloon"。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：text = "nlaebolko"
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：text = "loonbalxballpoon"
 * 输出：2
 *
 *
 * 示例 3：
 *
 * 输入：text = "leetcode"
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= text.length <= 10^4
 * text 全部由小写英文字母组成
 *
 *
 */
func maxNumberOfBalloons(text string) int {
	bytes := []byte(text)
	cb, ca, cl, co, cn := 0, 0, 0, 0, 0
	for _, b := range bytes {
		switch b {
		case 'b':
			cb++
		case 'a':
			ca++
		case 'l':
			cl++
		case 'o':
			co++
		case 'n':
			cn++
		}
	}
	ans := min(cb, min(min(ca, cl/2), min(co/2, cn)))
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
