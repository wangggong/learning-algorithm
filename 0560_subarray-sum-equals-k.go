/*
 * @lc app=leetcode.cn id=subarray-sum-equals-k lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [560] 和为 K 的子数组
 *
 * https://leetcode-cn.com/problems/subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (44.75%)
 * Total Accepted:    182K
 * Total Submissions: 406.7K
 * Testcase Example:  '[1,1,1]\n2'
 *
 * 给你一个整数数组 nums 和一个整数 k ，请你统计并返回该数组中和为 k 的连续子数组的个数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,1], k = 2
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3], k = 3
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * -1000 <= nums[i] <= 1000
 * -10^7 <= k <= 10^7
 *
 *
 */

/*
 * func subarraySum(nums []int, k int) int {
 * 	n := len(nums)
 * 	prefixSum := make([]int, n+1)
 * 	for i, v := range nums {
 * 		prefixSum[i+1] = prefixSum[i] + v
 * 	}
 * 	ans := 0
 * 	for i := 0; i <= n; i++ {
 * 		for j := 0; j < i; j++ {
 * 			if prefixSum[i]-prefixSum[j] == k {
 * 				ans++
 * 			}
 * 		}
 * 	}
 * 	return ans
 * }
 */

/*
 * func subarraySum(nums []int, k int) int {
 * 	n := len(nums)
 * 	m := make(map[int][]int)
 * 	prefixSum := make([]int, n+1)
 * 	m[0] = []int{0}
 * 	for i, v := range nums {
 * 		prefixSum[i+1] = prefixSum[i] + v
 * 		m[prefixSum[i+1]] = append(m[prefixSum[i+1]], i+1)
 * 	}
 * 	ans := 0
 * 	for i := 1; i <= n; i++ {
 * 		for _, j := range m[prefixSum[i]-k] {
 * 			if j >= i {
 * 				break
 * 			}
 * 			ans++
 * 		}
 * 	}
 * 	return ans
 * }
 */
func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	ans := 0
	preSum := 0
	m[0] = 1
	for _, v := range nums {
		preSum += v
		ans += m[preSum-k]
		m[preSum]++
	}
	return ans
}
