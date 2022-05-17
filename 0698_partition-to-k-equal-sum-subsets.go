/*
 * @lc app=leetcode.cn id=partition-to-k-equal-sum-subsets lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [698] 划分为k个相等的子集
 *
 * https://leetcode.cn/problems/partition-to-k-equal-sum-subsets/description/
 *
 * algorithms
 * Medium (41.84%)
 * Total Accepted:    46.4K
 * Total Submissions: 110.8K
 * Testcase Example:  '[4,3,2,3,5,2,1]\n4'
 *
 * 给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
 * 输出： True
 * 说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
 *
 * 示例 2:
 *
 *
 * 输入: nums = [1,2,3,4], k = 3
 * 输出: false
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= len(nums) <= 16
 * 0 < nums[i] < 10000
 * 每个元素的频率在 [1,4] 范围内
 *
 *
 */

// 发现暴力不过是自己的问题, 那没事了.
func canPartitionKSubsets(nums []int, k int) bool {
	sort.Ints(nums)
	n, sum := len(nums), 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	visit := make([]bool, n)
	return backtrack(nums, visit, sum/k, n, k, 0, 0)
}

func backtrack(arr []int, visit []bool, target, rest, k, curr, index int) bool {
	// 没剩余数字了说明 OK 了. 直接返回 true.
	if rest == 0 {
		return true
	}
	// 如果 `curr == target`, 说明当前的子集凑齐了, 从 `0` 位置开始凑下一个子集即可.
	if curr == target {
		return backtrack(arr, visit, target, rest, k-1, 0, 0)
	}
	// 从 `index` 开始找到所有没用到的数字.
	for i, n := index, len(visit); i < n; i++ {
		if visit[i] {
			continue
		}
		// 剪枝: 如果爆了, 直接返回. 后面的数更大.
		if curr+arr[i] > target {
			break
		}
		visit[i] = true
		if backtrack(arr, visit, target, rest-1, k, curr+arr[i], i+1) {
			return true
		}
		visit[i] = false
		// 剪枝: **如果方案真的存在, 当前第一个没用到的数字已经在一个子集中充当最小值**
		if curr == 0 {
			break
		}
	}
	return false
}

/*
 * func canPartitionKSubsets(nums []int, k int) bool {
 * 	sort.Ints(nums)
 * 	n := len(nums)
 * 	sum := 0
 * 	for _, n := range nums {
 * 		sum += n
 * 	}
 * 	if sum%k != 0 {
 * 		return false
 * 	}
 * 	target := sum / k
 * 	if nums[n-1] > target {
 * 		return false
 * 	}
 * 	limit := 1 << n
 * 	dp := make([]bool, limit)
 * 	curSum := make([]int, limit)
 * 	dp[0] = true
 * 	for s := 0; s < limit; s++ {
 * 		if !dp[s] {
 * 			continue
 * 		}
 * 		for i := 0; i < n; i++ {
 * 			if s&(1<<i) != 0 {
 * 				continue
 * 			}
 * 			ns := s | (1 << i)
 * 			if dp[ns] {
 * 				continue
 * 			}
 * 			if (curSum[s]%target)+nums[i] <= target {
 * 				curSum[ns] = curSum[s] + nums[i]
 * 				dp[ns] = true
 * 			} else {
 * 				break
 * 			}
 * 		}
 * 	}
 * 	return dp[limit-1]
 * }
 */