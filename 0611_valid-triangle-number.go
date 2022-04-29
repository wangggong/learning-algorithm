import "sort"

/*
 * @lc app=leetcode.cn id=valid-triangle-number lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [611] 有效三角形的个数
 *
 * https://leetcode-cn.com/problems/valid-triangle-number/description/
 *
 * algorithms
 * Medium (53.33%)
 * Total Accepted:    63.2K
 * Total Submissions: 118.4K
 * Testcase Example:  '[2,2,3,4]'
 *
 * 给定一个包含非负整数的数组 nums ，返回其中可以组成三角形三条边的三元组个数。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [2,2,3,4]
 * 输出: 3
 * 解释:有效的组合是:
 * 2,3,4 (使用第一个 2)
 * 2,3,4 (使用第二个 2)
 * 2,2,3
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [4,2,3,4]
 * 输出: 4
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 1000
 * 0 <= nums[i] <= 1000
 *
 *
 */

/*
 * // 排序 + 二分, `O(n^2logn)`
 * func triangleNumber(nums []int) int {
 * 	// a <= b <= c
 * 	// a + b > c => b > c-a
 * 	sort.Ints(nums)
 * 	n := len(nums)
 * 	ans := 0
 * 	for p := 0; p < n; p++ {
 * 		for q := p + 2; q < n; q++ {
 * 			l, r := p+1, q
 * 			for l < r {
 * 				mid := (l + r) >> 1
 * 				if nums[mid] > nums[q]-nums[p] {
 * 					r = mid
 * 				} else {
 * 					l = mid + 1
 * 				}
 * 			}
 * 			// fmt.Println(nums, p, q, l)
 * 			ans += q - l
 * 		}
 * 	}
 * 	return ans
 * }
 */

// 排序 + 双指针, `O(n^2)`
func triangleNumber(nums []int) int {
	// a <= b <= c
	// a + b > c => b > c-a
	sort.Ints(nums)
	n := len(nums)
	ans := 0
	for i := 0; i < n; i++ {
		k := i + 2
		for j := i + 1; j < n; j++ {
			for ; k < n && nums[k]-nums[i] < nums[j]; k++ {
			}
			ans += max(k-j-1, 0)
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
