/*
 * @lc app=leetcode.cn id=valid-parentheses lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (44.54%)
 * Total Accepted:    930.2K
 * Total Submissions: 2.1M
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
 * 
 * 有效字符串需满足：
 * 
 * 
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "()"
 * 输出：true
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "()[]{}"
 * 输出：true
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：s = "(]"
 * 输出：false
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：s = "([)]"
 * 输出：false
 * 
 * 
 * 示例 5：
 * 
 * 
 * 输入：s = "{[]}"
 * 输出：true
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 1e4
 * s 仅由括号 '()[]{}' 组成
 * 
 * 
 */
func isValid(s string) bool {
	bytes := []byte(s)
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var S []byte
	for _, b := range bytes {
		if b == '(' || b == '[' || b == '{' {
			S = append(S, b)
			continue
		}
		if len(S) == 0 {
			return false
		}
		if S[len(S)-1] != m[b] {
			return false
		}
		S = S[:len(S)-1]
	}
	return len(S) == 0
}
