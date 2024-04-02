/*
 * @lc app=leetcode.cn id=maximum-score-of-spliced-array lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [5229] 拼接数组的最大分数
 *
 * https://leetcode.cn/problems/maximum-score-of-spliced-array/description/
 *
 * algorithms
 * Hard (44.76%)
 * Total Accepted:    3.8K
 * Total Submissions: 8.4K
 * Testcase Example:  '[60,60,60]\n[10,90,10]'
 *
 * 给你两个下标从 0 开始的整数数组 nums1 和 nums2 ，长度都是 n 。
 *
 * 你可以选择两个整数 left 和 right ，其中 0 <= left <= right < n ，接着 交换 两个子数组
 * nums1[left...right] 和 nums2[left...right] 。
 *
 *
 * 例如，设 nums1 = [1,2,3,4,5] 和 nums2 = [11,12,13,14,15] ，整数选择 left = 1 和 right =
 * 2，那么 nums1 会变为 [1,12,13,4,5] 而 nums2 会变为 [11,2,3,14,15] 。
 *
 *
 * 你可以选择执行上述操作 一次 或不执行任何操作。
 *
 * 数组的 分数 取 sum(nums1) 和 sum(nums2) 中的最大值，其中 sum(arr) 是数组 arr 中所有元素之和。
 *
 * 返回 可能的最大分数 。
 *
 * 子数组 是数组中连续的一个元素序列。arr[left...right] 表示子数组包含 nums 中下标 left 和 right 之间的元素（含 下标
 * left 和 right 对应元素）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [60,60,60], nums2 = [10,90,10]
 * 输出：210
 * 解释：选择 left = 1 和 right = 1 ，得到 nums1 = [60,90,60] 和 nums2 = [10,60,10] 。
 * 分数为 max(sum(nums1), sum(nums2)) = max(210, 80) = 210 。
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [20,40,20,70,30], nums2 = [50,20,50,40,20]
 * 输出：220
 * 解释：选择 left = 3 和 right = 4 ，得到 nums1 = [20,40,20,40,20] 和 nums2 =
 * [50,20,50,70,30] 。
 * 分数为 max(sum(nums1), sum(nums2)) = max(140, 220) = 220 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums1 = [7,11,13], nums2 = [1,1,1]
 * 输出：31
 * 解释：选择不交换任何子数组。
 * 分数为 max(sum(nums1), sum(nums2)) = max(31, 3) = 31 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums1.length == nums2.length
 * 1 <= n <= 10^5
 * 1 <= nums1[i], nums2[i] <= 10^4
 *
 *
 */

import "math"

func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	n := len(nums1)
	ps1 := make([]int, n+1)
	ps2 := make([]int, n+1)
	for i := 0; i < n; i++ {
		ps1[i+1] = ps1[i] + nums1[i]
		ps2[i+1] = ps2[i] + nums2[i]
	}
	minBma, minAmb := math.MaxInt32, math.MaxInt32
	ans := max(ps1[n]-ps1[0], ps2[n]-ps2[0])
	for i := 0; i < n; i++ {
		v := ps2[i] - ps1[i]
		v1 := ps1[n] - ps1[0] + v - minBma
		v2 := ps2[n] - ps2[0] - v - minAmb
		ans = max(ans, max(v1, v2))
		minBma = min(minBma, v)
		minAmb = min(minAmb, -v)
	}
	return ans
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

// \Sigma_{i=1}^n{a[i]} \Sigma_{i=1}^n{b[i]}
// => \Sigma{i=1}^{left}{a[i]} + \Sigma{i=right}^{n}{a[i]} + \Sigma{i=left}^{right}{b[i]} = psa[left]-psa[0] + psa[n]-psa[right] + psb[right] - psb[left] = (psa[n] - psa[0]) + (psb[right] - psa[right]) - (psb[left] - psa[left])
// => \Sigma{i=1}^{left}{b[i]} + \Sigma{i=right}^{n}{b[i]} + \Sigma{i=left}^{right}{a[i]} = psb[left]-psb[0] + psb[n]-psb[right] + psa[right] - psa[left] = (psb[n] - psb[0]) + (psa[right] - psb[right]) - (psa[left] - psb[left])
