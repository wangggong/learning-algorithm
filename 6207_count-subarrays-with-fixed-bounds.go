/*
 * @lc app=leetcode.cn id=count-subarrays-with-fixed-bounds lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6207] 统计定界子数组的数目
 *
 * https://leetcode.cn/problems/count-subarrays-with-fixed-bounds/description/
 *
 * algorithms
 * Hard (31.75%)
 * Total Accepted:    2.9K
 * Total Submissions: 8.8K
 * Testcase Example:  '[1,3,5,2,7,5]\n1\n5'
 *
 * 给你一个整数数组 nums 和两个整数 minK 以及 maxK 。
 *
 * nums 的定界子数组是满足下述条件的一个子数组：
 *
 *
 * 子数组中的 最小值 等于 minK 。
 * 子数组中的 最大值 等于 maxK 。
 *
 *
 * 返回定界子数组的数目。
 *
 * 子数组是数组中的一个连续部分。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [1,3,5,2,7,5], minK = 1, maxK = 5
 * 输出：2
 * 解释：定界子数组是 [1,3,5] 和 [1,3,5,2] 。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [1,1,1,1], minK = 1, maxK = 1
 * 输出：10
 * 解释：nums 的每个子数组都是一个定界子数组。共有 10 个子数组。
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= nums.length <= 10^5
 * 1 <= nums[i], minK, maxK <= 10^6
 *
 *
 */

// https://leetcode.cn/problems/count-subarrays-with-fixed-bounds/solution/jian-ji-xie-fa-pythonjavacgo-by-endlessc-gag2/
// 枚举右端点, 新的玩法出现了!
func countSubarrays(nums []int, minK int, maxK int) (ans int64) {
	minIdx, maxIdx, outIdx := -1, -1, -1
	for i, n := range nums {
		if n == minK {
			minIdx = i
		}
		if n == maxK {
			maxIdx = i
		}
		if n < minK || n > maxK {
			outIdx = i
		}
		ans += int64(max(min(minIdx, maxIdx)-outIdx, 0))
	}
	return
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
