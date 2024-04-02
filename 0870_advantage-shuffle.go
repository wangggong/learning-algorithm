/*
 * @lc app=leetcode.cn id=advantage-shuffle lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [870] 优势洗牌
 *
 * https://leetcode.cn/problems/advantage-shuffle/description/
 *
 * algorithms
 * Medium (47.61%)
 * Total Accepted:    47.3K
 * Total Submissions: 96.9K
 * Testcase Example:  '[2,7,11,15]\n[1,10,4,11]'
 *
 * 给定两个大小相等的数组 nums1 和 nums2，nums1 相对于 nums2 的优势可以用满足 nums1[i] > nums2[i] 的索引 i
 * 的数目来描述。
 *
 * 返回 nums1 的任意排列，使其相对于 nums2 的优势最大化。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [2,7,11,15], nums2 = [1,10,4,11]
 * 输出：[2,11,7,15]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [12,24,8,32], nums2 = [13,25,32,11]
 * 输出：[24,32,8,12]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums1.length <= 10^5
 * nums2.length == nums1.length
 * 0 <= nums1[i], nums2[i] <= 10^9
 *
 *
 */
import "sort"

func advantageCount(nums1 []int, nums2 []int) (ans []int) {
	n := len(nums1)
	ans = make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	infos := make([][2]int, 0, n)
	for i, n := range nums2 {
		infos = append(infos, [2]int{n, i})
	}
	sort.Slice(infos, func(p, q int) bool { return infos[p][0] > infos[q][0] })
	sort.Ints(nums1)
	k := n - 1
	for i := 0; i < n; i++ {
		for ; i < n && infos[i][0] >= nums1[k]; i++ {
		}
		if i >= n {
			break
		}
		ans[infos[i][1]] = nums1[k]
		k--
	}
	for i, a := range ans {
		if a < 0 {
			ans[i] = nums1[k]
			k--
		}
	}
	return
}
