/*
 * @lc app=leetcode.cn id=number-of-subarrays-with-lcm-equal-to-k lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6234] 最小公倍数为 K 的子数组数目
 *
 * https://leetcode.cn/problems/number-of-subarrays-with-lcm-equal-to-k/description/
 *
 * algorithms
 * Medium (33.31%)
 * Total Accepted:    3.4K
 * Total Submissions: 10.2K
 * Testcase Example:  '[3,6,2,7,1]\n6'
 *
 * 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 nums 的 子数组 中满足 元素最小公倍数为 k 的子数组数目。
 *
 * 子数组 是数组中一个连续非空的元素序列。
 *
 * 数组的最小公倍数 是可被所有数组元素整除的最小正整数。
 *
 *
 *
 * 示例 1 ：
 *
 * 输入：nums = [3,6,2,7,1], k = 6
 * 输出：4
 * 解释：以 6 为最小公倍数的子数组是：
 * - [3,6,2,7,1]
 * - [3,6,2,7,1]
 * - [3,6,2,7,1]
 * - [3,6,2,7,1]
 *
 *
 * 示例 2 ：
 *
 * 输入：nums = [3], k = 2
 * 输出：0
 * 解释：不存在以 2 为最小公倍数的子数组。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 1000
 * 1 <= nums[i], k <= 1000
 *
 *
 */
func subarrayLCM(nums []int, k int) (ans int) {
	for p, n := 0, len(nums); p < n; p++ {
		lcm := nums[p]
		for q := p; q < n; q++ {
			lcm = LCM(lcm, nums[q])
			if lcm == k {
				ans++
			}
		}
	}
	return
}

func LCM(x, y int) int { return x * y / GCD(x, y) }

func GCD(x, y int) int {
	if y == 0 {
		return x
	}
	return GCD(y, x%y)
}
