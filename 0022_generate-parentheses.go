/*
 * @lc app=leetcode.cn id=generate-parentheses lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (77.29%)
 * Total Accepted:    419.4K
 * Total Submissions: 542.3K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：["((()))","(()())","(())()","()(())","()()()"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：["()"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 8
 *
 *
 */
var m map[int][]string

func generateParenthesis(n int) []string {
	// assert n > 0;
	m = make(map[int][]string)
	return dfs(n)
}

func dfs(n int) []string {
	if n == 0 {
		return []string{""}
	}
	if v, ok := m[n]; ok {
		return v
	}
	var ans []string
	defer func() { m[n] = ans }()
	if n == 1 {
		ans = []string{"()"}
		return ans
	}
	for k := 0; k < n; k++ {
		left := dfs(k)
		right := dfs(n - k - 1)
		for _, l := range left {
			for _, r := range right {
				ans = append(ans, l+"("+r+")")
			}
		}
	}
	return ans
}
