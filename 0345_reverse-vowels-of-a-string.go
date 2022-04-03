/*
 * @lc app=leetcode.cn id=reverse-vowels-of-a-string lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [345] 反转字符串中的元音字母
 *
 * https://leetcode-cn.com/problems/reverse-vowels-of-a-string/description/
 *
 * algorithms
 * Easy (54.01%)
 * Total Accepted:    116.3K
 * Total Submissions: 215.3K
 * Testcase Example:  '"hello"'
 *
 * 给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
 *
 * 元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "hello"
 * 输出："holle"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "leetcode"
 * 输出："leotcede"
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 3 * 10^5
 * s 由 可打印的 ASCII 字符组成
 *
 *
 */

var m = map[byte]struct{}{
	'a': struct{}{},
	'e': struct{}{},
	'i': struct{}{},
	'o': struct{}{},
	'u': struct{}{},
	'A': struct{}{},
	'E': struct{}{},
	'I': struct{}{},
	'O': struct{}{},
	'U': struct{}{},
}

func prev(b []byte, n int) int {
	for i := n - 1; i >= 0; i-- {
		if _, ok := m[b[i]]; ok {
			return i
		}
	}
	return -1
}

func next(b []byte, n int) int {
	for i := n + 1; i < len(b); i++ {
		if _, ok := m[b[i]]; ok {
			return i
		}
	}
	return n
}

func reverseVowels(s string) string {
	b := []byte(s)
	n := len(b)
	p, q := next(b, -1), prev(b, n)

	for p < q {
		b[p], b[q] = b[q], b[p]
		p, q = next(b, p), prev(b, q)
	}

	return string(b)
}
