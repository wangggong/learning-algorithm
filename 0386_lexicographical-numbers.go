/*
 * @lc app=leetcode.cn id=lexicographical-numbers lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [386] 字典序排数
 *
 * https://leetcode-cn.com/problems/lexicographical-numbers/description/
 *
 * algorithms
 * Medium (75.35%)
 * Total Accepted:    29.3K
 * Total Submissions: 38.9K
 * Testcase Example:  '13'
 *
 * 给你一个整数 n ，按字典序返回范围 [1, n] 内所有整数。
 *
 * 你必须设计一个时间复杂度为 O(n) 且使用 O(1) 额外空间的算法。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 13
 * 输出：[1,10,11,12,13,2,3,4,5,6,7,8,9]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 2
 * 输出：[1,2]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 5 * 10^4
 *
 *
 */

// 参考: https://leetcode-cn.com/problems/lexicographical-numbers/solution/386-zi-dian-xu-pai-shu-o1-kong-jian-fu-z-aea2/
//
// 用 DFS 模拟遍历 Trie 的过程...
var ans []int

/*
 * func lexicalOrder(n int) []int {
 * 	ans = nil
 * 	for i := 1; i <= 9; i++ {
 * 		dfs(n, i)
 * 	}
 * 	return ans
 * }
 */

func dfs(n, k int) {
	if k > n {
		return
	}
	ans = append(ans, k)
	for i := k * 10; i <= k * 10 + 9; i++ {
		dfs(n, i)
	}
}

// 迭代
func lexicalOrder(n int) []int {
	ans = nil
	k := 1
	for len(ans) < n {
		for k <= n {
			ans = append(ans, k)
			k *= 10
		}
		for k % 10 == 9 || k > n {
			k /= 10
		}
		k++
	}
	return ans
}
