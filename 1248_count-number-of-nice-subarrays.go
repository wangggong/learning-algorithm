/*
 * @lc app=leetcode.cn id=count-number-of-nice-subarrays lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1248] 统计「优美子数组」
 *
 * https://leetcode-cn.com/problems/count-number-of-nice-subarrays/description/
 *
 * algorithms
 * Medium (56.54%)
 * Total Accepted:    38.4K
 * Total Submissions: 67.8K
 * Testcase Example:  '[1,1,2,1,1]\n3'
 *
 * 给你一个整数数组 nums 和一个整数 k。如果某个连续子数组中恰好有 k 个奇数数字，我们就认为这个子数组是「优美子数组」。
 *
 * 请返回这个数组中 「优美子数组」 的数目。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,2,1,1], k = 3
 * 输出：2
 * 解释：包含 3 个奇数的子数组是 [1,1,2,1] 和 [1,2,1,1] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [2,4,6], k = 1
 * 输出：0
 * 解释：数列中不包含任何奇数，所以不存在优美子数组。
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [2,2,2,1,2,2,1,2,2,2], k = 2
 * 输出：16
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 50000
 * 1 <= nums[i] <= 10^5
 * 1 <= k <= nums.length
 *
 *
 */

var n int
var prefixSum []int

/*
 * func numberOfSubarrays(nums []int, k int) int {
 * 	n = len(nums)
 * 	// assert n > 0 && n <= 5e4;
 * 	// assert k > 0;
 * 	// C("恰好 k 个") <=> C("至少 k 个") - C("至少 k-1 个")
 * 	return _numberOfSubArrays(nums, k) - _numberOfSubArrays(nums, k-1)
 * }
 */

func numberOfSubarrays(nums []int, k int) int {
	n = len(nums)
	prefixSum = make([]int, n+1)
	prefixSum[0] = 0
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + (nums[i] & 1)
	}
	cntMap := make(map[int][]int)
	for i, p := range prefixSum {
		cntMap[p] = append(cntMap[p], i)
	}
	ans := 0
	for i := 0; i <= n; i++ {
		/*
		 * // 暴力会超时的... 所以加个 Hash 吧 (
		 * for j := i + 1; j <= n; j++ {
		 * 	if prefixSum[j]-prefixSum[i] == k {
		 * 		ans++
		 * 	}
		 * }
		 */
		 // 谢天谢地如果 Hash 还不行可能还得在 Hash 里面套个二分了...
		 for _, j := range cntMap[prefixSum[i] + k] {
			 if j > i {
				 ans++
			 }
		 }
	}
	return ans
}

// 返回最多 k 个奇数的子数组数目.
func _numberOfSubArrays(arr []int, k int) int {
	// [p, q] 的滑动窗口, 注意这里是闭区间 (0 <= p <= q < n)...
	p, q := 0, 0
	ans := 0
	c := 0
	// 枚举每个起点...
	for p < n {
		// 找到最多 k 个奇数的子数组的右边界:
		// 如果:
		// - p == n (到末尾了)
		// - or c > k (奇数个数爆了)
		// 退出.
		for ; q < n && c <= k; q++ {
			if arr[q]&1 != 0 {
				c++
			}
		}
		// 如果到末尾了也就不管了.
		// 如果奇数个数爆了, p <= n 一定成立, 而且 arr[p-1] 一定是奇数. 减一下.
		if c > k {
			q, c = q-1, c-1
		}
		// 统计当前起点的子数组数...
		ans += q - p + 1
		// 向前移动指针, 别忘了维护计数...
		if arr[p]&1 != 0 {
			c--
		}
		p++
	}
	return ans
}
