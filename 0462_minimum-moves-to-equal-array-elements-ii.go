/*
 * @lc app=leetcode.cn id=minimum-moves-to-equal-array-elements-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [462] 最少移动次数使数组元素相等 II
 *
 * https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements-ii/description/
 *
 * algorithms
 * Medium (61.42%)
 * Total Accepted:    26.4K
 * Total Submissions: 42.7K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个长度为 n 的整数数组 nums ，返回使所有数组元素相等需要的最少移动数。
 *
 * 在一步操作中，你可以使数组中的一个元素加 1 或者减 1 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：2
 * 解释：
 * 只需要两步操作（每步操作指南使一个元素加 1 或减 1）：
 * [1,2,3]  =>  [2,2,3]  =>  [2,2,2]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,10,2,9]
 * 输出：16
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 1 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 *
 *
 */
func minMoves2(nums []int) int {
	k := getMedian(nums)
	ans := 0
	for _, n := range nums {
		ans += abs(n - k)
	}
	return ans
}

func getMedian(arr []int) int {
	n := len(arr)
	ans := 0
	for p, q := 0, n; p < q; {
		if k := partition(arr, p, q); k > n/2 {
			q = k
		} else if k < n/2 {
			p = k + 1
		} else {
			ans = arr[k]
			break
		}
	}
	if n%2 != 0 {
		return ans
	}
	for p, q := 0, n; p < q; {
		if k := partition(arr, p, q); k > n/2-1 {
			q = k
		} else if k < n/2-1 {
			p = k + 1
		} else {
			ans += arr[k]
			break
		}
	}
	return ans / 2
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
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
