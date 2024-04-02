/*
 * @lc app=leetcode.cn id=solve-the-equation lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [640] 求解方程
 *
 * https://leetcode.cn/problems/solve-the-equation/description/
 *
 * algorithms
 * Medium (42.45%)
 * Total Accepted:    28K
 * Total Submissions: 63.2K
 * Testcase Example:  '"x+5-3+x=6+x-2"'
 *
 * 求解一个给定的方程，将x以字符串 "x=#value" 的形式返回。该方程仅包含 '+' ， '-' 操作，变量 x 和其对应系数。
 *
 * 如果方程没有解，请返回 "No solution" 。如果方程有无限解，则返回 “Infinite solutions” 。
 *
 * 题目保证，如果方程中只有一个解，则 'x' 的值是一个整数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: equation = "x+5-3+x=6+x-2"
 * 输出: "x=2"
 *
 *
 * 示例 2:
 *
 *
 * 输入: equation = "x=x"
 * 输出: "Infinite solutions"
 *
 *
 * 示例 3:
 *
 *
 * 输入: equation = "2x=x"
 * 输出: "x=0"
 *
 *
 *
 *
 * 提示:
 *
 *
 * 3 <= equation.length <= 1000
 * equation 只有一个 '='.
 * equation 方程由整数组成，其绝对值在 [0, 100] 范围内，不含前导零和变量 'x' 。 ​​​
 *
 *
 */

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	NoSolution   = "No solution"
	InfSolutions = "Infinite solutions"
)

func solveEquation(s string) string {
	k := strings.IndexByte(s, '=')
	p, q := getPoly(s[:k]), getPoly(s[k+1:])
	if p[1] == q[1] {
		if p[0] == q[0] {
			return InfSolutions
		}
		return NoSolution
	}
	return fmt.Sprintf("x=%v", (p[0]-q[0])/(q[1]-p[1]))
}

func getPoly(s string) (ans [2]int) {
	s = "+" + s
	s = strings.Replace(s, "+x", "+1x", -1)
	s = strings.Replace(s, "-x", "-1x", -1)
	s = strings.Replace(s, "-", "+-", -1)
	arr := strings.Split(s, "+")
	for _, a := range arr[1:] {
		k := 0
		if strings.IndexByte(a, 'x') >= 0 {
			k, a = 1, a[:len(a)-1]
		}
		v, _ := strconv.Atoi(a)
		ans[k] += v
	}
	return
}
