/*
 * @lc app=leetcode.cn id=basic-calculator lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [224] 基本计算器
 *
 * https://leetcode-cn.com/problems/basic-calculator/description/
 *
 * algorithms
 * Hard (41.78%)
 * Total Accepted:    80K
 * Total Submissions: 191.4K
 * Testcase Example:  '"1 + 1"'
 *
 * 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
 *
 * 注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "1 + 1"
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = " 2-1 + 2 "
 * 输出：3
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "(1+(4+5+2)-3)+(6+8)"
 * 输出：23
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 3 * 10^5
 * s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
 * s 表示一个有效的表达式
 * '+' 不能用作一元运算(例如， "+1" 和 "+(2 + 3)" 无效)
 * '-' 可以用作一元运算(即 "-1" 和 "-(2 + 3)" 是有效的)
 * 输入中不存在两个连续的操作符
 * 每个数字和运行的计算将适合于一个有符号的 32位 整数
 *
 *
 */

import "strings"

/*
 *
 * var Snum []int
 * var Sop []byte
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
 * 		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
 * 			v := 0
 * 			j := i
 * 			for j < n && '0' <= bytes[j] && bytes[j] <= '9' {
 * 				v *= 10
 * 				v += int(bytes[j] - '0')
 * 				j++
 * 			}
 * 			Snum = append(Snum, v)
 * 			i = j - 1
 * 		case '+', '-':
 * 			if i > 0 && (bytes[i-1] == '(' || bytes[i-1] == '+' || bytes[i-1] == '-') {
 * 				Snum = append(Snum, 0)
 * 			}
 * 			for len(Sop) > 0 && Sop[len(Sop)-1] != '(' {
 * 				_calculate()
 * 			}
 * 			Sop = append(Sop, b)
 * 		default:
 * 			// unreachable
 * 		}
 * 	}
 * 	for len(Sop) > 0 {
 * 		_calculate()
 * 	}
 * 	return Snum[len(Snum)-1]
 * }
 *
 *
 * func _calculate() {
 * 	n, m := len(Snum), len(Sop)
 * 	if n < 2 || m < 1 {
 * 		return
 * 	}
 * 	a, b := Snum[n-2], Snum[n-1]
 * 	op := Sop[m-1]
 * 	Snum, Sop = Snum[:n-2], Sop[:m-1]
 * 	if op == '+' {
 * 		Snum = append(Snum, a+b)
 * 	} else {
 * 		Snum = append(Snum, a-b)
 * 	}
 * }
 */

// 参考 https://leetcode-cn.com/problems/basic-calculator/solution/shuang-zhan-jie-jue-tong-yong-biao-da-sh-olym/
// 一些细节：
//
// - 由于第一个数可能是负数，为了减少边界判断。一个小技巧是先往 nums 添加一个 0
// - 为防止 () 内出现的首个字符为运算符，将所有的空格去掉，并将 (- 替换为 (0-，(+ 替换为 (0+（当然也可以不进行这样的预处理，将这个处理逻辑放到循环里去做）

func calculate(s string) int {
	s = strings.Replace(s, " ", "", -1)
	arr := []byte(s)
	return _calc(arr, 0, len(arr))
}

func _calc(arr []byte, p, q int) int {
	var Sop []byte
	var Snum []int
	for ind := p; ind < q; ind++ {
		switch arr[ind] {
		case '+', '-':
			for len(Sop) > 0 {
				Sop, Snum = _calculate(Sop, Snum)
			}
			if ind == p {
				Snum = append(Snum, 0)
			}
			Sop = append(Sop, arr[ind])
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			v := 0
			right := ind
			for ; right < q && '0' <= arr[right] && arr[right] <= '9'; right++ {
				v = v*10 + int(arr[right]-'0')
			}
			Snum = append(Snum, v)
			ind = right - 1
		case '(':
			v := 1
			right := ind + 1
			for ; right < q && v != 0; right++ {
				switch arr[right] {
				case '(':
					v++
				case ')':
					v--
				}
			}
			c := _calc(arr, ind+1, right-1)
			Snum = append(Snum, c)
			ind = right - 1
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
	default:
		// unreachable
	}
	Snum = append(Snum, p)
	return Sop, Snum
}
