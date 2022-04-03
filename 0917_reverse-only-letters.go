/*
 * @lc app=leetcode.cn id=reverse-only-letters lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [917] 仅仅反转字母
 *
 * https://leetcode-cn.com/problems/reverse-only-letters/description/
 *
 * algorithms
 * Easy (57.01%)
 * Total Accepted:    32.6K
 * Total Submissions: 57.1K
 * Testcase Example:  '"ab-cd"'
 *
 * 给你一个字符串 s ，根据下述规则反转字符串：
 *
 *
 * 所有非英文字母保留在原有位置。
 * 所有英文字母（小写或大写）位置反转。
 *
 *
 * 返回反转后的 s 。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "ab-cd"
 * 输出："dc-ba"
 *
 *
 *
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "a-bC-dEf-ghIj"
 * 输出："j-Ih-gfE-dCba"
 *
 *
 *
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "Test1ng-Leet=code-Q!"
 * 输出："Qedo1ct-eeLg=ntse-T!"
 *
 *
 *
 *
 * 提示
 *
 *
 * 1 <= s.length <= 100
 * s 仅由 ASCII 值在范围 [33, 122] 的字符组成
 * s 不含 '\"' 或 '\\'
 *
 *
 */
func reverseOnlyLetters(s string) string {
	bs := []byte(s)
	n := len(bs)
	for p, q := 0, n-1; p < q; {
		for p < n && !isAlpha(bs[p]) {
			p++
		}
		for q >= 0 && !isAlpha(bs[q]) {
			q--
		}
		if p < q {
			bs[p], bs[q] = bs[q], bs[p]
		}
		p, q = p+1, q-1
	}
	return string(bs)
}

func isAlpha(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}

