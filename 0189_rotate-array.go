/*
 * @lc app=leetcode.cn id=rotate-array lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [189] 轮转数组
 *
 * https://leetcode-cn.com/problems/rotate-array/description/
 *
 * algorithms
 * Medium (44.36%)
 * Total Accepted:    475.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3,4,5,6,7]\n3'
 *
 * 给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,2,3,4,5,6,7], k = 3
 * 输出: [5,6,7,1,2,3,4]
 * 解释:
 * 向右轮转 1 步: [7,1,2,3,4,5,6]
 * 向右轮转 2 步: [6,7,1,2,3,4,5]
 * 向右轮转 3 步: [5,6,7,1,2,3,4]
 *
 *
 * 示例 2:
 *
 *
 * 输入：nums = [-1,-100,3,99], k = 2
 * 输出：[3,99,-1,-100]
 * 解释:
 * 向右轮转 1 步: [99,-1,-100,3]
 * 向右轮转 2 步: [3,99,-1,-100]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * -2^31 <= nums[i] <= 2^31 - 1
 * 0 <= k <= 10^5
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
 * 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
 *
 *
 *
 *
 *
 *
 *
 *
 */
// 编程珠玑里面的技巧, 一时间没反应过来...
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	if k == 0 || n == 0 {
		return
	}
	_rotate(nums, 0, n)
	_rotate(nums, 0, k)
	_rotate(nums, k, n)
}

func _rotate(arr []int, p, q int) {
	for i, j := p, q-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

/*
 * // 额外空间的:
 * func rotate(nums []int, k int)  {
 *     n := len(nums)
 *     k %= n
 *     if k == 0 || n == 0 {
 *         return
 *     }
 *     arr := make([]int, n)
 *     copy(arr[:k], nums[n-k:])
 *     copy(arr[k:], nums[:n-k])
 *     copy(nums, arr)
 * }
 */
