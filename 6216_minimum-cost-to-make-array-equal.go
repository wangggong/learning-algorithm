/*
 * @lc app=leetcode.cn id=minimum-cost-to-make-array-equal lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6216] 使数组相等的最小开销
 *
 * https://leetcode.cn/problems/minimum-cost-to-make-array-equal/description/
 *
 * algorithms
 * Hard (27.64%)
 * Total Accepted:    2.4K
 * Total Submissions: 8.5K
 * Testcase Example:  '[1,3,5,2]\n[2,3,1,14]'
 *
 * 给你两个下标从 0 开始的数组 nums 和 cost ，分别包含 n 个 正 整数。
 *
 * 你可以执行下面操作 任意 次：
 *
 *
 * 将 nums 中 任意 元素增加或者减小 1 。
 *
 *
 * 对第 i 个元素执行一次操作的开销是 cost[i] 。
 *
 * 请你返回使 nums 中所有元素 相等 的 最少 总开销。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [1,3,5,2], cost = [2,3,1,14]
 * 输出：8
 * 解释：我们可以执行以下操作使所有元素变为 2 ：
 * - 增加第 0 个元素 1 次，开销为 2 。
 * - 减小第 1 个元素 1 次，开销为 3 。
 * - 减小第 2 个元素 3 次，开销为 1 + 1 + 1 = 3 。
 * 总开销为 2 + 3 + 3 = 8 。
 * 这是最小开销。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [2,2,2,2,2], cost = [4,2,8,1,3]
 * 输出：0
 * 解释：数组中所有元素已经全部相等，不需要执行额外的操作。
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length == cost.length
 * 1 <= n <= 10^5
 * 1 <= nums[i], cost[i] <= 10^6
 *
 *
 */

import "sort"

func minCost(nums []int, cost []int) (ans int64) {
	n := len(nums)
	infos := make([][2]int64, 0, n)
	tot := int64(0)
	for i := 0; i < n; i++ {
		infos = append(infos, [2]int64{int64(nums[i]), int64(cost[i])})
		tot += int64(cost[i])
	}
	sort.Slice(infos, func(p, q int) bool { return infos[p][0] < infos[q][0] })
	cur := int64(0)
	for _, info := range infos[1:] {
		cur += (info[0] - infos[0][0]) * info[1]
	}
	ans = cur
	curTot := infos[0][1]
	for i := 0; i+1 < n; i++ {
		cur += (infos[i+1][0] - infos[i][0]) * (curTot*2 - tot)
		if cur < ans {
			ans = cur
		}
		curTot += infos[i+1][1]
	}
	return
}
