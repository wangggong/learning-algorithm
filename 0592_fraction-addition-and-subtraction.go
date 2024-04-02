/*
 * @lc app=leetcode.cn id=fraction-addition-and-subtraction lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [592] 分数加减运算
 *
 * https://leetcode.cn/problems/fraction-addition-and-subtraction/description/
 *
 * algorithms
 * Medium (52.82%)
 * Total Accepted:    18.9K
 * Total Submissions: 31.9K
 * Testcase Example:  '"-1/2+1/2"'
 *
 * 给定一个表示分数加减运算的字符串 expression ，你需要返回一个字符串形式的计算结果。
 *
 * 这个结果应该是不可约分的分数，即最简分数。 如果最终结果是一个整数，例如 2，你需要将它转换成分数形式，其分母为 1。所以在上述例子中, 2
 * 应该被转换为 2/1。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: expression = "-1/2+1/2"
 * 输出: "0/1"
 *
 *
 * 示例 2:
 *
 *
 * 输入: expression = "-1/2+1/2+1/3"
 * 输出: "1/3"
 *
 *
 * 示例 3:
 *
 *
 * 输入: expression = "1/3-1/2"
 * 输出: "-1/6"
 *
 *
 *
 *
 * 提示:
 *
 *
 * 输入和输出字符串只包含 '0' 到 '9' 的数字，以及 '/', '+' 和 '-'。
 * 输入和输出分数格式均为 ±分子/分母。如果输入的第一个分数或者输出的分数是正数，则 '+' 会被省略掉。
 * 输入只包含合法的最简分数，每个分数的分子与分母的范围是  [1,10]。 如果分母是1，意味着这个分数实际上是一个整数。
 * 输入的分数个数范围是 [1,10]。
 * 最终结果的分子与分母保证是 32 位整数范围内的有效整数。
 *
 *
 */
import (
	"fmt"
	"strconv"
	"strings"
)

func fractionAddition(expr string) string {
	// assert len(expr) > 0
	if expr[0] == '-' {
		expr = "0/1" + expr
	}
	expr = strings.Replace(expr, "-", "+-", -1)
	vals := strings.Split(expr, "+")
	v := vals[0]
	for i, n := 1, len(vals); i < n; i++ {
		v = calc(v, vals[i])
	}
	return v
}

func calc(s, t string) string {
	sNum, sDen := getFrac(s)
	tNum, tDen := getFrac(t)
	num, den := sNum*tDen+tNum*sDen, sDen*tDen
	// fmt.Println(s, t, sNum, sDen, tNum, tDen, num, den)
	d := gcd(den, abs(num))
	return fmt.Sprintf("%v/%v", num/d, den/d)
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func getFrac(s string) (int, int) {
	ind := strings.IndexByte(s, '/')
	num, _ := strconv.Atoi(s[:ind])
	den, _ := strconv.Atoi(s[ind+1:])
	return num, den
}
