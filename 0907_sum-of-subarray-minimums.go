/*
 * @lc app=leetcode.cn id=sum-of-subarray-minimums lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [907] 子数组的最小值之和
 *
 * https://leetcode-cn.com/problems/sum-of-subarray-minimums/description/
 *
 * algorithms
 * Medium (32.99%)
 * Total Accepted:    15K
 * Total Submissions: 45.5K
 * Testcase Example:  '[3,1,2,4]'
 *
 * 给定一个整数数组 arr，找到 min(b) 的总和，其中 b 的范围为 arr 的每个（连续）子数组。
 *
 * 由于答案可能很大，因此 返回答案模 10^9 + 7 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：arr = [3,1,2,4]
 * 输出：17
 * 解释：
 * 子数组为 [3]，[1]，[2]，[4]，[3,1]，[1,2]，[2,4]，[3,1,2]，[1,2,4]，[3,1,2,4]。
 * 最小值为 3，1，2，4，1，1，2，1，1，1，和为 17。
 *
 * 示例 2：
 *
 *
 * 输入：arr = [11,81,94,43,3]
 * 输出：444
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= arr.length <= 3 * 1e4
 * 1 <= arr[i] <= 3 * 1e4
 *
 *
 *
 *
 */
import "math"

const mod int = 1e9 + 7

// 参考: https://leetcode-cn.com/problems/sum-of-subarray-minimums/solution/dan-diao-zhan-zuo-you-liang-bian-di-yi-g-ww3n/
//
// 就是接雨水套个壳子嘛, 傻逼了吧...
func sumSubarrayMins(arr []int) int {
	n := len(arr)
	ans := 0
	left, right := make([]int, n), make([]int, n)
	var S []int
	S = append(S, -1)
	for i := 0; i < n; i++ {
		for len(S) > 0 && get(arr, S[len(S)-1]) >= get(arr, i) {
			S = S[:len(S)-1]
		}
		left[i] = S[len(S)-1]
		S = append(S, i)
	}
	S = nil
	S = append(S, n)
	for i := n - 1; i >= 0; i-- {
		for len(S) > 0 && get(arr, S[len(S)-1]) > get(arr, i) {
			S = S[:len(S)-1]
		}
		right[i] = S[len(S)-1]
		S = append(S, i)
	}
	for i := 0; i < n; i++ {
		l, r := i-left[i], right[i]-i
		ans = (ans + l*r*arr[i]%mod) % mod
	}
	return ans
}

func get(arr []int, k int) int {
	if k < 0 || k >= len(arr) {
		return math.MinInt32
	}
	return arr[k]
}
