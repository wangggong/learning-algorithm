import "sort"

/*
 * @lc app=leetcode.cn id=largest-palindromic-number lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6166] 最大回文数字
 *
 * https://leetcode.cn/problems/largest-palindromic-number/description/
 *
 * algorithms
 * Medium (25.79%)
 * Total Accepted:    4.4K
 * Total Submissions: 17.2K
 * Testcase Example:  '"444947137"'
 *
 * 给你一个仅由数字（0 - 9）组成的字符串 num 。
 *
 * 请你找出能够使用 num 中数字形成的 最大回文 整数，并以字符串形式返回。该整数不含 前导零 。
 *
 * 注意：
 *
 *
 * 你 无需 使用 num 中的所有数字，但你必须使用 至少 一个数字。
 * 数字可以重新排序。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：num = "444947137"
 * 输出："7449447"
 * 解释：
 * 从 "444947137" 中选用数字 "4449477"，可以形成回文整数 "7449447" 。
 * 可以证明 "7449447" 是能够形成的最大回文整数。
 *
 *
 * 示例 2：
 *
 *
 * 输入：num = "00009"
 * 输出："9"
 * 解释：
 * 可以证明 "9" 能够形成的最大回文整数。
 * 注意返回的整数不应含前导零。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= num.length <= 10^5
 * num 由数字（0 - 9）组成
 *
 *
 */
func largestPalindromic(num string) string {
	var cnt [10]int
	for _, b := range num {
		cnt[int(b-'0')]++
	}
	single := byte(0)
	var bs, ans []byte
	for i, c := range cnt {
		for j := 0; j < c/2; j++ {
			bs = append(bs, byte(i)+'0')
		}
		if c%2 != 0 {
			single = byte(i) + '0'
		}
	}
	sort.Slice(bs, func(p, q int) bool { return bs[p] > bs[q] })
	ans = append(ans, bs...)
	if single != 0 {
		ans = append(ans, single)
	}
	for i := len(bs) - 1; i >= 0; i-- {
		ans = append(ans, bs[i])
	}
	k := 0
	for ; k < len(ans) && ans[k] == '0'; k++ {
	}
	if k == len(ans) {
		return "0"
	}
	return string(ans[k : len(ans)-k])
}
