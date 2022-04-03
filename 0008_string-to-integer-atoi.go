import "math"

/*
 * @lc app=leetcode.cn id=string-to-integer-atoi lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [8] 字符串转换整数 (atoi)
 *
 * https://leetcode-cn.com/problems/string-to-integer-atoi/description/
 *
 * algorithms
 * Medium (21.77%)
 * Total Accepted:    411.1K
 * Total Submissions: 1.9M
 * Testcase Example:  '"42"'
 *
 * 请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。
 *
 * 函数 myAtoi(string s) 的算法如下：
 *
 *
 * 读入字符串并丢弃无用的前导空格
 * 检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
 * 读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
 * 将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤
 * 2 开始）。
 * 如果整数数超过 32 位有符号整数范围 [−2^31,  2^31 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −2^31
 * 的整数应该被固定为 −2^31 ，大于 2^31 − 1 的整数应该被固定为 2^31 − 1 。
 * 返回整数作为最终结果。
 *
 *
 * 注意：
 *
 *
 * 本题中的空白字符只包括空格字符 ' ' 。
 * 除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "42"
 * 输出：42
 * 解释：加粗的字符串为已经读入的字符，插入符号是当前读取的字符。
 * 第 1 步："42"（当前没有读入字符，因为没有前导空格）
 * ⁠        ^
 * 第 2 步："42"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
 * ⁠        ^
 * 第 3 步："42"（读入 "42"）
 * ⁠          ^
 * 解析得到整数 42 。
 * 由于 "42" 在范围 [-2^31, 2^31 - 1] 内，最终结果为 42 。
 *
 * 示例 2：
 *
 *
 * 输入：s = "   -42"
 * 输出：-42
 * 解释：
 * 第 1 步："   -42"（读入前导空格，但忽视掉）
 * ⁠           ^
 * 第 2 步："   -42"（读入 '-' 字符，所以结果应该是负数）
 * ⁠            ^
 * 第 3 步："   -42"（读入 "42"）
 * ⁠              ^
 * 解析得到整数 -42 。
 * 由于 "-42" 在范围 [-2^31, 2^31 - 1] 内，最终结果为 -42 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "4193 with words"
 * 输出：4193
 * 解释：
 * 第 1 步："4193 with words"（当前没有读入字符，因为没有前导空格）
 * ⁠        ^
 * 第 2 步："4193 with words"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
 * ⁠        ^
 * 第 3 步："4193 with words"（读入 "4193"；由于下一个字符不是一个数字，所以读入停止）
 * ⁠            ^
 * 解析得到整数 4193 。
 * 由于 "4193" 在范围 [-2^31, 2^31 - 1] 内，最终结果为 4193 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 200
 * s 由英文字母（大写和小写）、数字（0-9）、' '、'+'、'-' 和 '.' 组成
 *
 *
 */

/*
 * func myAtoi(s string) int {
 * 	strs := strings.Split(s, " ")
 * 	for _, s := range strs {
 * 		if len(s) == 0 {
 * 			continue
 * 		}
 * 		return atoi([]byte(s))
 * 	}
 * 	return 0
 * }
 *
 * func atoi(arr []byte) int {
 * 	n := len(arr)
 * 	if n == 0 {
 * 		return 0
 * 	}
 * 	neg := false
 * 	ind := 0
 * 	if arr[0] == '-' || arr[0] == '+' {
 * 		neg = arr[0] == '-'
 * 		ind++
 * 	}
 * 	ans := 0
 * 	for i := ind; i < n && '0' <= arr[i] && arr[i] <= '9'; i++ {
 * 		ans = ans*10 + int(arr[i]-'0')
 * 		if ans > 1<<31 {
 * 			if neg {
 * 				return -(1 << 31)
 * 			}
 * 			return (1 << 31) - 1
 * 		}
 * 	}
 * 	if neg {
 * 		ans *= -1
 * 	}
 * 	if ans > (1<<31)-1 {
 * 		return (1 << 31) - 1
 * 	}
 * 	if ans < -(1 << 31) {
 * 		return -(1 << 31)
 * 	}
 * 	return ans
 * }
 */

type State int

const (
	StateStart State = iota
	StateSign
	StateDigit
	StateEnd
)

var dfa = map[State][]State{
	StateStart: {StateStart, StateSign, StateDigit, StateEnd},
	StateSign:  {StateEnd, StateEnd, StateDigit, StateEnd},
	StateDigit: {StateEnd, StateEnd, StateDigit, StateEnd},
	StateEnd:   {StateEnd, StateEnd, StateEnd, StateEnd},
}

var state State
var sign, ans int

func get(b byte) {
	getState := func(b byte) State {
		switch b {
		case ' ':
			return StateStart
		case '+', '-':
			return StateSign
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return StateDigit
		default:
		}
		return StateEnd
	}
	state = dfa[state][getState(b)]
	switch state {
	case StateSign:
		if b == '-' {
			sign *= -1
		}
	case StateDigit:
		ans = ans*10 + int(b-'0')
		v := sign * ans
		if v > math.MaxInt32 {
			ans = math.MaxInt32
		}
		if v < math.MinInt32 {
			ans = -math.MinInt32
		}
	}
	return
}

// DFA
func myAtoi(s string) int {
	bytes := []byte(s)
	state = StateStart
	sign, ans = 1, 0
	for _, b := range bytes {
		get(b)
	}
	return sign * ans
}
