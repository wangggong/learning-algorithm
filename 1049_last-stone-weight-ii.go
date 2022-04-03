/*
 * @lc app=leetcode.cn id=last-stone-weight-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1049] 最后一块石头的重量 II
 *
 * https://leetcode-cn.com/problems/last-stone-weight-ii/description/
 *
 * algorithms
 * Medium (66.22%)
 * Total Accepted:    58.3K
 * Total Submissions: 88K
 * Testcase Example:  '[2,7,4,1,8,1]'
 *
 * 有一堆石头，用整数数组 stones 表示。其中 stones[i] 表示第 i 块石头的重量。
 *
 * 每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x 。那么粉碎的可能结果如下：
 *
 *
 * 如果 x == y，那么两块石头都会被完全粉碎；
 * 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
 *
 *
 * 最后，最多只会剩下一块 石头。返回此石头 最小的可能重量 。如果没有石头剩下，就返回 0。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：stones = [2,7,4,1,8,1]
 * 输出：1
 * 解释：
 * 组合 2 和 4，得到 2，所以数组转化为 [2,7,1,8,1]，
 * 组合 7 和 8，得到 1，所以数组转化为 [2,1,1,1]，
 * 组合 2 和 1，得到 1，所以数组转化为 [1,1,1]，
 * 组合 1 和 1，得到 0，所以数组转化为 [1]，这就是最优值。
 *
 *
 * 示例 2：
 *
 *
 * 输入：stones = [31,26,33,21,40]
 * 输出：5
 *
 *
 * 示例 3：
 *
 *
 * 输入：stones = [1,2]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= stones.length <= 30
 * 1 <= stones[i] <= 100
 *
 *
 */

// 参考 https://leetcode-cn.com/problems/last-stone-weight-ii/solution/gong-shui-san-xie-xiang-jie-wei-he-neng-jgxik/
//
// 为啥能转成背包呢? 因为 **"操作石子重量的差值" 不会实际产生 "新的石子重量"**.
func lastStoneWeightII(S []int) int {
	// 直接看答案了:
	// `W = \Sigma{y_{i} - x_{i}} = \Sigma{y_{i}} - \Sigma{x_{i}} = (\Sigma{arr_{i}} - \Sigma{x_{i}}) - \Sigma{x_{i}}`
	// 因此求最小的 `W` 等价于求最大的 `\Sigma{x_{i}}`, 变背包了.
	n := len(S)
	if n == 0 {
		return 0
	}
	sum := 0
	for _, s := range S {
		sum += s
	}
	lim := sum >> 1
	dp := make([]bool, lim+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		for j := lim; j >= 0; j-- {
			if j-S[i] >= 0 {
				dp[j] = dp[j] || dp[j-S[i]]
			}
		}
	}
	for i := lim; i >= 0; i-- {
		if dp[i] {
			return sum - (i << 1)
		}
	}
	// unreachable
	return -1
}

/*
 * // 二刷还不会+1
 * func lastStoneWeightII(S []int) int {
 * 	sum := 0
 * 	for _, s := range S {
 * 		sum += s
 * 	}
 * 	target := sum >> 1
 * 	dp := make([]bool, target+1)
 * 	dp[0] = true
 * 	ans := 0
 * 	for _, s := range S {
 * 		for j := target; j >= 0; j-- {
 * 			if j >= s {
 * 				dp[j] = dp[j] || dp[j-s]
 * 				if dp[j] && j > ans {
 * 					ans = j
 * 				}
 * 			}
 * 		}
 * 	}
 * 	return sum - (ans << 1)
 * }
 */
