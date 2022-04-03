/*
 * @lc app=leetcode.cn id=combination-sum-iv lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [377] 组合总和 Ⅳ
 *
 * https://leetcode-cn.com/problems/combination-sum-iv/description/
 *
 * algorithms
 * Medium (51.57%)
 * Total Accepted:    75.8K
 * Total Submissions: 147K
 * Testcase Example:  '[1,2,3]\n4'
 *
 * 给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
 *
 * 题目数据保证答案符合 32 位整数范围。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3], target = 4
 * 输出：7
 * 解释：
 * 所有可能的组合为：
 * (1, 1, 1, 1)
 * (1, 1, 2)
 * (1, 2, 1)
 * (1, 3)
 * (2, 1, 1)
 * (2, 2)
 * (3, 1)
 * 请注意，顺序不同的序列被视作不同的组合。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [9], target = 3
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 200
 * 1 <= nums[i] <= 1000
 * nums 中的所有元素 互不相同
 * 1 <= target <= 1000
 *
 *
 *
 *
 * 进阶：如果给定的数组中含有负数会发生什么？问题会产生何种变化？如果允许负数出现，需要向题目中添加哪些限制条件？
 *
 */
const MAXN = 1e3 + 5

/*
 * var dp [MAXN][MAXN]int
 *
 * func combinationSum4(nums []int, target int) int {
 * 	// n := len(nums)
 * 	// assert n > 0 && n <= 200;
 * 	// assert target > 0 && target <= 1000;
 * 	for i := 0; i < MAXN; i++ {
 * 		for j := 0; j < MAXN; j++ {
 * 			dp[i][j] = 0
 * 		}
 * 	}
 * 	dp[0][0] = 1
 * 	// 枚举结果的总长度, `1 ~ target` (注意上限不是 `n`, 数字可以复用!)
 * 	for i := 1; i <= target; i++ {
 * 		// 枚举该长度下和为指定值的方案数目
 * 		for j := 1; j <= target; j++ {
 * 			for _, n := range nums {
 * 				if j-n < 0 {
 * 					continue
 * 				}
 * 				dp[i][j] += dp[i-1][j-n]
 * 			}
 * 		}
 * 	}
 * 	ans := 0
 * 	for i := 1; i <= target; i++ {
 * 		ans += dp[i][target]
 * 	}
 * 	return ans
 * }
 */

var dp [MAXN]int

func combinationSum4(nums []int, target int) int {
	for i := range dp {
		dp[i] = 0
	}
	dp[0] = 1
	/*
	 * 如果求组合数就是外层 for 循环遍历物品, 内层 for 循环遍历背包.
	 * 如果求排列数就是外层 for 循环遍历背包, 内层 for 循环遍历物品.
	 * 如果求组合数就是外层 for 循环遍历物品, 内层 for 循环遍历背包.
	 * 如果求排列数就是外层 for 循环遍历背包, 内层 for 循环遍历物品.
	 * 如果求组合数就是外层 for 循环遍历物品, 内层 for 循环遍历背包.
	 * 如果求排列数就是外层 for 循环遍历背包, 内层 for 循环遍历物品.
	 * 如果求组合数就是外层 for 循环遍历物品, 内层 for 循环遍历背包.
	 * 如果求排列数就是外层 for 循环遍历背包, 内层 for 循环遍历物品.
	 * 如果求组合数就是外层 for 循环遍历物品, 内层 for 循环遍历背包.
	 * 如果求排列数就是外层 for 循环遍历背包, 内层 for 循环遍历物品.
	 */
	// 循环内外层对换了, 这样就能出现 `[1,2,1]` 这种排列了.
	for i := 1; i <= target; i++ {
		for _, n := range nums {
			if i-n < 0 {
				continue
			}
			dp[i] += dp[i-n]
		}
	}
	return dp[target]
}

/*
 * // 二刷还不会, 欠刷.
 * func combinationSum4(N []int, target int) int {
 *     dp := make([]int, target+1)
 *     dp[0] = 1
 *     for j := 0; j <= target; j++ {
 *         for _, n := range N {
 *             if j >= n {
 *                 dp[j] += dp[j-n]
 *             }
 *         }
 *     }
 *     return dp[target]
 * }
 */
