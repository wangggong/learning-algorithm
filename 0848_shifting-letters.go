/*
 * @lc app=leetcode.cn id=shifting-letters lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [848] 字母移位
 *
 * https://leetcode-cn.com/problems/shifting-letters/description/
 *
 * algorithms
 * Medium (45.00%)
 * Total Accepted:    10.1K
 * Total Submissions: 22.5K
 * Testcase Example:  '"abc"\n[3,5,9]'
 *
 * 有一个由小写字母组成的字符串 s，和一个长度相同的整数数组 shifts。
 *
 * 我们将字母表中的下一个字母称为原字母的 移位 shift() （由于字母表是环绕的， 'z' 将会变成 'a'）。
 *
 *
 * 例如，shift('a') = 'b', shift('t') = 'u', 以及 shift('z') = 'a'。
 *
 *
 * 对于每个 shifts[i] = x ， 我们会将 s 中的前 i + 1 个字母移位 x 次。
 *
 * 返回 将所有这些移位都应用到 s 后最终得到的字符串 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "abc", shifts = [3,5,9]
 * 输出："rpl"
 * 解释：
 * 我们以 "abc" 开始。
 * 将 S 中的第 1 个字母移位 3 次后，我们得到 "dbc"。
 * 再将 S 中的前 2 个字母移位 5 次后，我们得到 "igc"。
 * 最后将 S 中的这 3 个字母移位 9 次后，我们得到答案 "rpl"。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "aaa", shifts = [1,2,3]
 * 输出: "gfd"
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= s.length <= 10^5
 * s 由小写英文字母组成
 * shifts.length == s.length
 * 0 <= shifts[i] <= 10^9
 *
 * ​​​​​​
 */
const ALPHA_COUNT = 26

func shiftingLetters(s string, shifts []int) string {
	// assert len(s) == len(shifts);
	// assert len(s) >= 1 && len(s) <= 1e5;
	bytes := []byte(s)
	n := len(bytes)
	suffixSum := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffixSum[i] = (suffixSum[i+1] + shifts[i]) % ALPHA_COUNT
	}
	for i, b := range bytes {
		// 注意类型转换! Golang 这方面很严格!
		bytes[i] = (((b - byte('a')) + byte(suffixSum[i])) % ALPHA_COUNT) + byte('a')
	}
	return string(bytes)
}
