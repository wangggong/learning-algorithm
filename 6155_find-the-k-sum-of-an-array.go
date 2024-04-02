/*
 * @lc app=leetcode.cn id=find-the-k-sum-of-an-array lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6155] 找出数组的第 K 大和
 *
 * https://leetcode.cn/problems/find-the-k-sum-of-an-array/description/
 *
 * algorithms
 * Hard (18.16%)
 * Total Accepted:    842
 * Total Submissions: 3.4K
 * Testcase Example:  '[2,4,-2]\n5'
 *
 * 给你一个整数数组 nums 和一个 正 整数 k 。你可以选择数组的任一 子序列 并且对其全部元素求和。
 *
 * 数组的 第 k 大和 定义为：可以获得的第 k 个 最大 子序列和（子序列和允许出现重复）
 *
 * 返回数组的 第 k 大和 。
 *
 * 子序列是一个可以由其他数组删除某些或不删除元素排生而来的数组，且派生过程不改变剩余元素的顺序。
 *
 * 注意：空子序列的和视作 0 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [2,4,-2], k = 5
 * 输出：2
 * 解释：所有可能获得的子序列和列出如下，按递减顺序排列：
 * - 6、4、4、2、2、0、0、-2
 * 数组的第 5 大和是 2 。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [1,-2,3,4,-10,12], k = 16
 * 输出：10
 * 解释：数组的第 16 大和是 10 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 1 <= n <= 10^5
 * -10^9 <= nums[i] <= 10^9
 * 1 <= k <= min(2000, 2^n)
 *
 *
 */

import (
	"container/heap"
	"sort"
)

type Heap [][2]int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(p, q int) bool  { return h[p][0] < h[q][0] }
func (h Heap) Swap(p, q int)       { h[p], h[q] = h[q], h[p] }
func (h *Heap) Push(v interface{}) { *h = append(*h, v.([2]int)) }
func (h *Heap) Pop() interface{}   { l := len(*h); v := (*h)[l-1]; *h = (*h)[:l-1]; return v }

// 参考:
// - https://leetcode.cn/problems/find-the-k-sum-of-an-array/solution/zhuan-huan-dui-by-endlesscheng-8yiq/
// - https://leetcode.cn/problems/find-the-k-sum-of-an-array/solution/by-tsreaper-ps7w/
func kSum(nums []int, k int) int64 {
	tot := int64(0)
	for i, n := range nums {
		if n >= 0 {
			tot += int64(n)
		} else {
			nums[i] = -n
		}
	}
	sort.Ints(nums)
	h := &Heap{}
	heap.Push(h, [2]int{})
	for ; k > 1; k-- {
		v := heap.Pop(h).([2]int)
		if v[1] >= len(nums) {
			continue
		}
		heap.Push(h, [2]int{v[0] + nums[v[1]], v[1] + 1})
		if v[1] > 0 {
			heap.Push(h, [2]int{v[0] + nums[v[1]] - nums[v[1]-1], v[1] + 1})
		}
	}
	return tot - int64((*h)[0][0])
}
