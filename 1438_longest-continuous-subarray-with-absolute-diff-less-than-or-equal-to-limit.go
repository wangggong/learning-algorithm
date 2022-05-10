/*
 * @lc app=leetcode.cn id=longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1438] 绝对差不超过限制的最长连续子数组
 *
 * https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/description/
 *
 * algorithms
 * Medium (48.64%)
 * Total Accepted:    37.2K
 * Total Submissions: 76.3K
 * Testcase Example:  '[8,2,4,7]\n4'
 *
 * 给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于
 * limit 。
 *
 * 如果不存在满足条件的子数组，则返回 0 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [8,2,4,7], limit = 4
 * 输出：2
 * 解释：所有子数组如下：
 * [8] 最大绝对差 |8-8| = 0 <= 4.
 * [8,2] 最大绝对差 |8-2| = 6 > 4.
 * [8,2,4] 最大绝对差 |8-2| = 6 > 4.
 * [8,2,4,7] 最大绝对差 |8-2| = 6 > 4.
 * [2] 最大绝对差 |2-2| = 0 <= 4.
 * [2,4] 最大绝对差 |2-4| = 2 <= 4.
 * [2,4,7] 最大绝对差 |2-7| = 5 > 4.
 * [4] 最大绝对差 |4-4| = 0 <= 4.
 * [4,7] 最大绝对差 |4-7| = 3 <= 4.
 * [7] 最大绝对差 |7-7| = 0 <= 4.
 * 因此，满足题意的最长子数组的长度为 2 。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [10,1,2,4,7,2], limit = 5
 * 输出：4
 * 解释：满足题意的最长子数组是 [2,4,7,2]，其最大绝对差 |2-7| = 5 <= 5 。
 *
 *
 * 示例 3：
 *
 * 输入：nums = [4,2,2,2,4,4,2,2], limit = 0
 * 输出：3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^9
 * 0 <= limit <= 10^9
 *
 *
 */

// 参考: https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/solution/jue-dui-chai-bu-chao-guo-xian-zhi-de-zui-5bki/
//
// 不用二刷都没啥思路的题. 滑动窗口+单调队列.
func longestSubarray(nums []int, limit int) int {
	n := len(nums)
	ans := 0
	var maxQ, minQ []int
	for p, q := 0, 0; q < n; q++ {
		// 更新区间最大 / 最小值
		for len(maxQ) > 0 && nums[maxQ[len(maxQ)-1]] < nums[q] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, q)
		for len(minQ) > 0 && nums[minQ[len(minQ)-1]] > nums[q] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, q)
		// 使区间满足条件
		for len(maxQ) > 0 && len(minQ) > 0 && nums[maxQ[0]]-nums[minQ[0]] > limit {
			if p == maxQ[0] {
				maxQ = maxQ[1:]
			}
			if p == minQ[0] {
				minQ = minQ[1:]
			}
			p++
		}
		ans = max(ans, q-p+1)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
