/*
 * @lc app=leetcode.cn id=count-number-of-maximum-bitwise-or-subsets lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2044] 统计按位或能得到最大值的子集数目
 *
 * https://leetcode-cn.com/problems/count-number-of-maximum-bitwise-or-subsets/description/
 *
 * algorithms
 * Medium (74.41%)
 * Total Accepted:    13.4K
 * Total Submissions: 16.6K
 * Testcase Example:  '[3,1]'
 *
 * 给你一个整数数组 nums ，请你找出 nums 子集 按位或 可能得到的 最大值 ，并返回按位或能得到最大值的 不同非空子集的数目 。
 *
 * 如果数组 a 可以由数组 b 删除一些元素（或不删除）得到，则认为数组 a 是数组 b 的一个 子集 。如果选中的元素下标位置不一样，则认为两个子集
 * 不同 。
 *
 * 对数组 a 执行 按位或 ，结果等于 a[0] OR a[1] OR ... OR a[a.length - 1]（下标从 0 开始）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [3,1]
 * 输出：2
 * 解释：子集按位或能得到的最大值是 3 。有 2 个子集按位或可以得到 3 ：
 * - [3]
 * - [3,1]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [2,2,2]
 * 输出：7
 * 解释：[2,2,2] 的所有非空子集的按位或都可以得到 2 。总共有 2^3 - 1 = 7 个子集。
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [3,2,1,5]
 * 输出：6
 * 解释：子集按位或可能的最大值是 7 。有 6 个子集按位或可以得到 7 ：
 * - [3,5]
 * - [3,1,5]
 * - [3,2,5]
 * - [3,2,1,5]
 * - [2,5]
 * - [2,1,5]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 16
 * 1 <= nums[i] <= 10^5
 *
 *
 */

const MAXN = 1 << 16

var n int
var maxOR int

/*
 * func countMaxOrSubsets(nums []int) int {
 * 	maxOR = 0
 * 	n = len(nums)
 * 	lim := 1 << n
 * 	ans := 0
 * 	for i := 0; i < lim; i++ {
 * 		subset := getSubset(nums, i)
 * 		OR := getOR(subset)
 * 		if OR > maxOR {
 * 			maxOR = OR
 * 			ans = 1
 * 		} else if OR == maxOR {
 * 			ans++
 * 		}
 * 	}
 * 	return ans
 * }
 */

var dp [MAXN]int

/*
 * // 状压 DP
 * func countMaxOrSubsets(nums []int) int {
 * 	for i := range dp {
 * 		dp[i] = 0
 * 	}
 * 	n = len(nums)
 * 	lim := 1 << n
 * 	for i, k := range nums {
 * 		dp[1<<i] = k
 * 	}
 * 	for s := 0; s < lim; s++ {
 * 		for i := 0; i < n; i++ {
 * 			if s&(1<<i) != 0 {
 * 				continue
 * 			}
 * 			dp[s|(1<<i)] = dp[s] | nums[i]
 * 		}
 * 	}
 * 	ans, maxOR := 0, 0
 * 	for i := 0; i < lim; i++ {
 * 		if dp[i] > maxOR {
 * 			maxOR = dp[i]
 * 			ans = 1
 * 		} else if dp[i] == maxOR {
 * 			ans++
 * 		}
 * 	}
 * 	return ans
 * }
 */

// Codeforces 式脑筋急转弯, 但好像时间复杂度还是 `O(2^n)`, 常数优化恐怖如斯.
func countMaxOrSubsets(nums []int) int {
	all := 0
	for _, n := range nums {
		all = all | n
	}
	lim := 1 << len(nums)
	dp := make([]int, lim)
	ans := 0
	for i, n := range nums {
		k := 1 << i
		for j := 0; j < k; j++ {
			dp[j|k] = dp[j] | n
			if dp[j|k] == all {
				ans++
			}
		}
	}
	return ans
}

func getOR(arr []int) int {
	ans := 0
	for _, a := range arr {
		ans = ans | a
	}
	return ans
}

func getSubset(arr []int, k int) []int {
	var ans []int
	for i := 0; i < n; i++ {
		if k&(1<<i) != 0 {
			ans = append(ans, arr[i])
		}
	}
	return ans
}
