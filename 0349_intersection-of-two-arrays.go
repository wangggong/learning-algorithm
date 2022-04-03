/*
 * @lc app=leetcode.cn id=intersection-of-two-arrays lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [349] 两个数组的交集
 *
 * https://leetcode-cn.com/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (73.88%)
 * Total Accepted:    267.9K
 * Total Submissions: 362.6K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出：[2]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出：[9,4]
 * 解释：[4,9] 也是可通过的
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums1.length, nums2.length <= 1000
 * 0 <= nums1[i], nums2[i] <= 1000
 *
 *
 */

const MAXN int = 1000

var count [MAXN + 5]bool

func intersection(nums1 []int, nums2 []int) []int {
	for i := range count {
		count[i] = false
	}
	for _, n := range nums1 {
		count[n] = true
	}
	var ans []int
	for _, n := range nums2 {
		if count[n] {
			ans = append(ans, n)
		}
		count[n] = false
	}
	return ans
}
