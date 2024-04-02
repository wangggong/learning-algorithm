/*
 * @lc app=leetcode.cn id=n-repeated-element-in-size-2n-array lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [961] 在长度 2N 的数组中找出重复 N 次的元素
 *
 * https://leetcode-cn.com/problems/n-repeated-element-in-size-2n-array/description/
 *
 * algorithms
 * Easy (67.66%)
 * Total Accepted:    50K
 * Total Submissions: 72.5K
 * Testcase Example:  '[1,2,3,3]'
 *
 * 给你一个整数数组 nums ，该数组具有以下属性：
 *
 *
 *
 *
 * nums.length == 2 * n.
 * nums 包含 n + 1 个 不同的 元素
 * nums 中恰有一个元素重复 n 次
 *
 *
 * 找出并返回重复了 n 次的那个元素。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3,3]
 * 输出：3
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [2,1,2,5,3,2]
 * 输出：2
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [5,1,5,2,5,3,5,4]
 * 输出：5
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= n <= 5000
 * nums.length == 2 * n
 * 0 <= nums[i] <= 10^4
 * nums 由 n + 1 个 不同的 元素组成，且其中一个元素恰好重复 n 次
 *
 *
 */
func repeatedNTimes(nums []int) int {
	n := len(nums)
	quickSelect(nums, n/2)
	quickSelect(nums, n/2-1)
	if nums[n/2-1] == nums[n/2] {
		return nums[n/2-1]
	}
	l, r := 0, 0
	for _, k := range nums {
		if k == nums[n/2-1] {
			l++
		} else if k == nums[n/2] {
			r++
		} else {
		}
	}
	if l >= n/2 {
		return nums[n/2-1]
	}
	return nums[n/2]
}

func quickSelect(arr []int, k int) {
	for p, q := 0, len(arr); p < q; {
		m := partition(arr, p, q)
		if m < k {
			p = m + 1
		} else if m > k {
			q = m
		} else {
			break
		}
	}
	return
}

func partition(arr []int, l, r int) int {
	start := l
	for i := l + 1; i < r; i++ {
		if arr[i] < arr[l] {
			start++
			arr[i], arr[start] = arr[start], arr[i]
		}
	}
	arr[l], arr[start] = arr[start], arr[l]
	return start
}
