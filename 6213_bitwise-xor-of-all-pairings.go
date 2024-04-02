/*
 * @lc app=leetcode.cn id=bitwise-xor-of-all-pairings lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6213] 所有数对的异或和
 *
 * https://leetcode.cn/problems/bitwise-xor-of-all-pairings/description/
 *
 * algorithms
 * Medium (60.06%)
 * Total Accepted:    2.5K
 * Total Submissions: 4.1K
 * Testcase Example:  '[2,1,3]\n[10,2,5,0]'
 *
 * 给你两个下标从 0 开始的数组 nums1 和 nums2 ，两个数组都只包含非负整数。请你求出另外一个数组 nums3 ，包含 nums1 和
 * nums2 中 所有数对 的异或和（nums1 中每个整数都跟 nums2 中每个整数 恰好 匹配一次）。
 *
 * 请你返回 nums3 中所有整数的 异或和 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums1 = [2,1,3], nums2 = [10,2,5,0]
 * 输出：13
 * 解释：
 * 一个可能的 nums3 数组是 [8,0,7,2,11,3,4,1,9,1,6,3] 。
 * 所有这些数字的异或和是 13 ，所以我们返回 13 。
 *
 *
 * 示例 2：
 *
 * 输入：nums1 = [1,2], nums2 = [3,4]
 * 输出：0
 * 解释：
 * 所有数对异或和的结果分别为 nums1[0] ^ nums2[0] ，nums1[0] ^ nums2[1] ，nums1[1] ^ nums2[0]
 * 和 nums1[1] ^ nums2[1] 。
 * 所以，一个可能的 nums3 数组是 [2,5,1,6] 。
 * 2 ^ 5 ^ 1 ^ 6 = 0 ，所以我们返回 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums1.length, nums2.length <= 10^5
 * 0 <= nums1[i], nums2[j] <= 10^9
 *
 *
 */
func xorAllNums(n1 []int, n2 []int) (ans int) {
	for d := 0; d < 32; d++ {
		c10, c11, c20, c21 := 0, 0, 0, 0
		for _, n := range n1 {
			if n&(1<<d) != 0 {
				c11++
			} else {
				c10++
			}
		}
		for _, n := range n2 {
			if n&(1<<d) != 0 {
				c21++
			} else {
				c20++
			}
		}
		c := c11*c20 + c21*c10
		if c&1 != 0 {
			ans |= 1 << d
		}
	}
	return
}
