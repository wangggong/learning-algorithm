/*
 * @lc app=leetcode.cn id=maximum-product-after-k-increments lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6039] K 次增加后的最大乘积
 *
 * https://leetcode-cn.com/problems/maximum-product-after-k-increments/description/
 *
 * algorithms
 * Medium (32.70%)
 * Total Accepted:    4.9K
 * Total Submissions: 14.9K
 * Testcase Example:  '[0,4]\n5'
 *
 * 给你一个非负整数数组 nums 和一个整数 k 。每次操作，你可以选择 nums 中 任一 元素并将它 增加 1 。
 *
 * 请你返回 至多 k 次操作后，能得到的 nums的 最大乘积 。由于答案可能很大，请你将答案对 10^9 + 7 取余后返回。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [0,4], k = 5
 * 输出：20
 * 解释：将第一个数增加 5 次。
 * 得到 nums = [5, 4] ，乘积为 5 * 4 = 20 。
 * 可以证明 20 是能得到的最大乘积，所以我们返回 20 。
 * 存在其他增加 nums 的方法，也能得到最大乘积。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [6,3,3,2], k = 2
 * 输出：216
 * 解释：将第二个数增加 1 次，将第四个数增加 1 次。
 * 得到 nums = [6, 4, 3, 3] ，乘积为 6 * 4 * 3 * 3 = 216 。
 * 可以证明 216 是能得到的最大乘积，所以我们返回 216 。
 * 存在其他增加 nums 的方法，也能得到最大乘积。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length, k <= 10^5
 * 0 <= nums[i] <= 10^6
 *
 *
 */
import (
	"container/heap"
	"sort"
)

const mod int = 1e9 + 7

type Heap struct {
	sort.IntSlice
}

func (h *Heap) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *Heap) Pop() interface{} {
	l := len(h.IntSlice)
	v := h.IntSlice[l-1]
	h.IntSlice = h.IntSlice[:l-1]
	return v
}

func maximumProduct(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0] + k
	}
	h := &Heap{}
	for _, n := range nums {
		heap.Push(h, n)
	}
	for k > 0 {
		p := heap.Pop(h).(int)
		q := heap.Pop(h).(int)
		diff := max(min(k, q-p), 1)
		p += diff
		k -= diff
		heap.Push(h, p)
		heap.Push(h, q)
	}
	ans := 1
	for h.Len() > 0 {
		ans = (ans * heap.Pop(h).(int)) % mod
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
