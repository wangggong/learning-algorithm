/*
 * @lc app=leetcode.cn id=smallest-subarrays-with-maximum-bitwise-or lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6186] 按位或最大的最小子数组长度
 *
 * https://leetcode.cn/problems/smallest-subarrays-with-maximum-bitwise-or/description/
 *
 * algorithms
 * Medium (33.26%)
 * Total Accepted:    2.3K
 * Total Submissions: 6.8K
 * Testcase Example:  '[1,0,2,1,3]'
 *
 * 给你一个长度为 n 下标从 0 开始的数组 nums ，数组中所有数字均为非负整数。对于 0 到 n - 1 之间的每一个下标 i ，你需要找出
 * nums 中一个 最小 非空子数组，它的起始位置为 i （包含这个位置），同时有 最大 的 按位或运算值 。
 *
 *
 * 换言之，令 Bij 表示子数组 nums[i...j] 的按位或运算的结果，你需要找到一个起始位置为 i 的最小子数组，这个子数组的按位或运算的结果等于
 * max(Bik) ，其中 i <= k <= n - 1 。
 *
 *
 * 一个数组的按位或运算值是这个数组里所有数字按位或运算的结果。
 *
 * 请你返回一个大小为 n 的整数数组 answer，其中 answer[i]是开始位置为 i ，按位或运算结果最大，且 最短 子数组的长度。
 *
 * 子数组 是数组里一段连续非空元素组成的序列。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [1,0,2,1,3]
 * 输出：[3,3,2,2,1]
 * 解释：
 * 任何位置开始，最大按位或运算的结果都是 3 。
 * - 下标 0 处，能得到结果 3 的最短子数组是 [1,0,2] 。
 * - 下标 1 处，能得到结果 3 的最短子数组是 [0,2,1] 。
 * - 下标 2 处，能得到结果 3 的最短子数组是 [2,1] 。
 * - 下标 3 处，能得到结果 3 的最短子数组是 [1,3] 。
 * - 下标 4 处，能得到结果 3 的最短子数组是 [3] 。
 * 所以我们返回 [3,3,2,2,1] 。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [1,2]
 * 输出：[2,1]
 * 解释：
 * 下标 0 处，能得到最大按位或运算值的最短子数组长度为 2 。
 * 下标 1 处，能得到最大按位或运算值的最短子数组长度为 1 。
 * 所以我们返回 [2,1] 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 1 <= n <= 10^5
 * 0 <= nums[i] <= 10^9
 *
 *
 */
const D int = 32

func smallestSubarrays(nums []int) (ans []int) {
	n := len(nums)
	cnt := make([][D]int, n+1)
	for i := range nums {
		copy(cnt[i+1][:], cnt[i][:])
		for d := 0; d < D; d++ {
			if nums[i]&(1<<d) != 0 {
				cnt[i+1][d]++
			}
		}
	}
	// fmt.Println(cnt)
	ans = make([]int, n)
	tot := 0
	for i := n - 1; i >= 0; i-- {
		tot = tot | nums[i]
		p, q := i, n-1
		for p < q {
			mid := (p + q) >> 1
			if check(cnt, tot, mid, i) {
				q = mid
			} else {
				p = mid + 1
			}
		}
		ans[i] = p - i + 1
	}
	return
}

func check(cnt [][D]int, target, k, from int) bool {
	for d := 0; d < D; d++ {
		if target&(1<<d) == 0 {
			continue
		}
		if cnt[k+1][d] == cnt[from][d] {
			return false
		}
	}
	return true
}
