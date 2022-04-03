/*
 * @lc app=leetcode.cn id=decode-string lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [394] 字符串解码
 *
 * https://leetcode-cn.com/problems/decode-string/description/
 *
 * algorithms
 * Medium (55.66%)
 * Total Accepted:    141.1K
 * Total Submissions: 253.4K
 * Testcase Example:  '"3[a]2[bc]"'
 *
 * 给定一个经过编码的字符串，返回它解码后的字符串。
 *
 * 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
 *
 * 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
 *
 * 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "3[a]2[bc]"
 * 输出："aaabcbc"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "3[a2[c]]"
 * 输出："accaccacc"
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "2[abc]3[cd]ef"
 * 输出："abcabccdcdcdef"
 *
 *
 * 示例 4：
 *
 *
 * 输入：s = "abc3[cd]xyz"
 * 输出："abccdcdcdxyz"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 30
 * s 由小写英文字母、数字和方括号 '[]' 组成
 * s 保证是一个 有效 的输入。
 * s 中所有整数的取值范围为 [1, 300]
 *
 *
 */
import "strings"

var (
	bytes []byte
	index, n int
)

func decodeString(s string) string {
	bytes = []byte(s)
	index, n = 0, len(bytes)
	return getString()
}

func getString() string {
	ans := ""
	// 如果是 ']' 也需要直接退出, 傻逼!
	if index == n || bytes[index] == ']' {
		return ans
	}
	switch b := bytes[index]; {
	case '0' <= b && b <= '9':
		k := getDigit()
		index++ // '['
		s := getString()
		index++ // ']'
		for ; k > 0; k-- {
			ans = ans + s
		}
	case 'a' <= b && b <= 'z':
		var cur []byte
		for ; index < n && 'a' <= bytes[index] && bytes[index] <= 'z'; index++ {
			cur = append(cur, bytes[index])
		}
		ans = ans + string(cur)
	default:
	}
	return ans + getString()
}

func getDigit() int {
	ans := 0
	for ; index < n && '0' <= bytes[index] && bytes[index] <= '9'; index++ {
		ans = ans * 10 + int(bytes[index] - '0')
	}
	return ans
}

