/*
 * @lc app=leetcode.cn id=find-k-th-smallest-pair-distance lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [719] 找出第 K 小的数对距离
 *
 * https://leetcode.cn/problems/find-k-th-smallest-pair-distance/description/
 *
 * algorithms
 * Hard (39.17%)
 * Total Accepted:    16.7K
 * Total Submissions: 39.7K
 * Testcase Example:  '[1,3,1]\n1'
 *
 * 数对 (a,b) 由整数 a 和 b 组成，其数对距离定义为 a 和 b 的绝对差值。
 *
 * 给你一个整数数组 nums 和一个整数 k ，数对由 nums[i] 和 nums[j] 组成且满足 0 <= i < j < nums.length
 * 。返回 所有数对距离中 第 k 小的数对距离。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,3,1], k = 1
 * 输出：0
 * 解释：数对和对应的距离如下：
 * (1,3) -> 2
 * (1,1) -> 0
 * (3,1) -> 2
 * 距离第 1 小的数对是 (1,1) ，距离为 0 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,1,1], k = 2
 * 输出：0
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1,6,1], k = 3
 * 输出：5
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 2 <= n <= 10^4
 * 0 <= nums[i] <= 10^6
 * 1 <= k <= n * (n - 1) / 2
 *
 *
 */

import "sort"

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	p, q := 0, nums[len(nums)-1]-nums[0]
	for p < q {
		mid := (p + q) >> 1
		if cntDist(nums, mid) >= k {
			q = mid
		} else {
			p = mid + 1
		}
	}
	return p
}

func cntDist(arr []int, k int) int {
	ans, n := 0, len(arr)
	for i, a := range arr {
		p, q := i, n
		for p < q {
			mid := (p + q) >> 1
			if arr[mid] > a+k {
				q = mid
			} else {
				p = mid + 1
			}
		}
		ans += q - i - 1
	}
	return ans
}
