/*
 * @lc app=leetcode.cn id=peak-index-in-a-mountain-array lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [852] 山脉数组的峰顶索引
 *
 * https://leetcode-cn.com/problems/peak-index-in-a-mountain-array/description/
 *
 * algorithms
 * Easy (71.32%)
 * Total Accepted:    92.4K
 * Total Submissions: 129.7K
 * Testcase Example:  '[0,1,0]'
 *
 * 符合下列属性的数组 arr 称为 山脉数组 ：
 *
 * arr.length >= 3
 * 存在 i（0 < i < arr.length - 1）使得：
 *
 * arr[0] < arr[1] < ... arr[i-1] < arr[i]
 * arr[i] > arr[i+1] > ... > arr[arr.length - 1]
 *
 *
 *
 *
 * 给你由整数组成的山脉数组 arr ，返回任何满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i +
 * 1] > ... > arr[arr.length - 1] 的下标 i 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：arr = [0,1,0]
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 * 输入：arr = [0,2,1,0]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：arr = [0,10,5,2]
 * 输出：1
 *
 *
 * 示例 4：
 *
 *
 * 输入：arr = [3,4,5,1]
 * 输出：2
 *
 *
 * 示例 5：
 *
 *
 * 输入：arr = [24,69,100,99,79,78,67,36,26,19]
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= arr.length <= 1e4
 * 0 <= arr[i] <= 1e6
 * 题目数据保证 arr 是一个山脉数组
 *
 *
 *
 *
 * 进阶：很容易想到时间复杂度 O(n) 的解决方案，你可以设计一个 O(log(n)) 的解决方案吗？
 *
 */

const MAXN = 1e6 + 10

func peakIndexInMountainArray(arr []int) int {
	n := len(arr)
	// assert n >= 3;
	p, q := 0, n-1
	for p <= q {
		mid := (p + q) >> 1
		if get(arr, mid) > get(arr, mid+1) && get(arr, mid) > get(arr, mid-1) {
			return mid
		}
		if get(arr, mid) < get(arr, mid+1) {
			p = mid + 1
			continue
		}
		q = mid - 1
		continue
	}
	// unreachable
	return -1
}

func get(arr []int, index int) int {
	n := len(arr)
	// assert index >= -1 && index <= n;
	if index == -1 || index == n {
		return -MAXN
	}
	return arr[index]
}
