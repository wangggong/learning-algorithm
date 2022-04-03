/*
 * @lc app=leetcode.cn id=kth-largest-element-in-an-array lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (64.71%)
 * Total Accepted:    544.9K
 * Total Submissions: 842.5K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
 *
 * 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: [3,2,1,5,6,4] 和 k = 2
 * 输出: 5
 *
 *
 * 示例 2:
 *
 *
 * 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
 * 输出: 4
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= nums.length <= 1e4
 * -1e4 <= nums[i] <= 1e4
 *
 *
 */

/*
 * import (
 * 	"container/heap"
 * 	"sort"
 * )
 *
 * type Heap struct {
 * 	sort.IntSlice
 * }
 *
 * func (h Heap) Less(p, q int) bool  { return h.IntSlice[p] > h.IntSlice[q] }
 * func (h *Heap) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
 * func (h *Heap) Pop() interface{} {
 * 	l := len(h.IntSlice)
 * 	v := h.IntSlice[l-1]
 * 	h.IntSlice = h.IntSlice[:l-1]
 * 	return v
 * }
 *
 * func findKthLargest(nums []int, k int) int {
 * 	h := &Heap{}
 * 	for _, n := range nums {
 * 		heap.Push(h, n)
 * 	}
 * 	for ; k-1 > 0; k-- {
 * 		heap.Pop(h)
 * 	}
 * 	return heap.Pop(h).(int)
 * }
 */

func findKthLargest(nums []int, k int) int {
	// partition := func(left, right int) int {
	// 	v := nums[left]
	// 	nums[left], nums[right] = nums[right], nums[left]
	// 	p, q := left, right
	// 	for p < q {
	// 		for p < q && nums[p] <= v {
	// 			p++
	// 		}
	// 		for p < q && nums[q] >= v {
	// 			q--
	// 		}
	// 		if p < q {
	// 			nums[p], nums[q] = nums[q], nums[p]
	// 		}
	// 	}
	// 	nums[right], nums[p] = nums[p], nums[right]
	// 	return p
	// }

	// 另一版本的 partition
	partition := func(left, right int) int {
		v := nums[left]
		start := left
		// 遍历 `[left+1, right]` 区间:
		for i := left + 1; i <= right; i++ {
			// 如果当前值小于左侧值 `v`, 交换. 保证 `[left, start]` 的值小于 `v`.
			if nums[i] < v {
				start++
				nums[i], nums[start] = nums[start], nums[i]
			}
		}
		// 最终把 `v` 放到应该在的位置.
		nums[start], nums[left] = nums[left], nums[start]
		return start
	}

	// 3 1 2 5 6 4
	// 2 1 3 5 6 4
	// ----- 5 6 4
	// ----- 4 5 6
	n := len(nums)
	p, q := 0, n-1
	for p < q {
		ind := partition(p, q)
		if ind == n-k {
			return nums[ind]
		} else if ind > n-k {
			q = ind - 1
		} else {
			p = ind + 1
		}
	}
	return nums[p]
}

