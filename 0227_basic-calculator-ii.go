import "strings"

/*
 * @lc app=leetcode.cn id=basic-calculator-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [227] 基本计算器 II
 *
 * https://leetcode-cn.com/problems/basic-calculator-ii/description/
 *
 * algorithms
 * Medium (43.83%)
 * Total Accepted:    99.5K
 * Total Submissions: 226.8K
 * Testcase Example:  '"3+2*2"'
 *
 * 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
 *
 * 整数除法仅保留整数部分。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "3+2*2"
 * 输出：7
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = " 3/2 "
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = " 3+5 / 2 "
 * 输出：5
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 3 * 1e5
 * s 由整数和算符 ('+', '-', '*', '/') 组成，中间由一些空格隔开
 * s 表示一个 有效表达式
 * 表达式中的所有整数都是非负整数，且在范围 [0, 1 << 31 - 1] 内
 * 题目数据保证答案是一个 32-bit 整数
 *
 *
 *
 *
 */

var Snum []int
var Sop []byte
var priority = map[byte]int{
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
}

func calculate(s string) int {
	s = strings.Replace(s, " ", "", -1)
	bytes := []byte(s)
	n := len(bytes)
	Snum, Sop = nil, nil
	Snum = append(Snum, 0)
	for i := 0; i < n; i++ {
		b := bytes[i]
		switch b {
		case '0', '1', '2', '3', '4',
			'5', '6', '7', '8', '9':
			v, j := 0, i
			for j < n && '0' <= bytes[j] && bytes[j] <= '9' {
				v *= 10
				v += int(bytes[j] - '0')
				j++
			}
			Snum = append(Snum, v)
			i = j - 1
		case '+', '-', '*', '/':
			if i > 0 && (bytes[i-1] == '+' || bytes[i-1] == '-') {
				Snum = append(Snum, 0)
			}
			for len(Sop) > 0 {
				if priority[Sop[len(Sop)-1]] >= priority[b] {
					_calculate()
				} else {
					break
				}
			}
			Sop = append(Sop, b)
		default:
			// unreachable
		}
	}
	for len(Sop) > 0 {
		_calculate()
	}
	return Snum[len(Snum)-1]
}

func _calculate() {
	n, m := len(Snum), len(Sop)
	if n < 2 || m < 1 {
		return
	}
	a, b := Snum[n-2], Snum[n-1]
	op := Sop[m-1]
	Snum, Sop = Snum[:n-2], Sop[:m-1]
	switch op {
	case '+':
		a += b
	case '-':
		a -= b
	case '*':
		a *= b
	case '/':
		a /= b
	default:
		// unreachable
	}
	Snum = append(Snum, a)

}
