/*
 * @lc app=leetcode.cn id=maximum-swap lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [670] 最大交换
 *
 * https://leetcode.cn/problems/maximum-swap/description/
 *
 * algorithms
 * Medium (46.42%)
 * Total Accepted:    37.8K
 * Total Submissions: 80K
 * Testcase Example:  '2736'
 *
 * 给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。
 *
 * 示例 1 :
 *
 *
 * 输入: 2736
 * 输出: 7236
 * 解释: 交换数字2和数字7。
 *
 *
 * 示例 2 :
 *
 *
 * 输入: 9973
 * 输出: 9973
 * 解释: 不需要交换。
 *
 *
 * 注意:
 *
 *
 * 给定数字的范围是 [0, 10^8]
 *
 *
 */
import "strconv"

func maximumSwap(num int) int {
	s := strconv.Itoa(num)
	bs := []byte(s)
	for i, n := 0, len(bs); i < n; i++ {
		mi := i
		for j := i + 1; j < n; j++ {
			if bs[j] >= bs[mi] {
				mi = j
			}
		}
		if bs[mi] != bs[i] {
			bs[i], bs[mi] = bs[mi], bs[i]
			break
		}
	}
	ans, _ := strconv.Atoi(string(bs))
	return ans
}
