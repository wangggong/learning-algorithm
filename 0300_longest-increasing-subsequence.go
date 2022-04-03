/*
 * @lc app=leetcode.cn id=longest-increasing-subsequence lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [300] 最长递增子序列
 *
 * https://leetcode-cn.com/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (52.38%)
 * Total Accepted:    420.7K
 * Total Submissions: 801.8K
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
 *
 * 子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7]
 * 的子序列。
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [10,9,2,5,3,7,101,18]
 * 输出：4
 * 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1,0,3,2,3]
 * 输出：4
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [7,7,7,7,7,7,7]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2500
 * -1e4 <= nums[i] <= 1e4
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 你可以设计时间复杂度为 O(n^2) 的解决方案吗？
 * 你能将算法的时间复杂度降低到 O(n log(n)) 吗?
 *
 *
 */
// import "fmt"

const MAXN = 2500 + 10

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
 * func lengthOfLIS(N []int) int {
 * 	var dp [MAXN]int
 * 	n := len(N)
 * 	ans := 0
 * 	for i := 0; i < n; i++ {
 * 		dp[i] = 1
 * 		for j := 0; j < i; j++ {
 * 			if N[i] <= N[j] {
 * 				continue
 * 			}
 * 			dp[i] = max(dp[i], dp[j]+1)
 * 		}
 * 		ans = max(ans, dp[i])
 * 	}
 * 	return ans
 * }
 */

func lengthOfLIS(N []int) int {
	n := len(N)
	ans := 1
	if n == 0 {
		return 0
	}
	var d [MAXN]int
	d[ans] = 0
	for i := 1; i < n; i++ {
		if N[i] > N[d[ans]] {
			ans++
			d[ans] = i
			continue
		}
		l, r := 1, ans
		pos := 0
		for l <= r {
			m := (l + r) >> 1
			if N[d[m]] < N[i] {
				pos = m
				l = m + 1
			} else {
				r = m - 1
			}
		}
		d[pos+1] = i
	}
	// fmt.Printf("%v\n", d)
	return ans
}
