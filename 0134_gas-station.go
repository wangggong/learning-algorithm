/*
 * @lc app=leetcode.cn id=gas-station lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [134] 加油站
 *
 * https://leetcode-cn.com/problems/gas-station/description/
 *
 * algorithms
 * Medium (55.09%)
 * Total Accepted:    157.8K
 * Total Submissions: 286.5K
 * Testcase Example:  '[1,2,3,4,5]\n[3,4,5,1,2]'
 *
 * 在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
 *
 * 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
 *
 * 给定两个整数数组 gas 和 cost ，如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一
 * 的。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
 * 输出: 3
 * 解释:
 * 从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
 * 开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
 * 开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
 * 开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
 * 开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
 * 开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
 * 因此，3 可为起始索引。
 *
 * 示例 2:
 *
 *
 * 输入: gas = [2,3,4], cost = [3,4,3]
 * 输出: -1
 * 解释:
 * 你不能从 0 号或 1 号加油站出发，因为没有足够的汽油可以让你行驶到下一个加油站。
 * 我们从 2 号加油站出发，可以获得 4 升汽油。 此时油箱有 = 0 + 4 = 4 升汽油
 * 开往 0 号加油站，此时油箱有 4 - 3 + 2 = 3 升汽油
 * 开往 1 号加油站，此时油箱有 3 - 3 + 3 = 3 升汽油
 * 你无法返回 2 号加油站，因为返程需要消耗 4 升汽油，但是你的油箱只有 3 升汽油。
 * 因此，无论怎样，你都不可能绕环路行驶一周。
 *
 *
 *
 * 提示:
 *
 *
 * gas.length == n
 * cost.length == n
 * 1 <= n <= 10^5
 * 0 <= gas[i], cost[i] <= 10^4
 *
 *
 */

/*
 * func canCompleteCircuit(G []int, C []int) int {
 * 	// assert len(G) == len(C);
 * 	n := len(G)
 * 	cur := 0
 * 	tot := 0
 * 	ind := 0
 * 	// 贪心点:
 * 	// 1. 如果 `\Sigma{G} < \Sigma{C}`, 无法绕一圈, 直接返回 -1.
 * 	// 2. 否则, 如果 x -> y 不能绕一圈, 那 x 到 y 的 ind 均不合法, 从下个 index 开始检查.
 * 	for i := 0; i < n; i++ {
 * 		cur += G[i]-C[i]
 * 		tot += G[i]-C[i]
 * 		if cur < 0 {
 * 			cur = 0
 * 			ind = i+1
 * 		}
 * 	}
 * 	if tot < 0 {
 * 		return -1
 * 	}
 * 	return ind
 * }
 */

// 瞎逼贪, 但也过了
func canCompleteCircuit(G []int, C []int) int {
	n := len(G)
	// assert len(G) == len(C);
	cur := 0
	start := 0
	for start < n {
		k := start
		for ; k < start+n; k++ {
			ind := k % n
			cur += G[ind] - C[ind]
			if cur < 0 {
				break
			}
		}
		if k == start+n {
			return start
		}
		cur = 0
		start = k + 1
	}
	return -1
}
