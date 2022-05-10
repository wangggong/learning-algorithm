/*
 * @lc app=leetcode.cn id=remove-invalid-parentheses lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [301] 删除无效的括号
 *
 * https://leetcode-cn.com/problems/remove-invalid-parentheses/description/
 *
 * algorithms
 * Hard (55.02%)
 * Total Accepted:    67.4K
 * Total Submissions: 122.5K
 * Testcase Example:  '"()())()"'
 *
 * 给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。
 *
 * 返回所有可能的结果。答案可以按 任意顺序 返回。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "()())()"
 * 输出：["(())()","()()()"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "(a)())()"
 * 输出：["(a())()","(a)()()"]
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = ")("
 * 输出：[""]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 25
 * s 由小写英文字母以及括号 '(' 和 ')' 组成
 * s 中至多含 20 个括号
 *
 *
 */

var ans []string

// 啥也不说了, y总流批!
//
// 本题分两步:
// 1. 确定需要删除几个括号;
// 2. 具体删除括号, 一次删一片 (下面详细说);
func removeInvalidParentheses(str string) []string {
	ans = nil
	l, r := 0, 0
	for _, s := range str {
		switch s {
		case '(':
			l++
		case ')':
			if l > 0 {
				l--
			} else {
				r++
			}
		default:
			// do nothing
		}
	}
	backtrack(str, "", 0, len(str), 0, l, r)
	return ans
}

// backtrack 暴力搜索+剪枝.
//
// cur: 当前字符串前缀;
// cnt: 当前字符串前缀左括号比右括号多多少;
// l: 左括号应该比右括号多多少;
// r: 右括号应该比左括号多多少;
func backtrack(s, cur string, i, n, cnt, l, r int) {
	if i == n {
		if cnt == 0 {
			ans = append(ans, cur)
		}
		return
	}
	switch s[i] {
	case '(':
		// 找到这一片 '(', 初始状态下全部删除, 一个一个添加. 直接遍历后面的子串.
		j := i
		for ; j < n && s[j] == '('; j++ {
		}
		l -= j - i
		for k := j - i; k >= 0; k-- {
			if l >= 0 {
				backtrack(s, cur, j, n, cnt, l, r)
			}
			cur, cnt, l = cur+"(", cnt+1, l+1
		}
	case ')':
		// 找到这一片 ')', 初始状态下全部删除, 一个一个添加. 直接遍历后面的子串.
		j := i
		for ; j < n && s[j] == ')'; j++ {
		}
		r -= j - i
		for k := j - i; k >= 0; k-- {
			if cnt >= 0 && r >= 0 {
				backtrack(s, cur, j, n, cnt, l, r)
			}
			cur, cnt, r = cur+")", cnt-1, r+1
		}
	default:
		// 添加字符, 跳过. 不能删.
		backtrack(s, cur+s[i:i+1], i+1, n, cnt, l, r)
	}
	return
}
