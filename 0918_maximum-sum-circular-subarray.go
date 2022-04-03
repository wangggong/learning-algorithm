/*
 * @lc app=leetcode.cn id=maximum-sum-circular-subarray lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [918] 环形子数组的最大和
 *
 * https://leetcode-cn.com/problems/maximum-sum-circular-subarray/description/
 *
 * algorithms
 * Medium (36.52%)
 * Total Accepted:    32.3K
 * Total Submissions: 88.4K
 * Testcase Example:  '[1,-2,3,-2]'
 *
 * 给定一个长度为 n 的环形整数数组 nums ，返回 nums 的非空 子数组 的最大可能和 。
 *
 * 环形数组 意味着数组的末端将会与开头相连呈环状。形式上， nums[i] 的下一个元素是 nums[(i + 1) % n] ， nums[i]
 * 的前一个元素是 nums[(i - 1 + n) % n] 。
 *
 * 子数组 最多只能包含固定缓冲区 nums 中的每个元素一次。形式上，对于子数组 nums[i], nums[i + 1], ..., nums[j]
 * ，不存在 i <= k1, k2 <= j 其中 k1 % n == k2 % n 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,-2,3,-2]
 * 输出：3
 * 解释：从子数组 [3] 得到最大和 3
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [5,-3,5]
 * 输出：10
 * 解释：从子数组 [5,5] 得到最大和 5 + 5 = 10
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [3,-2,2,-3]
 * 输出：3
 * 解释：从子数组 [3] 和 [3,-2,2] 都可以得到最大和 3
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 1 <= n <= 3 * 10^4
 * -3 * 10^4 <= nums[i] <= 3 * 10^4
 *
 *
 */

const MAXN = 3e4 + 10

/*
 * // TLE 了, 别说了...
 * // 如果想优化需要用到单调队列, 参考官解方法 2:
 * // https://leetcode-cn.com/problems/maximum-sum-circular-subarray/solution/huan-xing-zi-shu-zu-de-zui-da-he-by-leetcode/
 * func maxSubarraySumCircular(nums []int) int {
 * 	n := len(nums)
 * 	var curr []int
 * 	curr = append(curr, nums...)
 * 	curr = append(curr, nums...)
 * 	prefixSum := make([]int, len(curr)+1)
 * 	for i, c := range curr {
 * 		prefixSum[i+1] = prefixSum[i] + c
 * 	}
 * 	//	X	1	2	-3	2
 * 	//	X	1	2	-3	2	1	2	-3	2
 * 	//	0	1	3	0	2	3	5	2	4
 *
 * 	//	X	5	-3	5
 * 	//	X	5	-3	5	5	-3	5
 * 	//	0	5	2	7	12	9	14
 * 	ans := int(-MAXN)
 * 	for i := 0; i < n; i++ {
 * 		for j := 1; j <= n; j++ {
 * 			ans = max(ans, prefixSum[i+j]-prefixSum[i])
 * 		}
 * 	}
 * 	return ans
 * }
 */

// 参考: https://leetcode-cn.com/problems/maximum-sum-circular-subarray/solution/wo-hua-yi-bian-jiu-kan-dong-de-ti-jie-ni-892u/
//
// 分两种情况:
// 1. 返回结果是连续的, 有环没环一样, 直接求最大子数组;
// 2. 返回结果是不连续的 (出现了首尾交叉的部分), 此时由于整个环上的总和为固定值, 求这里的最大值等价于求中间没交叉部分的最小值;
func maxSubarraySumCircular(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prevMax, prevMin, maxVal, minVal := int(-MAXN), int(MAXN), int(-MAXN), int(MAXN)
	tot := 0
	for _, n := range nums {
		tot += n
		currMax := max(prevMax+n, n)
		maxVal = max(maxVal, currMax)
		currMin := min(prevMin+n, n)
		minVal = min(minVal, currMin)
		prevMax, prevMin = currMax, currMin
	}
	// NOTE: 关键点, 如果整个数组都是负数, 那直接返回最大值 (因为你不能一个都不选, 而此时 tot - minVal 代表的就是一个都不选那个 0)
	if maxVal < 0 {
		return maxVal
	}
	return max(maxVal, tot-minVal)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
