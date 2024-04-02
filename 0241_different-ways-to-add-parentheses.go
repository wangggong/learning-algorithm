/*
 * @lc app=leetcode.cn id=different-ways-to-add-parentheses lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [241] 为运算表达式设计优先级
 *
 * https://leetcode.cn/problems/different-ways-to-add-parentheses/description/
 *
 * algorithms
 * Medium (73.66%)
 * Total Accepted:    50K
 * Total Submissions: 67.2K
 * Testcase Example:  '"2-1-1"'
 *
 * 给你一个由数字和运算符组成的字符串 expression ，按不同优先级组合数字和运算符，计算并返回所有可能组合的结果。你可以 按任意顺序 返回答案。
 *
 * 生成的测试用例满足其对应输出值符合 32 位整数范围，不同结果的数量不超过 10^4 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：expression = "2-1-1"
 * 输出：[0,2]
 * 解释：
 * ((2-1)-1) = 0
 * (2-(1-1)) = 2
 *
 *
 * 示例 2：
 *
 *
 * 输入：expression = "2*3-4*5"
 * 输出：[-34,-14,-10,-10,10]
 * 解释：
 * (2*(3-(4*5))) = -34
 * ((2*3)-(4*5)) = -14
 * ((2*(3-4))*5) = -10
 * (2*((3-4)*5)) = -10
 * (((2*3)-4)*5) = 10
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= expression.length <= 20
 * expression 由数字和算符 '+'、'-' 和 '*' 组成。
 * 输入表达式中的所有整数值在范围 [0, 99]
 *
 *
 */
func diffWaysToCompute(expr string) []int {
	var arr []string
	for p, q, n := 0, 0, len(expr); p < n; p = q + 1 {
		for q = p; q < n && '0' <= expr[q] && expr[q] <= '9'; q++ {
		}
		arr = append(arr, expr[p:q])
		if q == n {
			break
		}
		arr = append(arr, expr[q:q+1])
	}
	return dfs(arr)
}

func dfs(arr []string) []int {
	n := len(arr)
	// assert len(arr) > 0
	var ans []int
	if n == 1 {
		x := 0
		for _, a := range arr[0] {
			x = x*10 + int(a-'0')
		}
		ans = append(ans, x)
		return ans
	}
	for i := 1; i < n; i += 2 {
		left, right := dfs(arr[:i]), dfs(arr[i+1:])
		for _, l := range left {
			for _, r := range right {
				switch arr[i] {
				case "+":
					ans = append(ans, l+r)
				case "-":
					ans = append(ans, l-r)
				case "*":
					ans = append(ans, l*r)
				default:
					// unreachable
				}
			}
		}
	}
	return ans
}
