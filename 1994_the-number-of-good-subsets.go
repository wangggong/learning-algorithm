/*
 * @lc app=leetcode.cn id=the-number-of-good-subsets lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1994] 好子集的数目
 *
 * https://leetcode-cn.com/problems/the-number-of-good-subsets/description/
 *
 * algorithms
 * Hard (37.37%)
 * Total Accepted:    3.8K
 * Total Submissions: 7.1K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给你一个整数数组 nums 。如果 nums 的一个子集中，所有元素的乘积可以表示为一个或多个 互不相同的质数 的乘积，那么我们称它为 好子集
 * 。
 *
 *
 * 比方说，如果 nums = [1, 2, 3, 4] ：
 *
 *
 * [2, 3] ，[1, 2, 3] 和 [1, 3] 是 好 子集，乘积分别为 6 = 2*3 ，6 = 2*3 和 3 = 3 。
 * [1, 4] 和 [4] 不是 好 子集，因为乘积分别为 4 = 2*2 和 4 = 2*2 。
 *
 *
 *
 *
 * 请你返回 nums 中不同的 好 子集的数目对 10^9 + 7 取余 的结果。
 *
 * nums 中的 子集 是通过删除 nums
 * 中一些（可能一个都不删除，也可能全部都删除）元素后剩余元素组成的数组。如果两个子集删除的下标不同，那么它们被视为不同的子集。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3,4]
 * 输出：6
 * 解释：好子集为：
 * - [1,2]：乘积为 2 ，可以表示为质数 2 的乘积。
 * - [1,2,3]：乘积为 6 ，可以表示为互不相同的质数 2 和 3 的乘积。
 * - [1,3]：乘积为 3 ，可以表示为质数 3 的乘积。
 * - [2]：乘积为 2 ，可以表示为质数 2 的乘积。
 * - [2,3]：乘积为 6 ，可以表示为互不相同的质数 2 和 3 的乘积。
 * - [3]：乘积为 3 ，可以表示为质数 3 的乘积。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [4,2,3,15]
 * 输出：5
 * 解释：好子集为：
 * - [2]：乘积为 2 ，可以表示为质数 2 的乘积。
 * - [2,3]：乘积为 6 ，可以表示为互不相同质数 2 和 3 的乘积。
 * - [2,15]：乘积为 30 ，可以表示为互不相同质数 2，3 和 5 的乘积。
 * - [3]：乘积为 3 ，可以表示为质数 3 的乘积。
 * - [15]：乘积为 15 ，可以表示为互不相同质数 3 和 5 的乘积。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 30
 *
 *
 */

const STATES = 1 << 10
const MOD = 1e9 + 7
const NUMS = 30 + 1

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
var dp [STATES]int
var count [NUMS]int

func numberOfGoodSubsets(nums []int) int {
	for i := range count {
		count[i] = 0
	}
	for i := range dp {
		dp[i] = 0
	}
	// 先计波数...
	for _, n := range nums {
		count[n]++
	}
	// 状态压缩 DP.
	// 1. 奠基: 如果只有 1 的子集数, 1 << count[1]
	dp[0] = 1
	for i := 0; i < count[1]; i++ {
		dp[0] = (dp[0] << 1) % MOD
	}
	// 2. 枚举每个值 n (范围 2 .. 30), 更新计数: dp[state] += dp[state \ getState(n)] * count[n]
	for n := 2; n < NUMS; n++ {
		// nums 里面没这玩意就跳过...
		if count[n] == 0 {
			continue
		}
		// ... 然后如果不能分解成互不相同的质因数就跳过...
		cs, ok := getGoodState(n)
		if !ok {
			continue
		}
		// ... 最后枚举状态, 更新计数:
		for s := STATES; s > 0; s-- {
			// 如果当前状态 cs 不能转化为目标状态 s, 跳过...
			if s & cs != cs {
				continue
			}
			dp[s] = (dp[s] + dp[s^cs]*count[n]) % MOD
		}
	}
	// 最终统计一波, 别把 0 状态统计进去...
	ans := 0
	for _, v := range dp[1:] {
		ans = (ans + v) % MOD
	}
	return ans
}

// 分解质因数, 并整理成 state. 如果有重复质因数直接 !ok.
func getGoodState(n int) (int, bool) {
	s := 0
	for i, p := range primes {
		if n%(p*p) == 0 {
			return s, false
		}
		if n%p == 0 {
			s |= 1 << i
		}
	}
	return s, true
}

// 也可以参考: https://leetcode-cn.com/problems/the-number-of-good-subsets/solution/hui-su-zhuang-tai-ya-suo-jian-zhi-wei-yu-1ruz/
