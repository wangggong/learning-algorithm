/*
 * @lc app=leetcode.cn id=target-sum lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [494] 目标和
 *
 * https://leetcode-cn.com/problems/target-sum/description/
 *
 * algorithms
 * Medium (49.10%)
 * Total Accepted:    198.8K
 * Total Submissions: 405K
 * Testcase Example:  '[1,1,1,1,1]\n3'
 *
 * 给你一个整数数组 nums 和一个整数 target 。
 *
 * 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
 *
 *
 * 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
 *
 *
 * 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,1,1,1], target = 3
 * 输出：5
 * 解释：一共有 5 种方法让最终目标和为 3 。
 * -1 + 1 + 1 + 1 + 1 = 3
 * +1 - 1 + 1 + 1 + 1 = 3
 * +1 + 1 - 1 + 1 + 1 = 3
 * +1 + 1 + 1 - 1 + 1 = 3
 * +1 + 1 + 1 + 1 - 1 = 3
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1], target = 1
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 20
 * 0 <= nums[i] <= 1000
 * 0 <= sum(nums[i]) <= 1000
 * -1000 <= target <= 1000
 *
 *
 */

const maxn int = 1000

var dp [2*maxn + 10]int

func findTargetSumWays(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		if target == 0 {
			return 1
		}
		return 0
	}
	for i := 0; i < 2*maxn+10; i++ {
		dp[i] = 0
	}
	dp[arr[0]+maxn] += 1
	dp[-arr[0]+maxn] += 1
	for i := 1; i < n; i++ {
		next := make([]int, 2*maxn+10)
		for j := 0; j < 2*maxn+10; j++ {
			if j+arr[i] < 2*maxn+10 {
				next[j+arr[i]] += dp[j]
			}
			if j-arr[i] >= 0 {
				next[j-arr[i]] += dp[j]
			}
		}
		// 傻逼, 别滚反了!
		for j := 0; j < 2*maxn+10; j++ {
			dp[j] = next[j]
		}
	}
	return dp[target+maxn]
}
