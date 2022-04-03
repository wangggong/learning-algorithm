/*
 * @lc app=leetcode.cn id=complex-number-multiplication lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [537] 复数乘法
 *
 * https://leetcode-cn.com/problems/complex-number-multiplication/description/
 *
 * algorithms
 * Medium (71.39%)
 * Total Accepted:    20.4K
 * Total Submissions: 27.6K
 * Testcase Example:  '"1+1i"\n"1+1i"'
 *
 * 复数 可以用字符串表示，遵循 "实部+虚部i" 的形式，并满足下述条件：
 *
 *
 * 实部 是一个整数，取值范围是 [-100, 100]
 * 虚部 也是一个整数，取值范围是 [-100, 100]
 * i^2 == -1
 *
 *
 * 给你两个字符串表示的复数 num1 和 num2 ，请你遵循复数表示形式，返回表示它们乘积的字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：num1 = "1+1i", num2 = "1+1i"
 * 输出："0+2i"
 * 解释：(1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i ，你需要将它转换为 0+2i 的形式。
 *
 *
 * 示例 2：
 *
 *
 * 输入：num1 = "1+-1i", num2 = "1+-1i"
 * 输出："0+-2i"
 * 解释：(1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i ，你需要将它转换为 0+-2i 的形式。
 *
 *
 *
 *
 * 提示：
 *
 *
 * num1 和 num2 都是有效的复数表示。
 *
 *
 */

import (
	"fmt"
	"strings"
)

type Complex struct {
	Real int
	Imag int
}

func (c Complex) String() string {
	return fmt.Sprintf("%v+%vi", c.Real, c.Imag)
}

func (c Complex) Mul(com Complex) Complex {
	return Complex{
		Real: c.Real*com.Real - c.Imag*com.Imag,
		Imag: c.Real*com.Imag + c.Imag*com.Real,
	}
}

func String2Complex(s string) Complex {
	strs := strings.Split(s, "+")
	// assert len(strs) == 2;
	strs[0], strs[1] = strings.Trim(strs[0], " "), strings.Trim(strs[1], " ")
	bytes := [][]byte{[]byte(strs[0]), []byte(strs[1])}
	if bytes[1][len(bytes[1])-1] == 'i' {
		return Complex{
			Real: atoi(bytes[0]),
			Imag: atoi(bytes[1][:len(bytes[1])-1]),
		}
	}
	return Complex{
		Real: atoi(bytes[1]),
		Imag: atoi(bytes[0][:len(bytes[0])-1]),
	}

}

func atoi(arr []byte) int {
	ans := 0
	neg := false
	for i, a := range arr {
		if i == 0 && (a == '-' || a == '+') {
			neg = a == '-'
			continue
		}
		ans = ans*10 + int(a-'0')
	}
	if neg {
		ans *= -1
	}
	return ans
}

func complexNumberMultiply(num1 string, num2 string) string {
	c1, c2 := String2Complex(num1), String2Complex(num2)
	// fmt.Printf("%+v %+v\n", c1, c2)
	c := c1.Mul(c2)
	return c.String()
}
