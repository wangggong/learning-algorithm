/*
 * Hard
 *
 * 实现一个基本的计算器来计算简单的表达式字符串。
 *
 * 表达式字符串只包含非负整数，算符 +、-、*、/ ，左括号 ( 和右括号 ) 。整数除法需要 向下截断 。
 *
 * 你可以假定给定的表达式总是有效的。所有的中间结果的范围为 [-231, 231 - 1] 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "1+1"
 * 输出：2
 * 示例 2：
 *
 * 输入：s = "6-4/2"
 * 输出：4
 * 示例 3：
 *
 * 输入：s = "2*(5+5*2)/3+(6/2+8)"
 * 输出：21
 * 示例 4：
 *
 * 输入：s = "(2+6*3+5-(3*14/7+2)*5)+3"
 * 输出：-12
 * 示例 5：
 *
 * 输入：s = "0"
 * 输出：0
 *
 *
 * 提示：
 *
 * 1 <= s <= 1e4
 * s 由整数、'+'、'-'、'*'、'/'、'(' 和 ')' 组成
 * s 是一个 有效的 表达式
 *
 * 通过次数7,399
 * 提交次数14,759
 */

import "strings"

var priority = map[byte]int{
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
}

/*
 * var Snum []int
 * var Sop []byte
 * }
 *
 * func calculate(s string) int {
 * 	s = strings.Replace(s, " ", "", -1)
 * 	bytes := []byte(s)
 * 	n := len(bytes)
 * 	Snum, Sop = nil, nil
 * 	Snum = append(Snum, 0)
 * 	for i := 0; i < n; i++ {
 * 		b := bytes[i]
 * 		switch b {
 * 		case '(':
 * 			Sop = append(Sop, b)
 * 		case ')':
 * 			for len(Sop) > 0 && Sop[len(Sop)-1] != '(' {
 * 				_calculate()
 * 			}
 * 			Sop = Sop[:len(Sop)-1]
 * 		case '0', '1', '2', '3', '4',
 * 			'5', '6', '7', '8', '9':
 * 			v, j := 0, i
 * 			for j < n && '0' <= bytes[j] && bytes[j] <= '9' {
 * 				v *= 10
 * 				v += int(bytes[j] - '0')
 * 				j++
 * 			}
 * 			Snum = append(Snum, v)
 * 			i = j - 1
 * 		case '+', '-', '*', '/':
 * 			if i > 0 && (bytes[i-1] == '(' || bytes[i-1] == '-' || bytes[i-1] == '+') {
 * 				Snum = append(Snum, 0)
 * 			}
 * 			for len(Sop) > 0 && Sop[len(Sop)-1] != '(' {
 * 				if priority[Sop[len(Sop)-1]] >= priority[b] {
 * 					_calculate()
 * 				} else {
 * 					break
 * 				}
 * 			}
 * 			// 忘了入栈了
 * 			Sop = append(Sop, b)
 * 		default:
 * 		}
 * 	}
 * 	for len(Sop) > 0 {
 * 		_calculate()
 * 	}
 * 	return Snum[len(Snum)-1]
 * }
 *
 * func _calculate() {
 * 	n, m := len(Snum), len(Sop)
 * 	if n < 2 || m < 1 {
 * 		return
 * 	}
 * 	a, b := Snum[n-2], Snum[n-1]
 * 	op := Sop[m-1]
 * 	Snum, Sop = Snum[:n-2], Sop[:m-1]
 * 	switch op {
 * 	case '+':
 * 		a += b
 * 	case '-':
 * 		a -= b
 * 	case '*':
 * 		a *= b
 * 	case '/':
 * 		a /= b
 * 	default:
 * 	}
 * 	Snum = append(Snum, a)
 * }
 */

func calculate(s string) int {
	s = strings.Replace(s, " ", "", -1)
	arr := []byte(s)
	return _calc(arr, 0, len(arr))
}

func _calc(arr []byte, p, q int) int {
	var Sop []byte
	var Snum []int
	for i := p; i < q; i++ {
		switch arr[i] {
		case '+', '-', '*', '/':
			for len(Sop) > 0 && priority[Sop[len(Sop)-1]] >= priority[arr[i]] {
				Sop, Snum = _calculate(Sop, Snum)
			}
			if i == p && (arr[i] == '+' || arr[i] == '-') {
				Snum = append(Snum, 0)
			}
			Sop = append(Sop, arr[i])
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			v := 0
			j := i
			for ; j < q && '0' <= arr[j] && arr[j] <= '9'; j++ {
				v = v*10 + int(arr[j]-'0')
			}
			Snum = append(Snum, v)
			i = j - 1
		case '(':
			v := 1
			j := i + 1
			for ; j < q && v != 0; j++ {
				switch arr[j] {
				case '(':
					v++
				case ')':
					v--
				default:
					// do nothing
				}
			}
			c := _calc(arr, i+1, j-1)
			Snum = append(Snum, c)
			i = j - 1
		default:
			// unreachable
		}
	}
	for len(Sop) > 0 {
		Sop, Snum = _calculate(Sop, Snum)
	}
	return Snum[len(Snum)-1]
}

func _calculate(Sop []byte, Snum []int) ([]byte, []int) {
	n, m := len(Sop), len(Snum)
	if n <= 0 || m <= 1 {
		return Sop, Snum
	}
	p, q := Snum[m-2], Snum[m-1]
	op := Sop[n-1]
	Sop, Snum = Sop[:n-1], Snum[:m-2]
	switch op {
	case '+':
		p += q
	case '-':
		p -= q
	case '*':
		p *= q
	case '/':
		p /= q
	default:
		// unreachable
	}
	Snum = append(Snum, p)
	return Sop, Snum
}
