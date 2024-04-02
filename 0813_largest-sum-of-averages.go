/*
 * @lc app=leetcode.cn id=largest-sum-of-averages lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [813] 最大平均值和的分组
 *
 * https://leetcode.cn/problems/largest-sum-of-averages/description/
 *
 * algorithms
 * Medium (56.23%)
 * Total Accepted:    9.2K
 * Total Submissions: 16.4K
 * Testcase Example:  '[9,1,2,3,9]\n3'
 *
 * 给定数组 nums 和一个整数 k 。我们将给定的数组 nums 分成 最多 k 个相邻的非空子数组 。 分数 由每个子数组内的平均值的总和构成。
 *
 * 注意我们必须使用 nums 数组中的每一个数进行分组，并且分数不一定需要是整数。
 *
 * 返回我们所能得到的最大 分数 是多少。答案误差在 10^-6 内被视为是正确的。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [9,1,2,3,9], k = 3
 * 输出: 20.00000
 * 解释:
 * nums 的最优分组是[9], [1, 2, 3], [9]. 得到的分数是 9 + (1 + 2 + 3) / 3 + 9 = 20.
 * 我们也可以把 nums 分成[9, 1], [2], [3, 9].
 * 这样的分组得到的分数为 5 + 2 + 6 = 13, 但不是最大值.
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [1,2,3,4,5,6,7], k = 4
 * 输出: 20.50000
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 100
 * 1 <= nums[i] <= 10^4
 * 1 <= k <= nums.length
 *
 *
 */
import "math"

func largestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	ps := make([]int, n+1)
	for i := 0; i < n; i++ {
		ps[i+1] = ps[i] + nums[i]
	}
	dp := make([][]float64, 2)
	for i := range dp {
		dp[i] = make([]float64, n)
	}
	for i := 0; i < n; i++ {
		dp[0][i] = float64(ps[i+1]-ps[0]) / float64(i+1)
	}
	for i := 1; i < k; i++ {
		for j := i; j < n; j++ {
			for k := 0; k < j; k++ {
				dp[i%2][j] = math.Max(dp[i%2][j], dp[(i-1)%2][k]+float64(ps[j+1]-ps[k+1])/float64(j-k))
			}
		}
	}
	return dp[(k-1)%2][n-1]
}
