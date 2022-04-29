/*
 * @lc app=leetcode.cn id=minimize-result-by-adding-parentheses-to-expression lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6038] 向表达式添加括号后的最小结果
 *
 * https://leetcode-cn.com/problems/minimize-result-by-adding-parentheses-to-expression/description/
 *
 * algorithms
 * Medium (59.20%)
 * Total Accepted:    5K
 * Total Submissions: 8.4K
 * Testcase Example:  '"247+38"'
 *
 * 给你一个下标从 0 开始的字符串 expression ，格式为 "<num1>+<num2>" ，其中 <num1> 和 <num2> 表示正整数。
 *
 * 请你向 expression 中添加一对括号，使得在添加之后， expression 仍然是一个有效的数学表达式，并且计算后可以得到 最小
 * 可能值。左括号 必须 添加在 '+' 的左侧，而右括号必须添加在 '+' 的右侧。
 *
 * 返回添加一对括号后形成的表达式 expression ，且满足 expression 计算得到 最小
 * 可能值。如果存在多个答案都能产生相同结果，返回任意一个答案。
 *
 * 生成的输入满足：expression 的原始值和添加满足要求的任一对括号之后 expression 的值，都符合 32-bit 带符号整数范围。
 *
 *
 *
 * 示例 1：
 *
 * 输入：expression = "247+38"
 * 输出："2(47+38)"
 * 解释：表达式计算得到 2 * (47 + 38) = 2 * 85 = 170 。
 * 注意 "2(4)7+38" 不是有效的结果，因为右括号必须添加在 '+' 的右侧。
 * 可以证明 170 是最小可能值。
 *
 *
 * 示例 2：
 *
 * 输入：expression = "12+34"
 * 输出："1(2+3)4"
 * 解释：表达式计算得到 1 * (2 + 3) * 4 = 1 * 5 * 4 = 20 。
 *
 *
 * 示例 3：
 *
 * 输入：expression = "999+999"
 * 输出："(999+999)"
 * 解释：表达式计算得到 999 + 999 = 1998 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= expression.length <= 10
 * expression 仅由数字 '1' 到 '9' 和 '+' 组成
 * expression 由数字开始和结束
 * expression 恰好仅含有一个 '+'.
 * expression 的原始值和添加满足要求的任一对括号之后 expression 的值，都符合 32-bit 带符号整数范围
 *
 *
 */
import "math"

func minimizeResult(s string) string {
	arr := []byte(s)
	n := len(arr)
	cur := math.MaxInt32
	ans := ""
	pos := 0
	for ; pos < n; pos++ {
		if arr[pos] == '+' {
			break
		}
	}
	for i := 0; i < pos; i++ {
		for j := pos + 2; j <= n; j++ {
			if v := getResult(arr, i, j, pos); v < cur {
				ans = string(arr[:i]) + "(" + string(arr[i:j]) + ")" + string(arr[j:])
				cur = v
			}
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getResult(arr []byte, p, q, pos int) int {
	k1, k2 := atoi(arr[:p]), atoi(arr[q:])
	if k1 == 0 {
		k1 = 1
	}
	if k2 == 0 {
		k2 = 1
	}
	k := atoi(arr[p:pos]) + atoi(arr[pos+1:q])
	return k1 * k2 * k
}

func atoi(arr []byte) int {
	ans := 0
	for _, a := range arr {
		ans = ans*10 + int(a-'0')
	}
	return ans
}
