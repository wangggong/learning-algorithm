/*
 * @lc app=leetcode.cn id=valid-number lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [65] 有效数字
 *
 * https://leetcode-cn.com/problems/valid-number/description/
 *
 * algorithms
 * Hard (27.20%)
 * Total Accepted:    51.1K
 * Total Submissions: 187.8K
 * Testcase Example:  '"0"'
 *
 * 有效数字（按顺序）可以分成以下几个部分：
 *
 *
 * 一个 小数 或者 整数
 * （可选）一个 'e' 或 'E' ，后面跟着一个 整数
 *
 *
 * 小数（按顺序）可以分成以下几个部分：
 *
 *
 * （可选）一个符号字符（'+' 或 '-'）
 * 下述格式之一：
 *
 * 至少一位数字，后面跟着一个点 '.'
 * 至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
 * 一个点 '.' ，后面跟着至少一位数字
 *
 *
 *
 *
 * 整数（按顺序）可以分成以下几个部分：
 *
 *
 * （可选）一个符号字符（'+' 或 '-'）
 * 至少一位数字
 *
 *
 * 部分有效数字列举如下：["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3",
 * "3e+7", "+6e-1", "53.5e93", "-123.456e789"]
 *
 * 部分无效数字列举如下：["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"]
 *
 * 给你一个字符串 s ，如果 s 是一个 有效数字 ，请返回 true 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "0"
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "e"
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "."
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 20
 * s 仅含英文字母（大写和小写），数字（0-9），加号 '+' ，减号 '-' ，或者点 '.' 。
 *
 *
 */

/*
 * func isNumber(s string) bool {
 * 	bs := []byte(s)
 * 	n := len(bs)
 * 	if n == 0 {
 * 		return false
 * 	}
 * 	for i, b := range bs {
 * 		if b == 'e' || b == 'E' {
 * 			if i+1 == n {
 * 				return false
 * 			}
 * 			return validFloat(bs[:i]) && validInt(bs[i+1:], true)
 * 		}
 * 	}
 * 	return validFloat(bs)
 * }
 *
 * func validFloat(arr []byte) bool {
 * 	if len(arr) == 0 {
 * 		return false
 * 	}
 * 	ind, dot := 0, -1
 * 	for i, b := range arr {
 * 		if b == '.' {
 * 			dot = i
 * 			continue
 * 		}
 * 		if '0' <= b && b <= '9' {
 * 			continue
 * 		}
 * 		if i == 0 && (b == '+' || b == '-') {
 * 			ind++
 * 			continue
 * 		}
 * 		return false
 * 	}
 * 	if dot >= 0 {
 * 		intPart, floatPart := arr[ind:dot], arr[dot+1:]
 * 		if len(intPart) == 0 {
 * 			return validInt(floatPart, false)
 * 		}
 * 		if len(floatPart) == 0 {
 * 			return validInt(intPart, false)
 * 		}
 * 		return validInt(intPart, false) && validInt(floatPart, false)
 * 	}
 * 	return validInt(arr, true)
 * }
 *
 * func validInt(arr []byte, withSign bool) bool {
 * 	if len(arr) == 0 {
 * 		return false
 * 	}
 * 	hasSign := false
 * 	for i, b := range arr {
 * 		if i == 0 && (b == '+' || b == '-') && withSign {
 * 			hasSign = true
 * 			continue
 * 		}
 * 		if '0' <= b && b <= '9' {
 * 			continue
 * 		}
 * 		return false
 * 	}
 * 	if hasSign && len(arr) == 1 {
 * 		return false
 * 	}
 * 	return true
 * }
 */

// DFA: 相比于梳理状态与状态转移, 写逻辑可简单多了.
const (
	StateInvalid = iota
	StateStart
	StateSign
	StateInt
	StatePointWithInt
	StatePointWithoutInt
	StateFloat
	StateExp
	StateExpSign
	StateExpInt
)

const (
	TokenInvalid = iota
	TokenDigit
	TokenSign
	TokenPoint
	TokenExp
)

var dfa = [][]int{
	/* StateInvalid: */ {
		/* TokenInvalid: */ StateInvalid,
		/* TokenDigit: */ StateInvalid,
		/* TokenSign: */ StateInvalid,
		/* TokenPoint: */ StateInvalid,
		/* TokenExp: */ StateInvalid,
	},
	/* StateStart: */ {
		/* TokenInvalid: */ StateInvalid,
		/* TokenDigit: */ StateInt,
		/* TokenSign: */ StateSign,
		/* TokenPoint: */ StatePointWithoutInt,
		/* TokenExp: */ StateInvalid,
	},
	/* StateSign: */ {
		/* TokenInvalid: */ StateInvalid,
		/* TokenDigit: */ StateInt,
		/* TokenSign: */ StateInvalid,
		/* TokenPoint: */ StatePointWithoutInt,
		/* TokenExp: */ StateInvalid,
	},
	/* StateInt: */ {
		/* TokenInvalid: */ StateInvalid,
		/* TokenDigit: */ StateInt,
		/* TokenSign: */ StateInvalid,
		/* TokenPoint: */ StatePointWithInt,
		/* TokenExp: */ StateExp,
	},
	/* StatePointWithInt: */ {
		/* TokenInvalid: */ StateInvalid,
		/* TokenDigit: */ StateFloat,
		/* TokenSign: */ StateInvalid,
		/* TokenPoint: */ StateInvalid,
		/* TokenExp: */ StateExp,
	},
	/* StatePointWithoutInt: */ {
		/* TokenInvalid: */ StateInvalid,
		/* TokenDigit: */ StateFloat,
		/* TokenSign: */ StateInvalid,
		/* TokenPoint: */ StateInvalid,
		/* TokenExp: */ StateInvalid,
	},
	/* StateFloat: */ {
		/* TokenInvalid: */ StateInvalid,
		/* TokenDigit: */ StateFloat,
		/* TokenSign: */ StateInvalid,
		/* TokenPoint: */ StateInvalid,
		/* TokenExp: */ StateExp,
	},
	/* StateExp: */ {
		/* TokenInvalid: */ StateInvalid,
		/* TokenDigit: */ StateExpInt,
		/* TokenSign: */ StateExpSign,
		/* TokenPoint: */ StateInvalid,
		/* TokenExp: */ StateInvalid,
	},
	/* StateExpSign: */ {
		/* TokenInvalid: */ StateInvalid,
		/* TokenDigit: */ StateExpInt,
		/* TokenSign: */ StateInvalid,
		/* TokenPoint: */ StateInvalid,
		/* TokenExp: */ StateInvalid,
	},
	/* StateExpInt: */ {
		/* TokenInvalid: */ StateInvalid,
		/* TokenDigit: */ StateExpInt,
		/* TokenSign: */ StateInvalid,
		/* TokenPoint: */ StateInvalid,
		/* TokenExp: */ StateInvalid,
	},
}

func isNumber(s string) bool {
	arr := []byte(s)
	state := StateStart
	for _, a := range arr {
		state = dfa[state][tokenType(a)]
	}
	return validState(state)
}

func tokenType(ch byte) int {
	switch {
	case '0' <= ch && ch <= '9':
		return TokenDigit
	case ch == '+' || ch == '-':
		return TokenSign
	case ch == '.':
		return TokenPoint
	case ch == 'e' || ch == 'E':
		return TokenExp
	default:
		// do nothing
	}
	return TokenInvalid
}

func validState(s int) bool {
	switch s {
	case StateInt, StatePointWithInt, StateFloat, StateExpInt:
		return true
	default:
		// do nothing
	}
	return false
}
