/*
 * @lc app=leetcode.cn id=maximum-xor-after-operations lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6105] 操作后的最大异或和
 *
 * https://leetcode.cn/problems/maximum-xor-after-operations/description/
 *
 * algorithms
 * Medium (80.87%)
 * Total Accepted:    2.3K
 * Total Submissions: 2.9K
 * Testcase Example:  '[3,2,4,6]'
 *
 * 给你一个下标从 0 开始的整数数组 nums 。一次操作中，选择 任意 非负整数 x 和一个下标 i ，更新 nums[i] 为 nums[i] AND
 * (nums[i] XOR x) 。
 *
 * 注意，AND 是逐位与运算，XOR 是逐位异或运算。
 *
 * 请你执行 任意次 更新操作，并返回 nums 中所有元素 最大 逐位异或和。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [3,2,4,6]
 * 输出：7
 * 解释：选择 x = 4 和 i = 3 进行操作，num[3] = 6 AND (6 XOR 4) = 6 AND 2 = 2 。
 * 现在，nums = [3, 2, 4, 2] 且所有元素逐位异或得到 3 XOR 2 XOR 4 XOR 2 = 7 。
 * 可知 7 是能得到的最大逐位异或和。
 * 注意，其他操作可能也能得到逐位异或和 7 。
 *
 * 示例 2：
 *
 * 输入：nums = [1,2,3,9,2]
 * 输出：11
 * 解释：执行 0 次操作。
 * 所有元素的逐位异或和为 1 XOR 2 XOR 3 XOR 9 XOR 2 = 11 。
 * 可知 11 是能得到的最大逐位异或和。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * 0 <= nums[i] <= 10^8
 *
 *
 */
func maximumXOR(nums []int) int {
	ans := 0
	for i := 31; i >= 0; i-- {
		found := false
		for _, n := range nums {
			if n&(1<<i) != 0 {
				found = true
				break
			}
		}
		if found {
			ans |= 1 << i
		}
	}
	return ans
}
