/*
 * @lc app=leetcode.cn id=sort-colors lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [75] 颜色分类
 *
 * https://leetcode-cn.com/problems/sort-colors/description/
 *
 * algorithms
 * Medium (60.06%)
 * Total Accepted:    336.8K
 * Total Submissions: 560.9K
 * Testcase Example:  '[2,0,2,1,1,0]'
 *
 * 给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
 *
 * 我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
 *
 *
 *
 *
 * 必须在不使用库的sort函数的情况下解决这个问题。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,0,2,1,1,0]
 * 输出：[0,0,1,1,2,2]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [2,0,1]
 * 输出：[0,1,2]
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 1 <= n <= 300
 * nums[i] 为 0、1 或 2
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 你可以不使用代码库中的排序函数来解决这道题吗？
 * 你能想出一个仅使用常数空间的一趟扫描算法吗？
 *
 *
 */

func sortColors(nums []int) {
	// 3-way partition
	n := len(nums)
	// `i`, `j`, `k` 指向 `0` 的终点, `1` 的起点, `2` 的起点.
	i, j, k := 0, 0, n-1
	/*
	 * for j <= k {
	 * 	switch c := nums[j] - 1; {
	 * 	case c > 0:
	 * 		nums[j], nums[k] = nums[k], nums[j]
	 * 		k--
	 * 	case c < 0:
	 * 		nums[j], nums[i] = nums[i], nums[j]
	 * 		i++
	 * 		fallthrough
	 * 	case c == 0:
	 * 		j++
	 * 	default:
	 * 		// unreachable
	 * 	}
	 * }
	 */

	// 说人话:
	for j <= k {
		if nums[j] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else if nums[j] == 2 {
			nums[j], nums[k] = nums[k], nums[j]
			k--
		} else {
			j++
		}
	}
	return
}
