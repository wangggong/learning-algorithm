/*
 * @lc app=leetcode.cn id=4sum lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [18] 四数之和
 *
 * https://leetcode-cn.com/problems/4sum/description/
 *
 * algorithms
 * Medium (39.49%)
 * Total Accepted:    264.1K
 * Total Submissions: 668.7K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a],
 * nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
 *
 *
 * 0 <= a, b, c, d < n
 * a、b、c 和 d 互不相同
 * nums[a] + nums[b] + nums[c] + nums[d] == target
 *
 *
 * 你可以按 任意顺序 返回答案 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,0,-1,0,-2,2], target = 0
 * 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [2,2,2,2,2], target = 8
 * 输出：[[2,2,2,2]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 200
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 *
 *
 */

import "sort"

/*
 * func nSumTarget(arr []int, start, size, target, n int) (ans [][]int) {
 * 	if n < 2 || start > size {
 * 		return
 * 	}
 * 	if n == 2 {
 * 		p, q := start, size-1
 * 		for p < q {
 * 			left, right := arr[p], arr[q]
 * 			switch v := target - (left + right); {
 * 			case v == 0:
 * 				ans = append(ans, []int{left, right})
 * 				for p < q && arr[p] == left {
 * 					p++
 * 				}
 * 				for p < q && arr[q] == right {
 * 					q--
 * 				}
 * 			case v < 0:
 * 				for p < q && arr[q] == right {
 * 					q--
 * 				}
 * 			case v > 0:
 * 				for p < q && arr[p] == left {
 * 					p++
 * 				}
 * 			default:
 * 			}
 * 		}
 * 		return
 * 	}
 * 	for i := start; i < size; i++ {
 * 		as := nSumTarget(arr, i+1, size, target-arr[i], n-1)
 * 		for _, a := range as {
 * 			a = append(a, arr[i])
 * 			ans = append(ans, a)
 * 		}
 * 		for i+1 < size && arr[i+1] == arr[i] {
 * 			i++
 * 		}
 * 	}
 * 	return
 * }
 *
 * func fourSum(nums []int, target int) [][]int {
 * 	sort.Ints(nums)
 * 	return nSumTarget(nums, 0, len(nums), target, 4)
 * }
 */

func nSumTarget(arr []int, target, n int) [][]int {
	size := len(arr)
	var ans [][]int
	if size < n {
		return ans
	}
	if n == 2 {
		p, q := 0, len(arr)-1
		for p < q {
			if p > 0 && arr[p] == arr[p-1] {
				p++
				continue
			}
			switch c := arr[p] + arr[q] - target; {
			case c == 0:
				ans = append(ans, []int{arr[p], arr[q]})
				p, q = p+1, q-1
			case c > 0:
				q--
			case c < 0:
				p++
			default:
				// unreachable
			}
		}
		return ans
	}
	for i := 0; i < size; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		as := nSumTarget(arr[i+1:], target-arr[i], n-1)
		for _, a := range as {
			a = append([]int{arr[i]}, a...)
			ans = append(ans, a)
		}
	}
	return ans
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSumTarget(nums, target, 4)
}
