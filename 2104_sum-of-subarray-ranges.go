/*
 * @lc app=leetcode.cn id=sum-of-subarray-ranges lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2104] 子数组范围和
 *
 * https://leetcode-cn.com/problems/sum-of-subarray-ranges/description/
 *
 * algorithms
 * Medium (56.24%)
 * Total Accepted:    5.8K
 * Total Submissions: 10.3K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个整数数组 nums 。nums 中，子数组的 范围 是子数组中最大元素和最小元素的差值。
 *
 * 返回 nums 中 所有 子数组范围的 和 。
 *
 * 子数组是数组中一个连续 非空 的元素序列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：4
 * 解释：nums 的 6 个子数组如下所示：
 * [1]，范围 = 最大 - 最小 = 1 - 1 = 0
 * [2]，范围 = 2 - 2 = 0
 * [3]，范围 = 3 - 3 = 0
 * [1,2]，范围 = 2 - 1 = 1
 * [2,3]，范围 = 3 - 2 = 1
 * [1,2,3]，范围 = 3 - 1 = 2
 * 所有范围的和是 0 + 0 + 0 + 1 + 1 + 2 = 4
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,3,3]
 * 输出：4
 * 解释：nums 的 6 个子数组如下所示：
 * [1]，范围 = 最大 - 最小 = 1 - 1 = 0
 * [3]，范围 = 3 - 3 = 0
 * [3]，范围 = 3 - 3 = 0
 * [1,3]，范围 = 3 - 1 = 2
 * [3,3]，范围 = 3 - 3 = 0
 * [1,3,3]，范围 = 3 - 1 = 2
 * 所有范围的和是 0 + 0 + 0 + 2 + 0 + 2 = 4
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [4,-2,-3,4,1]
 * 输出：59
 * 解释：nums 中所有子数组范围的和是 59
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 1000
 * -10^9 <= nums[i] <= 10^9
 *
 *
 *
 *
 * 进阶：你可以设计一种时间复杂度为 O(n) 的解决方案吗？
 *
 */

/*
 * const MAXN int64 = 1000+5
 *
 * var maxVal [MAXN][MAXN]int64
 * var minVal [MAXN][MAXN]int64
 *
 * func subArrayRanges(nums []int) int64 {
 * 	n := len(nums)
 * 	for i := int64(0); i < MAXN; i++ {
 * 		for j := int64(0); j < MAXN; j++ {
 * 			maxVal[i][j] = -MAXN
 * 			minVal[i][j] = MAXN
 * 		}
 * 	}
 * 	for i, v := range nums {
 * 		maxVal[i][i] = int64(v)
 * 		minVal[i][i] = int64(v)
 * 	}
 * 	for i := 0; i < n; i++ {
 * 		for j := i+1; j < n; j++ {
 * 			v := int64(nums[j])
 * 			maxVal[i][j] = max(maxVal[i][j-1], v)
 * 			minVal[i][j] = min(minVal[i][j-1], v)
 * 		}
 * 	}
 * 	var ans int64
 * 	for i := 0; i < n; i++ {
 * 		for j := i; j < n; j++ {
 * 			ans += maxVal[i][j] - minVal[i][j]
 * 		}
 * 	}
 * 	return ans
 * }
 */

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

/*
 * func subArrayRanges(nums []int) int64 {
 * 	n := len(nums)
 * 	var ans int64
 * 	for i := 0; i < n; i++ {
 * 		maxVal := int64(nums[i])
 * 		minVal := int64(nums[i])
 * 		for j := i; j < n; j++ {
 * 			maxVal = max(maxVal, int64(nums[j]))
 * 			minVal = min(minVal, int64(nums[j]))
 * 			ans += maxVal - minVal
 * 		}
 * 	}
 * 	return ans
 * 	return ans
 * }
 */

// 官解: 单调栈 -- 最近最小 & 最近最大
func subArrayRanges(nums []int) int64 {
	n := len(nums)
	var S []int
	leftMin := make([]int, n)  // left min
	rightMin := make([]int, n) // right min
	leftMax := make([]int, n)  // left max
	rightMax := make([]int, n) // right max
	// left min
	S = nil
	for i, v := range nums {
		// NOTE: 注意这里不取等, 如果前一个值与该值相同的话在前一个值中会被算到.
		for len(S) > 0 && nums[S[len(S)-1]] > v {
			S = S[:len(S)-1]
		}
		leftMin[i] = -1
		if len(S) > 0 {
			leftMin[i] = S[len(S)-1]
		}
		S = append(S, i)
	}
	// left max
	S = nil
	for i, v := range nums {
		for len(S) > 0 && nums[S[len(S)-1]] <= v {
			S = S[:len(S)-1]
		}
		leftMax[i] = -1
		if len(S) > 0 {
			leftMax[i] = S[len(S)-1]
		}
		S = append(S, i)
	}
	// right min
	S = nil
	for i := n - 1; i >= 0; i-- {
		v := nums[i]
		for len(S) > 0 && nums[S[len(S)-1]] >= v {
			S = S[:len(S)-1]
		}
		rightMin[i] = n
		if len(S) > 0 {
			rightMin[i] = S[len(S)-1]
		}
		S = append(S, i)
	}
	// right max
	S = nil
	for i := n - 1; i >= 0; i-- {
		v := nums[i]
		for len(S) > 0 && nums[S[len(S)-1]] < v {
			S = S[:len(S)-1]
		}
		rightMax[i] = n
		if len(S) > 0 {
			rightMax[i] = S[len(S)-1]
		}
		S = append(S, i)
	}
	// fmt.Printf("%v\n%v\n%v\n%v\n", leftMax, rightMax, leftMin, rightMin)
	var sumMax, sumMin int64
	for i := 0; i < n; i++ {
		sumMax += int64(nums[i]) * int64(i-leftMax[i]) * int64(rightMax[i]-i)
		sumMin += int64(nums[i]) * int64(i-leftMin[i]) * int64(rightMin[i]-i)
	}
	return sumMax - sumMin
}
