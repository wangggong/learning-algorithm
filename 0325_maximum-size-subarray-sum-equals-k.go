/*
 * Medium
 *
 * 给定一个数组 nums 和一个目标值 k，找到和等于 k 的最长连续子数组长度。如果不存在任意一个符合要求的子数组，则返回 0。
 *
 *
 *
 * 示例 1:
 *
 * 输入: nums = [1,-1,5,-2,3], k = 3
 * 输出: 4
 * 解释: 子数组 [1, -1, 5, -2] 和等于 3，且长度最长。
 * 示例 2:
 *
 * 输入: nums = [-2,-1,2,1], k = 1
 * 输出: 2
 * 解释: 子数组 [-1, 2] 和等于 1，且长度最长。
 *
 *
 * 提示：
 *
 * 1 <= nums.length <= 2 * 105
 * -104 <= nums[i] <= 104
 * -109 <= k <= 109
 * 通过次数12,182
 * 提交次数22,807
 */

func maxSubArrayLen(nums []int, k int) int {
	right := make(map[int]int)
	right[0] = -1
	ans := 0
	cur := 0
	for i, n := range nums {
		cur += n
		if r, ok := right[cur-k]; ok {
			ans = max(ans, i-r)
		}
		if _, ok := right[cur]; !ok {
			right[cur] = i
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
