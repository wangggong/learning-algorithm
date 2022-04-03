/*
 * @lc app=leetcode.cn id=minimum-increment-to-make-array-unique lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [945] 使数组唯一的最小增量
 *
 * https://leetcode-cn.com/problems/minimum-increment-to-make-array-unique/description/
 *
 * algorithms
 * Medium (48.57%)
 * Total Accepted:    37.4K
 * Total Submissions: 77.2K
 * Testcase Example:  '[1,2,2]'
 *
 * 给你一个整数数组 nums 。每次 move 操作将会选择任意一个满足 0 <= i < nums.length 的下标 i，并将 nums[i] 递增
 * 1。
 *
 * 返回使 nums 中的每个值都变成唯一的所需要的最少操作次数。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,2]
 * 输出：1
 * 解释：经过一次 move 操作，数组将变为 [1, 2, 3]。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,2,1,2,1,7]
 * 输出：6
 * 解释：经过 6 次 move 操作，数组将变为 [3, 4, 1, 2, 5, 7]。
 * 可以看出 5 次或 5 次以下的 move 操作是不能让数组的每个值唯一的。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * 0 <= nums[i] <= 10^5
 *
 *
 */

const MAXN int = 1e5

var visit [MAXN*2 + 5]int

/*
 * func minIncrementForUnique(nums []int) int {
 * 	ans := 0
 * 	minVal := len(visit) - 1
 * 	for i := range visit {
 * 		visit[i] = 0
 * 	}
 * 	for _, n := range nums {
 * 		visit[n]++
 * 		if n < minVal {
 * 			minVal = n
 * 		}
 * 	}
 * 	for i := range visit {
 * 		if visit[i] > 1 && minVal < i {
 * 			minVal = i
 * 		}
 * 		for visit[i] > 1 {
 * 			for minVal < len(visit) && visit[minVal] > 0 {
 * 				minVal++
 * 			}
 * 			if minVal == len(visit) {
 * 				return ans
 * 			}
 * 			ans += minVal - i
 * 			visit[minVal]++
 * 			visit[i]--
 * 		}
 * 	}
 * 	return ans
 * }
 */

func minIncrementForUnique(nums []int) int {
	for i := range visit {
		visit[i] = 0
	}
	for _, n := range nums {
		visit[n]++
	}
	ans := 0
	taken := 0
	for x, v := range visit {
		if v >= 2 {
			// 如果出现次数超过 1, 先欠着 `v-1` 个 `x`;
			taken += v - 1
			ans -= x * (v - 1)
		} else if taken > 0 && v == 0 {
			// 如果还欠着数, 先补一个;
			taken--
			ans += x
		}
	}
	return ans
}
