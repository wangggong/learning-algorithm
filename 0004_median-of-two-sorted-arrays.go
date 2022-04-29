/*
 * @lc app=leetcode.cn id=median-of-two-sorted-arrays lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [4] 寻找两个正序数组的中位数
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (41.13%)
 * Total Accepted:    619.6K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
 *
 * 算法的时间复杂度应该为 O(log (m+n)) 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,3], nums2 = [2]
 * 输出：2.00000
 * 解释：合并数组 = [1,2,3] ，中位数 2
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [1,2], nums2 = [3,4]
 * 输出：2.50000
 * 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * nums1.length == m
 * nums2.length == n
 * 0 <= m <= 1000
 * 0 <= n <= 1000
 * 1 <= m + n <= 2000
 * -10^6 <= nums1[i], nums2[i] <= 10^6
 *
 *
 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	tot := len(nums1) + len(nums2)
	midIndex := tot / 2
	switch tot % 2 {
	case 0:
		// 偶数个: 第 `k` 小和第 `k+1` 小的数字的平均值, `k = tot / 2`.
		return float64(findKthElem(nums1, nums2, midIndex)+findKthElem(nums1, nums2, midIndex+1)) / 2
	case 1:
		// 奇数个: 第 `k+1` 小的数字, 同上.
		return float64(findKthElem(nums1, nums2, midIndex+1))
	default:
		// unreachable
	}
	// unreachable
	return 0
}

// 返回两数组第 `k` 小元素.
func findKthElem(arr1, arr2 []int, k int) int {
	n, m := len(arr1), len(arr2)
	if n == 0 {
		return arr2[k-1]
	}
	if m == 0 {
		return arr1[k-1]
	}
	// 返回两个数组中最小的元素.
	if k == 1 {
		return min(arr1[0], arr2[0])
	}
	// 分别在两数组内取 `k/2` 个元素.
	// 不够的话能取多少取多少.
	p, q := min(k/2, n)-1, min(k/2, m)-1
	if arr1[p] < arr2[q] {
		// `arr1[:p+1]` 的位置均可以抛弃了; 抛弃了 `p+1` 个元素.
		// 问题转化为剩下数据中第 `k-(p+1)` 小的元素.
		return findKthElem(arr1[p+1:], arr2, k-(p+1))
	} else {
		return findKthElem(arr1, arr2[q+1:], k-(q+1))
	}
	// unreachable
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}


// 二刷卡住了, 哭了. 瞟一眼答案接着写.
