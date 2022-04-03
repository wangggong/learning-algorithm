/*
 * @lc app=leetcode.cn id=find-the-minimum-number-of-fibonacci-numbers-whose-sum-is-k lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1414] 和为 K 的最少斐波那契数字数目
 *
 * https://leetcode-cn.com/problems/find-the-minimum-number-of-fibonacci-numbers-whose-sum-is-k/description/
 *
 * algorithms
 * Medium (60.67%)
 * Total Accepted:    14.8K
 * Total Submissions: 22.3K
 * Testcase Example:  '7'
 *
 * 给你数字 k ，请你返回和为 k 的斐波那契数字的最少数目，其中，每个斐波那契数字都可以被使用多次。
 *
 * 斐波那契数字定义为：
 *
 *
 * F1 = 1
 * F2 = 1
 * Fn = Fn-1 + Fn-2 ， 其中 n > 2 。
 *
 *
 * 数据保证对于给定的 k ，一定能找到可行解。
 *
 *
 *
 * 示例 1：
 *
 * 输入：k = 7
 * 输出：2
 * 解释：斐波那契数字为：1，1，2，3，5，8，13，……
 * 对于 k = 7 ，我们可以得到 2 + 5 = 7 。
 *
 * 示例 2：
 *
 * 输入：k = 10
 * 输出：2
 * 解释：对于 k = 10 ，我们可以得到 2 + 8 = 10 。
 *
 *
 * 示例 3：
 *
 * 输入：k = 19
 * 输出：3
 * 解释：对于 k = 19 ，我们可以得到 1 + 5 + 13 = 19 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= 10^9
 *
 *
 */
// import "fmt"

const MAXN = 1e9 + 10

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func findMinFibonacciNumbers(k int) int {
	F := []int{1}
	prev, curr := 1, 2
	for curr <= k {
		F = append(F, curr)
		next := curr + prev
		prev, curr = curr, next
	}
	// fmt.Printf("%v\n", F)
	ans, idx := 0, len(F)-1
	for k > 0 {
		for idx > 0 && F[idx] > k {
			idx--
		}
		// fmt.Printf("%v %v\n", k, F[idx])
		k -= F[idx]
		ans++
	}
	return ans
}
