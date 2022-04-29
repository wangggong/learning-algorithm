/*
 * @lc app=leetcode.cn id=find-triangular-sum-of-an-array lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6034] 数组的三角和
 *
 * https://leetcode-cn.com/problems/find-triangular-sum-of-an-array/description/
 *
 * algorithms
 * Medium (79.83%)
 * Total Accepted:    4.4K
 * Total Submissions: 5.5K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给你一个下标从 0 开始的整数数组 nums ，其中 nums[i] 是 0 到 9 之间（两者都包含）的一个数字。
 *
 * nums 的 三角和 是执行以下操作以后最后剩下元素的值：
 *
 *
 * nums 初始包含 n 个元素。如果 n == 1 ，终止 操作。否则，创建 一个新的下标从 0 开始的长度为 n - 1 的整数数组 newNums
 * 。
 * 对于满足 0 <= i < n - 1 的下标 i ，newNums[i] 赋值 为 (nums[i] + nums[i+1]) % 10 ，%
 * 表示取余运算。
 * 将 newNums 替换 数组 nums 。
 * 从步骤 1 开始 重复 整个过程。
 *
 *
 * 请你返回 nums 的三角和。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：nums = [1,2,3,4,5]
 * 输出：8
 * 解释：
 * 上图展示了得到数组三角和的过程。
 *
 * 示例 2：
 *
 *
 * 输入：nums = [5]
 * 输出：5
 * 解释：
 * 由于 nums 中只有一个元素，数组的三角和为这个元素自己。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 1000
 * 0 <= nums[i] <= 9
 *
 *
 */
func triangularSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	for len(nums) > 1 {
		tmp := make([]int, 0, len(nums))
		for i := 0; i+1 < len(nums); i++ {
			tmp = append(tmp, (nums[i]+nums[i+1])%10)
		}
		nums = tmp
	}
	return nums[0]
}
