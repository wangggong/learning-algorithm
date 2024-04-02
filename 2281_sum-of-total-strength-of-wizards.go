/*
 * @lc app=leetcode.cn id=sum-of-total-strength-of-wizards lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2281] 巫师的总力量和
 *
 * https://leetcode-cn.com/problems/sum-of-total-strength-of-wizards/description/
 *
 * algorithms
 * Hard (18.02%)
 * Total Accepted:    1.4K
 * Total Submissions: 7.9K
 * Testcase Example:  '[1,3,1,2]'
 *
 * 作为国王的统治者，你有一支巫师军队听你指挥。
 *
 * 给你一个下标从 0 开始的整数数组 strength ，其中 strength[i] 表示第 i
 * 位巫师的力量值。对于连续的一组巫师（也就是这些巫师的力量值是 strength 的 子数组），总力量 定义为以下两个值的 乘积 ：
 *
 *
 * 巫师中 最弱 的能力值。
 * 组中所有巫师的个人力量值 之和 。
 *
 *
 * 请你返回 所有 巫师组的 总 力量之和。由于答案可能很大，请将答案对 10^9 + 7 取余 后返回。
 *
 * 子数组 是一个数组里 非空 连续子序列。
 *
 *
 *
 * 示例 1：
 *
 * 输入：strength = [1,3,1,2]
 * 输出：44
 * 解释：以下是所有连续巫师组：
 * - [1,3,1,2] 中 [1] ，总力量值为 min([1]) * sum([1]) = 1 * 1 = 1
 * - [1,3,1,2] 中 [3] ，总力量值为 min([3]) * sum([3]) = 3 * 3 = 9
 * - [1,3,1,2] 中 [1] ，总力量值为 min([1]) * sum([1]) = 1 * 1 = 1
 * - [1,3,1,2] 中 [2] ，总力量值为 min([2]) * sum([2]) = 2 * 2 = 4
 * - [1,3,1,2] 中 [1,3] ，总力量值为 min([1,3]) * sum([1,3]) = 1 * 4 = 4
 * - [1,3,1,2] 中 [3,1] ，总力量值为 min([3,1]) * sum([3,1]) = 1 * 4 = 4
 * - [1,3,1,2] 中 [1,2] ，总力量值为 min([1,2]) * sum([1,2]) = 1 * 3 = 3
 * - [1,3,1,2] 中 [1,3,1] ，总力量值为 min([1,3,1]) * sum([1,3,1]) = 1 * 5 = 5
 * - [1,3,1,2] 中 [3,1,2] ，总力量值为 min([3,1,2]) * sum([3,1,2]) = 1 * 6 = 6
 * - [1,3,1,2] 中 [1,3,1,2] ，总力量值为 min([1,3,1,2]) * sum([1,3,1,2]) = 1 * 7 = 7
 * 所有力量值之和为 1 + 9 + 1 + 4 + 4 + 4 + 3 + 5 + 6 + 7 = 44 。
 *
 *
 * 示例 2：
 *
 * 输入：strength = [5,4,6]
 * 输出：213
 * 解释：以下是所有连续巫师组：
 * - [5,4,6] 中 [5] ，总力量值为 min([5]) * sum([5]) = 5 * 5 = 25
 * - [5,4,6] 中 [4] ，总力量值为 min([4]) * sum([4]) = 4 * 4 = 16
 * - [5,4,6] 中 [6] ，总力量值为 min([6]) * sum([6]) = 6 * 6 = 36
 * - [5,4,6] 中 [5,4] ，总力量值为 min([5,4]) * sum([5,4]) = 4 * 9 = 36
 * - [5,4,6] 中 [4,6] ，总力量值为 min([4,6]) * sum([4,6]) = 4 * 10 = 40
 * - [5,4,6] 中 [5,4,6] ，总力量值为 min([5,4,6]) * sum([5,4,6]) = 4 * 15 = 60
 * 所有力量值之和为 25 + 16 + 36 + 36 + 40 + 60 = 213 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= strength.length <= 10^5
 * 1 <= strength[i] <= 10^9
 *
 *
 */
import "math"

const mod int = 1e9 + 7

func totalStrength(arr []int) int {
	n := len(arr)

	ps := make([]int, n+1)
	for i := 0; i < n; i++ {
		ps[i+1] = (ps[i] + arr[i]) % mod
	}
	pss := make([]int, n+2)
	for i := 0; i <= n; i++ {
		pss[i+1] = (pss[i] + ps[i]) % mod
	}
	pss = pss

	left, right := make([]int, n), make([]int, n)
	{
		var S []int
		S = append(S, -1)
		for i := 0; i < n; i++ {
			for get(arr, S[len(S)-1]) > get(arr, i) {
				S = S[:len(S)-1]
			}
			left[i] = S[len(S)-1]
			S = append(S, i)
		}
	}
	{
		var S []int
		S = append(S, n)
		for i := n - 1; i >= 0; i-- {
			for get(arr, S[len(S)-1]) >= get(arr, i) {
				S = S[:len(S)-1]
			}
			right[i] = S[len(S)-1]
			S = append(S, i)
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		l, r := left[i], right[i]
		sum := (((pss[r+1]-pss[i+1])*(i-l))%mod - ((pss[i+1]-pss[l+1])*(r-i))%mod + mod) % mod
		cur := (arr[i] * sum) % mod
		ans = (ans + cur) % mod
	}

	return ans
}

func get(arr []int, k int) int {
	if k < 0 || k >= len(arr) {
		return math.MinInt32
	}
	return arr[k]
}

// \Sigma_{p=l+1}^{p<=k}\Sigma_{q=k}^{q<=r}{S[q]-S[p-1]} = \Sigma_{p=l+1}^{p<=k}{(\Sigma_{q=k}^{q<=r}{S[q]-S[p-1]})}
// = \Sigma_{q=k}^{q<=r}{S[q]} * (k-l) - \Sigma_{p=l+1}^{p<=k}{S[p-1]} * (r-k+1)
// = (SS[r] - SS[k-1]) * (k-l) - (SS[k-1] - SS[l]) * (r-k+1)
//      1   3  *1*  2
// 0    1   4   5   7
// 0    1   5   10  17
// l = 0 k = 2 r = 4
// (17 - 5) * (2) - (5 - 0) * (2) = 14
